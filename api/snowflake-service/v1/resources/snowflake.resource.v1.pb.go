// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/snowflake-service/v1/resources/snowflake.resource.v1.proto

package snowflakev1

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

// SnowflakeNodeID ENGINE InnoDB CHARSET utf8mb4 COMMENT '雪花算法节点ID'
type SnowflakeNodeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id 数据库ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// instance_launch_time 实例启动时间
	//
	//	string instance_launch_time = 2;
	//
	// instance_extend_time 实例续期时间
	InstanceExtendTime string `protobuf:"bytes,3,opt,name=instance_extend_time,json=instanceExtendTime,proto3" json:"instance_extend_time,omitempty"`
	// instance_id 实例ID
	InstanceId string `protobuf:"bytes,4,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// node_id 雪花算法节点id
	NodeId int64 `protobuf:"varint,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// instance_name 实例名称
	InstanceName string `protobuf:"bytes,6,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	// instance_endpoint_list 实例端点数组
	InstanceEndpointList string `protobuf:"bytes,7,opt,name=instance_endpoint_list,json=instanceEndpointList,proto3" json:"instance_endpoint_list,omitempty"`
	// instance_metadata 实例元数据
	InstanceMetadata string `protobuf:"bytes,8,opt,name=instance_metadata,json=instanceMetadata,proto3" json:"instance_metadata,omitempty"`
}

func (x *SnowflakeNodeID) Reset() {
	*x = SnowflakeNodeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnowflakeNodeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakeNodeID) ProtoMessage() {}

func (x *SnowflakeNodeID) ProtoReflect() protoreflect.Message {
	mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakeNodeID.ProtoReflect.Descriptor instead.
func (*SnowflakeNodeID) Descriptor() ([]byte, []int) {
	return file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *SnowflakeNodeID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SnowflakeNodeID) GetInstanceExtendTime() string {
	if x != nil {
		return x.InstanceExtendTime
	}
	return ""
}

func (x *SnowflakeNodeID) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *SnowflakeNodeID) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *SnowflakeNodeID) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *SnowflakeNodeID) GetInstanceEndpointList() string {
	if x != nil {
		return x.InstanceEndpointList
	}
	return ""
}

func (x *SnowflakeNodeID) GetInstanceMetadata() string {
	if x != nil {
		return x.InstanceMetadata
	}
	return ""
}

// GetNodeIdReq 获取节点请求
type GetNodeIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// instance_id 实例ID
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// instance_name 实例名称
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	// endpoints 实例端点列表
	Endpoints []string `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// metadata 实例元数据
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetNodeIdReq) Reset() {
	*x = GetNodeIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeIdReq) ProtoMessage() {}

func (x *GetNodeIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeIdReq.ProtoReflect.Descriptor instead.
func (*GetNodeIdReq) Descriptor() ([]byte, []int) {
	return file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodeIdReq) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GetNodeIdReq) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *GetNodeIdReq) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *GetNodeIdReq) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ExtendNodeIdReq 节点需求请求
type ExtendNodeIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id 数据库ID
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// instance_id 实例ID
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// node_id 雪花算法节点id
	NodeId int64 `protobuf:"varint,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *ExtendNodeIdReq) Reset() {
	*x = ExtendNodeIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendNodeIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendNodeIdReq) ProtoMessage() {}

func (x *ExtendNodeIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendNodeIdReq.ProtoReflect.Descriptor instead.
func (*ExtendNodeIdReq) Descriptor() ([]byte, []int) {
	return file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ExtendNodeIdReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExtendNodeIdReq) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *ExtendNodeIdReq) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

// ExtendNodeIdResp process result
type ExtendNodeIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// success true or false
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ExtendNodeIdResp) Reset() {
	*x = ExtendNodeIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendNodeIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendNodeIdResp) ProtoMessage() {}

func (x *ExtendNodeIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendNodeIdResp.ProtoReflect.Descriptor instead.
func (*ExtendNodeIdResp) Descriptor() ([]byte, []int) {
	return file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ExtendNodeIdResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_snowflake_service_v1_resources_snowflake_resource_v1_proto protoreflect.FileDescriptor

var file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x76, 0x31, 0x22, 0x95, 0x02, 0x0a, 0x0f, 0x53, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a,
	0x14, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34,
	0x0a, 0x16, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x80, 0x02, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x22, 0x2c, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x87, 0x01, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x76, 0x31, 0x42, 0x15, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72,
	0x76, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x6e,
	0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescOnce sync.Once
	file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescData = file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDesc
)

func file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescGZIP() []byte {
	file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescData)
	})
	return file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDescData
}

var file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_goTypes = []interface{}{
	(*SnowflakeNodeID)(nil),  // 0: service.api.snowflakev1.SnowflakeNodeID
	(*GetNodeIdReq)(nil),     // 1: service.api.snowflakev1.GetNodeIdReq
	(*ExtendNodeIdReq)(nil),  // 2: service.api.snowflakev1.ExtendNodeIdReq
	(*ExtendNodeIdResp)(nil), // 3: service.api.snowflakev1.ExtendNodeIdResp
	nil,                      // 4: service.api.snowflakev1.GetNodeIdReq.MetadataEntry
}
var file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_depIdxs = []int32{
	4, // 0: service.api.snowflakev1.GetNodeIdReq.metadata:type_name -> service.api.snowflakev1.GetNodeIdReq.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_init() }
func file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_init() {
	if File_api_snowflake_service_v1_resources_snowflake_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnowflakeNodeID); i {
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
		file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeIdReq); i {
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
		file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendNodeIdReq); i {
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
		file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendNodeIdResp); i {
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
			RawDescriptor: file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_msgTypes,
	}.Build()
	File_api_snowflake_service_v1_resources_snowflake_resource_v1_proto = out.File
	file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_rawDesc = nil
	file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_goTypes = nil
	file_api_snowflake_service_v1_resources_snowflake_resource_v1_proto_depIdxs = nil
}
