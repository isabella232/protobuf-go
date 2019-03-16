// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3/fields.proto

package proto3

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

type FieldTestMessage_Enum int32

const (
	FieldTestMessage_ZERO FieldTestMessage_Enum = 0
)

func (e FieldTestMessage_Enum) Type() protoreflect.EnumType {
	return xxx_File_proto3_fields_proto_enumTypes[0]
}
func (e FieldTestMessage_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

// Deprecated: Use FieldTestMessage_Enum.Type.Values instead.
var FieldTestMessage_Enum_name = map[int32]string{
	0: "ZERO",
}

// Deprecated: Use FieldTestMessage_Enum.Type.Values instead.
var FieldTestMessage_Enum_value = map[string]int32{
	"ZERO": 0,
}

func (x FieldTestMessage_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

// Deprecated: Use FieldTestMessage_Enum.Type instead.
func (FieldTestMessage_Enum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_proto3_fields_proto_rawdesc_gzipped, []int{0, 0}
}

type FieldTestMessage struct {
	OptionalBool         string                               `protobuf:"bytes,1,opt,name=optional_bool,json=optionalBool,proto3" json:"optional_bool,omitempty"`
	OptionalEnum         FieldTestMessage_Enum                `protobuf:"varint,2,opt,name=optional_enum,json=optionalEnum,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum" json:"optional_enum,omitempty"`
	OptionalInt32        int32                                `protobuf:"varint,3,opt,name=optional_int32,json=optionalInt32,proto3" json:"optional_int32,omitempty"`
	OptionalSint32       int32                                `protobuf:"zigzag32,4,opt,name=optional_sint32,json=optionalSint32,proto3" json:"optional_sint32,omitempty"`
	OptionalUint32       uint32                               `protobuf:"varint,5,opt,name=optional_uint32,json=optionalUint32,proto3" json:"optional_uint32,omitempty"`
	OptionalInt64        int64                                `protobuf:"varint,6,opt,name=optional_int64,json=optionalInt64,proto3" json:"optional_int64,omitempty"`
	OptionalSint64       int64                                `protobuf:"zigzag64,7,opt,name=optional_sint64,json=optionalSint64,proto3" json:"optional_sint64,omitempty"`
	OptionalUint64       uint64                               `protobuf:"varint,8,opt,name=optional_uint64,json=optionalUint64,proto3" json:"optional_uint64,omitempty"`
	OptionalSfixed32     int32                                `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32,proto3" json:"optional_sfixed32,omitempty"`
	OptionalFixed32      uint32                               `protobuf:"fixed32,10,opt,name=optional_fixed32,json=optionalFixed32,proto3" json:"optional_fixed32,omitempty"`
	OptionalFloat        float32                              `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat,proto3" json:"optional_float,omitempty"`
	OptionalSfixed64     int64                                `protobuf:"fixed64,12,opt,name=optional_sfixed64,json=optionalSfixed64,proto3" json:"optional_sfixed64,omitempty"`
	OptionalFixed64      uint64                               `protobuf:"fixed64,13,opt,name=optional_fixed64,json=optionalFixed64,proto3" json:"optional_fixed64,omitempty"`
	OptionalDouble       float64                              `protobuf:"fixed64,14,opt,name=optional_double,json=optionalDouble,proto3" json:"optional_double,omitempty"`
	OptionalString       string                               `protobuf:"bytes,15,opt,name=optional_string,json=optionalString,proto3" json:"optional_string,omitempty"`
	OptionalBytes        []byte                               `protobuf:"bytes,16,opt,name=optional_bytes,json=optionalBytes,proto3" json:"optional_bytes,omitempty"`
	Optional_Message     *FieldTestMessage_Message            `protobuf:"bytes,17,opt,name=optional_Message,json=optionalMessage,proto3" json:"optional_Message,omitempty"`
	RepeatedBool         []bool                               `protobuf:"varint,201,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedEnum         []FieldTestMessage_Enum              `protobuf:"varint,202,rep,packed,name=repeated_enum,json=repeatedEnum,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum" json:"repeated_enum,omitempty"`
	RepeatedInt32        []int32                              `protobuf:"varint,203,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedSint32       []int32                              `protobuf:"zigzag32,204,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedUint32       []uint32                             `protobuf:"varint,205,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedInt64        []int64                              `protobuf:"varint,206,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedSint64       []int64                              `protobuf:"zigzag64,207,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedUint64       []uint64                             `protobuf:"varint,208,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSfixed32     []int32                              `protobuf:"fixed32,209,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedFixed32      []uint32                             `protobuf:"fixed32,210,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFloat        []float32                            `protobuf:"fixed32,211,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedSfixed64     []int64                              `protobuf:"fixed64,212,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedFixed64      []uint64                             `protobuf:"fixed64,213,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedDouble       []float64                            `protobuf:"fixed64,214,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedString       []string                             `protobuf:"bytes,215,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes        [][]byte                             `protobuf:"bytes,216,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	Repeated_Message     []*FieldTestMessage_Message          `protobuf:"bytes,217,rep,name=repeated_Message,json=repeatedMessage,proto3" json:"repeated_Message,omitempty"`
	MapInt32Int64        map[int32]int64                      `protobuf:"bytes,500,rep,name=map_int32_int64,json=mapInt32Int64,proto3" json:"map_int32_int64,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapStringMessage     map[string]*FieldTestMessage_Message `protobuf:"bytes,501,rep,name=map_string_message,json=mapStringMessage,proto3" json:"map_string_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapFixed64Enum       map[uint64]FieldTestMessage_Enum     `protobuf:"bytes,502,rep,name=map_fixed64_enum,json=mapFixed64Enum,proto3" json:"map_fixed64_enum,omitempty" protobuf_key:"fixed64,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=goproto.protoc.proto3.FieldTestMessage_Enum"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *FieldTestMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_proto3_fields_proto_messageTypes[0].MessageOf(m)
}
func (m *FieldTestMessage) Reset()         { *m = FieldTestMessage{} }
func (m *FieldTestMessage) String() string { return proto.CompactTextString(m) }
func (*FieldTestMessage) ProtoMessage()    {}

// Deprecated: Use FieldTestMessage.ProtoReflect.Type instead.
func (*FieldTestMessage) Descriptor() ([]byte, []int) {
	return xxx_File_proto3_fields_proto_rawdesc_gzipped, []int{0}
}

func (m *FieldTestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldTestMessage.Unmarshal(m, b)
}
func (m *FieldTestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldTestMessage.Marshal(b, m, deterministic)
}
func (m *FieldTestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldTestMessage.Merge(m, src)
}
func (m *FieldTestMessage) XXX_Size() int {
	return xxx_messageInfo_FieldTestMessage.Size(m)
}
func (m *FieldTestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldTestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FieldTestMessage proto.InternalMessageInfo

func (m *FieldTestMessage) GetOptionalBool() string {
	if m != nil {
		return m.OptionalBool
	}
	return ""
}

func (m *FieldTestMessage) GetOptionalEnum() FieldTestMessage_Enum {
	if m != nil {
		return m.OptionalEnum
	}
	return FieldTestMessage_ZERO
}

func (m *FieldTestMessage) GetOptionalInt32() int32 {
	if m != nil {
		return m.OptionalInt32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSint32() int32 {
	if m != nil {
		return m.OptionalSint32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalUint32() uint32 {
	if m != nil {
		return m.OptionalUint32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalInt64() int64 {
	if m != nil {
		return m.OptionalInt64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSint64() int64 {
	if m != nil {
		return m.OptionalSint64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalUint64() uint64 {
	if m != nil {
		return m.OptionalUint64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSfixed32() int32 {
	if m != nil {
		return m.OptionalSfixed32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalFixed32() uint32 {
	if m != nil {
		return m.OptionalFixed32
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalFloat() float32 {
	if m != nil {
		return m.OptionalFloat
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalSfixed64() int64 {
	if m != nil {
		return m.OptionalSfixed64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalFixed64() uint64 {
	if m != nil {
		return m.OptionalFixed64
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalDouble() float64 {
	if m != nil {
		return m.OptionalDouble
	}
	return 0
}

func (m *FieldTestMessage) GetOptionalString() string {
	if m != nil {
		return m.OptionalString
	}
	return ""
}

func (m *FieldTestMessage) GetOptionalBytes() []byte {
	if m != nil {
		return m.OptionalBytes
	}
	return nil
}

func (m *FieldTestMessage) GetOptional_Message() *FieldTestMessage_Message {
	if m != nil {
		return m.Optional_Message
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedEnum() []FieldTestMessage_Enum {
	if m != nil {
		return m.RepeatedEnum
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSint32() []int32 {
	if m != nil {
		return m.RepeatedSint32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSint64() []int64 {
	if m != nil {
		return m.RepeatedSint64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *FieldTestMessage) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

func (m *FieldTestMessage) GetRepeated_Message() []*FieldTestMessage_Message {
	if m != nil {
		return m.Repeated_Message
	}
	return nil
}

func (m *FieldTestMessage) GetMapInt32Int64() map[int32]int64 {
	if m != nil {
		return m.MapInt32Int64
	}
	return nil
}

func (m *FieldTestMessage) GetMapStringMessage() map[string]*FieldTestMessage_Message {
	if m != nil {
		return m.MapStringMessage
	}
	return nil
}

func (m *FieldTestMessage) GetMapFixed64Enum() map[uint64]FieldTestMessage_Enum {
	if m != nil {
		return m.MapFixed64Enum
	}
	return nil
}

type FieldTestMessage_Message struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldTestMessage_Message) ProtoReflect() protoreflect.Message {
	return xxx_File_proto3_fields_proto_messageTypes[4].MessageOf(m)
}
func (m *FieldTestMessage_Message) Reset()         { *m = FieldTestMessage_Message{} }
func (m *FieldTestMessage_Message) String() string { return proto.CompactTextString(m) }
func (*FieldTestMessage_Message) ProtoMessage()    {}

// Deprecated: Use FieldTestMessage_Message.ProtoReflect.Type instead.
func (*FieldTestMessage_Message) Descriptor() ([]byte, []int) {
	return xxx_File_proto3_fields_proto_rawdesc_gzipped, []int{0, 3}
}

func (m *FieldTestMessage_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldTestMessage_Message.Unmarshal(m, b)
}
func (m *FieldTestMessage_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldTestMessage_Message.Marshal(b, m, deterministic)
}
func (m *FieldTestMessage_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldTestMessage_Message.Merge(m, src)
}
func (m *FieldTestMessage_Message) XXX_Size() int {
	return xxx_messageInfo_FieldTestMessage_Message.Size(m)
}
func (m *FieldTestMessage_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldTestMessage_Message.DiscardUnknown(m)
}

var xxx_messageInfo_FieldTestMessage_Message proto.InternalMessageInfo

func init() {
	proto.RegisterFile("proto3/fields.proto", xxx_File_proto3_fields_proto_rawdesc_gzipped)
	proto.RegisterEnum("goproto.protoc.proto3.FieldTestMessage_Enum", FieldTestMessage_Enum_name, FieldTestMessage_Enum_value)
	proto.RegisterType((*FieldTestMessage)(nil), "goproto.protoc.proto3.FieldTestMessage")
	proto.RegisterMapType((map[uint64]FieldTestMessage_Enum)(nil), "goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry")
	proto.RegisterMapType((map[int32]int64)(nil), "goproto.protoc.proto3.FieldTestMessage.MapInt32Int64Entry")
	proto.RegisterMapType((map[string]*FieldTestMessage_Message)(nil), "goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry")
	proto.RegisterType((*FieldTestMessage_Message)(nil), "goproto.protoc.proto3.FieldTestMessage.Message")
}

var xxx_File_proto3_fields_proto_rawdesc = []byte{
	// 2378 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x22, 0xd0, 0x11, 0x0a,
	0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x51, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x2b, 0x0a, 0x11, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x07, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x06, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x10, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0xc9, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52,
	0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x52, 0x0a,
	0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0xca,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0xcb, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0xcc, 0x01, 0x20,
	0x03, 0x28, 0x11, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0xcd, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x26, 0x0a,
	0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18,
	0xce, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0xcf, 0x01, 0x20, 0x03, 0x28, 0x12, 0x52,
	0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x18, 0xd0, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0xd1,
	0x01, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0xd2, 0x01, 0x20, 0x03,
	0x28, 0x07, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0xd3, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0d, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0xd4, 0x01, 0x20, 0x03, 0x28, 0x10, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0xd5, 0x01,
	0x20, 0x03, 0x28, 0x06, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0xd6, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0xd7, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0xd8, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x5b, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xd9, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x63,
	0x0a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0xf4, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x6c, 0x0a, 0x12, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xf5, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x10, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x66, 0x0a, 0x10, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0xf6, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x45, 0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x6d, 0x61, 0x70, 0x46, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x40, 0x0a, 0x12, 0x4d, 0x61, 0x70,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x74, 0x0a, 0x15, 0x4d,
	0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x6f, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x45,
	0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a,
	0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_proto3_fields_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_proto3_fields_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_proto3_fields_proto protoreflect.FileDescriptor

var xxx_File_proto3_fields_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_proto3_fields_proto_messageTypes = make([]protoimpl.MessageType, 5)
var xxx_File_proto3_fields_proto_goTypes = []interface{}{
	(FieldTestMessage_Enum)(0),       // 0: goproto.protoc.proto3.FieldTestMessage.Enum
	(*FieldTestMessage)(nil),         // 1: goproto.protoc.proto3.FieldTestMessage
	nil,                              // 2: goproto.protoc.proto3.FieldTestMessage.MapInt32Int64Entry
	nil,                              // 3: goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry
	nil,                              // 4: goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry
	(*FieldTestMessage_Message)(nil), // 5: goproto.protoc.proto3.FieldTestMessage.Message
}
var xxx_File_proto3_fields_proto_depIdxs = []int32{
	0, // goproto.protoc.proto3.FieldTestMessage.optional_enum:type_name -> goproto.protoc.proto3.FieldTestMessage.Enum
	5, // goproto.protoc.proto3.FieldTestMessage.optional_Message:type_name -> goproto.protoc.proto3.FieldTestMessage.Message
	0, // goproto.protoc.proto3.FieldTestMessage.repeated_enum:type_name -> goproto.protoc.proto3.FieldTestMessage.Enum
	5, // goproto.protoc.proto3.FieldTestMessage.repeated_Message:type_name -> goproto.protoc.proto3.FieldTestMessage.Message
	2, // goproto.protoc.proto3.FieldTestMessage.map_int32_int64:type_name -> goproto.protoc.proto3.FieldTestMessage.MapInt32Int64Entry
	3, // goproto.protoc.proto3.FieldTestMessage.map_string_message:type_name -> goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry
	4, // goproto.protoc.proto3.FieldTestMessage.map_fixed64_enum:type_name -> goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry
	5, // goproto.protoc.proto3.FieldTestMessage.MapStringMessageEntry.value:type_name -> goproto.protoc.proto3.FieldTestMessage.Message
	0, // goproto.protoc.proto3.FieldTestMessage.MapFixed64EnumEntry.value:type_name -> goproto.protoc.proto3.FieldTestMessage.Enum
}

func init() { xxx_File_proto3_fields_proto_init() }
func xxx_File_proto3_fields_proto_init() {
	if File_proto3_fields_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 5)
	File_proto3_fields_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_proto3_fields_proto_rawdesc,
		GoTypes:            xxx_File_proto3_fields_proto_goTypes,
		DependencyIndexes:  xxx_File_proto3_fields_proto_depIdxs,
		EnumOutputTypes:    xxx_File_proto3_fields_proto_enumTypes,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_proto3_fields_proto_goTypes[1:][:5]
	for i, mt := range messageTypes {
		xxx_File_proto3_fields_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_proto3_fields_proto_messageTypes[i].PBType = mt
	}
	xxx_File_proto3_fields_proto_goTypes = nil
	xxx_File_proto3_fields_proto_depIdxs = nil
}
