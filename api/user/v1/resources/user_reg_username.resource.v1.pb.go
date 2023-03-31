// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: api/user/v1/resources/user_reg_username.resource.v1.proto

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

// UserRegUsername ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户-注册-用户名'
type UserRegUsername struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_name 用户名
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// user_uuid uuid
	UserUuid string `protobuf:"bytes,5,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *UserRegUsername) Reset() {
	*x = UserRegUsername{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsername) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsername) ProtoMessage() {}

func (x *UserRegUsername) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsername.ProtoReflect.Descriptor instead.
func (*UserRegUsername) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegUsername) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRegUsername) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserRegUsername) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *UserRegUsername) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *UserRegUsername) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

// UserRegUsernameSaveReq save
type UserRegUsernameSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_name 用户名
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// user_uuid uuid
	UserUuid string `protobuf:"bytes,5,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *UserRegUsernameSaveReq) Reset() {
	*x = UserRegUsernameSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsernameSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsernameSaveReq) ProtoMessage() {}

func (x *UserRegUsernameSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsernameSaveReq.ProtoReflect.Descriptor instead.
func (*UserRegUsernameSaveReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegUsernameSaveReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserRegUsernameSaveReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserRegUsernameSaveReq) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *UserRegUsernameSaveReq) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *UserRegUsernameSaveReq) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

// UserRegUsernameIdReq id
type UserRegUsernameIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_reg_username_id id
	UserRegUsernameId int64 `protobuf:"varint,1,opt,name=user_reg_username_id,json=userRegUsernameId,proto3" json:"user_reg_username_id,omitempty"`
}

func (x *UserRegUsernameIdReq) Reset() {
	*x = UserRegUsernameIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsernameIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsernameIdReq) ProtoMessage() {}

func (x *UserRegUsernameIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsernameIdReq.ProtoReflect.Descriptor instead.
func (*UserRegUsernameIdReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegUsernameIdReq) GetUserRegUsernameId() int64 {
	if x != nil {
		return x.UserRegUsernameId
	}
	return 0
}

// UserRegUsernameIdsReq ids
type UserRegUsernameIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_reg_username_ids ids
	UserRegUsernameIds []int64 `protobuf:"varint,1,rep,packed,name=user_reg_username_ids,json=userRegUsernameIds,proto3" json:"user_reg_username_ids,omitempty"`
}

func (x *UserRegUsernameIdsReq) Reset() {
	*x = UserRegUsernameIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsernameIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsernameIdsReq) ProtoMessage() {}

func (x *UserRegUsernameIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsernameIdsReq.ProtoReflect.Descriptor instead.
func (*UserRegUsernameIdsReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *UserRegUsernameIdsReq) GetUserRegUsernameIds() []int64 {
	if x != nil {
		return x.UserRegUsernameIds
	}
	return nil
}

// UserRegUsernameListReq list
type UserRegUsernameListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_request paging request
	PageRequest *v1.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (x *UserRegUsernameListReq) Reset() {
	*x = UserRegUsernameListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsernameListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsernameListReq) ProtoMessage() {}

func (x *UserRegUsernameListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsernameListReq.ProtoReflect.Descriptor instead.
func (*UserRegUsernameListReq) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{4}
}

func (x *UserRegUsernameListReq) GetPageRequest() *v1.PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

// UserRegUsernameListResp list
type UserRegUsernameListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list UserRegUsername-array
	List []*UserRegUsername `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// page_info paging response
	PageInfo *v1.PageResponse `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *UserRegUsernameListResp) Reset() {
	*x = UserRegUsernameListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsernameListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsernameListResp) ProtoMessage() {}

func (x *UserRegUsernameListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsernameListResp.ProtoReflect.Descriptor instead.
func (*UserRegUsernameListResp) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{5}
}

func (x *UserRegUsernameListResp) GetList() []*UserRegUsername {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UserRegUsernameListResp) GetPageInfo() *v1.PageResponse {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

// UserRegUsernameProcessResult process result
type UserRegUsernameProcessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success true or false
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserRegUsernameProcessResult) Reset() {
	*x = UserRegUsernameProcessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegUsernameProcessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegUsernameProcessResult) ProtoMessage() {}

func (x *UserRegUsernameProcessResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegUsernameProcessResult.ProtoReflect.Descriptor instead.
func (*UserRegUsernameProcessResult) Descriptor() ([]byte, []int) {
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP(), []int{6}
}

func (x *UserRegUsernameProcessResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_user_v1_resources_user_reg_username_resource_v1_proto protoreflect.FileDescriptor

var file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x39, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x1a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69,
	0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0f, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0xa8,
	0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x2f, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x22, 0x4a, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x15, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x12, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x73, 0x22, 0x58,
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x38, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x6b, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x42, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x70, 0x69, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61,
	0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescOnce sync.Once
	file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescData = file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDesc
)

func file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescGZIP() []byte {
	file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescData)
	})
	return file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDescData
}

var file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_user_v1_resources_user_reg_username_resource_v1_proto_goTypes = []interface{}{
	(*UserRegUsername)(nil),              // 0: service.api.userv1.UserRegUsername
	(*UserRegUsernameSaveReq)(nil),       // 1: service.api.userv1.UserRegUsernameSaveReq
	(*UserRegUsernameIdReq)(nil),         // 2: service.api.userv1.UserRegUsernameIdReq
	(*UserRegUsernameIdsReq)(nil),        // 3: service.api.userv1.UserRegUsernameIdsReq
	(*UserRegUsernameListReq)(nil),       // 4: service.api.userv1.UserRegUsernameListReq
	(*UserRegUsernameListResp)(nil),      // 5: service.api.userv1.UserRegUsernameListResp
	(*UserRegUsernameProcessResult)(nil), // 6: service.api.userv1.UserRegUsernameProcessResult
	(*v1.PageRequest)(nil),               // 7: kit.api.pagev1.PageRequest
	(*v1.PageResponse)(nil),              // 8: kit.api.pagev1.PageResponse
}
var file_api_user_v1_resources_user_reg_username_resource_v1_proto_depIdxs = []int32{
	7, // 0: service.api.userv1.UserRegUsernameListReq.page_request:type_name -> kit.api.pagev1.PageRequest
	0, // 1: service.api.userv1.UserRegUsernameListResp.list:type_name -> service.api.userv1.UserRegUsername
	8, // 2: service.api.userv1.UserRegUsernameListResp.page_info:type_name -> kit.api.pagev1.PageResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_user_v1_resources_user_reg_username_resource_v1_proto_init() }
func file_api_user_v1_resources_user_reg_username_resource_v1_proto_init() {
	if File_api_user_v1_resources_user_reg_username_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsername); i {
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
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsernameSaveReq); i {
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
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsernameIdReq); i {
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
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsernameIdsReq); i {
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
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsernameListReq); i {
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
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsernameListResp); i {
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
		file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegUsernameProcessResult); i {
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
			RawDescriptor: file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_v1_resources_user_reg_username_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_user_v1_resources_user_reg_username_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_user_v1_resources_user_reg_username_resource_v1_proto_msgTypes,
	}.Build()
	File_api_user_v1_resources_user_reg_username_resource_v1_proto = out.File
	file_api_user_v1_resources_user_reg_username_resource_v1_proto_rawDesc = nil
	file_api_user_v1_resources_user_reg_username_resource_v1_proto_goTypes = nil
	file_api_user_v1_resources_user_reg_username_resource_v1_proto_depIdxs = nil
}
