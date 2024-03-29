// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeburner/total_burned_furyas_amount.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// TotalBurnedFuryasAmount defines total amount of burned furya fees
type TotalBurnedFuryasAmount struct {
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin" yaml:"coin"`
}

func (m *TotalBurnedFuryasAmount) Reset()         { *m = TotalBurnedFuryasAmount{} }
func (m *TotalBurnedFuryasAmount) String() string { return proto.CompactTextString(m) }
func (*TotalBurnedFuryasAmount) ProtoMessage()    {}
func (*TotalBurnedFuryasAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a29e81c9c2bc4df, []int{0}
}
func (m *TotalBurnedFuryasAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalBurnedFuryasAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalBurnedFuryasAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalBurnedFuryasAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalBurnedFuryasAmount.Merge(m, src)
}
func (m *TotalBurnedFuryasAmount) XXX_Size() int {
	return m.Size()
}
func (m *TotalBurnedFuryasAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalBurnedFuryasAmount.DiscardUnknown(m)
}

var xxx_messageInfo_TotalBurnedFuryasAmount proto.InternalMessageInfo

func (m *TotalBurnedFuryasAmount) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*TotalBurnedFuryasAmount)(nil), "furya.feeburner.TotalBurnedFuryasAmount")
}

func init() {
	proto.RegisterFile("feeburner/total_burned_furyas_amount.proto", fileDescriptor_3a29e81c9c2bc4df)
}

var fileDescriptor_3a29e81c9c2bc4df = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x4b, 0x4d, 0x4d,
	0x2a, 0x2d, 0xca, 0x4b, 0x2d, 0xd2, 0x2f, 0xc9, 0x2f, 0x49, 0xcc, 0x89, 0x07, 0x73, 0x52, 0xe2,
	0xf3, 0x52, 0x4b, 0x4b, 0x8a, 0xf2, 0xf3, 0x8a, 0xe3, 0x13, 0x73, 0xf3, 0x4b, 0xf3, 0x4a, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xa1, 0xc2, 0x7a, 0x70, 0x5d, 0x52, 0x72, 0xc9, 0xf9,
	0xc5, 0xb9, 0xf9, 0xc5, 0xfa, 0x49, 0x89, 0xc5, 0xa9, 0xfa, 0x65, 0x86, 0x49, 0xa9, 0x25, 0x89,
	0x86, 0xfa, 0xc9, 0xf9, 0x99, 0x79, 0x10, 0x2d, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6,
	0x3e, 0x88, 0x05, 0x11, 0x55, 0x8a, 0xe7, 0x92, 0x0c, 0x01, 0x59, 0xe7, 0x04, 0xb6, 0xcd, 0x0f,
	0x6a, 0x99, 0x23, 0xd8, 0x2e, 0x21, 0x27, 0x2e, 0x16, 0x90, 0x01, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0x92, 0x7a, 0x10, 0x1b, 0xf4, 0x40, 0x36, 0xe8, 0x41, 0x6d, 0xd0, 0x73, 0xce, 0xcf,
	0xcc, 0x73, 0x12, 0x3e, 0x71, 0x4f, 0x9e, 0xe1, 0xd3, 0x3d, 0x79, 0xee, 0xca, 0xc4, 0xdc, 0x1c,
	0x2b, 0x25, 0x90, 0x26, 0xa5, 0x20, 0xb0, 0x5e, 0x27, 0xaf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x87, 0x7a, 0x47, 0x37, 0xbf, 0x28, 0x1d, 0xc6, 0xd6, 0xaf, 0xd0, 0x47, 0x0a, 0x92, 0xca, 0x82,
	0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x9b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xb4, 0xe9,
	0x66, 0x2c, 0x01, 0x00, 0x00,
}

func (m *TotalBurnedFuryasAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalBurnedFuryasAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalBurnedFuryasAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTotalBurnedFuryasAmount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTotalBurnedFuryasAmount(dAtA []byte, offset int, v uint64) int {
	offset -= sovTotalBurnedFuryasAmount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TotalBurnedFuryasAmount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovTotalBurnedFuryasAmount(uint64(l))
	return n
}

func sovTotalBurnedFuryasAmount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTotalBurnedFuryasAmount(x uint64) (n int) {
	return sovTotalBurnedFuryasAmount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TotalBurnedFuryasAmount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTotalBurnedFuryasAmount
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
			return fmt.Errorf("proto: TotalBurnedFuryasAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalBurnedFuryasAmount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTotalBurnedFuryasAmount
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
				return ErrInvalidLengthTotalBurnedFuryasAmount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTotalBurnedFuryasAmount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTotalBurnedFuryasAmount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTotalBurnedFuryasAmount
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
func skipTotalBurnedFuryasAmount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTotalBurnedFuryasAmount
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
					return 0, ErrIntOverflowTotalBurnedFuryasAmount
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
					return 0, ErrIntOverflowTotalBurnedFuryasAmount
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
				return 0, ErrInvalidLengthTotalBurnedFuryasAmount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTotalBurnedFuryasAmount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTotalBurnedFuryasAmount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTotalBurnedFuryasAmount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTotalBurnedFuryasAmount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTotalBurnedFuryasAmount = fmt.Errorf("proto: unexpected end of group")
)
