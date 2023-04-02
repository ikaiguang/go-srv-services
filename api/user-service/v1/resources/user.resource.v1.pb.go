// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: api/user-service/v1/resources/user.resource.v1.proto

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

// User ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户'
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_uuid uuid
	UserUuid string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// deleted_time 删除时间
	DeletedTime string `protobuf:"bytes,5,opt,name=deleted_time,json=deletedTime,proto3" json:"deleted_time,omitempty"`
	// is_deleted 是否已删除
	IsDeleted bool `protobuf:"varint,6,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// user_email 邮箱
	UserEmail string `protobuf:"bytes,7,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// user_nickname 昵称
	UserNickname string `protobuf:"bytes,8,opt,name=user_nickname,json=userNickname,proto3" json:"user_nickname,omitempty"`
	// user_avatar 头像
	UserAvatar string `protobuf:"bytes,9,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	// user_gender 性别
	UserGender string `protobuf:"bytes,10,opt,name=user_gender,json=userGender,proto3" json:"user_gender,omitempty"`
	// user_status 状态
	UserStatus string `protobuf:"bytes,11,opt,name=user_status,json=userStatus,proto3" json:"user_status,omitempty"`
	// active_begin_time 激活开始时间
	ActiveBeginTime string `protobuf:"bytes,12,opt,name=active_begin_time,json=activeBeginTime,proto3" json:"active_begin_time,omitempty"`
	// active_end_time 激活结束时间
	ActiveEndTime string `protobuf:"bytes,13,opt,name=active_end_time,json=activeEndTime,proto3" json:"active_end_time,omitempty"`
	// disable_time 禁用时间
	DisableTime string `protobuf:"bytes,14,opt,name=disable_time,json=disableTime,proto3" json:"disable_time,omitempty"`
	// blacklist_time 黑名单时间
	BlacklistTime string `protobuf:"bytes,15,opt,name=blacklist_time,json=blacklistTime,proto3" json:"blacklist_time,omitempty"`
	// password_hash 密码
	// string password_hash = 16;
	// register_from 注册来源
	RegisterFrom string `protobuf:"bytes,17,opt,name=register_from,json=registerFrom,proto3" json:"register_from,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *User) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *User) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *User) GetDeletedTime() string {
	if x != nil {
		return x.DeletedTime
	}
	return ""
}

func (x *User) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *User) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *User) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *User) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *User) GetUserGender() string {
	if x != nil {
		return x.UserGender
	}
	return ""
}

func (x *User) GetUserStatus() string {
	if x != nil {
		return x.UserStatus
	}
	return ""
}

func (x *User) GetActiveBeginTime() string {
	if x != nil {
		return x.ActiveBeginTime
	}
	return ""
}

func (x *User) GetActiveEndTime() string {
	if x != nil {
		return x.ActiveEndTime
	}
	return ""
}

func (x *User) GetDisableTime() string {
	if x != nil {
		return x.DisableTime
	}
	return ""
}

func (x *User) GetBlacklistTime() string {
	if x != nil {
		return x.BlacklistTime
	}
	return ""
}

func (x *User) GetRegisterFrom() string {
	if x != nil {
		return x.RegisterFrom
	}
	return ""
}

// UserSaveReq save
type UserSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// user_uuid uuid
	UserUuid string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	// created_time 创建时间
	CreatedTime string `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// updated_time 最后修改时间
	UpdatedTime string `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// deleted_time 删除时间
	DeletedTime string `protobuf:"bytes,5,opt,name=deleted_time,json=deletedTime,proto3" json:"deleted_time,omitempty"`
	// is_deleted 是否已删除
	IsDeleted bool `protobuf:"varint,6,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// user_email 邮箱
	UserEmail string `protobuf:"bytes,7,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// user_nickname 昵称
	UserNickname string `protobuf:"bytes,8,opt,name=user_nickname,json=userNickname,proto3" json:"user_nickname,omitempty"`
	// user_avatar 头像
	UserAvatar string `protobuf:"bytes,9,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	// user_gender 性别
	UserGender string `protobuf:"bytes,10,opt,name=user_gender,json=userGender,proto3" json:"user_gender,omitempty"`
	// user_status 状态
	UserStatus string `protobuf:"bytes,11,opt,name=user_status,json=userStatus,proto3" json:"user_status,omitempty"`
	// active_begin_time 激活开始时间
	ActiveBeginTime string `protobuf:"bytes,12,opt,name=active_begin_time,json=activeBeginTime,proto3" json:"active_begin_time,omitempty"`
	// active_end_time 激活结束时间
	ActiveEndTime string `protobuf:"bytes,13,opt,name=active_end_time,json=activeEndTime,proto3" json:"active_end_time,omitempty"`
	// disable_time 禁用时间
	DisableTime string `protobuf:"bytes,14,opt,name=disable_time,json=disableTime,proto3" json:"disable_time,omitempty"`
	// blacklist_time 黑名单时间
	BlacklistTime string `protobuf:"bytes,15,opt,name=blacklist_time,json=blacklistTime,proto3" json:"blacklist_time,omitempty"`
}

func (x *UserSaveReq) Reset() {
	*x = UserSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSaveReq) ProtoMessage() {}

func (x *UserSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSaveReq.ProtoReflect.Descriptor instead.
func (*UserSaveReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *UserSaveReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserSaveReq) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *UserSaveReq) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *UserSaveReq) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *UserSaveReq) GetDeletedTime() string {
	if x != nil {
		return x.DeletedTime
	}
	return ""
}

func (x *UserSaveReq) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *UserSaveReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UserSaveReq) GetUserNickname() string {
	if x != nil {
		return x.UserNickname
	}
	return ""
}

func (x *UserSaveReq) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *UserSaveReq) GetUserGender() string {
	if x != nil {
		return x.UserGender
	}
	return ""
}

func (x *UserSaveReq) GetUserStatus() string {
	if x != nil {
		return x.UserStatus
	}
	return ""
}

func (x *UserSaveReq) GetActiveBeginTime() string {
	if x != nil {
		return x.ActiveBeginTime
	}
	return ""
}

func (x *UserSaveReq) GetActiveEndTime() string {
	if x != nil {
		return x.ActiveEndTime
	}
	return ""
}

func (x *UserSaveReq) GetDisableTime() string {
	if x != nil {
		return x.DisableTime
	}
	return ""
}

func (x *UserSaveReq) GetBlacklistTime() string {
	if x != nil {
		return x.BlacklistTime
	}
	return ""
}

// UserIdReq id
type UserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id id
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserIdReq) Reset() {
	*x = UserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReq.ProtoReflect.Descriptor instead.
func (*UserIdReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// UserIdsReq ids
type UserIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_ids ids
	UserIds []int64 `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *UserIdsReq) Reset() {
	*x = UserIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdsReq) ProtoMessage() {}

func (x *UserIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdsReq.ProtoReflect.Descriptor instead.
func (*UserIdsReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *UserIdsReq) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// UserListReq list
type UserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_request paging request
	PageRequest *v1.PageRequest `protobuf:"bytes,1,opt,name=page_request,json=pageRequest,proto3" json:"page_request,omitempty"`
}

func (x *UserListReq) Reset() {
	*x = UserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReq) ProtoMessage() {}

func (x *UserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListReq.ProtoReflect.Descriptor instead.
func (*UserListReq) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{4}
}

func (x *UserListReq) GetPageRequest() *v1.PageRequest {
	if x != nil {
		return x.PageRequest
	}
	return nil
}

// UserListResp list
type UserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list User-array
	List []*User `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	// page_info paging response
	PageInfo *v1.PageResponse `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *UserListResp) Reset() {
	*x = UserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResp) ProtoMessage() {}

func (x *UserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResp.ProtoReflect.Descriptor instead.
func (*UserListResp) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{5}
}

func (x *UserListResp) GetList() []*User {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *UserListResp) GetPageInfo() *v1.PageResponse {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

// UserProcessResult process result
type UserProcessResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success true or false
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UserProcessResult) Reset() {
	*x = UserProcessResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProcessResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProcessResult) ProtoMessage() {}

func (x *UserProcessResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProcessResult.ProtoReflect.Descriptor instead.
func (*UserProcessResult) Descriptor() ([]byte, []int) {
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP(), []int{6}
}

func (x *UserProcessResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_user_service_v1_resources_user_resource_v1_proto protoreflect.FileDescriptor

var file_api_user_service_v1_resources_user_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6c, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x22, 0x87, 0x04,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x27, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3e, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2d,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x73, 0x0a,
	0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x76, 0x31, 0x42, 0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x72, 0x76, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_service_v1_resources_user_resource_v1_proto_rawDescOnce sync.Once
	file_api_user_service_v1_resources_user_resource_v1_proto_rawDescData = file_api_user_service_v1_resources_user_resource_v1_proto_rawDesc
)

func file_api_user_service_v1_resources_user_resource_v1_proto_rawDescGZIP() []byte {
	file_api_user_service_v1_resources_user_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_user_service_v1_resources_user_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_service_v1_resources_user_resource_v1_proto_rawDescData)
	})
	return file_api_user_service_v1_resources_user_resource_v1_proto_rawDescData
}

var file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_user_service_v1_resources_user_resource_v1_proto_goTypes = []interface{}{
	(*User)(nil),              // 0: service.api.userv1.User
	(*UserSaveReq)(nil),       // 1: service.api.userv1.UserSaveReq
	(*UserIdReq)(nil),         // 2: service.api.userv1.UserIdReq
	(*UserIdsReq)(nil),        // 3: service.api.userv1.UserIdsReq
	(*UserListReq)(nil),       // 4: service.api.userv1.UserListReq
	(*UserListResp)(nil),      // 5: service.api.userv1.UserListResp
	(*UserProcessResult)(nil), // 6: service.api.userv1.UserProcessResult
	(*v1.PageRequest)(nil),    // 7: kit.api.pagev1.PageRequest
	(*v1.PageResponse)(nil),   // 8: kit.api.pagev1.PageResponse
}
var file_api_user_service_v1_resources_user_resource_v1_proto_depIdxs = []int32{
	7, // 0: service.api.userv1.UserListReq.page_request:type_name -> kit.api.pagev1.PageRequest
	0, // 1: service.api.userv1.UserListResp.list:type_name -> service.api.userv1.User
	8, // 2: service.api.userv1.UserListResp.page_info:type_name -> kit.api.pagev1.PageResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_user_service_v1_resources_user_resource_v1_proto_init() }
func file_api_user_service_v1_resources_user_resource_v1_proto_init() {
	if File_api_user_service_v1_resources_user_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSaveReq); i {
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
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdReq); i {
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
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdsReq); i {
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
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListReq); i {
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
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResp); i {
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
		file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProcessResult); i {
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
			RawDescriptor: file_api_user_service_v1_resources_user_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_user_service_v1_resources_user_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_user_service_v1_resources_user_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_user_service_v1_resources_user_resource_v1_proto_msgTypes,
	}.Build()
	File_api_user_service_v1_resources_user_resource_v1_proto = out.File
	file_api_user_service_v1_resources_user_resource_v1_proto_rawDesc = nil
	file_api_user_service_v1_resources_user_resource_v1_proto_goTypes = nil
	file_api_user_service_v1_resources_user_resource_v1_proto_depIdxs = nil
}