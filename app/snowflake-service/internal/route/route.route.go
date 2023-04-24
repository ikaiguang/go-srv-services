package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/setup"
	stdlog "log"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// snowflake
	RegisterSnowflakeRoutes(engineHandler, hs, gs)

	return err
}
