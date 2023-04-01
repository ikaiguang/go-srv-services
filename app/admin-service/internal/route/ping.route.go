package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"

	services "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/service"
)

// RegisterPingRoutes 注册路由
func RegisterPingRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) {

	ping := services.NewPingService(logger)
	stdlog.Println("|*** 注册路由：NewPingService")
	pingservicev1.RegisterSrvPingHTTPServer(hs, ping)
	pingservicev1.RegisterSrvPingServer(gs, ping)
}
