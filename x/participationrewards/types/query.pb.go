// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/query.proto

package types

import (
	context "context"
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc16b3ccc632b3de, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc16b3ccc632b3de, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryProtocolDataRequest is the request type for querying Protocol Data.
type QueryProtocolDataRequest struct {
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *QueryProtocolDataRequest) Reset()         { *m = QueryProtocolDataRequest{} }
func (m *QueryProtocolDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProtocolDataRequest) ProtoMessage()    {}
func (*QueryProtocolDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc16b3ccc632b3de, []int{2}
}
func (m *QueryProtocolDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProtocolDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProtocolDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProtocolDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProtocolDataRequest.Merge(m, src)
}
func (m *QueryProtocolDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProtocolDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProtocolDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProtocolDataRequest proto.InternalMessageInfo

func (m *QueryProtocolDataRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *QueryProtocolDataRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// QueryProtocolDataResponse is the response type for querying Protocol Data.
type QueryProtocolDataResponse struct {
	// data defines the underlying protocol data.
	Data []encoding_json.RawMessage `protobuf:"bytes,1,rep,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty" yaml:"data"`
}

func (m *QueryProtocolDataResponse) Reset()         { *m = QueryProtocolDataResponse{} }
func (m *QueryProtocolDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProtocolDataResponse) ProtoMessage()    {}
func (*QueryProtocolDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc16b3ccc632b3de, []int{3}
}
func (m *QueryProtocolDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProtocolDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProtocolDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProtocolDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProtocolDataResponse.Merge(m, src)
}
func (m *QueryProtocolDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProtocolDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProtocolDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProtocolDataResponse proto.InternalMessageInfo

func (m *QueryProtocolDataResponse) GetData() []encoding_json.RawMessage {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "quicksilver.participationrewards.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "quicksilver.participationrewards.v1.QueryParamsResponse")
	proto.RegisterType((*QueryProtocolDataRequest)(nil), "quicksilver.participationrewards.v1.QueryProtocolDataRequest")
	proto.RegisterType((*QueryProtocolDataResponse)(nil), "quicksilver.participationrewards.v1.QueryProtocolDataResponse")
}

func init() {
	proto.RegisterFile("quicksilver/participationrewards/v1/query.proto", fileDescriptor_bc16b3ccc632b3de)
}

var fileDescriptor_bc16b3ccc632b3de = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x6a, 0xd4, 0x40,
	0x1c, 0xc6, 0x33, 0xdd, 0x75, 0xc1, 0x69, 0x0f, 0x32, 0xf6, 0x10, 0x83, 0x64, 0x4b, 0xbc, 0x14,
	0x96, 0x66, 0x68, 0x7b, 0x50, 0x04, 0xd7, 0xb2, 0x8a, 0xe0, 0x41, 0xd0, 0x1c, 0x45, 0xc1, 0xd9,
	0x64, 0x18, 0xc7, 0xcd, 0xce, 0x64, 0x33, 0x93, 0xad, 0xa1, 0xf4, 0xe2, 0x13, 0x88, 0xbe, 0x88,
	0x8f, 0xd1, 0x63, 0x41, 0x04, 0x4f, 0x8b, 0xec, 0xfa, 0x04, 0x1e, 0x3d, 0xc9, 0x4c, 0x72, 0x68,
	0xd9, 0x1c, 0xb2, 0xde, 0xfe, 0xfc, 0x93, 0xef, 0xfb, 0x7d, 0xdf, 0xfc, 0x21, 0x9e, 0x15, 0x3c,
	0x9e, 0x28, 0x9e, 0xce, 0x69, 0x8e, 0x33, 0x92, 0x6b, 0x1e, 0xf3, 0x8c, 0x68, 0x2e, 0x45, 0x4e,
	0x4f, 0x49, 0x9e, 0x28, 0x3c, 0x3f, 0xc4, 0xb3, 0x82, 0xe6, 0x65, 0x98, 0xe5, 0x52, 0x4b, 0x74,
	0xef, 0x8a, 0x20, 0x6c, 0x12, 0x84, 0xf3, 0x43, 0x6f, 0x97, 0x49, 0x26, 0xed, 0xff, 0xd8, 0x4c,
	0x95, 0xd4, 0xbb, 0xcb, 0xa4, 0x64, 0x29, 0xc5, 0x24, 0xe3, 0x98, 0x08, 0x21, 0xb5, 0x95, 0xa9,
	0xfa, 0xeb, 0xb0, 0x4d, 0x92, 0x46, 0xa0, 0xd5, 0x07, 0xbb, 0x10, 0xbd, 0x32, 0x39, 0x5f, 0x92,
	0x9c, 0x4c, 0x55, 0x44, 0x67, 0x05, 0x55, 0x3a, 0x78, 0x07, 0x6f, 0x5f, 0xdb, 0xaa, 0x4c, 0x0a,
	0x45, 0xd1, 0x73, 0xd8, 0xcb, 0xec, 0xc6, 0x05, 0x7b, 0x60, 0x7f, 0xfb, 0x68, 0x10, 0xb6, 0xa8,
	0x15, 0x56, 0x26, 0xa3, 0xee, 0xc5, 0xa2, 0xef, 0x44, 0xb5, 0x41, 0x70, 0x02, 0xdd, 0x8a, 0x60,
	0x52, 0xc4, 0x32, 0x7d, 0x4a, 0x34, 0xa9, 0xe9, 0x08, 0xc1, 0xae, 0x2e, 0x33, 0x6a, 0x21, 0x37,
	0x23, 0x3b, 0xa3, 0x5b, 0xb0, 0x33, 0xa1, 0xa5, 0xbb, 0x65, 0x57, 0x66, 0x0c, 0xde, 0xc0, 0x3b,
	0x0d, 0x0e, 0x75, 0xd2, 0xc7, 0xb0, 0x9b, 0x10, 0x4d, 0x5c, 0xb0, 0xd7, 0xd9, 0xdf, 0x19, 0x0d,
	0xfe, 0x2c, 0xfa, 0xdb, 0x25, 0x99, 0xa6, 0x0f, 0x03, 0xb3, 0x0d, 0xfe, 0x2e, 0xfa, 0x2e, 0x15,
	0xb1, 0x4c, 0xb8, 0x60, 0xf8, 0x83, 0x92, 0x22, 0x8c, 0xc8, 0xe9, 0x0b, 0xaa, 0x14, 0x61, 0x34,
	0xb2, 0xc2, 0xa3, 0x2f, 0x1d, 0x78, 0xc3, 0xda, 0xa3, 0x6f, 0x00, 0xf6, 0xaa, 0x0a, 0xe8, 0x7e,
	0xab, 0xbe, 0xeb, 0xef, 0xe9, 0x3d, 0xd8, 0x5c, 0x58, 0x15, 0x09, 0x8e, 0x3f, 0x7d, 0xff, 0xfd,
	0x75, 0xeb, 0x00, 0x0d, 0x70, 0xcb, 0x43, 0x9b, 0x9c, 0x3f, 0x00, 0xdc, 0xb9, 0xfa, 0x2c, 0xe8,
	0xd1, 0x06, 0xfc, 0xf5, 0x83, 0x78, 0xc3, 0xff, 0x95, 0xd7, 0x25, 0x9e, 0xd9, 0x12, 0x27, 0x68,
	0xd8, 0xae, 0x44, 0x6d, 0x61, 0xee, 0x80, 0xcf, 0xcc, 0xf5, 0xcf, 0xf1, 0xd9, 0x84, 0x96, 0xe7,
	0xa3, 0xb7, 0x17, 0x4b, 0x1f, 0x5c, 0x2e, 0x7d, 0xf0, 0x6b, 0xe9, 0x83, 0xcf, 0x2b, 0xdf, 0xb9,
	0x5c, 0xf9, 0xce, 0xcf, 0x95, 0xef, 0xbc, 0x7e, 0xc2, 0xb8, 0x7e, 0x5f, 0x8c, 0xc3, 0x58, 0x4e,
	0x31, 0x17, 0x8c, 0x8a, 0x82, 0xeb, 0xf2, 0x60, 0x5c, 0xf0, 0x34, 0xb9, 0xc6, 0xfc, 0xd8, 0x4c,
	0x35, 0x10, 0x35, 0xee, 0x59, 0xf2, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xcf, 0xde,
	0xf0, 0xde, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the total set of participation rewards parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ProtocolData returns the requested protocol data.
	ProtocolData(ctx context.Context, in *QueryProtocolDataRequest, opts ...grpc.CallOption) (*QueryProtocolDataResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.participationrewards.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProtocolData(ctx context.Context, in *QueryProtocolDataRequest, opts ...grpc.CallOption) (*QueryProtocolDataResponse, error) {
	out := new(QueryProtocolDataResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.participationrewards.v1.Query/ProtocolData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of participation rewards parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ProtocolData returns the requested protocol data.
	ProtocolData(context.Context, *QueryProtocolDataRequest) (*QueryProtocolDataResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ProtocolData(ctx context.Context, req *QueryProtocolDataRequest) (*QueryProtocolDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProtocolData not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.participationrewards.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProtocolData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProtocolDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProtocolData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.participationrewards.v1.Query/ProtocolData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProtocolData(ctx, req.(*QueryProtocolDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.participationrewards.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ProtocolData",
			Handler:    _Query_ProtocolData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/participationrewards/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryProtocolDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProtocolDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProtocolDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProtocolDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProtocolDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProtocolDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Data[iNdEx])
			copy(dAtA[i:], m.Data[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Data[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryProtocolDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProtocolDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, b := range m.Data {
			l = len(b)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryProtocolDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryProtocolDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProtocolDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryProtocolDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryProtocolDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProtocolDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, make([]byte, postIndex-iNdEx))
			copy(m.Data[len(m.Data)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
