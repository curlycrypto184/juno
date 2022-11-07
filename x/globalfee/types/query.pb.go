// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: juno/globalfee/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// TODO:
type QueryWhiteListRequest struct {
}

func (m *QueryWhiteListRequest) Reset()         { *m = QueryWhiteListRequest{} }
func (m *QueryWhiteListRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWhiteListRequest) ProtoMessage()    {}
func (*QueryWhiteListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_613026c22387ef9c, []int{0}
}
func (m *QueryWhiteListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhiteListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhiteListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhiteListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhiteListRequest.Merge(m, src)
}
func (m *QueryWhiteListRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhiteListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhiteListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhiteListRequest proto.InternalMessageInfo

// TODO:
type QueryWhiteListResponse struct {
	AccountRecord []AccountRecord `protobuf:"bytes,1,rep,name=account_record,json=accountRecord,proto3" json:"account_record"`
}

func (m *QueryWhiteListResponse) Reset()         { *m = QueryWhiteListResponse{} }
func (m *QueryWhiteListResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWhiteListResponse) ProtoMessage()    {}
func (*QueryWhiteListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_613026c22387ef9c, []int{1}
}
func (m *QueryWhiteListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhiteListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhiteListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhiteListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhiteListResponse.Merge(m, src)
}
func (m *QueryWhiteListResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhiteListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhiteListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhiteListResponse proto.InternalMessageInfo

func (m *QueryWhiteListResponse) GetAccountRecord() []AccountRecord {
	if m != nil {
		return m.AccountRecord
	}
	return nil
}

// QueryMinimumGasPricesRequest is the request type for the
// Query/MinimumGasPrices RPC method.
type QueryMinimumGasPricesRequest struct {
}

func (m *QueryMinimumGasPricesRequest) Reset()         { *m = QueryMinimumGasPricesRequest{} }
func (m *QueryMinimumGasPricesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMinimumGasPricesRequest) ProtoMessage()    {}
func (*QueryMinimumGasPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_613026c22387ef9c, []int{2}
}
func (m *QueryMinimumGasPricesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinimumGasPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinimumGasPricesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinimumGasPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinimumGasPricesRequest.Merge(m, src)
}
func (m *QueryMinimumGasPricesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinimumGasPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinimumGasPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinimumGasPricesRequest proto.InternalMessageInfo

// QueryMinimumGasPricesResponse is the response type for the
// Query/MinimumGasPrices RPC method.
type QueryMinimumGasPricesResponse struct {
	MinimumGasPrices github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=minimum_gas_prices,json=minimumGasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"minimum_gas_prices,omitempty" yaml:"minimum_gas_prices"`
}

func (m *QueryMinimumGasPricesResponse) Reset()         { *m = QueryMinimumGasPricesResponse{} }
func (m *QueryMinimumGasPricesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMinimumGasPricesResponse) ProtoMessage()    {}
func (*QueryMinimumGasPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_613026c22387ef9c, []int{3}
}
func (m *QueryMinimumGasPricesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMinimumGasPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMinimumGasPricesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMinimumGasPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMinimumGasPricesResponse.Merge(m, src)
}
func (m *QueryMinimumGasPricesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMinimumGasPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMinimumGasPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMinimumGasPricesResponse proto.InternalMessageInfo

func (m *QueryMinimumGasPricesResponse) GetMinimumGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.MinimumGasPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryWhiteListRequest)(nil), "juno.globalfee.QueryWhiteListRequest")
	proto.RegisterType((*QueryWhiteListResponse)(nil), "juno.globalfee.QueryWhiteListResponse")
	proto.RegisterType((*QueryMinimumGasPricesRequest)(nil), "juno.globalfee.QueryMinimumGasPricesRequest")
	proto.RegisterType((*QueryMinimumGasPricesResponse)(nil), "juno.globalfee.QueryMinimumGasPricesResponse")
}

func init() { proto.RegisterFile("juno/globalfee/query.proto", fileDescriptor_613026c22387ef9c) }

var fileDescriptor_613026c22387ef9c = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x6c, 0x96, 0x1f, 0x09, 0x23, 0x56, 0x2b, 0x8b, 0xbf, 0x0d, 0xdd, 0x14, 0x22, 0x40, 0x08,
	0x76, 0x6d, 0x75, 0xb9, 0x71, 0xa3, 0x45, 0x42, 0xe2, 0x47, 0x82, 0x5e, 0x90, 0xb8, 0x54, 0x8e,
	0x6b, 0xb2, 0x86, 0xd8, 0x5f, 0x36, 0x76, 0x80, 0x1e, 0xb8, 0xf0, 0x04, 0x48, 0x1c, 0xb9, 0x72,
	0xe2, 0x19, 0x78, 0x80, 0x3d, 0xae, 0xc4, 0x85, 0x53, 0x40, 0x2d, 0x27, 0x8e, 0x3c, 0x01, 0x8a,
	0x93, 0xae, 0x4a, 0x4a, 0x61, 0x4f, 0x89, 0x3c, 0x63, 0x7f, 0x33, 0xf3, 0x0d, 0xf2, 0x9f, 0xe7,
	0x1a, 0x68, 0x9c, 0x40, 0xc4, 0x92, 0x67, 0x42, 0xd0, 0xdd, 0x5c, 0x64, 0x63, 0x92, 0x66, 0x60,
	0x01, 0xaf, 0x96, 0x18, 0x39, 0xc0, 0xfc, 0xd3, 0x31, 0xc4, 0xe0, 0x20, 0x5a, 0xfe, 0x55, 0x2c,
	0xbf, 0x1d, 0x03, 0xc4, 0x89, 0xa0, 0x2c, 0x95, 0x94, 0x69, 0x0d, 0x96, 0x59, 0x09, 0xda, 0xd4,
	0x68, 0xc0, 0xc1, 0x28, 0x30, 0x34, 0x62, 0x46, 0xd0, 0x97, 0xdd, 0x48, 0x58, 0xd6, 0xa5, 0x1c,
	0xa4, 0x9e, 0xe1, 0x8d, 0xf9, 0xaf, 0x76, 0xa4, 0x15, 0x89, 0x34, 0xb6, 0xc2, 0xc3, 0x73, 0xe8,
	0xcc, 0xe3, 0x52, 0xd2, 0x93, 0xf2, 0xfc, 0x81, 0x34, 0x76, 0x20, 0x76, 0x73, 0x61, 0x6c, 0x38,
	0x42, 0x67, 0x9b, 0x80, 0x49, 0x41, 0x1b, 0x81, 0xef, 0xa1, 0x55, 0xc6, 0x39, 0xe4, 0xda, 0x0e,
	0x33, 0xc1, 0x21, 0x1b, 0x9d, 0xf7, 0x2e, 0x1e, 0xb9, 0x76, 0x72, 0x7b, 0x83, 0xfc, 0xe9, 0x87,
	0xdc, 0xae, 0x58, 0x03, 0x47, 0xea, 0x1d, 0xdd, 0x2b, 0x3a, 0xad, 0xc1, 0x29, 0x36, 0x7f, 0x18,
	0x06, 0xa8, 0xed, 0xa6, 0x3c, 0x94, 0x5a, 0xaa, 0x5c, 0xdd, 0x65, 0xe6, 0x51, 0x26, 0xb9, 0x30,
	0x33, 0x15, 0x85, 0x87, 0x36, 0x96, 0x10, 0x6a, 0x35, 0x9f, 0x3d, 0x84, 0x55, 0x05, 0x0e, 0x63,
	0x66, 0x86, 0xa9, 0x83, 0x6b, 0x49, 0x6d, 0x52, 0xc5, 0x43, 0xca, 0x78, 0x48, 0x1d, 0x0f, 0xb9,
	0x23, 0x78, 0x1f, 0xa4, 0xee, 0xa5, 0xa5, 0xa2, 0x9f, 0x45, 0xa7, 0xbd, 0x78, 0x7f, 0x13, 0x94,
	0xb4, 0x42, 0xa5, 0x76, 0xfc, 0xab, 0xe8, 0xac, 0x8f, 0x99, 0x4a, 0x6e, 0x85, 0x8b, 0xac, 0xf0,
	0xd3, 0xb7, 0xce, 0x8d, 0x58, 0xda, 0x9d, 0x3c, 0x22, 0x1c, 0x14, 0xad, 0x77, 0x51, 0x7d, 0xb6,
	0xcc, 0xe8, 0x05, 0xb5, 0xe3, 0x54, 0x98, 0xd9, 0x40, 0x33, 0x58, 0x53, 0x0d, 0x1b, 0xdb, 0x1f,
	0x57, 0xd0, 0x31, 0x67, 0x10, 0x7f, 0xf0, 0xd0, 0x5a, 0xd3, 0x25, 0xde, 0x6c, 0x66, 0xfa, 0xaf,
	0xb4, 0xfc, 0xad, 0x43, 0xb2, 0xab, 0xe8, 0xc2, 0xeb, 0x6f, 0xbf, 0xfc, 0x78, 0xbf, 0x72, 0x19,
	0x87, 0xb4, 0x51, 0x92, 0x45, 0xa7, 0xf8, 0x0d, 0x3a, 0x71, 0xd0, 0x04, 0x7c, 0xe5, 0xaf, 0x73,
	0x9a, 0x15, 0xf2, 0xaf, 0xfe, 0x8f, 0x56, 0xeb, 0xb8, 0xe4, 0x74, 0x5c, 0xc0, 0xeb, 0x74, 0x59,
	0x59, 0x7b, 0xf7, 0xf7, 0x26, 0x81, 0xb7, 0x3f, 0x09, 0xbc, 0xef, 0x93, 0xc0, 0x7b, 0x37, 0x0d,
	0x5a, 0xfb, 0xd3, 0xa0, 0xf5, 0x75, 0x1a, 0xb4, 0x9e, 0x76, 0xe7, 0xf2, 0xef, 0xbb, 0xe0, 0xfb,
	0xa0, 0x6d, 0xc6, 0xb8, 0x35, 0xd5, 0x73, 0xaf, 0xe7, 0x1e, 0x74, 0xeb, 0x88, 0x8e, 0xbb, 0xea,
	0xdf, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x99, 0xff, 0xd4, 0x79, 0x9c, 0x03, 0x00, 0x00,
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
	// Params queries the parameters of the globalfee module.
	MinimumGasPrices(ctx context.Context, in *QueryMinimumGasPricesRequest, opts ...grpc.CallOption) (*QueryMinimumGasPricesResponse, error)
	// returns addresses which can override the global fee in case of an attack
	WhiteList(ctx context.Context, in *QueryWhiteListRequest, opts ...grpc.CallOption) (*QueryWhiteListResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MinimumGasPrices(ctx context.Context, in *QueryMinimumGasPricesRequest, opts ...grpc.CallOption) (*QueryMinimumGasPricesResponse, error) {
	out := new(QueryMinimumGasPricesResponse)
	err := c.cc.Invoke(ctx, "/juno.globalfee.Query/MinimumGasPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WhiteList(ctx context.Context, in *QueryWhiteListRequest, opts ...grpc.CallOption) (*QueryWhiteListResponse, error) {
	out := new(QueryWhiteListResponse)
	err := c.cc.Invoke(ctx, "/juno.globalfee.Query/WhiteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of the globalfee module.
	MinimumGasPrices(context.Context, *QueryMinimumGasPricesRequest) (*QueryMinimumGasPricesResponse, error)
	// returns addresses which can override the global fee in case of an attack
	WhiteList(context.Context, *QueryWhiteListRequest) (*QueryWhiteListResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MinimumGasPrices(ctx context.Context, req *QueryMinimumGasPricesRequest) (*QueryMinimumGasPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinimumGasPrices not implemented")
}
func (*UnimplementedQueryServer) WhiteList(ctx context.Context, req *QueryWhiteListRequest) (*QueryWhiteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhiteList not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MinimumGasPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMinimumGasPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinimumGasPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juno.globalfee.Query/MinimumGasPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinimumGasPrices(ctx, req.(*QueryMinimumGasPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWhiteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juno.globalfee.Query/WhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WhiteList(ctx, req.(*QueryWhiteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "juno.globalfee.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MinimumGasPrices",
			Handler:    _Query_MinimumGasPrices_Handler,
		},
		{
			MethodName: "WhiteList",
			Handler:    _Query_WhiteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "juno/globalfee/query.proto",
}

func (m *QueryWhiteListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhiteListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhiteListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryWhiteListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhiteListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhiteListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountRecord) > 0 {
		for iNdEx := len(m.AccountRecord) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountRecord[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryMinimumGasPricesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinimumGasPricesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinimumGasPricesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMinimumGasPricesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMinimumGasPricesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMinimumGasPricesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for iNdEx := len(m.MinimumGasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinimumGasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryWhiteListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryWhiteListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccountRecord) > 0 {
		for _, e := range m.AccountRecord {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryMinimumGasPricesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMinimumGasPricesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for _, e := range m.MinimumGasPrices {
			l = e.Size()
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
func (m *QueryWhiteListRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWhiteListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhiteListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryWhiteListResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWhiteListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhiteListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountRecord", wireType)
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
			m.AccountRecord = append(m.AccountRecord, AccountRecord{})
			if err := m.AccountRecord[len(m.AccountRecord)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryMinimumGasPricesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMinimumGasPricesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinimumGasPricesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryMinimumGasPricesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMinimumGasPricesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMinimumGasPricesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrices", wireType)
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
			m.MinimumGasPrices = append(m.MinimumGasPrices, types.DecCoin{})
			if err := m.MinimumGasPrices[len(m.MinimumGasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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