// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: freight_class.proto

package carrier

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

type FreightClass int32

const (
	FreightClass_CLASS_NONE FreightClass = 0
	FreightClass_CLASS50    FreightClass = 1
	FreightClass_CLASS55    FreightClass = 2
	FreightClass_CLASS60    FreightClass = 3
	FreightClass_CLASS65    FreightClass = 4
	FreightClass_CLASS70    FreightClass = 5
	FreightClass_CLASS775   FreightClass = 6
	FreightClass_CLASS85    FreightClass = 7
	FreightClass_CLASS925   FreightClass = 8
	FreightClass_CLASS100   FreightClass = 9
	FreightClass_CLASS110   FreightClass = 10
	FreightClass_CLASS125   FreightClass = 11
	FreightClass_CLASS150   FreightClass = 12
	FreightClass_CLASS175   FreightClass = 13
	FreightClass_CLASS200   FreightClass = 14
	FreightClass_CLASS250   FreightClass = 15
	FreightClass_CLASS300   FreightClass = 16
	FreightClass_CLASS400   FreightClass = 17
	FreightClass_CLASS500   FreightClass = 18
)

// Enum value maps for FreightClass.
var (
	FreightClass_name = map[int32]string{
		0:  "CLASS_NONE",
		1:  "CLASS50",
		2:  "CLASS55",
		3:  "CLASS60",
		4:  "CLASS65",
		5:  "CLASS70",
		6:  "CLASS775",
		7:  "CLASS85",
		8:  "CLASS925",
		9:  "CLASS100",
		10: "CLASS110",
		11: "CLASS125",
		12: "CLASS150",
		13: "CLASS175",
		14: "CLASS200",
		15: "CLASS250",
		16: "CLASS300",
		17: "CLASS400",
		18: "CLASS500",
	}
	FreightClass_value = map[string]int32{
		"CLASS_NONE": 0,
		"CLASS50":    1,
		"CLASS55":    2,
		"CLASS60":    3,
		"CLASS65":    4,
		"CLASS70":    5,
		"CLASS775":   6,
		"CLASS85":    7,
		"CLASS925":   8,
		"CLASS100":   9,
		"CLASS110":   10,
		"CLASS125":   11,
		"CLASS150":   12,
		"CLASS175":   13,
		"CLASS200":   14,
		"CLASS250":   15,
		"CLASS300":   16,
		"CLASS400":   17,
		"CLASS500":   18,
	}
)

func (x FreightClass) Enum() *FreightClass {
	p := new(FreightClass)
	*p = x
	return p
}

func (x FreightClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FreightClass) Descriptor() protoreflect.EnumDescriptor {
	return file_freight_class_proto_enumTypes[0].Descriptor()
}

func (FreightClass) Type() protoreflect.EnumType {
	return &file_freight_class_proto_enumTypes[0]
}

func (x FreightClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FreightClass.Descriptor instead.
func (FreightClass) EnumDescriptor() ([]byte, []int) {
	return file_freight_class_proto_rawDescGZIP(), []int{0}
}

var File_freight_class_proto protoreflect.FileDescriptor

var file_freight_class_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x2a, 0x95, 0x02, 0x0a, 0x0d, 0x66, 0x72,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x35, 0x30, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53,
	0x53, 0x35, 0x35, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x36, 0x30,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x36, 0x35, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x37, 0x30, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x4c, 0x41, 0x53, 0x53, 0x37, 0x37, 0x35, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x38, 0x35, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x39, 0x32, 0x35, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x31, 0x30,
	0x30, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x31, 0x31, 0x30, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x31, 0x32, 0x35, 0x10, 0x0b, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x31, 0x35, 0x30, 0x10, 0x0c, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x31, 0x37, 0x35, 0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x32, 0x30, 0x30, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41,
	0x53, 0x53, 0x32, 0x35, 0x30, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53,
	0x33, 0x30, 0x30, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x34, 0x30,
	0x30, 0x10, 0x11, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x35, 0x30, 0x30, 0x10,
	0x12, 0x42, 0x70, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x46, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d,
	0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0xa2, 0x02, 0x03,
	0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freight_class_proto_rawDescOnce sync.Once
	file_freight_class_proto_rawDescData = file_freight_class_proto_rawDesc
)

func file_freight_class_proto_rawDescGZIP() []byte {
	file_freight_class_proto_rawDescOnce.Do(func() {
		file_freight_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_freight_class_proto_rawDescData)
	})
	return file_freight_class_proto_rawDescData
}

var file_freight_class_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_freight_class_proto_goTypes = []interface{}{
	(FreightClass)(0), // 0: v1.freight_class
}
var file_freight_class_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_freight_class_proto_init() }
func file_freight_class_proto_init() {
	if File_freight_class_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_freight_class_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freight_class_proto_goTypes,
		DependencyIndexes: file_freight_class_proto_depIdxs,
		EnumInfos:         file_freight_class_proto_enumTypes,
	}.Build()
	File_freight_class_proto = out.File
	file_freight_class_proto_rawDesc = nil
	file_freight_class_proto_goTypes = nil
	file_freight_class_proto_depIdxs = nil
}
