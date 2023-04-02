// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: api/user-service/v1/resources/user_reg_mobile.resource.v1.proto

package userv1

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

// UserRegMobile ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户-注册-手机号'
type UserRegMobile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_mobile 用户手机号码
	UserMobile string `protobuf:"bytes,2,opt,name=user_mobile,json=userMobile,proto3" json:"user_mobile,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// user_uuid uuid
	UserUuid string `protobuf:"bytes,5,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *UserRegMobile) Reset() {
	*x = UserRegMobile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobile) ProtoMessage() {}

func (x *UserRegMobile) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobile.ProtoReflect.Descriptor instead.
func (*UserRegMobile) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegMobile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRegMobile) GetUserMobile() string {
	if x != nil {
		return x.UserMobile
	}
	return ""
}

func (x *UserRegMobile) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *UserRegMobile) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *UserRegMobile) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

// UserRegMobileSaveReq save
type UserRegMobileSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_mobile 用户手机号码
	UserMobile string `protobuf:"bytes,2,opt,name=user_mobile,json=userMobile,proto3" json:"user_mobile,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// user_uuid uuid
	UserUuid string `protobuf:"bytes,5,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *UserRegMobileSaveReq) Reset() {
	*x = UserRegMobileSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobileSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobileSaveReq) ProtoMessage() {}

func (x *UserRegMobileSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobileSaveReq.ProtoReflect.Descriptor instead.
func (*UserRegMobileSaveReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegMobileSaveReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRegMobileSaveReq) GetUserMobile() string {
	if x != nil {
		return x.UserMobile
	}
	return ""
}

func (x *UserRegMobileSaveReq) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *UserRegMobileSaveReq) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *UserRegMobileSaveReq) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

// UserRegMobileIdReq id
type UserRegMobileIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_reg_mobile_id id
	UserRegMobileId int64 `protobuf:"varint,1,opt,name=user_reg_mobile_id,json=userRegMobileId,proto3" json:"user_reg_mobile_id,omitempty"`
}

func (x *UserRegMobileIdReq) Reset() {
	*x = UserRegMobileIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobileIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobileIdReq) ProtoMessage() {}

func (x *UserRegMobileIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobileIdReq.ProtoReflect.Descriptor instead.
func (*UserRegMobileIdReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegMobileIdReq) GetUserRegMobileId() int64 {
	if x != nil {
		return x.UserRegMobileId
	}
	return 0
}

// UserRegMobileIdsReq ids
type UserRegMobileIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_reg_mobile_ids ids
	UserRegMobileIds []int64 `protobuf:"varint,1,rep,packed,name=user_reg_mobile_ids,json=userRegMobileIds,proto3" json:"user_reg_mobile_ids,omitempty"`
}

func (x *UserRegMobileIdsReq) Reset() {
	*x = UserRegMobileIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobileIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobileIdsReq) ProtoMessage() {}

func (x *UserRegMobileIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobileIdsReq.ProtoReflect.Descriptor instead.
func (*UserRegMobileIdsReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *UserRegMobileIdsReq) GetUserRegMobileIds() []int64 {
	if x != nil {
		return x.UserRegMobileIds
	}
	return nil
}

// UserRegMobileListReq list
type UserRegMobileListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_request paging request
	PageRequest *v1.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (x *UserRegMobileListReq) Reset() {
	*x = UserRegMobileListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobileListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobileListReq) ProtoMessage() {}

func (x *UserRegMobileListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobileListReq.ProtoReflect.Descriptor instead.
func (*UserRegMobileListReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{4}
}

func (x *UserRegMobileListReq) GetPageRequest() *v1.PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

// UserRegMobileListResp list
type UserRegMobileListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list UserRegMobile-array
	List []*UserRegMobile `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// page_info paging response
	PageInfo *v1.PageResponse `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *UserRegMobileListResp) Reset() {
	*x = UserRegMobileListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobileListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobileListResp) ProtoMessage() {}

func (x *UserRegMobileListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobileListResp.ProtoReflect.Descriptor instead.
func (*UserRegMobileListResp) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{5}
}

func (x *UserRegMobileListResp) GetList() []*UserRegMobile {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UserRegMobileListResp) GetPageInfo() *v1.PageResponse {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

// UserRegMobileProcessResult process result
type UserRegMobileProcessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success true or false
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserRegMobileProcessResult) Reset() {
	*x = UserRegMobileProcessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegMobileProcessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegMobileProcessResult) ProtoMessage() {}

func (x *UserRegMobileProcessResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegMobileProcessResult.ProtoReflect.Descriptor instead.
func (*UserRegMobileProcessResult) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP(), []int{6}
}

func (x *UserRegMobileProcessResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto protoreflect.FileDescriptor

var file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x76, 0x31, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73,
	0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x75, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x12, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a,
	0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x14,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x36, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x73, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x42, 0x10,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31,
	0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescOnce sync.Once
	file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescData = file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDesc
)

func file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescGZIP() []byte {
	file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescData)
	})
	return file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDescData
}

var file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_goTypes = []interface{}{
	(*UserRegMobile)(nil),              // 0: service.api.userv1.UserRegMobile
	(*UserRegMobileSaveReq)(nil),       // 1: service.api.userv1.UserRegMobileSaveReq
	(*UserRegMobileIdReq)(nil),         // 2: service.api.userv1.UserRegMobileIdReq
	(*UserRegMobileIdsReq)(nil),        // 3: service.api.userv1.UserRegMobileIdsReq
	(*UserRegMobileListReq)(nil),       // 4: service.api.userv1.UserRegMobileListReq
	(*UserRegMobileListResp)(nil),      // 5: service.api.userv1.UserRegMobileListResp
	(*UserRegMobileProcessResult)(nil), // 6: service.api.userv1.UserRegMobileProcessResult
	(*v1.PageRequest)(nil),             // 7: kit.api.pagev1.PageRequest
	(*v1.PageResponse)(nil),            // 8: kit.api.pagev1.PageResponse
}
var file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_depIdxs = []int32{
	7, // 0: service.api.userv1.UserRegMobileListReq.page_request:type_name -> kit.api.pagev1.PageRequest
	0, // 1: service.api.userv1.UserRegMobileListResp.list:type_name -> service.api.userv1.UserRegMobile
	8, // 2: service.api.userv1.UserRegMobileListResp.page_info:type_name -> kit.api.pagev1.PageResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_init() }
func file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_init() {
	if File_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobile); i {
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
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobileSaveReq); i {
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
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobileIdReq); i {
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
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobileIdsReq); i {
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
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobileListReq); i {
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
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobileListResp); i {
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
		file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegMobileProcessResult); i {
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
			RawDescriptor: file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_msgTypes,
	}.Build()
	File_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto = out.File
	file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_rawDesc = nil
	file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_goTypes = nil
	file_api_user_service_v1_resources_user_reg_mobile_resource_v1_proto_depIdxs = nil
}