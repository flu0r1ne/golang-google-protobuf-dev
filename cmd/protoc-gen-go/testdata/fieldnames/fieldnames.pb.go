// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fieldnames/fieldnames.proto

package fieldnames

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

// Assorted edge cases in field name conflict resolution.
//
// Not all (or possibly any) of these behave in an easily-understood fashion.
// This exists to demonstrate the current behavior and catch unintended
// changes in it.
type Message struct {
	// Various CamelCase conversions.
	FieldOne   *string `protobuf:"bytes,1,opt,name=field_one,json=fieldOne" json:"field_one,omitempty"`
	FieldTwo   *string `protobuf:"bytes,2,opt,name=FieldTwo" json:"FieldTwo,omitempty"`
	FieldThree *string `protobuf:"bytes,3,opt,name=fieldThree" json:"fieldThree,omitempty"`
	Field_Four *string `protobuf:"bytes,4,opt,name=field__four,json=fieldFour" json:"field__four,omitempty"`
	// Field names that conflict with standard methods on the message struct.
	Descriptor_   *string `protobuf:"bytes,10,opt,name=descriptor" json:"descriptor,omitempty"`
	Marshal_      *string `protobuf:"bytes,11,opt,name=marshal" json:"marshal,omitempty"`
	Unmarshal_    *string `protobuf:"bytes,12,opt,name=unmarshal" json:"unmarshal,omitempty"`
	ProtoMessage_ *string `protobuf:"bytes,13,opt,name=proto_message,json=protoMessage" json:"proto_message,omitempty"`
	// Field names that conflict with each other after CamelCasing.
	CamelCase    *string `protobuf:"bytes,20,opt,name=CamelCase" json:"CamelCase,omitempty"`
	CamelCase_   *string `protobuf:"bytes,21,opt,name=CamelCase_,json=CamelCase" json:"CamelCase_,omitempty"`
	CamelCase__  *string `protobuf:"bytes,22,opt,name=camel_case,json=camelCase" json:"camel_case,omitempty"`
	CamelCase___ *string `protobuf:"bytes,23,opt,name=CamelCase__,json=CamelCase" json:"CamelCase__,omitempty"`
	// Field with a getter that conflicts with another field.
	GetName *string `protobuf:"bytes,30,opt,name=get_name,json=getName" json:"get_name,omitempty"`
	Name_   *string `protobuf:"bytes,31,opt,name=name" json:"name,omitempty"`
	// Oneof that conflicts with its first field: The oneof is renamed.
	//
	// Types that are valid to be assigned to OneofConflictA_:
	//	*Message_OneofConflictA
	OneofConflictA_ isMessage_OneofConflictA_ `protobuf_oneof:"oneof_conflict_a"`
	// Oneof that conflicts with its second field: The field is renamed.
	//
	// Types that are valid to be assigned to OneofConflictB:
	//	*Message_OneofNoConflict
	//	*Message_OneofConflictB_
	OneofConflictB isMessage_OneofConflictB `protobuf_oneof:"oneof_conflict_b"`
	// Oneof with a field name that conflicts with a nested message.
	//
	// Types that are valid to be assigned to OneofConflictC:
	//	*Message_OneofMessageConflict_
	OneofConflictC       isMessage_OneofConflictC `protobuf_oneof:"oneof_conflict_c"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Message) ProtoReflect() protoreflect.Message {
	return xxx_File_fieldnames_fieldnames_proto_messageTypes[0].MessageOf(m)
}
func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Message) ProtoMessage()    {}

// Deprecated: Use Message.ProtoReflect.Type instead.
func (*Message) Descriptor() ([]byte, []int) {
	return xxx_File_fieldnames_fieldnames_proto_rawDescGZIP(), []int{0}
}

func (m *Message) GetFieldOne() string {
	if m != nil && m.FieldOne != nil {
		return *m.FieldOne
	}
	return ""
}

func (m *Message) GetFieldTwo() string {
	if m != nil && m.FieldTwo != nil {
		return *m.FieldTwo
	}
	return ""
}

func (m *Message) GetFieldThree() string {
	if m != nil && m.FieldThree != nil {
		return *m.FieldThree
	}
	return ""
}

func (m *Message) GetField_Four() string {
	if m != nil && m.Field_Four != nil {
		return *m.Field_Four
	}
	return ""
}

func (m *Message) GetDescriptor_() string {
	if m != nil && m.Descriptor_ != nil {
		return *m.Descriptor_
	}
	return ""
}

func (m *Message) GetMarshal_() string {
	if m != nil && m.Marshal_ != nil {
		return *m.Marshal_
	}
	return ""
}

func (m *Message) GetUnmarshal_() string {
	if m != nil && m.Unmarshal_ != nil {
		return *m.Unmarshal_
	}
	return ""
}

func (m *Message) GetProtoMessage_() string {
	if m != nil && m.ProtoMessage_ != nil {
		return *m.ProtoMessage_
	}
	return ""
}

func (m *Message) GetCamelCase() string {
	if m != nil && m.CamelCase != nil {
		return *m.CamelCase
	}
	return ""
}

func (m *Message) GetCamelCase_() string {
	if m != nil && m.CamelCase_ != nil {
		return *m.CamelCase_
	}
	return ""
}

func (m *Message) GetCamelCase__() string {
	if m != nil && m.CamelCase__ != nil {
		return *m.CamelCase__
	}
	return ""
}

func (m *Message) GetCamelCase___() string {
	if m != nil && m.CamelCase___ != nil {
		return *m.CamelCase___
	}
	return ""
}

func (m *Message) GetGetName() string {
	if m != nil && m.GetName != nil {
		return *m.GetName
	}
	return ""
}

func (m *Message) GetName_() string {
	if m != nil && m.Name_ != nil {
		return *m.Name_
	}
	return ""
}

type isMessage_OneofConflictA_ interface {
	isMessage_OneofConflictA_()
}

type Message_OneofConflictA struct {
	OneofConflictA string `protobuf:"bytes,40,opt,name=OneofConflictA,oneof"`
}

func (*Message_OneofConflictA) isMessage_OneofConflictA_() {}

func (m *Message) GetOneofConflictA_() isMessage_OneofConflictA_ {
	if m != nil {
		return m.OneofConflictA_
	}
	return nil
}

func (m *Message) GetOneofConflictA() string {
	if x, ok := m.GetOneofConflictA_().(*Message_OneofConflictA); ok {
		return x.OneofConflictA
	}
	return ""
}

type isMessage_OneofConflictB interface {
	isMessage_OneofConflictB()
}

type Message_OneofNoConflict struct {
	OneofNoConflict string `protobuf:"bytes,50,opt,name=oneof_no_conflict,json=oneofNoConflict,oneof"`
}

type Message_OneofConflictB_ struct {
	OneofConflictB_ string `protobuf:"bytes,51,opt,name=OneofConflictB,oneof"`
}

func (*Message_OneofNoConflict) isMessage_OneofConflictB() {}

func (*Message_OneofConflictB_) isMessage_OneofConflictB() {}

func (m *Message) GetOneofConflictB() isMessage_OneofConflictB {
	if m != nil {
		return m.OneofConflictB
	}
	return nil
}

func (m *Message) GetOneofNoConflict() string {
	if x, ok := m.GetOneofConflictB().(*Message_OneofNoConflict); ok {
		return x.OneofNoConflict
	}
	return ""
}

func (m *Message) GetOneofConflictB_() string {
	if x, ok := m.GetOneofConflictB().(*Message_OneofConflictB_); ok {
		return x.OneofConflictB_
	}
	return ""
}

type isMessage_OneofConflictC interface {
	isMessage_OneofConflictC()
}

type Message_OneofMessageConflict_ struct {
	OneofMessageConflict string `protobuf:"bytes,60,opt,name=oneof_message_conflict,json=oneofMessageConflict,oneof"`
}

func (*Message_OneofMessageConflict_) isMessage_OneofConflictC() {}

func (m *Message) GetOneofConflictC() isMessage_OneofConflictC {
	if m != nil {
		return m.OneofConflictC
	}
	return nil
}

func (m *Message) GetOneofMessageConflict() string {
	if x, ok := m.GetOneofConflictC().(*Message_OneofMessageConflict_); ok {
		return x.OneofMessageConflict
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_OneofConflictA)(nil),
		(*Message_OneofNoConflict)(nil),
		(*Message_OneofConflictB_)(nil),
		(*Message_OneofMessageConflict_)(nil),
	}
}

type Message_OneofMessageConflict struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_OneofMessageConflict) ProtoReflect() protoreflect.Message {
	return xxx_File_fieldnames_fieldnames_proto_messageTypes[1].MessageOf(m)
}
func (m *Message_OneofMessageConflict) Reset()         { *m = Message_OneofMessageConflict{} }
func (m *Message_OneofMessageConflict) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Message_OneofMessageConflict) ProtoMessage()    {}

// Deprecated: Use Message_OneofMessageConflict.ProtoReflect.Type instead.
func (*Message_OneofMessageConflict) Descriptor() ([]byte, []int) {
	return xxx_File_fieldnames_fieldnames_proto_rawDescGZIP(), []int{0, 0}
}

var File_fieldnames_fieldnames_proto protoreflect.FileDescriptor

var xxx_File_fieldnames_fieldnames_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xb8, 0x05, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x6e,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x77, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x68, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x68, 0x72, 0x65, 0x65, 0x12, 0x1e, 0x0a,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x5f, 0x66, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x6f, 0x75, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x72,
	0x73, 0x68, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x6d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61,
	0x6d, 0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x61, 0x6d, 0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x43, 0x61, 0x6d, 0x65,
	0x6c, 0x43, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x61,
	0x6d, 0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x6d, 0x65, 0x6c,
	0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6d,
	0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x43, 0x61, 0x6d, 0x65, 0x6c, 0x43,
	0x61, 0x73, 0x65, 0x5f, 0x5f, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x61, 0x6d,
	0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f,
	0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x41, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x41, 0x12,
	0x2c, 0x0a, 0x11, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6e, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x69, 0x63, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x4e, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x12, 0x28, 0x0a,
	0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x42, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f,
	0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x42, 0x12, 0x36, 0x0a, 0x16, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63,
	0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x14, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x1a,
	0x16, 0x0a, 0x14, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x5f, 0x61, 0x42, 0x12, 0x0a, 0x10, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x5f, 0x62, 0x42,
	0x12, 0x0a, 0x10, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63,
	0x74, 0x5f, 0x63, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73,
}

var (
	xxx_File_fieldnames_fieldnames_proto_rawDesc_once sync.Once
	xxx_File_fieldnames_fieldnames_proto_rawDesc_data = xxx_File_fieldnames_fieldnames_proto_rawDesc
)

func xxx_File_fieldnames_fieldnames_proto_rawDescGZIP() []byte {
	xxx_File_fieldnames_fieldnames_proto_rawDesc_once.Do(func() {
		xxx_File_fieldnames_fieldnames_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_fieldnames_fieldnames_proto_rawDesc_data)
	})
	return xxx_File_fieldnames_fieldnames_proto_rawDesc_data
}

var xxx_File_fieldnames_fieldnames_proto_messageTypes = make([]protoimpl.MessageType, 2)
var xxx_File_fieldnames_fieldnames_proto_goTypes = []interface{}{
	(*Message)(nil),                      // 0: goproto.protoc.fieldnames.Message
	(*Message_OneofMessageConflict)(nil), // 1: goproto.protoc.fieldnames.Message.OneofMessageConflict
}
var xxx_File_fieldnames_fieldnames_proto_depIdxs = []int32{}

func init() { xxx_File_fieldnames_fieldnames_proto_init() }
func xxx_File_fieldnames_fieldnames_proto_init() {
	if File_fieldnames_fieldnames_proto != nil {
		return
	}
	File_fieldnames_fieldnames_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_fieldnames_fieldnames_proto_rawDesc,
		GoTypes:            xxx_File_fieldnames_fieldnames_proto_goTypes,
		DependencyIndexes:  xxx_File_fieldnames_fieldnames_proto_depIdxs,
		MessageOutputTypes: xxx_File_fieldnames_fieldnames_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_fieldnames_fieldnames_proto_rawDesc = nil
	xxx_File_fieldnames_fieldnames_proto_goTypes = nil
	xxx_File_fieldnames_fieldnames_proto_depIdxs = nil
}
