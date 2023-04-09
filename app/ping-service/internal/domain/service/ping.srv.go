package srvs

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

// PingSrv ...
type PingSrv struct {
	log *log.Helper

	adminPingHTTPClient pingservicev1.SrvPingHTTPClient
	userPingHTTPClient  pingservicev1.SrvPingHTTPClient
	adminPingGRPCClient pingservicev1.SrvPingClient
	userPingGRPCClient  pingservicev1.SrvPingClient
}

// NewPingSrv ...
func NewPingSrv(
	logger log.Logger,
	adminPingHTTPClient, userPingHTTPClient pingservicev1.SrvPingHTTPClient,
	adminPingGRPCClient, userPingGRPCClient pingservicev1.SrvPingClient,
) *PingSrv {
	return &PingSrv{
		log: log.NewHelper(log.With(logger, "module", "ping/domain/service/ping")),

		adminPingHTTPClient: adminPingHTTPClient,
		userPingHTTPClient:  userPingHTTPClient,
		adminPingGRPCClient: adminPingGRPCClient,
		userPingGRPCClient:  userPingGRPCClient,
	}
}

// PingHTTP ...
func (s *PingSrv) PingHTTP(ctx context.Context) error {
	adminPingReq := &pingv1.PingReq{Message: "[GRPC] admin ping message"}
	adminPingResp, err := s.adminPingHTTPClient.Ping(ctx, adminPingReq)
	if err != nil {
		return err
	}
	s.log.Info("[HTTP] ping admin res: ", adminPingResp.Message)

	userPingReq := &pingv1.PingReq{Message: "[GRPC] user ping message"}
	userPingResp, err := s.userPingHTTPClient.Ping(ctx, userPingReq)
	if err != nil {
		return err
	}
	s.log.Info("[HTTP] ping user res: ", userPingResp.Message)

	return nil
}

// PingGRPC ...
func (s *PingSrv) PingGRPC(ctx context.Context) error {
	adminPingReq := &pingv1.PingReq{Message: "[GRPC] admin ping message"}
	adminPingResp, err := s.adminPingGRPCClient.Ping(ctx, adminPingReq)
	if err != nil {
		return err
	}
	s.log.Info("[GRPC] ping admin res: ", adminPingResp.Message)

	userPingReq := &pingv1.PingReq{Message: "[GRPC] user ping message"}
	userPingResp, err := s.userPingGRPCClient.Ping(ctx, userPingReq)
	if err != nil {
		return err
	}
	s.log.Info("[GRPC] ping user res: ", userPingResp.Message)

	return nil
}
