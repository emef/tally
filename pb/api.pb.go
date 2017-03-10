// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api.proto
	records.proto

It has these top-level messages:
	CounterValues
	RecordCounterRequest
	RecordCounterResponse
	BulkRecordCounterRequest
	BulkRecordCounterResponse
	GetCounterRequest
	GetCounterResponse
	RecordKey
	RecordEntry
	RecordBlock
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CounterValues struct {
	Count int32   `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Sum   float32 `protobuf:"fixed32,2,opt,name=sum" json:"sum,omitempty"`
	Min   float32 `protobuf:"fixed32,3,opt,name=min" json:"min,omitempty"`
	Max   float32 `protobuf:"fixed32,4,opt,name=max" json:"max,omitempty"`
}

func (m *CounterValues) Reset()                    { *m = CounterValues{} }
func (m *CounterValues) String() string            { return proto.CompactTextString(m) }
func (*CounterValues) ProtoMessage()               {}
func (*CounterValues) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CounterValues) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *CounterValues) GetSum() float32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *CounterValues) GetMin() float32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *CounterValues) GetMax() float32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type RecordCounterRequest struct {
	Name        string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Source      string         `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	EpochMinute int32          `protobuf:"varint,3,opt,name=epochMinute" json:"epochMinute,omitempty"`
	Values      *CounterValues `protobuf:"bytes,4,opt,name=values" json:"values,omitempty"`
}

func (m *RecordCounterRequest) Reset()                    { *m = RecordCounterRequest{} }
func (m *RecordCounterRequest) String() string            { return proto.CompactTextString(m) }
func (*RecordCounterRequest) ProtoMessage()               {}
func (*RecordCounterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RecordCounterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RecordCounterRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RecordCounterRequest) GetEpochMinute() int32 {
	if m != nil {
		return m.EpochMinute
	}
	return 0
}

func (m *RecordCounterRequest) GetValues() *CounterValues {
	if m != nil {
		return m.Values
	}
	return nil
}

type RecordCounterResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *RecordCounterResponse) Reset()                    { *m = RecordCounterResponse{} }
func (m *RecordCounterResponse) String() string            { return proto.CompactTextString(m) }
func (*RecordCounterResponse) ProtoMessage()               {}
func (*RecordCounterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RecordCounterResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *RecordCounterResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type BulkRecordCounterRequest struct {
	Requests []*RecordCounterRequest `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
}

func (m *BulkRecordCounterRequest) Reset()                    { *m = BulkRecordCounterRequest{} }
func (m *BulkRecordCounterRequest) String() string            { return proto.CompactTextString(m) }
func (*BulkRecordCounterRequest) ProtoMessage()               {}
func (*BulkRecordCounterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BulkRecordCounterRequest) GetRequests() []*RecordCounterRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

type BulkRecordCounterResponse struct {
	Ok    bool   `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *BulkRecordCounterResponse) Reset()                    { *m = BulkRecordCounterResponse{} }
func (m *BulkRecordCounterResponse) String() string            { return proto.CompactTextString(m) }
func (*BulkRecordCounterResponse) ProtoMessage()               {}
func (*BulkRecordCounterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BulkRecordCounterResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *BulkRecordCounterResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetCounterRequest struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Source           string `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	StartEpochMinute int32  `protobuf:"varint,3,opt,name=startEpochMinute" json:"startEpochMinute,omitempty"`
	EndEpochMinute   int32  `protobuf:"varint,4,opt,name=endEpochMinute" json:"endEpochMinute,omitempty"`
}

func (m *GetCounterRequest) Reset()                    { *m = GetCounterRequest{} }
func (m *GetCounterRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCounterRequest) ProtoMessage()               {}
func (*GetCounterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetCounterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetCounterRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *GetCounterRequest) GetStartEpochMinute() int32 {
	if m != nil {
		return m.StartEpochMinute
	}
	return 0
}

func (m *GetCounterRequest) GetEndEpochMinute() int32 {
	if m != nil {
		return m.EndEpochMinute
	}
	return 0
}

type GetCounterResponse struct {
	Ok     bool                     `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
	Error  string                   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Values map[int32]*CounterValues `protobuf:"bytes,3,rep,name=values" json:"values,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetCounterResponse) Reset()                    { *m = GetCounterResponse{} }
func (m *GetCounterResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCounterResponse) ProtoMessage()               {}
func (*GetCounterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetCounterResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *GetCounterResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetCounterResponse) GetValues() map[int32]*CounterValues {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*CounterValues)(nil), "pb.CounterValues")
	proto.RegisterType((*RecordCounterRequest)(nil), "pb.RecordCounterRequest")
	proto.RegisterType((*RecordCounterResponse)(nil), "pb.RecordCounterResponse")
	proto.RegisterType((*BulkRecordCounterRequest)(nil), "pb.BulkRecordCounterRequest")
	proto.RegisterType((*BulkRecordCounterResponse)(nil), "pb.BulkRecordCounterResponse")
	proto.RegisterType((*GetCounterRequest)(nil), "pb.GetCounterRequest")
	proto.RegisterType((*GetCounterResponse)(nil), "pb.GetCounterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RecordCounterService service

type RecordCounterServiceClient interface {
	RecordCounter(ctx context.Context, in *RecordCounterRequest, opts ...grpc.CallOption) (*RecordCounterResponse, error)
	BulkRecordCounter(ctx context.Context, in *BulkRecordCounterRequest, opts ...grpc.CallOption) (*BulkRecordCounterResponse, error)
}

type recordCounterServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecordCounterServiceClient(cc *grpc.ClientConn) RecordCounterServiceClient {
	return &recordCounterServiceClient{cc}
}

func (c *recordCounterServiceClient) RecordCounter(ctx context.Context, in *RecordCounterRequest, opts ...grpc.CallOption) (*RecordCounterResponse, error) {
	out := new(RecordCounterResponse)
	err := grpc.Invoke(ctx, "/pb.RecordCounterService/RecordCounter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordCounterServiceClient) BulkRecordCounter(ctx context.Context, in *BulkRecordCounterRequest, opts ...grpc.CallOption) (*BulkRecordCounterResponse, error) {
	out := new(BulkRecordCounterResponse)
	err := grpc.Invoke(ctx, "/pb.RecordCounterService/BulkRecordCounter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecordCounterService service

type RecordCounterServiceServer interface {
	RecordCounter(context.Context, *RecordCounterRequest) (*RecordCounterResponse, error)
	BulkRecordCounter(context.Context, *BulkRecordCounterRequest) (*BulkRecordCounterResponse, error)
}

func RegisterRecordCounterServiceServer(s *grpc.Server, srv RecordCounterServiceServer) {
	s.RegisterService(&_RecordCounterService_serviceDesc, srv)
}

func _RecordCounterService_RecordCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordCounterServiceServer).RecordCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordCounterService/RecordCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordCounterServiceServer).RecordCounter(ctx, req.(*RecordCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordCounterService_BulkRecordCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkRecordCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordCounterServiceServer).BulkRecordCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RecordCounterService/BulkRecordCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordCounterServiceServer).BulkRecordCounter(ctx, req.(*BulkRecordCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecordCounterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RecordCounterService",
	HandlerType: (*RecordCounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordCounter",
			Handler:    _RecordCounterService_RecordCounter_Handler,
		},
		{
			MethodName: "BulkRecordCounter",
			Handler:    _RecordCounterService_BulkRecordCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for QueryCounterService service

type QueryCounterServiceClient interface {
	GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error)
}

type queryCounterServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueryCounterServiceClient(cc *grpc.ClientConn) QueryCounterServiceClient {
	return &queryCounterServiceClient{cc}
}

func (c *queryCounterServiceClient) GetCounter(ctx context.Context, in *GetCounterRequest, opts ...grpc.CallOption) (*GetCounterResponse, error) {
	out := new(GetCounterResponse)
	err := grpc.Invoke(ctx, "/pb.QueryCounterService/GetCounter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QueryCounterService service

type QueryCounterServiceServer interface {
	GetCounter(context.Context, *GetCounterRequest) (*GetCounterResponse, error)
}

func RegisterQueryCounterServiceServer(s *grpc.Server, srv QueryCounterServiceServer) {
	s.RegisterService(&_QueryCounterService_serviceDesc, srv)
}

func _QueryCounterService_GetCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryCounterServiceServer).GetCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.QueryCounterService/GetCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryCounterServiceServer).GetCounter(ctx, req.(*GetCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryCounterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.QueryCounterService",
	HandlerType: (*QueryCounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCounter",
			Handler:    _QueryCounterService_GetCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xc7, 0x65, 0xa7, 0x89, 0x9a, 0x17, 0xb5, 0x6a, 0x5e, 0xbf, 0x9c, 0x28, 0x95, 0xa2, 0x59,
	0x40, 0xe8, 0x22, 0x11, 0x81, 0x05, 0x8a, 0xc4, 0x02, 0x50, 0xc5, 0x06, 0x24, 0x18, 0xa4, 0x6e,
	0x58, 0xa0, 0x89, 0x3b, 0x2a, 0x56, 0xec, 0x19, 0x33, 0x1e, 0x47, 0xcd, 0x96, 0x0b, 0xb0, 0xe8,
	0x95, 0xb8, 0x01, 0x57, 0xe0, 0x02, 0xdc, 0x00, 0xcd, 0x8c, 0x1b, 0x5c, 0xc7, 0x59, 0x54, 0xec,
	0x9e, 0xff, 0xef, 0xaf, 0xf9, 0xbd, 0x2f, 0x19, 0xda, 0x2c, 0x8d, 0xc6, 0xa9, 0x92, 0x5a, 0xa2,
	0x9f, 0xce, 0xfb, 0x83, 0x6b, 0x29, 0xaf, 0x63, 0x3e, 0x61, 0x69, 0x34, 0x61, 0x42, 0x48, 0xcd,
	0x74, 0x24, 0x45, 0xe6, 0x1c, 0xe4, 0x33, 0xec, 0xbd, 0x91, 0xb9, 0xd0, 0x5c, 0x5d, 0xb2, 0x38,
	0xe7, 0x19, 0x1e, 0x41, 0x33, 0x34, 0x42, 0xe0, 0x0d, 0xbd, 0x51, 0x93, 0xba, 0x0f, 0x3c, 0x80,
	0x46, 0x96, 0x27, 0x81, 0x3f, 0xf4, 0x46, 0x3e, 0x35, 0xa1, 0x51, 0x92, 0x48, 0x04, 0x0d, 0xa7,
	0x24, 0x91, 0xb0, 0x0a, 0xbb, 0x09, 0x76, 0x0a, 0x85, 0xdd, 0x90, 0x1f, 0x1e, 0x1c, 0x51, 0x1e,
	0x4a, 0x75, 0x55, 0x30, 0x28, 0xff, 0x96, 0xf3, 0x4c, 0x23, 0xc2, 0x8e, 0x60, 0x09, 0xb7, 0x8c,
	0x36, 0xb5, 0x31, 0x9e, 0x40, 0x2b, 0x93, 0xb9, 0x0a, 0xb9, 0xa5, 0xb4, 0x69, 0xf1, 0x85, 0x43,
	0xe8, 0xf0, 0x54, 0x86, 0x5f, 0xdf, 0x47, 0x22, 0xd7, 0xdc, 0x02, 0x9b, 0xb4, 0x2c, 0xe1, 0x13,
	0x68, 0x2d, 0x6d, 0xf1, 0x96, 0xdd, 0x99, 0x76, 0xc7, 0xe9, 0x7c, 0x7c, 0xaf, 0x2b, 0x5a, 0x18,
	0xc8, 0x4b, 0x38, 0xae, 0x14, 0x94, 0xa5, 0x52, 0x64, 0x1c, 0xf7, 0xc1, 0x97, 0x0b, 0x5b, 0xcf,
	0x2e, 0xf5, 0xe5, 0xc2, 0x8c, 0x81, 0x2b, 0x25, 0x55, 0x51, 0x8c, 0xfb, 0x20, 0x1f, 0x20, 0x78,
	0x9d, 0xc7, 0x8b, 0xda, 0x9e, 0x9e, 0xc3, 0xae, 0x72, 0x61, 0x16, 0x78, 0xc3, 0xc6, 0xa8, 0x33,
	0x0d, 0x4c, 0x1d, 0x75, 0x5e, 0xba, 0x76, 0x92, 0x57, 0xd0, 0xab, 0x79, 0xf1, 0x41, 0x45, 0xdd,
	0x7a, 0xd0, 0x7d, 0xcb, 0xf5, 0x7f, 0x8c, 0xf8, 0x1c, 0x0e, 0x32, 0xcd, 0x94, 0xbe, 0xd8, 0x98,
	0xf3, 0x86, 0x8e, 0x8f, 0x60, 0x9f, 0x8b, 0xab, 0xb2, 0x73, 0xc7, 0x3a, 0x2b, 0x2a, 0xf9, 0xe9,
	0x01, 0x96, 0xab, 0x7a, 0x48, 0x4b, 0x38, 0x5b, 0x6f, 0xb4, 0x61, 0x27, 0x49, 0xcc, 0x24, 0x37,
	0x5f, 0x1b, 0xbb, 0xed, 0x5e, 0x08, 0xad, 0x56, 0x77, 0x2b, 0xee, 0xbf, 0x83, 0x4e, 0x49, 0x36,
	0x57, 0xb9, 0xe0, 0xab, 0xe2, 0x9a, 0x4d, 0x88, 0x8f, 0xa1, 0x69, 0xad, 0x16, 0x59, 0x7b, 0x2d,
	0x2e, 0x3f, 0xf3, 0x5f, 0x78, 0xd3, 0x3f, 0xd5, 0x13, 0xfe, 0xc4, 0xd5, 0x32, 0x0a, 0x39, 0x86,
	0xb0, 0x77, 0x4f, 0xc7, 0xad, 0xdb, 0xee, 0xf7, 0x6a, 0x32, 0xae, 0x01, 0x72, 0xf6, 0xfd, 0xd7,
	0xef, 0x5b, 0xff, 0x94, 0xe0, 0x64, 0xf9, 0x74, 0x12, 0xba, 0xe4, 0x44, 0x59, 0xeb, 0xcc, 0x3b,
	0x47, 0x0d, 0xdd, 0x8d, 0xeb, 0xc0, 0x81, 0x79, 0x6e, 0xdb, 0x19, 0xf6, 0xcf, 0xb6, 0x64, 0x0b,
	0x20, 0xb1, 0xc0, 0x01, 0x39, 0x2d, 0x03, 0xe7, 0x79, 0xbc, 0xf8, 0xb2, 0xa6, 0x4e, 0x13, 0x38,
	0xfc, 0x98, 0x73, 0xb5, 0xaa, 0x74, 0x7c, 0x09, 0xf0, 0x6f, 0x05, 0x78, 0x5c, 0x5d, 0x89, 0xc3,
	0x9f, 0xd4, 0x6f, 0x8a, 0xf4, 0x2c, 0xf7, 0x10, 0xbb, 0x86, 0xab, 0x59, 0x1c, 0xaf, 0xee, 0xe8,
	0xf3, 0x96, 0xfd, 0x13, 0x3d, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x44, 0x4f, 0x93, 0x91, 0xb8,
	0x04, 0x00, 0x00,
}
