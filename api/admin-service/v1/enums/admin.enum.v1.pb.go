// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/admin-service/v1/enums/admin.enum.v1.proto

package adminenumv1

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

// AdminInit AdminInit enum
type AdminInitEnum_AdminInit int32

const (
	// UNSPECIFIED 未指定
	AdminInitEnum_UNSPECIFIED AdminInitEnum_AdminInit = 0
)

// Enum value maps for AdminInitEnum_AdminInit.
var (
	AdminInitEnum_AdminInit_name = map[int32]string{
		0: "UNSPECIFIED",
	}
	AdminInitEnum_AdminInit_value = map[string]int32{
		"UNSPECIFIED": 0,
	}
)

func (x AdminInitEnum_AdminInit) Enum() *AdminInitEnum_AdminInit {
	p := new(AdminInitEnum_AdminInit)
	*p = x
	return p
}

func (x AdminInitEnum_AdminInit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminInitEnum_AdminInit) Descriptor() protoreflect.EnumDescriptor {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes[0].Descriptor()
}

func (AdminInitEnum_AdminInit) Type() protoreflect.EnumType {
	return &file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes[0]
}

func (x AdminInitEnum_AdminInit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminInitEnum_AdminInit.Descriptor instead.
func (AdminInitEnum_AdminInit) EnumDescriptor() ([]byte, []int) {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP(), []int{0, 0}
}

// UserStatus 枚举值
type AdminStatusEnum_AdminStatus int32

const (
	// UNSPECIFIED 未指定
	AdminStatusEnum_UNSPECIFIED AdminStatusEnum_AdminStatus = 0
	// INACTIVATED 未激活
	AdminStatusEnum_INACTIVATED AdminStatusEnum_AdminStatus = 1
	// ENABLE 启用
	AdminStatusEnum_ENABLE AdminStatusEnum_AdminStatus = 2
	// DISABLE 禁用
	AdminStatusEnum_DISABLE AdminStatusEnum_AdminStatus = 3
	// BLACKLIST 黑名单
	AdminStatusEnum_BLACKLIST AdminStatusEnum_AdminStatus = 4
	// DELETED 已删除
	AdminStatusEnum_DELETED AdminStatusEnum_AdminStatus = 5
)

// Enum value maps for AdminStatusEnum_AdminStatus.
var (
	AdminStatusEnum_AdminStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "INACTIVATED",
		2: "ENABLE",
		3: "DISABLE",
		4: "BLACKLIST",
		5: "DELETED",
	}
	AdminStatusEnum_AdminStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"INACTIVATED": 1,
		"ENABLE":      2,
		"DISABLE":     3,
		"BLACKLIST":   4,
		"DELETED":     5,
	}
)

func (x AdminStatusEnum_AdminStatus) Enum() *AdminStatusEnum_AdminStatus {
	p := new(AdminStatusEnum_AdminStatus)
	*p = x
	return p
}

func (x AdminStatusEnum_AdminStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminStatusEnum_AdminStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes[1].Descriptor()
}

func (AdminStatusEnum_AdminStatus) Type() protoreflect.EnumType {
	return &file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes[1]
}

func (x AdminStatusEnum_AdminStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminStatusEnum_AdminStatus.Descriptor instead.
func (AdminStatusEnum_AdminStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP(), []int{1, 0}
}

// UserGender 枚举值
type AdminGenderEnum_AdminGender int32

const (
	// UNSPECIFIED 未指定
	AdminGenderEnum_UNSPECIFIED AdminGenderEnum_AdminGender = 0
	// MALE 男
	AdminGenderEnum_MALE AdminGenderEnum_AdminGender = 1
	// FEMALE 女
	AdminGenderEnum_FEMALE AdminGenderEnum_AdminGender = 2
	// SECRET 密码
	AdminGenderEnum_SECRET AdminGenderEnum_AdminGender = 3
)

// Enum value maps for AdminGenderEnum_AdminGender.
var (
	AdminGenderEnum_AdminGender_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "MALE",
		2: "FEMALE",
		3: "SECRET",
	}
	AdminGenderEnum_AdminGender_value = map[string]int32{
		"UNSPECIFIED": 0,
		"MALE":        1,
		"FEMALE":      2,
		"SECRET":      3,
	}
)

func (x AdminGenderEnum_AdminGender) Enum() *AdminGenderEnum_AdminGender {
	p := new(AdminGenderEnum_AdminGender)
	*p = x
	return p
}

func (x AdminGenderEnum_AdminGender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminGenderEnum_AdminGender) Descriptor() protoreflect.EnumDescriptor {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes[2].Descriptor()
}

func (AdminGenderEnum_AdminGender) Type() protoreflect.EnumType {
	return &file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes[2]
}

func (x AdminGenderEnum_AdminGender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminGenderEnum_AdminGender.Descriptor instead.
func (AdminGenderEnum_AdminGender) EnumDescriptor() ([]byte, []int) {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP(), []int{2, 0}
}

// AdminInitEnum AdminInitEnum enum
type AdminInitEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminInitEnum) Reset() {
	*x = AdminInitEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminInitEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminInitEnum) ProtoMessage() {}

func (x *AdminInitEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminInitEnum.ProtoReflect.Descriptor instead.
func (*AdminInitEnum) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP(), []int{0}
}

// AdminStatusEnum 用户状态
type AdminStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminStatusEnum) Reset() {
	*x = AdminStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminStatusEnum) ProtoMessage() {}

func (x *AdminStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminStatusEnum.ProtoReflect.Descriptor instead.
func (*AdminStatusEnum) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP(), []int{1}
}

// AdminGenderEnum 用户性别
type AdminGenderEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminGenderEnum) Reset() {
	*x = AdminGenderEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGenderEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGenderEnum) ProtoMessage() {}

func (x *AdminGenderEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGenderEnum.ProtoReflect.Descriptor instead.
func (*AdminGenderEnum) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP(), []int{2}
}

var File_api_admin_service_v1_enums_admin_enum_v1_proto protoreflect.FileDescriptor

var file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x22, 0x2d, 0x0a, 0x0d, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x1c, 0x0a, 0x09, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x22, 0x77, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x64, 0x0a, 0x0b, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49,
	0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x4c, 0x49,
	0x53, 0x54, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x22, 0x53, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0x40, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45,
	0x43, 0x52, 0x45, 0x54, 0x10, 0x03, 0x42, 0x7f, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x76,
	0x31, 0x42, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescOnce sync.Once
	file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescData = file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDesc
)

func file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescGZIP() []byte {
	file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescOnce.Do(func() {
		file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescData)
	})
	return file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDescData
}

var file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_admin_service_v1_enums_admin_enum_v1_proto_goTypes = []interface{}{
	(AdminInitEnum_AdminInit)(0),     // 0: service.api.adminenumv1.AdminInitEnum.AdminInit
	(AdminStatusEnum_AdminStatus)(0), // 1: service.api.adminenumv1.AdminStatusEnum.AdminStatus
	(AdminGenderEnum_AdminGender)(0), // 2: service.api.adminenumv1.AdminGenderEnum.AdminGender
	(*AdminInitEnum)(nil),            // 3: service.api.adminenumv1.AdminInitEnum
	(*AdminStatusEnum)(nil),          // 4: service.api.adminenumv1.AdminStatusEnum
	(*AdminGenderEnum)(nil),          // 5: service.api.adminenumv1.AdminGenderEnum
}
var file_api_admin_service_v1_enums_admin_enum_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_admin_service_v1_enums_admin_enum_v1_proto_init() }
func file_api_admin_service_v1_enums_admin_enum_v1_proto_init() {
	if File_api_admin_service_v1_enums_admin_enum_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminInitEnum); i {
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
		file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminStatusEnum); i {
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
		file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGenderEnum); i {
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
			RawDescriptor: file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_admin_service_v1_enums_admin_enum_v1_proto_goTypes,
		DependencyIndexes: file_api_admin_service_v1_enums_admin_enum_v1_proto_depIdxs,
		EnumInfos:         file_api_admin_service_v1_enums_admin_enum_v1_proto_enumTypes,
		MessageInfos:      file_api_admin_service_v1_enums_admin_enum_v1_proto_msgTypes,
	}.Build()
	File_api_admin_service_v1_enums_admin_enum_v1_proto = out.File
	file_api_admin_service_v1_enums_admin_enum_v1_proto_rawDesc = nil
	file_api_admin_service_v1_enums_admin_enum_v1_proto_goTypes = nil
	file_api_admin_service_v1_enums_admin_enum_v1_proto_depIdxs = nil
}
