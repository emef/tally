// Code generated by protoc-gen-go.
// source: records.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RecordKey struct {
	NameCode    int32 `protobuf:"varint,1,opt,name=nameCode" json:"nameCode,omitempty"`
	SourceCode  int32 `protobuf:"varint,2,opt,name=sourceCode" json:"sourceCode,omitempty"`
	EpochMinute int32 `protobuf:"varint,3,opt,name=epochMinute" json:"epochMinute,omitempty"`
}

func (m *RecordKey) Reset()                    { *m = RecordKey{} }
func (m *RecordKey) String() string            { return proto.CompactTextString(m) }
func (*RecordKey) ProtoMessage()               {}
func (*RecordKey) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RecordKey) GetNameCode() int32 {
	if m != nil {
		return m.NameCode
	}
	return 0
}

func (m *RecordKey) GetSourceCode() int32 {
	if m != nil {
		return m.SourceCode
	}
	return 0
}

func (m *RecordKey) GetEpochMinute() int32 {
	if m != nil {
		return m.EpochMinute
	}
	return 0
}

type RecordEntry struct {
	Key    *RecordKey     `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Values *CounterValues `protobuf:"bytes,2,opt,name=values" json:"values,omitempty"`
}

func (m *RecordEntry) Reset()                    { *m = RecordEntry{} }
func (m *RecordEntry) String() string            { return proto.CompactTextString(m) }
func (*RecordEntry) ProtoMessage()               {}
func (*RecordEntry) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *RecordEntry) GetKey() *RecordKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RecordEntry) GetValues() *CounterValues {
	if m != nil {
		return m.Values
	}
	return nil
}

type RecordBlock struct {
	NameCodeMapping   map[int32]string `protobuf:"bytes,1,rep,name=nameCodeMapping" json:"nameCodeMapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SourceCodeMapping map[int32]string `protobuf:"bytes,2,rep,name=sourceCodeMapping" json:"sourceCodeMapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Entries           []*RecordEntry   `protobuf:"bytes,3,rep,name=entries" json:"entries,omitempty"`
}

func (m *RecordBlock) Reset()                    { *m = RecordBlock{} }
func (m *RecordBlock) String() string            { return proto.CompactTextString(m) }
func (*RecordBlock) ProtoMessage()               {}
func (*RecordBlock) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RecordBlock) GetNameCodeMapping() map[int32]string {
	if m != nil {
		return m.NameCodeMapping
	}
	return nil
}

func (m *RecordBlock) GetSourceCodeMapping() map[int32]string {
	if m != nil {
		return m.SourceCodeMapping
	}
	return nil
}

func (m *RecordBlock) GetEntries() []*RecordEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordKey)(nil), "pb.RecordKey")
	proto.RegisterType((*RecordEntry)(nil), "pb.RecordEntry")
	proto.RegisterType((*RecordBlock)(nil), "pb.RecordBlock")
}

func init() { proto.RegisterFile("records.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x42, 0xab, 0x9d, 0x50, 0x6a, 0x97, 0x22, 0x21, 0x07, 0x0d, 0x41, 0xa4, 0xbd,
	0xe4, 0x10, 0x2f, 0xe2, 0xb1, 0xd5, 0x93, 0xb4, 0x87, 0x55, 0x04, 0x8f, 0x49, 0x3a, 0x68, 0x68,
	0xdd, 0x5d, 0x36, 0x89, 0x90, 0x7f, 0xec, 0xcf, 0x90, 0xce, 0x9a, 0xa6, 0xc6, 0x5e, 0x3c, 0xce,
	0xbe, 0x37, 0xdf, 0xbc, 0x3c, 0x02, 0x43, 0x8d, 0x99, 0xd4, 0xeb, 0x22, 0x52, 0x5a, 0x96, 0x92,
	0xd9, 0x2a, 0xf5, 0x07, 0x89, 0xca, 0xcd, 0x18, 0xe6, 0x30, 0xe0, 0xa4, 0x3f, 0x62, 0xcd, 0x7c,
	0x38, 0x15, 0xc9, 0x07, 0x2e, 0xe4, 0x1a, 0x3d, 0x2b, 0xb0, 0xa6, 0x3d, 0xbe, 0x9f, 0xd9, 0x05,
	0x40, 0x21, 0x2b, 0x9d, 0x19, 0xd5, 0x26, 0xf5, 0xe0, 0x85, 0x05, 0xe0, 0xa2, 0x92, 0xd9, 0xfb,
	0x32, 0x17, 0x55, 0x89, 0x9e, 0x43, 0x86, 0xc3, 0xa7, 0xf0, 0x15, 0x5c, 0x73, 0xea, 0x41, 0x94,
	0xba, 0x66, 0x97, 0xe0, 0x6c, 0xb0, 0xa6, 0x3b, 0x6e, 0x3c, 0x8c, 0x54, 0x1a, 0xed, 0x83, 0xf0,
	0x9d, 0xc2, 0x66, 0xd0, 0xff, 0x4c, 0xb6, 0x15, 0x16, 0x74, 0xcd, 0x8d, 0xc7, 0x3b, 0xcf, 0x42,
	0x56, 0xa2, 0x44, 0xfd, 0x42, 0x02, 0xff, 0x31, 0x84, 0x5f, 0x76, 0xc3, 0x9e, 0x6f, 0x65, 0xb6,
	0x61, 0x2b, 0x18, 0x35, 0xc1, 0x97, 0x89, 0x52, 0xb9, 0x78, 0xf3, 0xac, 0xc0, 0x99, 0xba, 0xf1,
	0x55, 0x7b, 0x87, 0x9c, 0xd1, 0xea, 0xb7, 0x8d, 0xa2, 0xf1, 0xee, 0x32, 0x7b, 0x86, 0x71, 0xfb,
	0xa9, 0x0d, 0xd1, 0x26, 0xe2, 0x75, 0x97, 0xf8, 0xd4, 0x35, 0x1a, 0xe6, 0x5f, 0x00, 0x9b, 0xc1,
	0x09, 0x8a, 0x52, 0xe7, 0x58, 0x78, 0x0e, 0xb1, 0x46, 0x2d, 0xcb, 0x2c, 0x35, 0xba, 0x3f, 0x87,
	0xc9, 0xb1, 0xa4, 0xec, 0xac, 0x2d, 0xb1, 0x67, 0x5a, 0x9b, 0x40, 0x8f, 0x4a, 0xa1, 0xd2, 0x06,
	0xdc, 0x0c, 0x77, 0xf6, 0xad, 0xe5, 0xdf, 0xc3, 0xf9, 0xf1, 0x6c, 0xff, 0xa1, 0xa4, 0x7d, 0xfa,
	0x6f, 0x6e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x13, 0x55, 0x59, 0x98, 0x57, 0x02, 0x00, 0x00,
}