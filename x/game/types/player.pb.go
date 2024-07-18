// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babychain/game/player.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Player struct {
	Address         string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CurrentLevel    uint64     `protobuf:"varint,2,opt,name=currentLevel,proto3" json:"currentLevel,omitempty"`
	CurrentIncome   types.Coin `protobuf:"bytes,3,opt,name=currentIncome,proto3" json:"currentIncome"`
	NextIncome      types.Coin `protobuf:"bytes,4,opt,name=nextIncome,proto3" json:"nextIncome"`
	NextLevelPrice  types.Coin `protobuf:"bytes,5,opt,name=nextLevelPrice,proto3" json:"nextLevelPrice"`
	LastIncomeBlock uint64     `protobuf:"varint,6,opt,name=lastIncomeBlock,proto3" json:"lastIncomeBlock,omitempty"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ee082fcb113b15, []int{0}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Player.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return m.Size()
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Player) GetCurrentLevel() uint64 {
	if m != nil {
		return m.CurrentLevel
	}
	return 0
}

func (m *Player) GetCurrentIncome() types.Coin {
	if m != nil {
		return m.CurrentIncome
	}
	return types.Coin{}
}

func (m *Player) GetNextIncome() types.Coin {
	if m != nil {
		return m.NextIncome
	}
	return types.Coin{}
}

func (m *Player) GetNextLevelPrice() types.Coin {
	if m != nil {
		return m.NextLevelPrice
	}
	return types.Coin{}
}

func (m *Player) GetLastIncomeBlock() uint64 {
	if m != nil {
		return m.LastIncomeBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "babychain.game.Player")
}

func init() { proto.RegisterFile("babychain/game/player.proto", fileDescriptor_94ee082fcb113b15) }

var fileDescriptor_94ee082fcb113b15 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x2b, 0x31,
	0x18, 0xc5, 0x27, 0xbd, 0xbd, 0x15, 0xa3, 0x56, 0x08, 0x2e, 0xc6, 0x0a, 0xb1, 0x74, 0x35, 0xb8,
	0x48, 0xa8, 0x3e, 0x80, 0x30, 0x2a, 0x22, 0xb8, 0x28, 0x5d, 0xba, 0x4b, 0xd2, 0x8f, 0xe9, 0xe0,
	0xcc, 0xa4, 0x24, 0x69, 0xe9, 0xbc, 0x85, 0xaf, 0xe4, 0xae, 0xcb, 0x2e, 0x5d, 0x89, 0xb4, 0x2f,
	0x22, 0xf3, 0xa7, 0x6a, 0xbb, 0xea, 0x6e, 0xce, 0x77, 0xe6, 0xfc, 0x38, 0xe1, 0xe0, 0x0b, 0x29,
	0x64, 0xae, 0xc6, 0x22, 0xce, 0x78, 0x24, 0x52, 0xe0, 0x93, 0x44, 0xe4, 0x60, 0xd8, 0xc4, 0x68,
	0xa7, 0x49, 0xfb, 0xc7, 0x64, 0x85, 0xd9, 0x39, 0x8b, 0x74, 0xa4, 0x4b, 0x8b, 0x17, 0x5f, 0xd5,
	0x5f, 0x1d, 0xaa, 0xb4, 0x4d, 0xb5, 0xe5, 0x52, 0x58, 0xe0, 0xb3, 0xbe, 0x04, 0x27, 0xfa, 0x5c,
	0xe9, 0x38, 0xab, 0xfc, 0xde, 0x7b, 0x03, 0xb7, 0x06, 0x25, 0x96, 0xf8, 0xf8, 0x40, 0x8c, 0x46,
	0x06, 0xac, 0xf5, 0x51, 0x17, 0x05, 0x87, 0xc3, 0x8d, 0x24, 0x3d, 0x7c, 0xac, 0xa6, 0xc6, 0x40,
	0xe6, 0x9e, 0x61, 0x06, 0x89, 0xdf, 0xe8, 0xa2, 0xa0, 0x39, 0xdc, 0xba, 0x91, 0x07, 0x7c, 0x52,
	0xeb, 0xa7, 0x4c, 0xe9, 0x14, 0xfc, 0x7f, 0x5d, 0x14, 0x1c, 0x5d, 0x9f, 0xb3, 0xaa, 0x00, 0x2b,
	0x0a, 0xb0, 0xba, 0x00, 0xbb, 0xd3, 0x71, 0x16, 0x36, 0x17, 0x9f, 0x97, 0xde, 0x70, 0x3b, 0x45,
	0x6e, 0x31, 0xce, 0x60, 0xbe, 0x61, 0x34, 0xf7, 0x63, 0xfc, 0x89, 0x90, 0x47, 0xdc, 0x2e, 0x54,
	0x59, 0x6a, 0x60, 0x62, 0x05, 0xfe, 0xff, 0xfd, 0x20, 0x3b, 0x31, 0x12, 0xe0, 0xd3, 0x44, 0xd8,
	0x1a, 0x1b, 0x26, 0x5a, 0xbd, 0xfa, 0xad, 0xf2, 0xdd, 0xbb, 0xe7, 0xf0, 0x7e, 0xb1, 0xa2, 0x68,
	0xb9, 0xa2, 0xe8, 0x6b, 0x45, 0xd1, 0xdb, 0x9a, 0x7a, 0xcb, 0x35, 0xf5, 0x3e, 0xd6, 0xd4, 0x7b,
	0xb9, 0x8a, 0x62, 0x37, 0x9e, 0x4a, 0xa6, 0x74, 0xca, 0x43, 0x21, 0x65, 0x1e, 0xc6, 0x8e, 0xff,
	0x8e, 0x3a, 0xaf, 0x66, 0x75, 0xf9, 0x04, 0xac, 0x6c, 0x95, 0x83, 0xdc, 0x7c, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xe5, 0x15, 0x20, 0xe6, 0xf5, 0x01, 0x00, 0x00,
}

func (m *Player) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Player) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Player) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastIncomeBlock != 0 {
		i = encodeVarintPlayer(dAtA, i, uint64(m.LastIncomeBlock))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.NextLevelPrice.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlayer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.NextIncome.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlayer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.CurrentIncome.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlayer(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.CurrentLevel != 0 {
		i = encodeVarintPlayer(dAtA, i, uint64(m.CurrentLevel))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPlayer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPlayer(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlayer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Player) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPlayer(uint64(l))
	}
	if m.CurrentLevel != 0 {
		n += 1 + sovPlayer(uint64(m.CurrentLevel))
	}
	l = m.CurrentIncome.Size()
	n += 1 + l + sovPlayer(uint64(l))
	l = m.NextIncome.Size()
	n += 1 + l + sovPlayer(uint64(l))
	l = m.NextLevelPrice.Size()
	n += 1 + l + sovPlayer(uint64(l))
	if m.LastIncomeBlock != 0 {
		n += 1 + sovPlayer(uint64(m.LastIncomeBlock))
	}
	return n
}

func sovPlayer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlayer(x uint64) (n int) {
	return sovPlayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Player) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlayer
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
			return fmt.Errorf("proto: Player: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Player: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
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
				return ErrInvalidLengthPlayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentLevel", wireType)
			}
			m.CurrentLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentLevel |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentIncome", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
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
				return ErrInvalidLengthPlayer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentIncome.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextIncome", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
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
				return ErrInvalidLengthPlayer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NextIncome.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextLevelPrice", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
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
				return ErrInvalidLengthPlayer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NextLevelPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIncomeBlock", wireType)
			}
			m.LastIncomeBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIncomeBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPlayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlayer
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
func skipPlayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlayer
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
					return 0, ErrIntOverflowPlayer
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
					return 0, ErrIntOverflowPlayer
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
				return 0, ErrInvalidLengthPlayer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlayer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlayer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlayer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlayer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlayer = fmt.Errorf("proto: unexpected end of group")
)
