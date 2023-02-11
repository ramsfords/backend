// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ok.proto

package v1

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

type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"success,omitempty"
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty" dynamodbav:"success,omitempty"`
	// @gotags: dynamodbav:"code,omitempty"
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty" dynamodbav:"code,omitempty"`
	// @gotags: dynamodbav:"message,omitempty"
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty" dynamodbav:"message,omitempty"`
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ok_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_ok_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_ok_proto_rawDescGZIP(), []int{0}
}

func (x *Ok) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Ok) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Ok) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_ok_proto protoreflect.FileDescriptor

var file_ok_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x4c,
	0x0a, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x5c, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x07, 0x4f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca,
	0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ok_proto_rawDescOnce sync.Once
	file_ok_proto_rawDescData = file_ok_proto_rawDesc
)

func file_ok_proto_rawDescGZIP() []byte {
	file_ok_proto_rawDescOnce.Do(func() {
		file_ok_proto_rawDescData = protoimpl.X.CompressGZIP(file_ok_proto_rawDescData)
	})
	return file_ok_proto_rawDescData
}

var file_ok_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ok_proto_goTypes = []interface{}{
	(*Ok)(nil), // 0: v1.ok
}
var file_ok_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ok_proto_init() }
func file_ok_proto_init() {
	if File_ok_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ok_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
			RawDescriptor: file_ok_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ok_proto_goTypes,
		DependencyIndexes: file_ok_proto_depIdxs,
		MessageInfos:      file_ok_proto_msgTypes,
	}.Build()
	File_ok_proto = out.File
	file_ok_proto_rawDesc = nil
	file_ok_proto_goTypes = nil
	file_ok_proto_depIdxs = nil
}