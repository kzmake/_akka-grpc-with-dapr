// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: agwd/email/v1/email.proto

package emailv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agwd_email_v1_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agwd_email_v1_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_agwd_email_v1_email_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agwd_email_v1_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agwd_email_v1_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_agwd_email_v1_email_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agwd_email_v1_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_agwd_email_v1_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_agwd_email_v1_email_proto_rawDescGZIP(), []int{2}
}

func (x *Email) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_agwd_email_v1_email_proto protoreflect.FileDescriptor

var file_agwd_email_v1_email_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x67, 0x77, 0x64, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x67, 0x77,
	0x64, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x67, 0x77, 0x64, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x55, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x61, 0x67, 0x77, 0x64, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x61, 0x67, 0x77, 0x64, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc7, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x67, 0x77, 0x64, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x7a,
	0x6d, 0x61, 0x6b, 0x65, 0x2f, 0x5f, 0x61, 0x6b, 0x6b, 0x61, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x77, 0x69, 0x74, 0x68, 0x2d, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x7a,
	0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x67,
	0x77, 0x64, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x0d, 0x41, 0x67, 0x77, 0x64,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x41, 0x67, 0x77, 0x64,
	0x5c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x41, 0x67, 0x77, 0x64,
	0x5c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x41, 0x67, 0x77, 0x64, 0x3a, 0x3a, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agwd_email_v1_email_proto_rawDescOnce sync.Once
	file_agwd_email_v1_email_proto_rawDescData = file_agwd_email_v1_email_proto_rawDesc
)

func file_agwd_email_v1_email_proto_rawDescGZIP() []byte {
	file_agwd_email_v1_email_proto_rawDescOnce.Do(func() {
		file_agwd_email_v1_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_agwd_email_v1_email_proto_rawDescData)
	})
	return file_agwd_email_v1_email_proto_rawDescData
}

var file_agwd_email_v1_email_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_agwd_email_v1_email_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: agwd.email.v1.CreateRequest
	(*CreateResponse)(nil), // 1: agwd.email.v1.CreateResponse
	(*Email)(nil),          // 2: agwd.email.v1.Email
}
var file_agwd_email_v1_email_proto_depIdxs = []int32{
	2, // 0: agwd.email.v1.CreateResponse.email:type_name -> agwd.email.v1.Email
	0, // 1: agwd.email.v1.EmailService.Create:input_type -> agwd.email.v1.CreateRequest
	1, // 2: agwd.email.v1.EmailService.Create:output_type -> agwd.email.v1.CreateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_agwd_email_v1_email_proto_init() }
func file_agwd_email_v1_email_proto_init() {
	if File_agwd_email_v1_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agwd_email_v1_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_agwd_email_v1_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_agwd_email_v1_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
			RawDescriptor: file_agwd_email_v1_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agwd_email_v1_email_proto_goTypes,
		DependencyIndexes: file_agwd_email_v1_email_proto_depIdxs,
		MessageInfos:      file_agwd_email_v1_email_proto_msgTypes,
	}.Build()
	File_agwd_email_v1_email_proto = out.File
	file_agwd_email_v1_email_proto_rawDesc = nil
	file_agwd_email_v1_email_proto_goTypes = nil
	file_agwd_email_v1_email_proto_depIdxs = nil
}
