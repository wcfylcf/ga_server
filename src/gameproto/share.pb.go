// Code generated by protoc-gen-gogo.
// source: share.proto
// DO NOT EDIT!

package gameproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 错误类型
type ErrorCode int32

const (
	OK            ErrorCode = 0
	Fail          ErrorCode = 1
	Error         ErrorCode = 2
	ServerFull    ErrorCode = 3
	KeyError      ErrorCode = 4
	NoFoundTarget ErrorCode = 5
	// old code
	IMPORTANT_WRONG_HEAD      ErrorCode = -1000
	RESOURCE_VITALITY_ERROR   ErrorCode = 1002
	RESOURCE_GOLD_ERROR       ErrorCode = 1003
	RESOURCE_RMB_ERROR        ErrorCode = 1004
	GUILD_EXIT_CHAIRMAN_ERROR ErrorCode = 1022
	UNKNOWN_ERROR             ErrorCode = -9999
)

var ErrorCode_name = map[int32]string{
	0:     "OK",
	1:     "Fail",
	2:     "Error",
	3:     "ServerFull",
	4:     "KeyError",
	5:     "NoFoundTarget",
	-1000: "IMPORTANT_WRONG_HEAD",
	1002:  "RESOURCE_VITALITY_ERROR",
	1003:  "RESOURCE_GOLD_ERROR",
	1004:  "RESOURCE_RMB_ERROR",
	1022:  "GUILD_EXIT_CHAIRMAN_ERROR",
	-9999: "UNKNOWN_ERROR",
}
var ErrorCode_value = map[string]int32{
	"OK":                        0,
	"Fail":                      1,
	"Error":                     2,
	"ServerFull":                3,
	"KeyError":                  4,
	"NoFoundTarget":             5,
	"IMPORTANT_WRONG_HEAD":      -1000,
	"RESOURCE_VITALITY_ERROR":   1002,
	"RESOURCE_GOLD_ERROR":       1003,
	"RESOURCE_RMB_ERROR":        1004,
	"GUILD_EXIT_CHAIRMAN_ERROR": 1022,
	"UNKNOWN_ERROR":             -9999,
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorShare, []int{0} }

type C2S_TestMsg struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *C2S_TestMsg) Reset()                    { *m = C2S_TestMsg{} }
func (*C2S_TestMsg) ProtoMessage()               {}
func (*C2S_TestMsg) Descriptor() ([]byte, []int) { return fileDescriptorShare, []int{0} }

func (m *C2S_TestMsg) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type S2C_TestMsg struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *S2C_TestMsg) Reset()                    { *m = S2C_TestMsg{} }
func (*S2C_TestMsg) ProtoMessage()               {}
func (*S2C_TestMsg) Descriptor() ([]byte, []int) { return fileDescriptorShare, []int{1} }

func (m *S2C_TestMsg) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type S2C_ConfirmInfo struct {
	MsgHead int32 `protobuf:"varint,1,opt,name=msgHead,proto3" json:"msgHead,omitempty"`
	Code    int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *S2C_ConfirmInfo) Reset()                    { *m = S2C_ConfirmInfo{} }
func (*S2C_ConfirmInfo) ProtoMessage()               {}
func (*S2C_ConfirmInfo) Descriptor() ([]byte, []int) { return fileDescriptorShare, []int{2} }

func (m *S2C_ConfirmInfo) GetMsgHead() int32 {
	if m != nil {
		return m.MsgHead
	}
	return 0
}

func (m *S2C_ConfirmInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*C2S_TestMsg)(nil), "gameproto.C2S_TestMsg")
	proto.RegisterType((*S2C_TestMsg)(nil), "gameproto.S2C_TestMsg")
	proto.RegisterType((*S2C_ConfirmInfo)(nil), "gameproto.S2C_ConfirmInfo")
	proto.RegisterEnum("gameproto.ErrorCode", ErrorCode_name, ErrorCode_value)
}
func (x ErrorCode) String() string {
	s, ok := ErrorCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *C2S_TestMsg) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*C2S_TestMsg)
	if !ok {
		that2, ok := that.(C2S_TestMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	return true
}
func (this *S2C_TestMsg) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*S2C_TestMsg)
	if !ok {
		that2, ok := that.(S2C_TestMsg)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	return true
}
func (this *S2C_ConfirmInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*S2C_ConfirmInfo)
	if !ok {
		that2, ok := that.(S2C_ConfirmInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.MsgHead != that1.MsgHead {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	return true
}
func (this *C2S_TestMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&gameproto.C2S_TestMsg{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *S2C_TestMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&gameproto.S2C_TestMsg{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *S2C_ConfirmInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&gameproto.S2C_ConfirmInfo{")
	s = append(s, "MsgHead: "+fmt.Sprintf("%#v", this.MsgHead)+",\n")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringShare(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *C2S_TestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2S_TestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShare(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *S2C_TestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_TestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShare(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *S2C_ConfirmInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_ConfirmInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MsgHead != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShare(dAtA, i, uint64(m.MsgHead))
	}
	if m.Code != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShare(dAtA, i, uint64(m.Code))
	}
	return i, nil
}

func encodeFixed64Share(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Share(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintShare(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *C2S_TestMsg) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShare(uint64(m.Id))
	}
	return n
}

func (m *S2C_TestMsg) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovShare(uint64(m.Id))
	}
	return n
}

func (m *S2C_ConfirmInfo) Size() (n int) {
	var l int
	_ = l
	if m.MsgHead != 0 {
		n += 1 + sovShare(uint64(m.MsgHead))
	}
	if m.Code != 0 {
		n += 1 + sovShare(uint64(m.Code))
	}
	return n
}

func sovShare(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozShare(x uint64) (n int) {
	return sovShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *C2S_TestMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&C2S_TestMsg{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *S2C_TestMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&S2C_TestMsg{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`}`,
	}, "")
	return s
}
func (this *S2C_ConfirmInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&S2C_ConfirmInfo{`,
		`MsgHead:` + fmt.Sprintf("%v", this.MsgHead) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringShare(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *C2S_TestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: C2S_TestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2S_TestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShare
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
func (m *S2C_TestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: S2C_TestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_TestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShare
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
func (m *S2C_ConfirmInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: S2C_ConfirmInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_ConfirmInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgHead", wireType)
			}
			m.MsgHead = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgHead |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShare
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
func skipShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShare
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthShare
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShare
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipShare(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthShare = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShare   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("share.proto", fileDescriptorShare) }

var fileDescriptorShare = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0x3f, 0xcf, 0xd2, 0x40,
	0x1c, 0x07, 0xf0, 0x5e, 0x85, 0xa7, 0x0f, 0xbf, 0x47, 0xb0, 0xfe, 0x34, 0xa1, 0x1a, 0xbd, 0xa8,
	0x93, 0x31, 0xc6, 0x01, 0x5f, 0x80, 0x29, 0xa5, 0x40, 0x03, 0x6d, 0xcd, 0xb5, 0x88, 0x4e, 0x4d,
	0xb5, 0x47, 0x6d, 0x02, 0x9c, 0xb9, 0x82, 0x89, 0x9b, 0xab, 0x9b, 0xa3, 0x2f, 0xc1, 0xd5, 0x77,
	0xe1, 0xc8, 0xe8, 0x28, 0x75, 0x31, 0xea, 0x80, 0x9b, 0x93, 0x1a, 0xca, 0x9f, 0xed, 0xb9, 0xe9,
	0x7b, 0xdf, 0xcf, 0x77, 0xb8, 0xe4, 0xe0, 0x2c, 0x7f, 0x19, 0x4b, 0xfe, 0xe0, 0x95, 0x14, 0x0b,
	0x81, 0xb5, 0x34, 0x9e, 0xf1, 0x32, 0xde, 0xb9, 0x09, 0x67, 0x56, 0x2b, 0x88, 0x42, 0x9e, 0x2f,
	0xdc, 0x3c, 0xc5, 0x06, 0xa8, 0x59, 0x62, 0x90, 0x5b, 0xe4, 0x6e, 0x9d, 0xa9, 0x59, 0xb2, 0xe5,
	0xa0, 0x65, 0x9d, 0xcb, 0x8f, 0xe0, 0xd2, 0x96, 0x2d, 0x31, 0x9f, 0x64, 0x72, 0xe6, 0xcc, 0x27,
	0x02, 0x0d, 0xd0, 0x66, 0x79, 0xda, 0xe7, 0xf1, 0x6e, 0x57, 0x65, 0x87, 0x2b, 0x22, 0x54, 0x5e,
	0x88, 0x84, 0x1b, 0x6a, 0x59, 0x97, 0xf9, 0xde, 0x3b, 0x15, 0x6a, 0xb6, 0x94, 0x42, 0x5a, 0x22,
	0xe1, 0x78, 0x02, 0xaa, 0x3f, 0xd0, 0x15, 0x3c, 0x85, 0x4a, 0x37, 0xce, 0xa6, 0x3a, 0xc1, 0x1a,
	0x54, 0x4b, 0xd6, 0x55, 0x6c, 0x00, 0x04, 0x5c, 0xbe, 0xe6, 0xb2, 0xbb, 0x9c, 0x4e, 0xf5, 0x0b,
	0x78, 0x11, 0x4e, 0x07, 0xfc, 0xcd, 0x4e, 0x2b, 0x78, 0x19, 0xea, 0x9e, 0xe8, 0x8a, 0xe5, 0x3c,
	0x09, 0x63, 0x99, 0xf2, 0x85, 0x5e, 0xc5, 0xdb, 0x70, 0xd5, 0x71, 0x1f, 0xfb, 0x2c, 0x34, 0xbd,
	0x30, 0x1a, 0x33, 0xdf, 0xeb, 0x45, 0x7d, 0xdb, 0xec, 0xe8, 0x1f, 0xfe, 0xfc, 0xdb, 0x1d, 0x82,
	0x37, 0xa0, 0xc9, 0xec, 0xc0, 0x1f, 0x31, 0xcb, 0x8e, 0x9e, 0x38, 0xa1, 0x39, 0x74, 0xc2, 0x67,
	0x91, 0xcd, 0x98, 0xcf, 0xf4, 0x1f, 0x1a, 0x1a, 0x70, 0xe5, 0xa8, 0x3d, 0x7f, 0xd8, 0xd9, 0xcb,
	0x4f, 0x0d, 0x9b, 0x80, 0x47, 0x61, 0x6e, 0x7b, 0x0f, 0xbf, 0x34, 0xa4, 0x70, 0xad, 0x37, 0x72,
	0xb6, 0xd3, 0xa7, 0x4e, 0x18, 0x59, 0x7d, 0xd3, 0x61, 0xae, 0xe9, 0xed, 0xfd, 0xaf, 0x86, 0xd7,
	0xa1, 0x3e, 0xf2, 0x06, 0x9e, 0x3f, 0x3e, 0x74, 0xbf, 0x3f, 0x1d, 0x1e, 0xd3, 0xbe, 0xbf, 0x5a,
	0x53, 0xe5, 0xcb, 0x9a, 0x2a, 0x9b, 0x35, 0x25, 0x6f, 0x0b, 0x4a, 0x3e, 0x16, 0x94, 0x7c, 0x2e,
	0x28, 0x59, 0x15, 0x94, 0x7c, 0x2d, 0x28, 0xf9, 0x5e, 0x50, 0x65, 0x53, 0x50, 0xf2, 0xfe, 0x1b,
	0x55, 0x9e, 0x9f, 0x94, 0xff, 0xf7, 0xf0, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xa3, 0x96,
	0xfe, 0xd9, 0x01, 0x00, 0x00,
}
