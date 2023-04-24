// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: api/user-service/v1/services/user_auth.service.v1.proto

package userservicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	resources "github.com/ikaiguang/go-srv-services/api/user-service/v1/resources"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSrvUserAuthLoginByEmail = "/service.api.userservicev1.SrvUserAuth/LoginByEmail"
const OperationSrvUserAuthPing = "/service.api.userservicev1.SrvUserAuth/Ping"

type SrvUserAuthHTTPServer interface {
	// LoginByEmail LoginByEmail Email登录
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	// Ping Ping ping pong
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
}

func RegisterSrvUserAuthHTTPServer(s *http.Server, srv SrvUserAuthHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/user-auth/login-by-email", _SrvUserAuth_LoginByEmail0_HTTP_Handler(srv))
	r.GET("/api/v1/user/user-auth/ping", _SrvUserAuth_Ping0_HTTP_Handler(srv))
}

func _SrvUserAuth_LoginByEmail0_HTTP_Handler(srv SrvUserAuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.LoginByEmailReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthLoginByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.LoginResp)
		return ctx.Result(200, reply)
	}
}

func _SrvUserAuth_Ping0_HTTP_Handler(srv SrvUserAuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.PingReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSrvUserAuthPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*resources.PingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.PingResp)
		return ctx.Result(200, reply)
	}
}

type SrvUserAuthHTTPClient interface {
	LoginByEmail(ctx context.Context, req *resources.LoginByEmailReq, opts ...http.CallOption) (rsp *resources.LoginResp, err error)
	Ping(ctx context.Context, req *resources.PingReq, opts ...http.CallOption) (rsp *resources.PingResp, err error)
}

type SrvUserAuthHTTPClientImpl struct {
	cc *http.Client
}

func NewSrvUserAuthHTTPClient(client *http.Client) SrvUserAuthHTTPClient {
	return &SrvUserAuthHTTPClientImpl{client}
}

func (c *SrvUserAuthHTTPClientImpl) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...http.CallOption) (*resources.LoginResp, error) {
	var out resources.LoginResp
	pattern := "/api/v1/user/user-auth/login-by-email"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSrvUserAuthLoginByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SrvUserAuthHTTPClientImpl) Ping(ctx context.Context, in *resources.PingReq, opts ...http.CallOption) (*resources.PingResp, error) {
	var out resources.PingResp
	pattern := "/api/v1/user/user-auth/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSrvUserAuthPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
