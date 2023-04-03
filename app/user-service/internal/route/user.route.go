package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"

	userservicev1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/services"
	services "github.com/ikaiguang/go-srv-services/app/user-service/internal/application/service"
	srvs "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/service"
	datas "github.com/ikaiguang/go-srv-services/app/user-service/internal/infra/data"
	"github.com/ikaiguang/go-srv-services/app/user-service/internal/setup"
)

// RegisterUserRoutes 注册路由
func RegisterUserRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	// 数据库
	dbConn, err := engineHandler.GetPostgresGormDB()
	if err != nil {
		logutil.Fatal(err)
		return
	}
	redisCC, err := engineHandler.GetRedisClient()
	if err != nil {
		logutil.Fatal(err)
		return
	}
	authTokenRepo := engineHandler.GetAuthTokenRepo(redisCC)

	// user
	userRepo := datas.NewUserRepo(dbConn)
	userRegEmailRepo := datas.NewUserRegEmailRepo(dbConn)

	authSrv := srvs.NewUserAuthSrv(
		authTokenRepo,
		logger,
		userRepo,
		userRegEmailRepo,
	)

	// oauth 授权
	userAuthSrv := services.NewUserAuthService(logger, authSrv)
	stdlog.Println("|*** 注册路由：UserAuth")
	userservicev1.RegisterSrvUserAuthHTTPServer(hs, userAuthSrv)
	userservicev1.RegisterSrvUserAuthServer(gs, userAuthSrv)

	return
}
