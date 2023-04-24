package clientutil

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	"github.com/ikaiguang/go-srv-services/business-kit/setup"

	userservicev1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/services"
	serviceutil "github.com/ikaiguang/go-srv-services/business-kit/service"
)

// NewUserPingHTTPClient ...
func NewUserPingHTTPClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingHTTPClient(conn), nil
}

// NewUserPingGRPCClient ...
func NewUserPingGRPCClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (pingservicev1.SrvPingClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return pingservicev1.NewSrvPingClient(conn), nil
}

// NewUserAuthHTTPClient ...
func NewUserAuthHTTPClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (userservicev1.SrvUserAuthHTTPClient, error) {
	conn, err := NewHTTPConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return userservicev1.NewSrvUserAuthHTTPClient(conn), nil
}

// NewUserAuthGRPCClient ...
func NewUserAuthGRPCClient(engineHandler setuputil.Engine, serviceName serviceutil.ServiceName) (userservicev1.SrvUserAuthClient, error) {
	conn, err := NewGRPCConnection(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	return userservicev1.NewSrvUserAuthClient(conn), nil
}
