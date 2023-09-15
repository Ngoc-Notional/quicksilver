// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/claimsmanager/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryClaimsRequest is the request type for the Query/Claims RPC method.
type QueryClaimsRequest struct {
	ChainId    string             `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	Address    string             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryClaimsRequest) Reset()         { *m = QueryClaimsRequest{} }
func (m *QueryClaimsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClaimsRequest) ProtoMessage()    {}
func (*QueryClaimsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a524187a5a706bf7, []int{0}
}
func (m *QueryClaimsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimsRequest.Merge(m, src)
}
func (m *QueryClaimsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimsRequest proto.InternalMessageInfo

func (m *QueryClaimsRequest) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *QueryClaimsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryClaimsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryClaimsResponse is the response type for the Query/Claims RPC method.
type QueryClaimsResponse struct {
	Claims     []Claim             `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryClaimsResponse) Reset()         { *m = QueryClaimsResponse{} }
func (m *QueryClaimsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClaimsResponse) ProtoMessage()    {}
func (*QueryClaimsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a524187a5a706bf7, []int{1}
}
func (m *QueryClaimsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClaimsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClaimsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClaimsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClaimsResponse.Merge(m, src)
}
func (m *QueryClaimsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClaimsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClaimsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClaimsResponse proto.InternalMessageInfo

func (m *QueryClaimsResponse) GetClaims() []Claim {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *QueryClaimsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryClaimsRequest)(nil), "quicksilver.claimsmanager.v1.QueryClaimsRequest")
	proto.RegisterType((*QueryClaimsResponse)(nil), "quicksilver.claimsmanager.v1.QueryClaimsResponse")
}

func init() {
	proto.RegisterFile("quicksilver/claimsmanager/v1/query.proto", fileDescriptor_a524187a5a706bf7)
}

var fileDescriptor_a524187a5a706bf7 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x31, 0x6b, 0x93, 0x41,
	0x1c, 0xc6, 0x73, 0xa9, 0x4d, 0xf5, 0x3a, 0x14, 0x2e, 0x1d, 0x62, 0x28, 0x6f, 0x43, 0x04, 0x0d,
	0x42, 0xef, 0xfa, 0xa6, 0x88, 0xa0, 0x22, 0x36, 0x6a, 0x45, 0x71, 0xd0, 0x88, 0x0e, 0x2e, 0xe1,
	0xf2, 0xe6, 0xb8, 0x1c, 0x26, 0x77, 0x6f, 0xde, 0xbb, 0x37, 0x18, 0x4a, 0x17, 0x3f, 0x81, 0xe0,
	0x27, 0x70, 0x71, 0x74, 0x72, 0x16, 0x27, 0xe9, 0x58, 0x74, 0x71, 0x2a, 0x92, 0xb8, 0xb9, 0xf9,
	0x09, 0x24, 0x77, 0x17, 0x4c, 0xac, 0x44, 0xbb, 0xd4, 0xed, 0xc2, 0xf3, 0x3c, 0x3c, 0xbf, 0xff,
	0xff, 0x9f, 0x04, 0x56, 0x7a, 0xa9, 0x88, 0x9e, 0x69, 0xd1, 0xe9, 0xb3, 0x84, 0x44, 0x1d, 0x2a,
	0xba, 0xba, 0x4b, 0x25, 0xe5, 0x2c, 0x21, 0xfd, 0x90, 0xf4, 0x52, 0x96, 0x0c, 0x70, 0x9c, 0x28,
	0xa3, 0xd0, 0xda, 0x94, 0x13, 0xcf, 0x38, 0x71, 0x3f, 0x2c, 0xae, 0x72, 0xc5, 0x95, 0x35, 0x92,
	0xf1, 0xcb, 0x65, 0x8a, 0x6b, 0x5c, 0x29, 0xde, 0x61, 0x84, 0xc6, 0x82, 0x50, 0x29, 0x95, 0xa1,
	0x46, 0x28, 0xa9, 0xbd, 0x7a, 0x31, 0x52, 0xba, 0xab, 0x34, 0x69, 0x52, 0xcd, 0x5c, 0x15, 0xe9,
	0x87, 0x4d, 0x66, 0x68, 0x48, 0x62, 0xca, 0x85, 0xb4, 0x66, 0xef, 0x3d, 0xeb, 0xbc, 0x0d, 0x57,
	0xe1, 0x3e, 0x78, 0x69, 0x73, 0xee, 0x08, 0xb3, 0xa4, 0x36, 0x51, 0xfe, 0x00, 0x20, 0x7a, 0x38,
	0xee, 0xbb, 0x69, 0xc5, 0x3a, 0xeb, 0xa5, 0x4c, 0x1b, 0x84, 0xe1, 0xe9, 0xa8, 0x4d, 0x85, 0x6c,
	0x88, 0x56, 0x01, 0x94, 0x40, 0xe5, 0x4c, 0x2d, 0xff, 0xe3, 0x70, 0x7d, 0x65, 0x40, 0xbb, 0x9d,
	0x2b, 0xe5, 0x89, 0x52, 0xae, 0x2f, 0xd9, 0xe7, 0xdd, 0x16, 0xaa, 0xc2, 0x25, 0xda, 0x6a, 0x25,
	0x4c, 0xeb, 0x42, 0xd6, 0xda, 0x0b, 0x9f, 0xde, 0x6d, 0xac, 0x7a, 0xb6, 0x6d, 0xa7, 0x3c, 0x32,
	0x89, 0x90, 0xbc, 0x3e, 0x31, 0xa2, 0x1d, 0x08, 0x7f, 0xcd, 0x56, 0x58, 0x28, 0x81, 0xca, 0x72,
	0xf5, 0x3c, 0xf6, 0x99, 0xf1, 0x22, 0xb0, 0xdb, 0xb9, 0x5f, 0x04, 0x7e, 0x40, 0x39, 0xf3, 0x7c,
	0xf5, 0xa9, 0x64, 0xf9, 0x35, 0x80, 0xf9, 0x99, 0x11, 0x74, 0xac, 0xa4, 0x66, 0x68, 0x1b, 0xe6,
	0xdc, 0xc4, 0x05, 0x50, 0x5a, 0xa8, 0x2c, 0x57, 0xcf, 0xe1, 0x79, 0x67, 0xc3, 0x36, 0x5d, 0x3b,
	0xb5, 0x7f, 0xb8, 0x9e, 0xa9, 0xfb, 0x20, 0xba, 0x33, 0x83, 0x98, 0xb5, 0x88, 0x17, 0xfe, 0x8a,
	0xe8, 0xfa, 0xa7, 0x19, 0xab, 0xdf, 0x17, 0xe1, 0xa2, 0x65, 0x44, 0x6f, 0x00, 0xcc, 0x39, 0x50,
	0xb4, 0x39, 0x1f, 0xe8, 0xe8, 0x59, 0x8a, 0xe1, 0x31, 0x12, 0x8e, 0xa2, 0x7c, 0xf9, 0xc5, 0xe7,
	0x6f, 0xaf, 0xb2, 0x21, 0x22, 0xe4, 0x1f, 0xbe, 0x1b, 0x64, 0x77, 0x72, 0xdb, 0x3d, 0xf4, 0x1e,
	0xc0, 0x95, 0xfb, 0x54, 0x9b, 0xdb, 0xb1, 0x8a, 0xda, 0x27, 0x49, 0xbc, 0x63, 0x89, 0x6f, 0xa0,
	0xeb, 0xf3, 0x89, 0xe3, 0x84, 0xf5, 0x85, 0x4a, 0x75, 0x83, 0x8d, 0x01, 0x1b, 0x47, 0x07, 0x78,
	0x0b, 0x20, 0x7c, 0xac, 0x59, 0x72, 0x92, 0xec, 0x57, 0x2d, 0xfb, 0x25, 0xb4, 0x35, 0x9f, 0x3d,
	0xd5, 0x2c, 0x21, 0xbb, 0xfe, 0x87, 0xb0, 0xe7, 0x75, 0xf4, 0x11, 0xc0, 0xfc, 0x18, 0xf8, 0xbf,
	0x6c, 0xfd, 0x9e, 0x25, 0xbf, 0x85, 0x6a, 0xc7, 0x22, 0xff, 0xe3, 0x11, 0x6a, 0x4f, 0xf6, 0x87,
	0x01, 0x38, 0x18, 0x06, 0xe0, 0xeb, 0x30, 0x00, 0x2f, 0x47, 0x41, 0xe6, 0x60, 0x14, 0x64, 0xbe,
	0x8c, 0x82, 0xcc, 0xd3, 0x6b, 0x5c, 0x98, 0x76, 0xda, 0xc4, 0x91, 0xea, 0x12, 0x21, 0x39, 0x93,
	0xa9, 0x30, 0x83, 0x8d, 0x66, 0x2a, 0x3a, 0xad, 0x99, 0xde, 0xe7, 0xbf, 0x35, 0x9b, 0x41, 0xcc,
	0x74, 0x33, 0x67, 0xff, 0xb3, 0xb6, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x73, 0xf2, 0xfb, 0xe5,
	0xaa, 0x05, 0x00, 0x00,
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
	// Claims returns all zone claims from the current epoch.
	Claims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
	// LastEpochClaims returns all zone claims from the last epoch.
	LastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
	// UserClaims returns all zone claims for a given address from the current
	// epoch.
	UserClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
	// UserLastEpochClaims returns all zone claims for a given address from the
	// last epoch.
	UserLastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Claims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/Claims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/LastEpochClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UserClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/UserClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UserLastEpochClaims(ctx context.Context, in *QueryClaimsRequest, opts ...grpc.CallOption) (*QueryClaimsResponse, error) {
	out := new(QueryClaimsResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.claimsmanager.v1.Query/UserLastEpochClaims", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Claims returns all zone claims from the current epoch.
	Claims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
	// LastEpochClaims returns all zone claims from the last epoch.
	LastEpochClaims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
	// UserClaims returns all zone claims for a given address from the current
	// epoch.
	UserClaims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
	// UserLastEpochClaims returns all zone claims for a given address from the
	// last epoch.
	UserLastEpochClaims(context.Context, *QueryClaimsRequest) (*QueryClaimsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Claims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claims not implemented")
}
func (*UnimplementedQueryServer) LastEpochClaims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastEpochClaims not implemented")
}
func (*UnimplementedQueryServer) UserClaims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserClaims not implemented")
}
func (*UnimplementedQueryServer) UserLastEpochClaims(ctx context.Context, req *QueryClaimsRequest) (*QueryClaimsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLastEpochClaims not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Claims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Claims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/Claims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Claims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastEpochClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastEpochClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/LastEpochClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastEpochClaims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UserClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/UserClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserClaims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UserLastEpochClaims_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClaimsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UserLastEpochClaims(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.claimsmanager.v1.Query/UserLastEpochClaims",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UserLastEpochClaims(ctx, req.(*QueryClaimsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.claimsmanager.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Claims",
			Handler:    _Query_Claims_Handler,
		},
		{
			MethodName: "LastEpochClaims",
			Handler:    _Query_LastEpochClaims_Handler,
		},
		{
			MethodName: "UserClaims",
			Handler:    _Query_UserClaims_Handler,
		},
		{
			MethodName: "UserLastEpochClaims",
			Handler:    _Query_UserLastEpochClaims_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/claimsmanager/v1/query.proto",
}

func (m *QueryClaimsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryClaimsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClaimsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClaimsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
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
func (m *QueryClaimsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryClaimsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryClaimsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClaimsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
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
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryClaimsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClaimsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClaimsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
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
			m.Claims = append(m.Claims, Claim{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
