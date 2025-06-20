// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquiditypool/v1/position.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Position
type Position struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PoolId    uint64 `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	LowerTick int64  `protobuf:"varint,4,opt,name=lower_tick,json=lowerTick,proto3" json:"lower_tick,omitempty"`
	UpperTick int64  `protobuf:"varint,5,opt,name=upper_tick,json=upperTick,proto3" json:"upper_tick,omitempty"`
	Liquidity string `protobuf:"bytes,6,opt,name=liquidity,proto3" json:"liquidity,omitempty"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5ccdf3cf21e1c35, []int{0}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Position.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return m.Size()
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Position) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Position) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *Position) GetLowerTick() int64 {
	if m != nil {
		return m.LowerTick
	}
	return 0
}

func (m *Position) GetUpperTick() int64 {
	if m != nil {
		return m.UpperTick
	}
	return 0
}

func (m *Position) GetLiquidity() string {
	if m != nil {
		return m.Liquidity
	}
	return ""
}

func init() {
	proto.RegisterType((*Position)(nil), "sunrise.liquiditypool.v1.Position")
}

func init() {
	proto.RegisterFile("sunrise/liquiditypool/v1/position.proto", fileDescriptor_a5ccdf3cf21e1c35)
}

var fileDescriptor_a5ccdf3cf21e1c35 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x2e, 0xcd, 0x2b,
	0xca, 0x2c, 0x4e, 0xd5, 0xcf, 0xc9, 0x2c, 0x2c, 0xcd, 0x4c, 0xc9, 0x2c, 0xa9, 0x2c, 0xc8, 0xcf,
	0xcf, 0xd1, 0x2f, 0x33, 0xd4, 0x2f, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0x2a, 0xd4, 0x43, 0x51, 0xa8, 0x57, 0x66, 0x28, 0x25, 0x99,
	0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x1c, 0x0f, 0x56, 0xa7, 0x0f, 0xe1, 0x40, 0x34, 0x29, 0x5d, 0x66,
	0xe4, 0xe2, 0x08, 0x80, 0x9a, 0x23, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x12, 0xc4, 0x94, 0x99, 0x22, 0x64, 0xc4, 0xc5, 0x9e, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c,
	0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0x71, 0x69, 0x8b, 0xae, 0x08, 0x54, 0xbf, 0x23,
	0x44, 0x26, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xa6, 0x50, 0x48, 0x9c, 0x8b, 0x1d, 0x64,
	0x6d, 0x7c, 0x66, 0x8a, 0x04, 0x33, 0xd8, 0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0x48, 0x96, 0x8b,
	0x2b, 0x27, 0xbf, 0x3c, 0xb5, 0x28, 0xbe, 0x24, 0x33, 0x39, 0x5b, 0x82, 0x45, 0x81, 0x51, 0x83,
	0x39, 0x88, 0x13, 0x2c, 0x12, 0x92, 0x99, 0x9c, 0x0d, 0x92, 0x2e, 0x2d, 0x28, 0x80, 0x49, 0xb3,
	0x42, 0xa4, 0xc1, 0x22, 0x60, 0x69, 0x1d, 0x2e, 0x4e, 0xb8, 0xb7, 0x24, 0xd8, 0xc0, 0x8e, 0xe1,
	0xbb, 0xb4, 0x45, 0x97, 0x0b, 0xea, 0x18, 0x97, 0xd4, 0xe4, 0x20, 0x84, 0x02, 0x27, 0xff, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x86, 0x57, 0x4e, 0x62, 0x65, 0x6a, 0x11, 0x8c, 0xa3, 0x5f,
	0x81, 0x16, 0xce, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xd0, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x80, 0xc1, 0xe8, 0x13, 0x8d, 0x01, 0x00, 0x00,
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Position) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Liquidity) > 0 {
		i -= len(m.Liquidity)
		copy(dAtA[i:], m.Liquidity)
		i = encodeVarintPosition(dAtA, i, uint64(len(m.Liquidity)))
		i--
		dAtA[i] = 0x32
	}
	if m.UpperTick != 0 {
		i = encodeVarintPosition(dAtA, i, uint64(m.UpperTick))
		i--
		dAtA[i] = 0x28
	}
	if m.LowerTick != 0 {
		i = encodeVarintPosition(dAtA, i, uint64(m.LowerTick))
		i--
		dAtA[i] = 0x20
	}
	if m.PoolId != 0 {
		i = encodeVarintPosition(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPosition(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPosition(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovPosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Position) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPosition(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPosition(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovPosition(uint64(m.PoolId))
	}
	if m.LowerTick != 0 {
		n += 1 + sovPosition(uint64(m.LowerTick))
	}
	if m.UpperTick != 0 {
		n += 1 + sovPosition(uint64(m.UpperTick))
	}
	l = len(m.Liquidity)
	if l > 0 {
		n += 1 + l + sovPosition(uint64(l))
	}
	return n
}

func sovPosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPosition(x uint64) (n int) {
	return sovPosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPosition
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
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
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
				return ErrInvalidLengthPosition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LowerTick", wireType)
			}
			m.LowerTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LowerTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpperTick", wireType)
			}
			m.UpperTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpperTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPosition
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
				return ErrInvalidLengthPosition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Liquidity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPosition
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
func skipPosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPosition
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
					return 0, ErrIntOverflowPosition
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
					return 0, ErrIntOverflowPosition
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
				return 0, ErrInvalidLengthPosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPosition = fmt.Errorf("proto: unexpected end of group")
)
