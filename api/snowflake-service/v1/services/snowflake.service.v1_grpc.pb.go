// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/snowflake-service/v1/services/snowflake.service.v1.proto

package snowflakeservicev1

import (
	context "context"
	resources "github.com/ikaiguang/go-srv-services/api/snowflake-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SrvSnowflakeNodeClient is the client API for SrvSnowflakeNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvSnowflakeNodeClient interface {
	GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...grpc.CallOption) (*resources.GetNodeIdResp, error)
	ExtendNodeId(ctx context.Context, in *resources.ExtendNodeIdReq, opts ...grpc.CallOption) (*resources.ExtendNodeIdResp, error)
}

type srvSnowflakeNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvSnowflakeNodeClient(cc grpc.ClientConnInterface) SrvSnowflakeNodeClient {
	return &srvSnowflakeNodeClient{cc}
}

func (c *srvSnowflakeNodeClient) GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...grpc.CallOption) (*resources.GetNodeIdResp, error) {
	out := new(resources.GetNodeIdResp)
	err := c.cc.Invoke(ctx, "/service.api.snowflakeservicev1.SrvSnowflakeNode/GetNodeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvSnowflakeNodeClient) ExtendNodeId(ctx context.Context, in *resources.ExtendNodeIdReq, opts ...grpc.CallOption) (*resources.ExtendNodeIdResp, error) {
	out := new(resources.ExtendNodeIdResp)
	err := c.cc.Invoke(ctx, "/service.api.snowflakeservicev1.SrvSnowflakeNode/ExtendNodeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvSnowflakeNodeServer is the server API for SrvSnowflakeNode service.
// All implementations must embed UnimplementedSrvSnowflakeNodeServer
// for forward compatibility
type SrvSnowflakeNodeServer interface {
	GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error)
	ExtendNodeId(context.Context, *resources.ExtendNodeIdReq) (*resources.ExtendNodeIdResp, error)
	mustEmbedUnimplementedSrvSnowflakeNodeServer()
}

// UnimplementedSrvSnowflakeNodeServer must be embedded to have forward compatible implementations.
type UnimplementedSrvSnowflakeNodeServer struct {
}

func (UnimplementedSrvSnowflakeNodeServer) GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeId not implemented")
}
func (UnimplementedSrvSnowflakeNodeServer) ExtendNodeId(context.Context, *resources.ExtendNodeIdReq) (*resources.ExtendNodeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtendNodeId not implemented")
}
func (UnimplementedSrvSnowflakeNodeServer) mustEmbedUnimplementedSrvSnowflakeNodeServer() {}

// UnsafeSrvSnowflakeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvSnowflakeNodeServer will
// result in compilation errors.
type UnsafeSrvSnowflakeNodeServer interface {
	mustEmbedUnimplementedSrvSnowflakeNodeServer()
}

func RegisterSrvSnowflakeNodeServer(s grpc.ServiceRegistrar, srv SrvSnowflakeNodeServer) {
	s.RegisterService(&SrvSnowflakeNode_ServiceDesc, srv)
}

func _SrvSnowflakeNode_GetNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvSnowflakeNodeServer).GetNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.api.snowflakeservicev1.SrvSnowflakeNode/GetNodeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvSnowflakeNodeServer).GetNodeId(ctx, req.(*resources.GetNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvSnowflakeNode_ExtendNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ExtendNodeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvSnowflakeNodeServer).ExtendNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.api.snowflakeservicev1.SrvSnowflakeNode/ExtendNodeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvSnowflakeNodeServer).ExtendNodeId(ctx, req.(*resources.ExtendNodeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvSnowflakeNode_ServiceDesc is the grpc.ServiceDesc for SrvSnowflakeNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvSnowflakeNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.api.snowflakeservicev1.SrvSnowflakeNode",
	HandlerType: (*SrvSnowflakeNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeId",
			Handler:    _SrvSnowflakeNode_GetNodeId_Handler,
		},
		{
			MethodName: "ExtendNodeId",
			Handler:    _SrvSnowflakeNode_ExtendNodeId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/snowflake-service/v1/services/snowflake.service.v1.proto",
}
