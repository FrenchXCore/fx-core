// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fx/erc20/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the module's genesis state.
type GenesisState struct {
	// module parameters
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// registered token pairs
	TokenPairs []TokenPair `protobuf:"bytes,2,rep,name=token_pairs,json=tokenPairs,proto3" json:"token_pairs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1892ef01bf6175af, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTokenPairs() []TokenPair {
	if m != nil {
		return m.TokenPairs
	}
	return nil
}

// Params defines the erc20 module params
type Params struct {
	// parameter to enable the intrarelaying of Cosmos coins <--> ERC20 tokens.
	EnableErc20 bool `protobuf:"varint,1,opt,name=enable_erc20,json=enableErc20,proto3" json:"enable_erc20,omitempty" yaml:"enable_erc20"`
	// parameter to enable the EVM hook to convert an ERC20 token to a Cosmos
	// Coin by transferring the Tokens through a MsgEthereumTx to the
	// ModuleAddress Ethereum address.
	EnableEVMHook bool `protobuf:"varint,2,opt,name=enable_evm_hook,json=enableEvmHook,proto3" json:"enable_evm_hook,omitempty" yaml:"enable_evm_hook"`
	// parameter to set ibc timeout
	IbcTimeout time.Duration `protobuf:"bytes,3,opt,name=ibc_timeout,json=ibcTimeout,proto3,stdduration" json:"ibc_timeout" yaml:"ibc_timeout"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1892ef01bf6175af, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnableErc20() bool {
	if m != nil {
		return m.EnableErc20
	}
	return false
}

func (m *Params) GetEnableEVMHook() bool {
	if m != nil {
		return m.EnableEVMHook
	}
	return false
}

func (m *Params) GetIbcTimeout() time.Duration {
	if m != nil {
		return m.IbcTimeout
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fx.erc20.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "fx.erc20.v1.Params")
}

func init() { proto.RegisterFile("fx/erc20/v1/genesis.proto", fileDescriptor_1892ef01bf6175af) }

var fileDescriptor_1892ef01bf6175af = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0xcf, 0x6a, 0x9c, 0x40,
	0x18, 0x77, 0x92, 0xb2, 0x94, 0x31, 0xa1, 0x60, 0x4a, 0x62, 0x72, 0x18, 0x83, 0xa7, 0xd0, 0xd2,
	0x99, 0xee, 0xf6, 0x16, 0xe8, 0x45, 0x12, 0xda, 0x4b, 0x21, 0x98, 0xd0, 0x43, 0x2e, 0x32, 0xda,
	0xd1, 0x0c, 0xbb, 0x3a, 0xa2, 0xa3, 0x98, 0x5b, 0x1f, 0xa1, 0xc7, 0x3e, 0xd2, 0x1e, 0xf7, 0xd8,
	0x93, 0x2d, 0xee, 0x1b, 0x2c, 0xf4, 0x1e, 0x9c, 0x51, 0x70, 0x6f, 0xf3, 0xfd, 0xfe, 0x7d, 0x9f,
	0x3f, 0xe1, 0x79, 0xdc, 0x10, 0x56, 0x44, 0x8b, 0x8f, 0xa4, 0x9e, 0x93, 0x84, 0x65, 0xac, 0xe4,
	0x25, 0xce, 0x0b, 0x21, 0x85, 0x65, 0xc6, 0x0d, 0x56, 0x14, 0xae, 0xe7, 0x17, 0x67, 0x53, 0x9d,
	0x46, 0x95, 0xea, 0xe2, 0x6d, 0x22, 0x12, 0xa1, 0x9e, 0xa4, 0x7f, 0x0d, 0x28, 0x4a, 0x84, 0x48,
	0x56, 0x8c, 0xa8, 0x29, 0xac, 0x62, 0xf2, 0xa3, 0x2a, 0xa8, 0xe4, 0x22, 0xd3, 0xbc, 0xfb, 0x13,
	0xc0, 0xa3, 0x2f, 0x7a, 0xdb, 0xbd, 0xa4, 0x92, 0x59, 0x73, 0x38, 0xcb, 0x69, 0x41, 0xd3, 0xd2,
	0x06, 0x97, 0xe0, 0xca, 0x5c, 0x9c, 0xe0, 0xc9, 0x76, 0x7c, 0xa7, 0x28, 0xef, 0xd5, 0xba, 0x75,
	0x0c, 0x7f, 0x10, 0x5a, 0x9f, 0xa1, 0x29, 0xc5, 0x92, 0x65, 0x41, 0x4e, 0x79, 0x51, 0xda, 0x07,
	0x97, 0x87, 0x57, 0xe6, 0xe2, 0x74, 0xcf, 0xf7, 0xd0, 0xf3, 0x77, 0x94, 0x17, 0x83, 0x15, 0xca,
	0x11, 0x28, 0xdd, 0xff, 0x00, 0xce, 0x74, 0xae, 0x75, 0x0d, 0x8f, 0x58, 0x46, 0xc3, 0x15, 0x0b,
	0x94, 0x53, 0x9d, 0xf0, 0xda, 0x3b, 0xdb, 0xb5, 0xce, 0xc9, 0x33, 0x4d, 0x57, 0xd7, 0xee, 0x94,
	0x75, 0x7d, 0x53, 0x8f, 0xb7, 0xfd, 0x64, 0xdd, 0xc3, 0x37, 0x23, 0x5b, 0xa7, 0xc1, 0x93, 0x10,
	0x4b, 0xfb, 0x40, 0xd9, 0xdf, 0x77, 0xad, 0x73, 0x7c, 0xab, 0x95, 0xdf, 0xbf, 0x7d, 0x15, 0x62,
	0xb9, 0x6b, 0x9d, 0xd3, 0xfd, 0xbc, 0xc1, 0xe1, 0xfa, 0xc7, 0x43, 0x64, 0x9d, 0xf6, 0x42, 0xeb,
	0x11, 0x9a, 0x3c, 0x8c, 0x02, 0xc9, 0x53, 0x26, 0x2a, 0x69, 0x1f, 0xaa, 0x4a, 0xce, 0xb1, 0x2e,
	0x15, 0x8f, 0xa5, 0xe2, 0x9b, 0xa1, 0x54, 0x0f, 0xf5, 0x5f, 0xb7, 0x6b, 0x1d, 0x4b, 0xc7, 0x4f,
	0xbc, 0xee, 0xef, 0xbf, 0x0e, 0xf0, 0x21, 0x0f, 0xa3, 0x07, 0x0d, 0x78, 0x37, 0xeb, 0x0e, 0x81,
	0x4d, 0x87, 0xc0, 0xbf, 0x0e, 0x81, 0x5f, 0x5b, 0x64, 0x6c, 0xb6, 0xc8, 0xf8, 0xb3, 0x45, 0xc6,
	0xe3, 0xbb, 0x84, 0xcb, 0xa7, 0x2a, 0xc4, 0x91, 0x48, 0x49, 0x5c, 0x65, 0x51, 0x1f, 0xdd, 0x90,
	0xb8, 0xf9, 0x10, 0x89, 0x82, 0x91, 0xf1, 0xff, 0xcb, 0xe7, 0x9c, 0x95, 0xe1, 0x4c, 0x1d, 0xf1,
	0xe9, 0x25, 0x00, 0x00, 0xff, 0xff, 0xec, 0xcc, 0xac, 0x6b, 0x40, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenPairs) > 0 {
		for iNdEx := len(m.TokenPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.IbcTimeout, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.IbcTimeout):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if m.EnableEVMHook {
		i--
		if m.EnableEVMHook {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.EnableErc20 {
		i--
		if m.EnableErc20 {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TokenPairs) > 0 {
		for _, e := range m.TokenPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableErc20 {
		n += 2
	}
	if m.EnableEVMHook {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.IbcTimeout)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPairs = append(m.TokenPairs, TokenPair{})
			if err := m.TokenPairs[len(m.TokenPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableErc20", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableErc20 = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableEVMHook", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableEVMHook = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.IbcTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
