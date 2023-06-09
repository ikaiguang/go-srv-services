// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/admin-service/v1/resources/admin_reg_mobile.resource.v1.proto

package adminv1

import (
	v1 "github.com/ikaiguang/go-srv-kit/api/page/v1"
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

// AdminRegMobile ENGINE InnoDB CHARSET utf8mb4 COMMENT '管理员-注册-手机号'
type AdminRegMobile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// admin_mobile 管理员手机号码
	AdminMobile string `protobuf:"bytes,2,opt,name=admin_mobile,json=adminMobile,proto3" json:"admin_mobile,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// admin_uuid uuid
	AdminUuid string `protobuf:"bytes,5,opt,name=admin_uuid,json=adminUuid,proto3" json:"admin_uuid,omitempty"`
}

func (x *AdminRegMobile) Reset() {
	*x = AdminRegMobile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobile) ProtoMessage() {}

func (x *AdminRegMobile) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobile.ProtoReflect.Descriptor instead.
func (*AdminRegMobile) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *AdminRegMobile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdminRegMobile) GetAdminMobile() string {
	if x != nil {
		return x.AdminMobile
	}
	return ""
}

func (x *AdminRegMobile) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *AdminRegMobile) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *AdminRegMobile) GetAdminUuid() string {
	if x != nil {
		return x.AdminUuid
	}
	return ""
}

// AdminRegMobileSaveReq save
type AdminRegMobileSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// admin_mobile 管理员手机号码
	AdminMobile string `protobuf:"bytes,2,opt,name=admin_mobile,json=adminMobile,proto3" json:"admin_mobile,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// admin_uuid uuid
	AdminUuid string `protobuf:"bytes,5,opt,name=admin_uuid,json=adminUuid,proto3" json:"admin_uuid,omitempty"`
}

func (x *AdminRegMobileSaveReq) Reset() {
	*x = AdminRegMobileSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobileSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobileSaveReq) ProtoMessage() {}

func (x *AdminRegMobileSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobileSaveReq.ProtoReflect.Descriptor instead.
func (*AdminRegMobileSaveReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *AdminRegMobileSaveReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AdminRegMobileSaveReq) GetAdminMobile() string {
	if x != nil {
		return x.AdminMobile
	}
	return ""
}

func (x *AdminRegMobileSaveReq) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *AdminRegMobileSaveReq) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *AdminRegMobileSaveReq) GetAdminUuid() string {
	if x != nil {
		return x.AdminUuid
	}
	return ""
}

// AdminRegMobileIdReq id
type AdminRegMobileIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin_reg_mobile_id id
	AdminRegMobileId int64 `protobuf:"varint,1,opt,name=admin_reg_mobile_id,json=adminRegMobileId,proto3" json:"admin_reg_mobile_id,omitempty"`
}

func (x *AdminRegMobileIdReq) Reset() {
	*x = AdminRegMobileIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobileIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobileIdReq) ProtoMessage() {}

func (x *AdminRegMobileIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobileIdReq.ProtoReflect.Descriptor instead.
func (*AdminRegMobileIdReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *AdminRegMobileIdReq) GetAdminRegMobileId() int64 {
	if x != nil {
		return x.AdminRegMobileId
	}
	return 0
}

// AdminRegMobileIdsReq ids
type AdminRegMobileIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// admin_reg_mobile_ids ids
	AdminRegMobileIds []int64 `protobuf:"varint,1,rep,packed,name=admin_reg_mobile_ids,json=adminRegMobileIds,proto3" json:"admin_reg_mobile_ids,omitempty"`
}

func (x *AdminRegMobileIdsReq) Reset() {
	*x = AdminRegMobileIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobileIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobileIdsReq) ProtoMessage() {}

func (x *AdminRegMobileIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobileIdsReq.ProtoReflect.Descriptor instead.
func (*AdminRegMobileIdsReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *AdminRegMobileIdsReq) GetAdminRegMobileIds() []int64 {
	if x != nil {
		return x.AdminRegMobileIds
	}
	return nil
}

// AdminRegMobileListReq list
type AdminRegMobileListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_request paging request
	PageRequest *v1.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (x *AdminRegMobileListReq) Reset() {
	*x = AdminRegMobileListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobileListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobileListReq) ProtoMessage() {}

func (x *AdminRegMobileListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobileListReq.ProtoReflect.Descriptor instead.
func (*AdminRegMobileListReq) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{4}
}

func (x *AdminRegMobileListReq) GetPageRequest() *v1.PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

// AdminRegMobileListResp list
type AdminRegMobileListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list AdminRegMobile-array
	List []*AdminRegMobile `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// page_info paging response
	PageInfo *v1.PageResponse `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *AdminRegMobileListResp) Reset() {
	*x = AdminRegMobileListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobileListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobileListResp) ProtoMessage() {}

func (x *AdminRegMobileListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobileListResp.ProtoReflect.Descriptor instead.
func (*AdminRegMobileListResp) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{5}
}

func (x *AdminRegMobileListResp) GetList() []*AdminRegMobile {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *AdminRegMobileListResp) GetPageInfo() *v1.PageResponse {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

// AdminRegMobileProcessResult process result
type AdminRegMobileProcessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success true or false
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AdminRegMobileProcessResult) Reset() {
	*x = AdminRegMobileProcessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminRegMobileProcessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminRegMobileProcessResult) ProtoMessage() {}

func (x *AdminRegMobileProcessResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminRegMobileProcessResult.ProtoReflect.Descriptor instead.
func (*AdminRegMobileProcessResult) Descriptor() ([]byte, []int) {
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{6}
}

func (x *AdminRegMobileProcessResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto protoreflect.FileDescriptor

var file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x41, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x22, 0xaf,
	0x01, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x75, 0x69, 0x64,
	0x22, 0x44, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2f,
	0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x11, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22,
	0x57, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x16, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x67,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x37, 0x0a, 0x1b, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x77, 0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x70, 0x69, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61,
	0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescOnce sync.Once
	file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescData = file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDesc
)

func file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescGZIP() []byte {
	file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescData)
	})
	return file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDescData
}

var file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_goTypes = []interface{}{
	(*AdminRegMobile)(nil),              // 0: service.api.adminv1.AdminRegMobile
	(*AdminRegMobileSaveReq)(nil),       // 1: service.api.adminv1.AdminRegMobileSaveReq
	(*AdminRegMobileIdReq)(nil),         // 2: service.api.adminv1.AdminRegMobileIdReq
	(*AdminRegMobileIdsReq)(nil),        // 3: service.api.adminv1.AdminRegMobileIdsReq
	(*AdminRegMobileListReq)(nil),       // 4: service.api.adminv1.AdminRegMobileListReq
	(*AdminRegMobileListResp)(nil),      // 5: service.api.adminv1.AdminRegMobileListResp
	(*AdminRegMobileProcessResult)(nil), // 6: service.api.adminv1.AdminRegMobileProcessResult
	(*v1.PageRequest)(nil),              // 7: kit.api.pagev1.PageRequest
	(*v1.PageResponse)(nil),             // 8: kit.api.pagev1.PageResponse
}
var file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_depIdxs = []int32{
	7, // 0: service.api.adminv1.AdminRegMobileListReq.page_request:type_name -> kit.api.pagev1.PageRequest
	0, // 1: service.api.adminv1.AdminRegMobileListResp.list:type_name -> service.api.adminv1.AdminRegMobile
	8, // 2: service.api.adminv1.AdminRegMobileListResp.page_info:type_name -> kit.api.pagev1.PageResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_init() }
func file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_init() {
	if File_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobile); i {
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
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobileSaveReq); i {
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
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobileIdReq); i {
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
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobileIdsReq); i {
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
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobileListReq); i {
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
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobileListResp); i {
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
		file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminRegMobileProcessResult); i {
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
			RawDescriptor: file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_msgTypes,
	}.Build()
	File_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto = out.File
	file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_rawDesc = nil
	file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_goTypes = nil
	file_api_admin_service_v1_resources_admin_reg_mobile_resource_v1_proto_depIdxs = nil
}
