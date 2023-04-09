package clientutil

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"

	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/services"
	serviceutil "github.com/ikaiguang/go-srv-services/business/service"
	setuppkg "github.com/ikaiguang/go-srv-services/pkg/setup"
)

// NewAdminPingHTTPClient ...
func NewAdminPingHTTPClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(conn), nil
}

// NewAdminPingGRPCClient ...
func NewAdminPingGRPCClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(conn), nil
}

// NewAdminAuthHTTPClient ...
func NewAdminAuthHTTPClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (adminservicev1.SrvAdminAuthHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return adminservicev1.NewSrvAdminAuthHTTPClient(conn), nil
}

// NewAdminAuthGRPCClient ...
func NewAdminAuthGRPCClient(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) (adminservicev1.SrvAdminAuthClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return adminservicev1.NewSrvAdminAuthClient(conn), nil
}
