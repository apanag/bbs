// Code generated by protoc-gen-gogo.
// source: error.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Error_Type int32

const (
	Error_UnknownError                            Error_Type = 0
	Error_InvalidDomain                           Error_Type = 1
	Error_UnkownVersion                           Error_Type = 2
	Error_InvalidRecord                           Error_Type = 3
	Error_InvalidRequest                          Error_Type = 4
	Error_InvalidResponse                         Error_Type = 5
	Error_InvalidProtobufMessage                  Error_Type = 6
	Error_InvalidJSON                             Error_Type = 7
	Error_FailedToOpenEnvelope                    Error_Type = 8
	Error_InvalidStateTransition                  Error_Type = 9
	Error_Unauthorized                            Error_Type = 10
	Error_ResourceConflict                        Error_Type = 11
	Error_ResourceExists                          Error_Type = 12
	Error_ResourceNotFound                        Error_Type = 13
	Error_RouterError                             Error_Type = 14
	Error_ActualLRPCannotBeClaimed                Error_Type = 15
	Error_ActualLRPCannotBeStarted                Error_Type = 16
	Error_ActualLRPCannotBeCrashed                Error_Type = 17
	Error_ActualLRPCannotBeFailed                 Error_Type = 18
	Error_ActualLRPCannotBeRemoved                Error_Type = 19
	Error_ActualLRPCannotBeStopped                Error_Type = 20
	Error_ActualLRPCannotBeUnclaimed              Error_Type = 21
	Error_ActualLRPCannotBeEvacuated              Error_Type = 22
	Error_DesiredLRPCannotBeUpdated               Error_Type = 23
	Error_RunningOnDifferentCell                  Error_Type = 24
	Error_DesiredLRPSchedulingInfoCannotBeUpdated Error_Type = 25
	Error_GUIDGeneration                          Error_Type = 26
	Error_Deserialize                             Error_Type = 27
	Error_Deadlock                                Error_Type = 28
	Error_Unrecoverable                           Error_Type = 29
)

var Error_Type_name = map[int32]string{
	0:  "UnknownError",
	1:  "InvalidDomain",
	2:  "UnkownVersion",
	3:  "InvalidRecord",
	4:  "InvalidRequest",
	5:  "InvalidResponse",
	6:  "InvalidProtobufMessage",
	7:  "InvalidJSON",
	8:  "FailedToOpenEnvelope",
	9:  "InvalidStateTransition",
	10: "Unauthorized",
	11: "ResourceConflict",
	12: "ResourceExists",
	13: "ResourceNotFound",
	14: "RouterError",
	15: "ActualLRPCannotBeClaimed",
	16: "ActualLRPCannotBeStarted",
	17: "ActualLRPCannotBeCrashed",
	18: "ActualLRPCannotBeFailed",
	19: "ActualLRPCannotBeRemoved",
	20: "ActualLRPCannotBeStopped",
	21: "ActualLRPCannotBeUnclaimed",
	22: "ActualLRPCannotBeEvacuated",
	23: "DesiredLRPCannotBeUpdated",
	24: "RunningOnDifferentCell",
	25: "DesiredLRPSchedulingInfoCannotBeUpdated",
	26: "GUIDGeneration",
	27: "Deserialize",
	28: "Deadlock",
	29: "Unrecoverable",
}
var Error_Type_value = map[string]int32{
	"UnknownError":                            0,
	"InvalidDomain":                           1,
	"UnkownVersion":                           2,
	"InvalidRecord":                           3,
	"InvalidRequest":                          4,
	"InvalidResponse":                         5,
	"InvalidProtobufMessage":                  6,
	"InvalidJSON":                             7,
	"FailedToOpenEnvelope":                    8,
	"InvalidStateTransition":                  9,
	"Unauthorized":                            10,
	"ResourceConflict":                        11,
	"ResourceExists":                          12,
	"ResourceNotFound":                        13,
	"RouterError":                             14,
	"ActualLRPCannotBeClaimed":                15,
	"ActualLRPCannotBeStarted":                16,
	"ActualLRPCannotBeCrashed":                17,
	"ActualLRPCannotBeFailed":                 18,
	"ActualLRPCannotBeRemoved":                19,
	"ActualLRPCannotBeStopped":                20,
	"ActualLRPCannotBeUnclaimed":              21,
	"ActualLRPCannotBeEvacuated":              22,
	"DesiredLRPCannotBeUpdated":               23,
	"RunningOnDifferentCell":                  24,
	"DesiredLRPSchedulingInfoCannotBeUpdated": 25,
	"GUIDGeneration":                          26,
	"Deserialize":                             27,
	"Deadlock":                                28,
	"Unrecoverable":                           29,
}

func (x Error_Type) Enum() *Error_Type {
	p := new(Error_Type)
	*p = x
	return p
}
func (x Error_Type) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(Error_Type_name, int32(x))
}
func (x *Error_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_Type_value, data, "Error_Type")
	if err != nil {
		return err
	}
	*x = Error_Type(value)
	return nil
}
func (Error_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorError, []int{0, 0} }

type Error struct {
	Type    Error_Type `protobuf:"varint,1,opt,name=type,enum=models.Error_Type" json:"type"`
	Message string     `protobuf:"bytes,2,opt,name=message" json:"message"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorError, []int{0} }

func (m *Error) GetType() Error_Type {
	if m != nil {
		return m.Type
	}
	return Error_UnknownError
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "models.Error")
	proto.RegisterEnum("models.Error_Type", Error_Type_name, Error_Type_value)
}
func (x Error_Type) String() string {
	s, ok := Error_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Error) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&models.Error{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringError(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Error) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Error) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintError(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x12
	i++
	i = encodeVarintError(dAtA, i, uint64(len(m.Message)))
	i += copy(dAtA[i:], m.Message)
	return i, nil
}

func encodeFixed64Error(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Error(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintError(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Error) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovError(uint64(m.Type))
	l = len(m.Message)
	n += 1 + l + sovError(uint64(l))
	return n
}

func sovError(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozError(x uint64) (n int) {
	return sovError(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Error{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringError(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Error) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowError
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
			return fmt.Errorf("proto: Error: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Error: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Error_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowError
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthError
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipError(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthError
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
func skipError(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
					return 0, ErrIntOverflowError
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
				return 0, ErrInvalidLengthError
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowError
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
				next, err := skipError(dAtA[start:])
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
	ErrInvalidLengthError = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowError   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("error.proto", fileDescriptorError) }

var fileDescriptorError = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x52, 0xdb, 0x3e,
	0x14, 0xc5, 0x63, 0xfe, 0xe1, 0x4b, 0x7c, 0x09, 0xc1, 0x1f, 0x42, 0x00, 0x97, 0x61, 0x53, 0x66,
	0x4a, 0xc3, 0x4c, 0xdf, 0xa0, 0x24, 0x81, 0xa1, 0xd3, 0x02, 0xe3, 0x90, 0xee, 0x85, 0x75, 0x93,
	0x68, 0xa2, 0xe8, 0xba, 0x92, 0x1c, 0x0a, 0xab, 0x3e, 0x42, 0x1f, 0xa3, 0xaf, 0xd0, 0x37, 0x60,
	0xc9, 0xb2, 0xab, 0x4e, 0xe3, 0x6e, 0xba, 0xe4, 0x11, 0x3a, 0xb6, 0x03, 0xcd, 0x94, 0x74, 0x67,
	0x9d, 0xdf, 0xbd, 0xc7, 0x47, 0xf7, 0x8a, 0xcc, 0x81, 0x31, 0x68, 0x2a, 0x91, 0x41, 0x87, 0x6c,
	0xaa, 0x87, 0x02, 0x94, 0x2d, 0xbf, 0x6c, 0x4b, 0xd7, 0x89, 0x2f, 0x2b, 0x21, 0xf6, 0x0e, 0xda,
	0xd8, 0xc6, 0x83, 0x0c, 0x5f, 0xc6, 0xad, 0xec, 0x94, 0x1d, 0xb2, 0xaf, 0xbc, 0x6d, 0xf7, 0xeb,
	0x14, 0x99, 0xac, 0xa7, 0x36, 0x6c, 0x9f, 0x14, 0xdd, 0x75, 0x04, 0x25, 0x6f, 0xc7, 0xdb, 0x5b,
	0x7c, 0xc5, 0x2a, 0xb9, 0x5f, 0x25, 0x83, 0x95, 0x8b, 0xeb, 0x08, 0x0e, 0x8b, 0xb7, 0xdf, 0x9f,
	0x15, 0x82, 0xac, 0x8a, 0xf9, 0x64, 0xba, 0x07, 0xd6, 0xf2, 0x36, 0x94, 0x26, 0x76, 0xbc, 0xbd,
	0xd9, 0x21, 0x7c, 0x10, 0x77, 0x07, 0x93, 0xa4, 0x98, 0x36, 0x31, 0x4a, 0xe6, 0x9b, 0xba, 0xab,
	0xf1, 0x4a, 0x67, 0x4e, 0xb4, 0xc0, 0x96, 0xc9, 0xc2, 0x89, 0xee, 0x73, 0x25, 0x45, 0x0d, 0x7b,
	0x5c, 0x6a, 0xea, 0xa5, 0x52, 0x53, 0x77, 0xf1, 0x4a, 0xbf, 0x07, 0x63, 0x25, 0x6a, 0x3a, 0x31,
	0x52, 0x15, 0x40, 0x88, 0x46, 0xd0, 0xff, 0x18, 0x23, 0x8b, 0x8f, 0xd2, 0x87, 0x18, 0xac, 0xa3,
	0x45, 0xb6, 0x42, 0x96, 0x1e, 0x35, 0x1b, 0xa1, 0xb6, 0x40, 0x27, 0x59, 0x99, 0xac, 0x0d, 0xc5,
	0xf3, 0xe1, 0xe5, 0xdf, 0xe5, 0xb1, 0xe8, 0x14, 0x5b, 0x22, 0x73, 0x43, 0xf6, 0xa6, 0x71, 0x76,
	0x4a, 0xa7, 0x59, 0x89, 0xac, 0x1e, 0x71, 0xa9, 0x40, 0x5c, 0xe0, 0x59, 0x04, 0xba, 0xae, 0xfb,
	0xa0, 0x30, 0x02, 0x3a, 0x33, 0x62, 0xd3, 0x70, 0xdc, 0xc1, 0x85, 0xe1, 0xda, 0x4a, 0x97, 0xc6,
	0x9b, 0xcd, 0xaf, 0xc5, 0x63, 0xd7, 0x41, 0x23, 0x6f, 0x40, 0x50, 0xc2, 0x56, 0x09, 0x0d, 0xc0,
	0x62, 0x6c, 0x42, 0xa8, 0xa2, 0x6e, 0x29, 0x19, 0x3a, 0x3a, 0x97, 0x66, 0x7e, 0x50, 0xeb, 0x1f,
	0xa5, 0x75, 0x96, 0xce, 0x8f, 0x56, 0x9e, 0xa2, 0x3b, 0xc2, 0x58, 0x0b, 0xba, 0x90, 0x06, 0x0b,
	0x30, 0x76, 0x60, 0xf2, 0x39, 0x2d, 0xb2, 0x2d, 0x52, 0x7a, 0x1d, 0xba, 0x98, 0xab, 0xb7, 0xc1,
	0x79, 0x95, 0x6b, 0x8d, 0xee, 0x10, 0xaa, 0x8a, 0xcb, 0x1e, 0x08, 0xba, 0x34, 0x96, 0x36, 0x1c,
	0x37, 0x0e, 0x04, 0xa5, 0xe3, 0x7b, 0x0d, 0xb7, 0x1d, 0x10, 0x74, 0x99, 0x6d, 0x92, 0xf5, 0x27,
	0x34, 0x9f, 0x01, 0x65, 0x63, 0x5b, 0x03, 0xe8, 0x61, 0x1f, 0x04, 0x5d, 0xf9, 0xc7, 0x6f, 0x31,
	0x8a, 0x40, 0xd0, 0x55, 0xe6, 0x93, 0xf2, 0x13, 0xda, 0xd4, 0xe1, 0x30, 0xf4, 0xff, 0x63, 0x79,
	0xbd, 0xcf, 0xc3, 0x98, 0xa7, 0xb1, 0xd7, 0xd8, 0x36, 0xd9, 0xa8, 0x81, 0x95, 0x06, 0xc4, 0xa8,
	0x41, 0x24, 0x32, 0xbc, 0x9e, 0x2e, 0x24, 0x88, 0xb5, 0x96, 0xba, 0x7d, 0xa6, 0x6b, 0xb2, 0xd5,
	0x02, 0x03, 0xda, 0x55, 0x41, 0x29, 0x5a, 0x62, 0x2f, 0xc8, 0xf3, 0x3f, 0xad, 0x8d, 0xb0, 0x03,
	0x22, 0x56, 0x52, 0xb7, 0x4f, 0x74, 0x0b, 0xff, 0x36, 0xda, 0x48, 0xb7, 0x72, 0xdc, 0x3c, 0xa9,
	0x1d, 0x83, 0x06, 0xc3, 0xb3, 0x8d, 0x96, 0xd3, 0xf9, 0xd7, 0xc0, 0x82, 0x91, 0x5c, 0xc9, 0x1b,
	0xa0, 0x9b, 0x6c, 0x9e, 0xcc, 0xd4, 0x80, 0x0b, 0x85, 0x61, 0x97, 0x6e, 0xe5, 0x4f, 0xd4, 0x40,
	0x88, 0x7d, 0x30, 0xfc, 0x52, 0x01, 0xdd, 0x3e, 0xdc, 0xbf, 0x1b, 0xf8, 0xde, 0xb7, 0x81, 0x5f,
	0xb8, 0x1f, 0xf8, 0xde, 0xa7, 0xc4, 0xf7, 0xbe, 0x24, 0x7e, 0xe1, 0x36, 0xf1, 0xbd, 0xbb, 0xc4,
	0xf7, 0x7e, 0x24, 0xbe, 0xf7, 0x2b, 0xf1, 0x0b, 0xf7, 0x89, 0xef, 0x7d, 0xfe, 0xe9, 0x17, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x98, 0xc4, 0x03, 0xae, 0x03, 0x00, 0x00,
}
