package clientutil

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"

	userservicev1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/services"
	serviceutil "github.com/ikaiguang/go-srv-services/business/service"
	setuppkg "github.com/ikaiguang/go-srv-services/pkg/setup"
)

// NewUserPingHTTPClient ...
func NewUserPingHTTPClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(conn), nil
}

// NewUserPingGRPCClient ...
func NewUserPingGRPCClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(conn), nil
}

// NewUserAuthHTTPClient ...
func NewUserAuthHTTPClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (userservicev1.SrvUserAuthHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return userservicev1.NewSrvUserAuthHTTPClient(conn), nil
}

// NewUserAuthGRPCClient ...
func NewUserAuthGRPCClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (userservicev1.SrvUserAuthClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return userservicev1.NewSrvUserAuthClient(conn), nil
}
