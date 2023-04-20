// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package snowflakeservicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/ikaiguang/go-srv-services/api/snowflake-service/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvSnowflakeNodeExtendNodeId = "/service.api.snowflakeservicev1.SrvSnowflakeNode/ExtendNodeId"
const OperationSrvSnowflakeNodeGetNodeId = "/service.api.snowflakeservicev1.SrvSnowflakeNode/GetNodeId"

type SrvSnowflakeNodeHTTPServer interface {
	ExtendNodeId(context.Context, *resources.ExtendNodeIdReq) (*resources.ExtendNodeIdResp, error)
	GetNodeId(context.Context, *resources.GetNodeIdReq) (*resources.GetNodeIdResp, error)
}

func RegisterSrvSnowflakeNodeHTTPServer(s *http.Server, srv SrvSnowflakeNodeHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/snowflake/nodeid/get", _SrvSnowflakeNode_GetNodeId0_HTTP_Handler(srv))
	r.POST("/api/v1/snowflake/nodeid/extend", _SrvSnowflakeNode_ExtendNodeId0_HTTP_Handler(srv))
}

func _SrvSnowflakeNode_GetNodeId0_HTTP_Handler(srv SrvSnowflakeNodeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.GetNodeIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvSnowflakeNodeGetNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetNodeId(ctx, req.(*resources.GetNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.GetNodeIdResp)
		return ctx.Result(200, reply)
	}
}

func _SrvSnowflakeNode_ExtendNodeId0_HTTP_Handler(srv SrvSnowflakeNodeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.ExtendNodeIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvSnowflakeNodeExtendNodeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExtendNodeId(ctx, req.(*resources.ExtendNodeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.ExtendNodeIdResp)
		return ctx.Result(200, reply)
	}
}

type SrvSnowflakeNodeHTTPClient interface {
	ExtendNodeId(ctx context.Context, req *resources.ExtendNodeIdReq, opts ...http.CallOption) (rsp *resources.ExtendNodeIdResp, err error)
	GetNodeId(ctx context.Context, req *resources.GetNodeIdReq, opts ...http.CallOption) (rsp *resources.GetNodeIdResp, err error)
}

type SrvSnowflakeNodeHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvSnowflakeNodeHTTPClient(client *http.Client) SrvSnowflakeNodeHTTPClient {
	return &SrvSnowflakeNodeHTTPClientImpl{client}
}

func (c *SrvSnowflakeNodeHTTPClientImpl) ExtendNodeId(ctx context.Context, in *resources.ExtendNodeIdReq, opts ...http.CallOption) (*resources.ExtendNodeIdResp, error) {
	var out resources.ExtendNodeIdResp
	pattern := "/api/v1/snowflake/nodeid/extend"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvSnowflakeNodeExtendNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvSnowflakeNodeHTTPClientImpl) GetNodeId(ctx context.Context, in *resources.GetNodeIdReq, opts ...http.CallOption) (*resources.GetNodeIdResp, error) {
	var out resources.GetNodeIdResp
	pattern := "/api/v1/snowflake/nodeid/get"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvSnowflakeNodeGetNodeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
