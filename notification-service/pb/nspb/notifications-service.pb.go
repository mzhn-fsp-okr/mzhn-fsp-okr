// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/notifications-service.proto

package nspb

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

type LinkTelegramRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId string `protobuf:"bytes,2,opt,name=chatId,proto3" json:"chatId,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LinkTelegramRequest) Reset() {
	*x = LinkTelegramRequest{}
	mi := &file_proto_notifications_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinkTelegramRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkTelegramRequest) ProtoMessage() {}

func (x *LinkTelegramRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkTelegramRequest.ProtoReflect.Descriptor instead.
func (*LinkTelegramRequest) Descriptor() ([]byte, []int) {
	return file_proto_notifications_service_proto_rawDescGZIP(), []int{0}
}

func (x *LinkTelegramRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *LinkTelegramRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LinkTelegramResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LinkTelegramResponse) Reset() {
	*x = LinkTelegramResponse{}
	mi := &file_proto_notifications_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinkTelegramResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkTelegramResponse) ProtoMessage() {}

func (x *LinkTelegramResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notifications_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkTelegramResponse.ProtoReflect.Descriptor instead.
func (*LinkTelegramResponse) Descriptor() ([]byte, []int) {
	return file_proto_notifications_service_proto_rawDescGZIP(), []int{1}
}

var File_proto_notifications_service_proto protoreflect.FileDescriptor

var file_proto_notifications_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x6e, 0x6b, 0x54,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x6e, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x54,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x27, 0x5a, 0x25, 0x6d, 0x7a, 0x68, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x6e, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notifications_service_proto_rawDescOnce sync.Once
	file_proto_notifications_service_proto_rawDescData = file_proto_notifications_service_proto_rawDesc
)

func file_proto_notifications_service_proto_rawDescGZIP() []byte {
	file_proto_notifications_service_proto_rawDescOnce.Do(func() {
		file_proto_notifications_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notifications_service_proto_rawDescData)
	})
	return file_proto_notifications_service_proto_rawDescData
}

var file_proto_notifications_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_notifications_service_proto_goTypes = []any{
	(*LinkTelegramRequest)(nil),  // 0: notifications.LinkTelegramRequest
	(*LinkTelegramResponse)(nil), // 1: notifications.LinkTelegramResponse
}
var file_proto_notifications_service_proto_depIdxs = []int32{
	0, // 0: notifications.NotificationService.LinkTelegram:input_type -> notifications.LinkTelegramRequest
	1, // 1: notifications.NotificationService.LinkTelegram:output_type -> notifications.LinkTelegramResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_notifications_service_proto_init() }
func file_proto_notifications_service_proto_init() {
	if File_proto_notifications_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_notifications_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notifications_service_proto_goTypes,
		DependencyIndexes: file_proto_notifications_service_proto_depIdxs,
		MessageInfos:      file_proto_notifications_service_proto_msgTypes,
	}.Build()
	File_proto_notifications_service_proto = out.File
	file_proto_notifications_service_proto_rawDesc = nil
	file_proto_notifications_service_proto_goTypes = nil
	file_proto_notifications_service_proto_depIdxs = nil
}
