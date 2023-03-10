// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: newsletter/newsletter/newsletter.proto

package types

import (
	fmt "fmt"
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

type Newsletter struct {
	Index           string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Title           string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description     string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price           string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	SubscriberList  []string `protobuf:"bytes,5,rep,name=subscriberList,proto3" json:"subscriberList,omitempty"`
	Creator         string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	WithdrawCount   uint64   `protobuf:"varint,7,opt,name=withdrawCount,proto3" json:"withdrawCount,omitempty"`
	SubscriberCount uint64   `protobuf:"varint,8,opt,name=subscriberCount,proto3" json:"subscriberCount,omitempty"`
}

func (m *Newsletter) Reset()         { *m = Newsletter{} }
func (m *Newsletter) String() string { return proto.CompactTextString(m) }
func (*Newsletter) ProtoMessage()    {}
func (*Newsletter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d170dbf735559b8c, []int{0}
}
func (m *Newsletter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Newsletter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Newsletter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Newsletter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Newsletter.Merge(m, src)
}
func (m *Newsletter) XXX_Size() int {
	return m.Size()
}
func (m *Newsletter) XXX_DiscardUnknown() {
	xxx_messageInfo_Newsletter.DiscardUnknown(m)
}

var xxx_messageInfo_Newsletter proto.InternalMessageInfo

func (m *Newsletter) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Newsletter) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Newsletter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Newsletter) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Newsletter) GetSubscriberList() []string {
	if m != nil {
		return m.SubscriberList
	}
	return nil
}

func (m *Newsletter) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Newsletter) GetWithdrawCount() uint64 {
	if m != nil {
		return m.WithdrawCount
	}
	return 0
}

func (m *Newsletter) GetSubscriberCount() uint64 {
	if m != nil {
		return m.SubscriberCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Newsletter)(nil), "newsletter.newsletter.Newsletter")
}

func init() {
	proto.RegisterFile("newsletter/newsletter/newsletter.proto", fileDescriptor_d170dbf735559b8c)
}

var fileDescriptor_d170dbf735559b8c = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4b, 0x2d, 0x2f,
	0xce, 0x49, 0x2d, 0x29, 0x49, 0x2d, 0xd2, 0xc7, 0xca, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x45, 0x12, 0x41, 0x30, 0x95, 0x9a, 0x98, 0xb8, 0xb8, 0xfc, 0xe0, 0x5c, 0x21, 0x11, 0x2e,
	0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x07, 0x24,
	0x5a, 0x92, 0x59, 0x92, 0x93, 0x2a, 0xc1, 0x04, 0x11, 0x05, 0x73, 0x84, 0x14, 0xb8, 0xb8, 0x53,
	0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0xc1, 0x72, 0xc8, 0x42,
	0x20, 0x7d, 0x05, 0x45, 0x99, 0xc9, 0xa9, 0x12, 0x2c, 0x10, 0x7d, 0x60, 0x8e, 0x90, 0x1a, 0x17,
	0x5f, 0x71, 0x69, 0x12, 0x48, 0x55, 0x52, 0x6a, 0x91, 0x4f, 0x66, 0x71, 0x89, 0x04, 0xab, 0x02,
	0xb3, 0x06, 0x67, 0x10, 0x9a, 0xa8, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e,
	0x91, 0x04, 0x1b, 0x58, 0x3f, 0x8c, 0x2b, 0xa4, 0xc2, 0xc5, 0x5b, 0x9e, 0x59, 0x92, 0x91, 0x52,
	0x94, 0x58, 0xee, 0x9c, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0xae, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x2a,
	0x28, 0xa4, 0xc1, 0xc5, 0x8f, 0x30, 0x11, 0xa2, 0x8e, 0x03, 0xac, 0x0e, 0x5d, 0xd8, 0xc9, 0xfc,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x64, 0x91, 0x82, 0xb4, 0x02, 0x39,
	0x7c, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x61, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x87, 0x7f, 0x5f, 0x2b, 0x85, 0x01, 0x00, 0x00,
}

func (m *Newsletter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Newsletter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Newsletter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubscriberCount != 0 {
		i = encodeVarintNewsletter(dAtA, i, uint64(m.SubscriberCount))
		i--
		dAtA[i] = 0x40
	}
	if m.WithdrawCount != 0 {
		i = encodeVarintNewsletter(dAtA, i, uint64(m.WithdrawCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNewsletter(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SubscriberList) > 0 {
		for iNdEx := len(m.SubscriberList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SubscriberList[iNdEx])
			copy(dAtA[i:], m.SubscriberList[iNdEx])
			i = encodeVarintNewsletter(dAtA, i, uint64(len(m.SubscriberList[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintNewsletter(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintNewsletter(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintNewsletter(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNewsletter(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNewsletter(dAtA []byte, offset int, v uint64) int {
	offset -= sovNewsletter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Newsletter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNewsletter(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovNewsletter(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovNewsletter(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovNewsletter(uint64(l))
	}
	if len(m.SubscriberList) > 0 {
		for _, s := range m.SubscriberList {
			l = len(s)
			n += 1 + l + sovNewsletter(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNewsletter(uint64(l))
	}
	if m.WithdrawCount != 0 {
		n += 1 + sovNewsletter(uint64(m.WithdrawCount))
	}
	if m.SubscriberCount != 0 {
		n += 1 + sovNewsletter(uint64(m.SubscriberCount))
	}
	return n
}

func sovNewsletter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNewsletter(x uint64) (n int) {
	return sovNewsletter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Newsletter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNewsletter
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
			return fmt.Errorf("proto: Newsletter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Newsletter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
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
				return ErrInvalidLengthNewsletter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewsletter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
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
				return ErrInvalidLengthNewsletter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewsletter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
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
				return ErrInvalidLengthNewsletter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewsletter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
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
				return ErrInvalidLengthNewsletter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewsletter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriberList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
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
				return ErrInvalidLengthNewsletter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewsletter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubscriberList = append(m.SubscriberList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
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
				return ErrInvalidLengthNewsletter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewsletter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawCount", wireType)
			}
			m.WithdrawCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriberCount", wireType)
			}
			m.SubscriberCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewsletter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriberCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNewsletter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNewsletter
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
func skipNewsletter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNewsletter
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
					return 0, ErrIntOverflowNewsletter
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
					return 0, ErrIntOverflowNewsletter
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
				return 0, ErrInvalidLengthNewsletter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNewsletter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNewsletter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNewsletter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNewsletter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNewsletter = fmt.Errorf("proto: unexpected end of group")
)
