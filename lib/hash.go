package lib

import (
	"crypto/md5"
	"encoding/binary"
)

func HashCode(source string) int64 {
	strMd5 := md5.Sum([]byte(source))
	hashCode, _ := binary.Varint(strMd5[:])
	return hashCode
}

func HashToRange(source string, minValue int64, maxValue int64) int64 {
	hash := HashCode(source) & 0x7FFFFFFF
	delta := hash % (1 + maxValue - minValue)
	return minValue + delta
}
