// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fx/staking/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
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

type LPToken struct {
	ValidatorAddr string                                 `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
	Address       string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name          string                                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Symbol        string                                 `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal       uint32                                 `protobuf:"varint,5,opt,name=decimal,proto3" json:"decimal,omitempty"`
	TotalSupply   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=total_supply,json=totalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_supply"`
}

func (m *LPToken) Reset()         { *m = LPToken{} }
func (m *LPToken) String() string { return proto.CompactTextString(m) }
func (*LPToken) ProtoMessage()    {}
func (*LPToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b5fb49b8596c32e, []int{0}
}
func (m *LPToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LPToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LPToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LPToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LPToken.Merge(m, src)
}
func (m *LPToken) XXX_Size() int {
	return m.Size()
}
func (m *LPToken) XXX_DiscardUnknown() {
	xxx_messageInfo_LPToken.DiscardUnknown(m)
}

var xxx_messageInfo_LPToken proto.InternalMessageInfo

func (m *LPToken) GetValidatorAddr() string {
	if m != nil {
		return m.ValidatorAddr
	}
	return ""
}

func (m *LPToken) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LPToken) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LPToken) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *LPToken) GetDecimal() uint32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

// QueryParamsRequest is request type for the Query/ValidatorLPToken RPC method.
type QueryValidatorLPTokenRequest struct {
	// validator_addr defines the validator address to query for.
	ValidatorAddr string `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3" json:"validator_addr,omitempty"`
}

func (m *QueryValidatorLPTokenRequest) Reset()         { *m = QueryValidatorLPTokenRequest{} }
func (m *QueryValidatorLPTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorLPTokenRequest) ProtoMessage()    {}
func (*QueryValidatorLPTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b5fb49b8596c32e, []int{1}
}
func (m *QueryValidatorLPTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorLPTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorLPTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorLPTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorLPTokenRequest.Merge(m, src)
}
func (m *QueryValidatorLPTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorLPTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorLPTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorLPTokenRequest proto.InternalMessageInfo

func (m *QueryValidatorLPTokenRequest) GetValidatorAddr() string {
	if m != nil {
		return m.ValidatorAddr
	}
	return ""
}

// QueryValidatorLPTokenResponse is response type for the Query/ValidatorLPToken RPC method.
type QueryValidatorLPTokenResponse struct {
	LpToken *LPToken `protobuf:"bytes,1,opt,name=lp_token,json=lpToken,proto3" json:"lp_token,omitempty"`
}

func (m *QueryValidatorLPTokenResponse) Reset()         { *m = QueryValidatorLPTokenResponse{} }
func (m *QueryValidatorLPTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryValidatorLPTokenResponse) ProtoMessage()    {}
func (*QueryValidatorLPTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b5fb49b8596c32e, []int{2}
}
func (m *QueryValidatorLPTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryValidatorLPTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryValidatorLPTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryValidatorLPTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryValidatorLPTokenResponse.Merge(m, src)
}
func (m *QueryValidatorLPTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryValidatorLPTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryValidatorLPTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryValidatorLPTokenResponse proto.InternalMessageInfo

func (m *QueryValidatorLPTokenResponse) GetLpToken() *LPToken {
	if m != nil {
		return m.LpToken
	}
	return nil
}

func init() {
	proto.RegisterType((*LPToken)(nil), "fx.staking.v1.LPToken")
	proto.RegisterType((*QueryValidatorLPTokenRequest)(nil), "fx.staking.v1.QueryValidatorLPTokenRequest")
	proto.RegisterType((*QueryValidatorLPTokenResponse)(nil), "fx.staking.v1.QueryValidatorLPTokenResponse")
}

func init() { proto.RegisterFile("fx/staking/v1/query.proto", fileDescriptor_6b5fb49b8596c32e) }

var fileDescriptor_6b5fb49b8596c32e = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0x6a, 0xbb, 0x5b, 0xa7, 0xae, 0xc8, 0x20, 0x25, 0x2e, 0x35, 0x5b, 0x16, 0x95,
	0xa2, 0x6d, 0x86, 0xac, 0x9e, 0x7a, 0xb3, 0xa0, 0x20, 0x78, 0xb0, 0x51, 0x3c, 0x78, 0x59, 0x26,
	0xc9, 0x6c, 0x0c, 0x9b, 0xcc, 0x9b, 0x66, 0x26, 0xcb, 0x2e, 0xe2, 0xc5, 0xbf, 0x40, 0xf0, 0xe2,
	0xb1, 0x57, 0xcf, 0xfe, 0x13, 0x3d, 0x16, 0xbc, 0x88, 0x87, 0x22, 0xbb, 0x1e, 0xbc, 0xf8, 0x3f,
	0x48, 0x26, 0x89, 0xeb, 0x16, 0x94, 0x9e, 0xf2, 0x7e, 0x7e, 0x32, 0xef, 0xfb, 0x1e, 0xbe, 0x31,
	0x9c, 0x50, 0xa5, 0xd9, 0x28, 0x16, 0x11, 0x1d, 0xbb, 0xf4, 0x28, 0xe7, 0xd9, 0xd4, 0x91, 0x19,
	0x68, 0x20, 0xed, 0xe1, 0xc4, 0xa9, 0x52, 0xce, 0xd8, 0xed, 0xdc, 0x0d, 0x40, 0xa5, 0xa0, 0xa8,
	0xcf, 0x14, 0x2f, 0xeb, 0xe8, 0xd8, 0xf5, 0xb9, 0x66, 0x2e, 0x95, 0x2c, 0x8a, 0x05, 0xd3, 0x31,
	0x88, 0xb2, 0xb5, 0x73, 0xab, 0xaa, 0x5d, 0x90, 0xcb, 0xc2, 0x1a, 0x57, 0x56, 0x5d, 0x8f, 0x20,
	0x02, 0x63, 0xd2, 0xc2, 0xaa, 0xa2, 0x5b, 0x11, 0x40, 0x94, 0x70, 0xca, 0x64, 0x4c, 0x99, 0x10,
	0xa0, 0x0d, 0x58, 0x95, 0xd9, 0xde, 0x2f, 0x84, 0x5b, 0x4f, 0x9f, 0xbd, 0x80, 0x11, 0x17, 0xe4,
	0x36, 0xbe, 0x3a, 0x66, 0x49, 0x1c, 0x32, 0x0d, 0xd9, 0x80, 0x85, 0x61, 0x66, 0xa1, 0x6d, 0xb4,
	0x73, 0xd9, 0x6b, 0xff, 0x89, 0x3e, 0x0c, 0xc3, 0x8c, 0x58, 0xb8, 0x55, 0x24, 0xb9, 0x52, 0xd6,
	0x8a, 0xc9, 0xd7, 0x2e, 0x21, 0x78, 0x55, 0xb0, 0x94, 0x5b, 0x97, 0x4c, 0xd8, 0xd8, 0x64, 0x13,
	0x37, 0xd5, 0x34, 0xf5, 0x21, 0xb1, 0x56, 0x4d, 0xb4, 0xf2, 0x0a, 0x4a, 0xc8, 0x83, 0x38, 0x65,
	0x89, 0xb5, 0xb6, 0x8d, 0x76, 0xda, 0x5e, 0xed, 0x92, 0x43, 0x7c, 0x45, 0x83, 0x66, 0xc9, 0x40,
	0xe5, 0x52, 0x26, 0x53, 0xab, 0x59, 0xf4, 0x1d, 0x38, 0x27, 0x67, 0xdd, 0xc6, 0xb7, 0xb3, 0xee,
	0x9d, 0x28, 0xd6, 0xaf, 0x73, 0xdf, 0x09, 0x20, 0xa5, 0x95, 0x2a, 0xe5, 0x67, 0x4f, 0x85, 0x23,
	0xaa, 0xa7, 0x92, 0x2b, 0xe7, 0x89, 0xd0, 0xde, 0x86, 0x61, 0x3c, 0x37, 0x88, 0xfd, 0xf5, 0x8f,
	0xc7, 0x5d, 0xf4, 0xf3, 0xb8, 0x8b, 0x7a, 0x8f, 0xf0, 0xd6, 0x61, 0xa1, 0xf5, 0xcb, 0x7a, 0xa4,
	0x6a, 0x78, 0x8f, 0x1f, 0xe5, 0x5c, 0xe9, 0x0b, 0x6a, 0xd0, 0xf3, 0xf0, 0xcd, 0x7f, 0x60, 0x94,
	0x04, 0xa1, 0x38, 0x71, 0xf1, 0x7a, 0x22, 0x07, 0xba, 0x88, 0x19, 0xc2, 0x46, 0x7f, 0xd3, 0x59,
	0xda, 0xbf, 0x53, 0x77, 0xb4, 0x12, 0x69, 0x8c, 0xfe, 0x67, 0x84, 0xd7, 0x0c, 0x94, 0x7c, 0x42,
	0xf8, 0xda, 0x79, 0x32, 0xb9, 0x77, 0xae, 0xff, 0x7f, 0x63, 0x74, 0x76, 0x2f, 0x56, 0x5c, 0x3e,
	0xb6, 0xb7, 0xff, 0xee, 0xcb, 0x8f, 0x0f, 0x2b, 0x0f, 0x48, 0x9f, 0x2e, 0x5f, 0xef, 0x42, 0x89,
	0x7a, 0x16, 0xfa, 0x66, 0x59, 0x9d, 0xb7, 0x07, 0x8f, 0x4f, 0x66, 0x36, 0x3a, 0x9d, 0xd9, 0xe8,
	0xfb, 0xcc, 0x46, 0xef, 0xe7, 0x76, 0xe3, 0x74, 0x6e, 0x37, 0xbe, 0xce, 0xed, 0xc6, 0xab, 0xdd,
	0xbf, 0x36, 0x35, 0xcc, 0x45, 0x50, 0x9c, 0xdd, 0x84, 0x0e, 0x27, 0x7b, 0x01, 0x64, 0x9c, 0x2e,
	0x7e, 0x64, 0x76, 0xe6, 0x37, 0xcd, 0x3d, 0xde, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x16, 0xe3,
	0xe0, 0xf1, 0x41, 0x03, 0x00, 0x00,
}

func (this *LPToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LPToken)
	if !ok {
		that2, ok := that.(LPToken)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ValidatorAddr != that1.ValidatorAddr {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Decimal != that1.Decimal {
		return false
	}
	if !this.TotalSupply.Equal(that1.TotalSupply) {
		return false
	}
	return true
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
	// ValidatorLPToken queries the validator lp token.
	ValidatorLPToken(ctx context.Context, in *QueryValidatorLPTokenRequest, opts ...grpc.CallOption) (*QueryValidatorLPTokenResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ValidatorLPToken(ctx context.Context, in *QueryValidatorLPTokenRequest, opts ...grpc.CallOption) (*QueryValidatorLPTokenResponse, error) {
	out := new(QueryValidatorLPTokenResponse)
	err := c.cc.Invoke(ctx, "/fx.staking.v1.Query/ValidatorLPToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ValidatorLPToken queries the validator lp token.
	ValidatorLPToken(context.Context, *QueryValidatorLPTokenRequest) (*QueryValidatorLPTokenResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ValidatorLPToken(ctx context.Context, req *QueryValidatorLPTokenRequest) (*QueryValidatorLPTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorLPToken not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ValidatorLPToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorLPTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorLPToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.staking.v1.Query/ValidatorLPToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorLPToken(ctx, req.(*QueryValidatorLPTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fx.staking.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidatorLPToken",
			Handler:    _Query_ValidatorLPToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fx/staking/v1/query.proto",
}

func (m *LPToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LPToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LPToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalSupply.Size()
		i -= size
		if _, err := m.TotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Decimal != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Decimal))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
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
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorLPTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorLPTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorLPTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryValidatorLPTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryValidatorLPTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryValidatorLPTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LpToken != nil {
		{
			size, err := m.LpToken.MarshalToSizedBuffer(dAtA[:i])
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
func (m *LPToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Decimal != 0 {
		n += 1 + sovQuery(uint64(m.Decimal))
	}
	l = m.TotalSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryValidatorLPTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryValidatorLPTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LpToken != nil {
		l = m.LpToken.Size()
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
func (m *LPToken) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LPToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LPToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
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
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryValidatorLPTokenRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryValidatorLPTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorLPTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
			m.ValidatorAddr = string(dAtA[iNdEx:postIndex])
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
func (m *QueryValidatorLPTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryValidatorLPTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryValidatorLPTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LpToken", wireType)
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
			if m.LpToken == nil {
				m.LpToken = &LPToken{}
			}
			if err := m.LpToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
