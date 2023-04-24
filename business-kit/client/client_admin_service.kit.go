package clientutil

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	"github.com/ikaiguang/go-srv-services/business-kit/setup"

	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/services"
	serviceutil "github.com/ikaiguang/go-srv-services/business-kit/service"
)

// NewAdminPingHTTPClient ...
func NewAdminPingHTTPClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(conn), nil
}

// NewAdminPingGRPCClient ...
func NewAdminPingGRPCClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(conn), nil
}

// NewAdminAuthHTTPClient ...
func NewAdminAuthHTTPClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (adminservicev1.SrvAdminAuthHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return adminservicev1.NewSrvAdminAuthHTTPClient(conn), nil
}

// NewAdminAuthGRPCClient ...
func NewAdminAuthGRPCClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (adminservicev1.SrvAdminAuthClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return adminservicev1.NewSrvAdminAuthClient(conn), nil
}
