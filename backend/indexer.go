package backend

import (
	"sync"

	"github.com/emef/tally/lib"
	"github.com/emef/tally/pb"
	"github.com/orcaman/concurrent-map"
	"github.com/petar/GoLLRB/llrb"
)

type Indexer struct {
	blocks               chan *pb.RecordBlock
	done                 chan interface{}
	nameSourceToIndexMap nameSourceMap
}

func CreateAndStartIndexer(blocks chan *pb.RecordBlock) *Indexer {
	indexer := &Indexer{
		blocks:               blocks,
		done:                 make(chan interface{}),
		nameSourceToIndexMap: newNameSourceMap()}

	go indexer.start()

	return indexer
}

func (indexer *Indexer) Stop() {
	indexer.done <- nil
}

func (indexer *Indexer) Get(
	name, source string,
	startEpochMinute, endEpochMinute int32) map[int32]*pb.CounterValues {

	index := indexer.getIndex(name, source, false)
	if index == nil {
		return map[int32]*pb.CounterValues{}
	} else {
		return index.query(startEpochMinute, endEpochMinute)
	}
}

func (indexer *Indexer) start() {
	defer close(indexer.done)

	for {
		select {
		case block := <-indexer.blocks:
			indexer.index(block)

		case <-indexer.done:
			return
		}
	}
}

func (indexer *Indexer) index(block *pb.RecordBlock) {
	nameCodeMapping := block.NameCodeMapping
	sourceCodeMapping := block.SourceCodeMapping

	for _, entry := range block.Entries {
		name := nameCodeMapping[entry.Key.NameCode]
		source := sourceCodeMapping[entry.Key.SourceCode]
		epochMinute := entry.Key.EpochMinute

		index := indexer.getIndex(name, source, true)
		index.insert(epochMinute, entry.Values)
	}
}

func (indexer *Indexer) getIndex(name, source string, create bool) *treeIndex {
	return indexer.nameSourceToIndexMap.getOrCreate(name, source, create)
}

// TODO: should use name/source code mapping
// 2-level map of {name: {source: treeIndex}}
type nameSourceMap cmap.ConcurrentMap

func newNameSourceMap() nameSourceMap {
	return nameSourceMap(cmap.New())
}

func (nameMap *nameSourceMap) getOrCreate(name, source string, create bool) *treeIndex {
	nameCmap := (*cmap.ConcurrentMap)(nameMap)

	var sourceCmap cmap.ConcurrentMap
	nameLookup, ok := nameCmap.Get(name)
	if ok {
		sourceCmap = nameLookup.(cmap.ConcurrentMap)
	} else {
		sourceCmap = cmap.New()
		nameCmap.Set(name, sourceCmap)
	}

	var index *treeIndex
	sourceLookup, ok := sourceCmap.Get(source)
	if ok {
		index = sourceLookup.(*treeIndex)
	} else if create {
		index = newTreeIndex()
		sourceCmap.Set(source, index)
	}

	return index
}

type treeItem struct {
	epochMinute int32
	values      *pb.CounterValues
}

func (item treeItem) Less(other llrb.Item) bool {
	return item.epochMinute < other.(treeItem).epochMinute
}

type treeIndex struct {
	tree *llrb.LLRB
	lock sync.RWMutex
}

func newTreeIndex() *treeIndex {
	return &treeIndex{tree: llrb.New()}
}

func (index *treeIndex) insert(epochMinute int32, values *pb.CounterValues) {
	index.lock.Lock()
	defer index.lock.Unlock()

	item := treeItem{epochMinute, values}
	replaced := index.tree.ReplaceOrInsert(item)
	if replaced != nil {
		lib.AddCountersInPlace(values, replaced.(treeItem).values)
	}
}

func (index *treeIndex) query(
	leftEpochMinute, rightEpochMinute int32) map[int32]*pb.CounterValues {

	valuesMap := make(map[int32]*pb.CounterValues)

	index.lock.RLock()
	defer index.lock.RUnlock()

	iterator := func(item llrb.Item) bool {
		asTreeItem := item.(treeItem)
		valuesMap[asTreeItem.epochMinute] = asTreeItem.values
		return true
	}

	leftKey := treeItem{leftEpochMinute, nil}
	rightKey := treeItem{rightEpochMinute + 1, nil}

	index.tree.AscendRange(leftKey, rightKey, iterator)

	return valuesMap
}
