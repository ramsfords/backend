// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: add_staff.proto

package user

import (
	_ "github.com/ramsfords/user_gen/v1/user/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddStaffData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"token"
	Token []string `protobuf:"bytes,1,rep,name=token,proto3" json:"token,omitempty" dynamodbav:"token"`
	// @gotags: dynamodbav:"roles"
	Roles []Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=user.Role" json:"roles,omitempty" dynamodbav:"roles"`
	// @gotags: dynamodbav:"new_staff_email"
	NewStaffEmail []string `protobuf:"bytes,3,rep,name=new_staff_email,json=newStaffEmail,proto3" json:"new_staff_email,omitempty" dynamodbav:"new_staff_email"`
}

func (x *AddStaffData) Reset() {
	*x = AddStaffData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStaffData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStaffData) ProtoMessage() {}

func (x *AddStaffData) ProtoReflect() protoreflect.Message {
	mi := &file_add_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStaffData.ProtoReflect.Descriptor instead.
func (*AddStaffData) Descriptor() ([]byte, []int) {
	return file_add_staff_proto_rawDescGZIP(), []int{0}
}

func (x *AddStaffData) GetToken() []string {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *AddStaffData) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AddStaffData) GetNewStaffEmail() []string {
	if x != nil {
		return x.NewStaffEmail
	}
	return nil
}

var File_add_staff_proto protoreflect.FileDescriptor

var file_add_staff_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x0c, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xe9,
	0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2c, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03,
	0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x0f, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x60,
	0x01, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x66, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x70, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0d, 0x41, 0x64,
	0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f,
	0x72, 0x64, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65,
	0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_add_staff_proto_rawDescOnce sync.Once
	file_add_staff_proto_rawDescData = file_add_staff_proto_rawDesc
)

func file_add_staff_proto_rawDescGZIP() []byte {
	file_add_staff_proto_rawDescOnce.Do(func() {
		file_add_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_add_staff_proto_rawDescData)
	})
	return file_add_staff_proto_rawDescData
}

var file_add_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_add_staff_proto_goTypes = []interface{}{
	(*AddStaffData)(nil), // 0: user.AddStaffData
	(Role)(0),            // 1: user.Role
}
var file_add_staff_proto_depIdxs = []int32{
	1, // 0: user.AddStaffData.roles:type_name -> user.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_add_staff_proto_init() }
func file_add_staff_proto_init() {
	if File_add_staff_proto != nil {
		return
	}
	file_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_add_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStaffData); i {
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
			RawDescriptor: file_add_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_add_staff_proto_goTypes,
		DependencyIndexes: file_add_staff_proto_depIdxs,
		MessageInfos:      file_add_staff_proto_msgTypes,
	}.Build()
	File_add_staff_proto = out.File
	file_add_staff_proto_rawDesc = nil
	file_add_staff_proto_goTypes = nil
	file_add_staff_proto_depIdxs = nil
}
