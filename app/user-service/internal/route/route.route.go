package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/ikaiguang/go-srv-services/app/user-service/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// root
	RegisterRootRoutes(engineHandler, hs, gs)

	// websocket
	//RegisterWssRoutes(engineHandler, hs, gs)

	// testdata
	RegisterPingRoutes(engineHandler, hs, gs)
	//RegisterTestdataRoutes(engineHandler, hs, gs)

	// user
	RegisterUserRoutes(engineHandler, hs, gs)

	return err
}
