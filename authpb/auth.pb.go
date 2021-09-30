// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: authpb/auth.proto

package authpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthMsg_MsgType int32

const (
	AuthMsg_CHALLENGE AuthMsg_MsgType = 0
	AuthMsg_RESPONSE  AuthMsg_MsgType = 1
	AuthMsg_CERT      AuthMsg_MsgType = 2
)

// Enum value maps for AuthMsg_MsgType.
var (
	AuthMsg_MsgType_name = map[int32]string{
		0: "CHALLENGE",
		1: "RESPONSE",
		2: "CERT",
	}
	AuthMsg_MsgType_value = map[string]int32{
		"CHALLENGE": 0,
		"RESPONSE":  1,
		"CERT":      2,
	}
)

func (x AuthMsg_MsgType) Enum() *AuthMsg_MsgType {
	p := new(AuthMsg_MsgType)
	*p = x
	return p
}

func (x AuthMsg_MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthMsg_MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_authpb_auth_proto_enumTypes[0].Descriptor()
}

func (AuthMsg_MsgType) Type() protoreflect.EnumType {
	return &file_authpb_auth_proto_enumTypes[0]
}

func (x AuthMsg_MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthMsg_MsgType.Descriptor instead.
func (AuthMsg_MsgType) EnumDescriptor() ([]byte, []int) {
	return file_authpb_auth_proto_rawDescGZIP(), []int{0, 0}
}

type AuthMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgtype AuthMsg_MsgType `protobuf:"varint,1,opt,name=msgtype,proto3,enum=auth.AuthMsg_MsgType" json:"msgtype,omitempty"`
	Msg     []byte          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AuthMsg) Reset() {
	*x = AuthMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authpb_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMsg) ProtoMessage() {}

func (x *AuthMsg) ProtoReflect() protoreflect.Message {
	mi := &file_authpb_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMsg.ProtoReflect.Descriptor instead.
func (*AuthMsg) Descriptor() ([]byte, []int) {
	return file_authpb_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthMsg) GetMsgtype() AuthMsg_MsgType {
	if x != nil {
		return x.Msgtype
	}
	return AuthMsg_CHALLENGE
}

func (x *AuthMsg) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_authpb_auth_proto protoreflect.FileDescriptor

var file_authpb_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x7e, 0x0a, 0x07, 0x41, 0x75, 0x74,
	0x68, 0x4d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x45, 0x52, 0x54, 0x10, 0x02, 0x32, 0x31, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x29, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x73, 0x67, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x73, 0x67, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6c, 0x70, 0x6f, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authpb_auth_proto_rawDescOnce sync.Once
	file_authpb_auth_proto_rawDescData = file_authpb_auth_proto_rawDesc
)

func file_authpb_auth_proto_rawDescGZIP() []byte {
	file_authpb_auth_proto_rawDescOnce.Do(func() {
		file_authpb_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_authpb_auth_proto_rawDescData)
	})
	return file_authpb_auth_proto_rawDescData
}

var file_authpb_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authpb_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_authpb_auth_proto_goTypes = []interface{}{
	(AuthMsg_MsgType)(0), // 0: auth.AuthMsg.MsgType
	(*AuthMsg)(nil),      // 1: auth.AuthMsg
}
var file_authpb_auth_proto_depIdxs = []int32{
	0, // 0: auth.AuthMsg.msgtype:type_name -> auth.AuthMsg.MsgType
	1, // 1: auth.Auth.Login:input_type -> auth.AuthMsg
	1, // 2: auth.Auth.Login:output_type -> auth.AuthMsg
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_authpb_auth_proto_init() }
func file_authpb_auth_proto_init() {
	if File_authpb_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authpb_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authpb_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authpb_auth_proto_goTypes,
		DependencyIndexes: file_authpb_auth_proto_depIdxs,
		EnumInfos:         file_authpb_auth_proto_enumTypes,
		MessageInfos:      file_authpb_auth_proto_msgTypes,
	}.Build()
	File_authpb_auth_proto = out.File
	file_authpb_auth_proto_rawDesc = nil
	file_authpb_auth_proto_goTypes = nil
	file_authpb_auth_proto_depIdxs = nil
}
