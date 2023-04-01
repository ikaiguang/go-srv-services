package routes

import (
	pkgerrors "github.com/pkg/errors"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/ikaiguang/go-srv-services/app/user-service/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		return err
	}

	// root
	err = RegisterRootRoutes(hs, gs, logger)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	// websocket
	//err = RegisterWssRoutes(hs, gs, logger)
	//if err != nil {
	//	return pkgerrors.WithStack(err)
	//}

	// testdata
	RegisterPingRoutes(hs, gs, logger)
	//RegisterTestdataRoutes(hs, gs, logger)

	// user
	if err = RegisterUserRoutes(engineHandler, hs, gs); err != nil {
		return pkgerrors.WithStack(err)
	}

	return err
}
