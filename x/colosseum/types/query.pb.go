// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: colosseum/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb5cf95a58cc59b, []int{0}
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

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb5cf95a58cc59b, []int{1}
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

type QueryGetCoinSymbolRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetCoinSymbolRequest) Reset()         { *m = QueryGetCoinSymbolRequest{} }
func (m *QueryGetCoinSymbolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCoinSymbolRequest) ProtoMessage()    {}
func (*QueryGetCoinSymbolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb5cf95a58cc59b, []int{2}
}
func (m *QueryGetCoinSymbolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCoinSymbolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCoinSymbolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCoinSymbolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCoinSymbolRequest.Merge(m, src)
}
func (m *QueryGetCoinSymbolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCoinSymbolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCoinSymbolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCoinSymbolRequest proto.InternalMessageInfo

func (m *QueryGetCoinSymbolRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetCoinSymbolResponse struct {
	CoinSymbol CoinSymbol `protobuf:"bytes,1,opt,name=CoinSymbol,proto3" json:"CoinSymbol"`
}

func (m *QueryGetCoinSymbolResponse) Reset()         { *m = QueryGetCoinSymbolResponse{} }
func (m *QueryGetCoinSymbolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCoinSymbolResponse) ProtoMessage()    {}
func (*QueryGetCoinSymbolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb5cf95a58cc59b, []int{3}
}
func (m *QueryGetCoinSymbolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCoinSymbolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCoinSymbolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCoinSymbolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCoinSymbolResponse.Merge(m, src)
}
func (m *QueryGetCoinSymbolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCoinSymbolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCoinSymbolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCoinSymbolResponse proto.InternalMessageInfo

func (m *QueryGetCoinSymbolResponse) GetCoinSymbol() CoinSymbol {
	if m != nil {
		return m.CoinSymbol
	}
	return CoinSymbol{}
}

type QueryAllCoinSymbolRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCoinSymbolRequest) Reset()         { *m = QueryAllCoinSymbolRequest{} }
func (m *QueryAllCoinSymbolRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCoinSymbolRequest) ProtoMessage()    {}
func (*QueryAllCoinSymbolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb5cf95a58cc59b, []int{4}
}
func (m *QueryAllCoinSymbolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCoinSymbolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCoinSymbolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCoinSymbolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCoinSymbolRequest.Merge(m, src)
}
func (m *QueryAllCoinSymbolRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCoinSymbolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCoinSymbolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCoinSymbolRequest proto.InternalMessageInfo

func (m *QueryAllCoinSymbolRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCoinSymbolResponse struct {
	CoinSymbol []CoinSymbol        `protobuf:"bytes,1,rep,name=CoinSymbol,proto3" json:"CoinSymbol"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCoinSymbolResponse) Reset()         { *m = QueryAllCoinSymbolResponse{} }
func (m *QueryAllCoinSymbolResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCoinSymbolResponse) ProtoMessage()    {}
func (*QueryAllCoinSymbolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb5cf95a58cc59b, []int{5}
}
func (m *QueryAllCoinSymbolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCoinSymbolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCoinSymbolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCoinSymbolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCoinSymbolResponse.Merge(m, src)
}
func (m *QueryAllCoinSymbolResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCoinSymbolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCoinSymbolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCoinSymbolResponse proto.InternalMessageInfo

func (m *QueryAllCoinSymbolResponse) GetCoinSymbol() []CoinSymbol {
	if m != nil {
		return m.CoinSymbol
	}
	return nil
}

func (m *QueryAllCoinSymbolResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "colosseum.colosseum.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "colosseum.colosseum.QueryParamsResponse")
	proto.RegisterType((*QueryGetCoinSymbolRequest)(nil), "colosseum.colosseum.QueryGetCoinSymbolRequest")
	proto.RegisterType((*QueryGetCoinSymbolResponse)(nil), "colosseum.colosseum.QueryGetCoinSymbolResponse")
	proto.RegisterType((*QueryAllCoinSymbolRequest)(nil), "colosseum.colosseum.QueryAllCoinSymbolRequest")
	proto.RegisterType((*QueryAllCoinSymbolResponse)(nil), "colosseum.colosseum.QueryAllCoinSymbolResponse")
}

func init() { proto.RegisterFile("colosseum/query.proto", fileDescriptor_bdb5cf95a58cc59b) }

var fileDescriptor_bdb5cf95a58cc59b = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x6e, 0xf4, 0x60, 0x04, 0x07, 0x6f, 0x20, 0x48, 0x21, 0x9b, 0x8c, 0xa0, 0x15,
	0x08, 0x5b, 0x1b, 0xe2, 0xc0, 0x71, 0x9b, 0x60, 0xd7, 0x12, 0x6e, 0x5c, 0x90, 0x93, 0x5a, 0x91,
	0xa5, 0x24, 0x2f, 0xab, 0x53, 0x44, 0x85, 0x90, 0x10, 0x9f, 0x00, 0x89, 0x1b, 0xe2, 0x23, 0xf0,
	0x39, 0xd0, 0x8e, 0x93, 0xb8, 0x70, 0x42, 0xa8, 0xe5, 0x83, 0xa0, 0xda, 0x66, 0x49, 0x99, 0xd7,
	0xa2, 0xdd, 0xac, 0xe7, 0xf7, 0xff, 0xbf, 0xdf, 0xf3, 0x7b, 0x09, 0xbe, 0x96, 0x40, 0x06, 0x5a,
	0xcb, 0x71, 0xce, 0x8f, 0xc6, 0x72, 0x34, 0x61, 0xe5, 0x08, 0x2a, 0x20, 0x1b, 0xa7, 0x61, 0x76,
	0x7a, 0x0a, 0x36, 0x53, 0x48, 0xc1, 0xdc, 0xf3, 0xf9, 0xc9, 0xa6, 0x06, 0xb7, 0x52, 0x80, 0x34,
	0x93, 0x5c, 0x94, 0x8a, 0x8b, 0xa2, 0x80, 0x4a, 0x54, 0x0a, 0x0a, 0xed, 0x6e, 0xef, 0x27, 0xa0,
	0x73, 0xd0, 0x3c, 0x16, 0x5a, 0xda, 0x0a, 0xfc, 0xf5, 0x4e, 0x2c, 0x2b, 0xb1, 0xc3, 0x4b, 0x91,
	0xaa, 0xc2, 0x24, 0xbb, 0xdc, 0xeb, 0x35, 0x4b, 0x29, 0x46, 0x22, 0xff, 0xeb, 0xd1, 0xad, 0xe3,
	0x09, 0xa8, 0xe2, 0x95, 0x9e, 0xe4, 0x31, 0x64, 0xf6, 0x92, 0x6e, 0x62, 0xf2, 0x7c, 0x6e, 0x3b,
	0x30, 0x8a, 0x48, 0x1e, 0x8d, 0xa5, 0xae, 0xe8, 0x00, 0x6f, 0x2c, 0x44, 0x75, 0x09, 0x85, 0x96,
	0xe4, 0x09, 0xee, 0x58, 0xe7, 0x1b, 0x68, 0x1b, 0xf5, 0x2f, 0xef, 0x76, 0x99, 0xa7, 0x4f, 0x66,
	0x45, 0xfb, 0xeb, 0xc7, 0x3f, 0xb7, 0x5a, 0x91, 0x13, 0xd0, 0x07, 0xf8, 0xa6, 0x71, 0x3c, 0x94,
	0xd5, 0x01, 0xa8, 0xe2, 0x85, 0x61, 0x70, 0xe5, 0xc8, 0x55, 0xdc, 0x56, 0x43, 0xe3, 0xb9, 0x1e,
	0xb5, 0xd5, 0x90, 0x26, 0x38, 0xf0, 0x25, 0x3b, 0x8a, 0xa7, 0x18, 0xd7, 0x51, 0x47, 0xb2, 0xe5,
	0x25, 0xa9, 0xd3, 0x1c, 0x4d, 0x43, 0x48, 0x13, 0x47, 0xb4, 0x97, 0x65, 0x67, 0x89, 0x9e, 0x61,
	0x5c, 0xbf, 0xaf, 0xab, 0x71, 0x8f, 0xd9, 0x61, 0xb0, 0xf9, 0x30, 0x98, 0x1d, 0xb7, 0x1b, 0x06,
	0x1b, 0x88, 0x54, 0x3a, 0x6d, 0xd4, 0x50, 0xd2, 0xaf, 0xc8, 0xb5, 0xf2, 0x4f, 0x95, 0x73, 0x5a,
	0x59, 0xbb, 0x50, 0x2b, 0xe4, 0x70, 0x81, 0xb6, 0x6d, 0x68, 0x7b, 0x2b, 0x69, 0x2d, 0x43, 0x13,
	0x77, 0xf7, 0xdb, 0x1a, 0xbe, 0x64, 0x70, 0xc9, 0x7b, 0x84, 0x3b, 0x76, 0x90, 0xa4, 0xe7, 0x05,
	0x3a, 0xbb, 0x35, 0x41, 0x7f, 0x75, 0xa2, 0xad, 0x49, 0xef, 0x7c, 0xf8, 0xfe, 0xfb, 0x53, 0xfb,
	0x36, 0xe9, 0xf2, 0x83, 0xc6, 0x6e, 0x2e, 0x6e, 0x2f, 0xf9, 0x82, 0x9a, 0xaf, 0x43, 0xd8, 0xf9,
	0xee, 0xbe, 0xa5, 0x0a, 0xf8, 0x7f, 0xe7, 0x3b, 0xa8, 0x87, 0x06, 0xaa, 0x47, 0xee, 0x7a, 0xa1,
	0x1a, 0x9f, 0x0e, 0x7f, 0xab, 0x86, 0xef, 0xc8, 0x67, 0x84, 0xaf, 0xd4, 0x2e, 0x7b, 0xd9, 0x52,
	0x42, 0xdf, 0x92, 0x2d, 0x23, 0xf4, 0xae, 0x0b, 0xed, 0x1b, 0x42, 0x4a, 0xb6, 0x57, 0x11, 0xee,
	0x3f, 0x3e, 0x9e, 0x86, 0xe8, 0x64, 0x1a, 0xa2, 0x5f, 0xd3, 0x10, 0x7d, 0x9c, 0x85, 0xad, 0x93,
	0x59, 0xd8, 0xfa, 0x31, 0x0b, 0x5b, 0x2f, 0xbb, 0xb5, 0xf4, 0x4d, 0x43, 0x5c, 0x4d, 0x4a, 0xa9,
	0xe3, 0x8e, 0xf9, 0x29, 0x3c, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x20, 0xaa, 0x27, 0xad, 0xd7,
	0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a CoinSymbol by id.
	CoinSymbol(ctx context.Context, in *QueryGetCoinSymbolRequest, opts ...grpc.CallOption) (*QueryGetCoinSymbolResponse, error)
	// Queries a list of CoinSymbol items.
	CoinSymbolAll(ctx context.Context, in *QueryAllCoinSymbolRequest, opts ...grpc.CallOption) (*QueryAllCoinSymbolResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/colosseum.colosseum.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CoinSymbol(ctx context.Context, in *QueryGetCoinSymbolRequest, opts ...grpc.CallOption) (*QueryGetCoinSymbolResponse, error) {
	out := new(QueryGetCoinSymbolResponse)
	err := c.cc.Invoke(ctx, "/colosseum.colosseum.Query/CoinSymbol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CoinSymbolAll(ctx context.Context, in *QueryAllCoinSymbolRequest, opts ...grpc.CallOption) (*QueryAllCoinSymbolResponse, error) {
	out := new(QueryAllCoinSymbolResponse)
	err := c.cc.Invoke(ctx, "/colosseum.colosseum.Query/CoinSymbolAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a CoinSymbol by id.
	CoinSymbol(context.Context, *QueryGetCoinSymbolRequest) (*QueryGetCoinSymbolResponse, error)
	// Queries a list of CoinSymbol items.
	CoinSymbolAll(context.Context, *QueryAllCoinSymbolRequest) (*QueryAllCoinSymbolResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CoinSymbol(ctx context.Context, req *QueryGetCoinSymbolRequest) (*QueryGetCoinSymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinSymbol not implemented")
}
func (*UnimplementedQueryServer) CoinSymbolAll(ctx context.Context, req *QueryAllCoinSymbolRequest) (*QueryAllCoinSymbolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinSymbolAll not implemented")
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
		FullMethod: "/colosseum.colosseum.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CoinSymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCoinSymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinSymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colosseum.colosseum.Query/CoinSymbol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinSymbol(ctx, req.(*QueryGetCoinSymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CoinSymbolAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCoinSymbolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinSymbolAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/colosseum.colosseum.Query/CoinSymbolAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinSymbolAll(ctx, req.(*QueryAllCoinSymbolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "colosseum.colosseum.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CoinSymbol",
			Handler:    _Query_CoinSymbol_Handler,
		},
		{
			MethodName: "CoinSymbolAll",
			Handler:    _Query_CoinSymbolAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "colosseum/query.proto",
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

func (m *QueryGetCoinSymbolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCoinSymbolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCoinSymbolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCoinSymbolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCoinSymbolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCoinSymbolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CoinSymbol.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllCoinSymbolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCoinSymbolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCoinSymbolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCoinSymbolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCoinSymbolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCoinSymbolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.CoinSymbol) > 0 {
		for iNdEx := len(m.CoinSymbol) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoinSymbol[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGetCoinSymbolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetCoinSymbolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CoinSymbol.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCoinSymbolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCoinSymbolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CoinSymbol) > 0 {
		for _, e := range m.CoinSymbol {
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
func (m *QueryGetCoinSymbolRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetCoinSymbolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCoinSymbolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryGetCoinSymbolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetCoinSymbolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCoinSymbolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinSymbol", wireType)
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
			if err := m.CoinSymbol.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllCoinSymbolRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllCoinSymbolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCoinSymbolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryAllCoinSymbolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllCoinSymbolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCoinSymbolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinSymbol", wireType)
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
			m.CoinSymbol = append(m.CoinSymbol, CoinSymbol{})
			if err := m.CoinSymbol[len(m.CoinSymbol)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
