package routes

import (
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"

	services "github.com/ikaiguang/go-srv-services/app/ping-service/internal/application/service"
	srvs "github.com/ikaiguang/go-srv-services/app/ping-service/internal/domain/service"
	"github.com/ikaiguang/go-srv-services/app/ping-service/internal/setup"
	clientutil "github.com/ikaiguang/go-srv-services/business-kit/client"
	serviceutil "github.com/ikaiguang/go-srv-services/business-kit/service"
)

// RegisterPingRoutes 注册路由
func RegisterPingRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	adminPingHTTPClient, err := clientutil.NewAdminPingHTTPClient(engineHandler, serviceutil.AdminService)
	if err != nil {
		logutil.Fatal(err)
		return
	}
	adminPingGRPCClient, err := clientutil.NewAdminPingGRPCClient(engineHandler, serviceutil.AdminService)
	if err != nil {
		logutil.Fatal(err)
		return
	}

	userPingHTTPClient, err := clientutil.NewUserPingHTTPClient(engineHandler, serviceutil.UserService)
	if err != nil {
		logutil.Fatal(err)
		return
	}
	userPingGRPCClient, err := clientutil.NewUserPingGRPCClient(engineHandler, serviceutil.UserService)
	if err != nil {
		logutil.Fatal(err)
		return
	}

	pingSrv := srvs.NewPingSrv(
		logger,
		adminPingHTTPClient, userPingHTTPClient,
		adminPingGRPCClient, userPingGRPCClient,
	)
	ping := services.NewPingService(logger, pingSrv)
	stdlog.Println("|*** 注册路由：NewPingService")
	pingservicev1.RegisterSrvPingHTTPServer(hs, ping)
	pingservicev1.RegisterSrvPingServer(gs, ping)
}
