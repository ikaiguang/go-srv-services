package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdlog "log"

	userservicev1 "github.com/ikaiguang/go-srv-services/api/user/v1/services"
	services "github.com/ikaiguang/go-srv-services/app/user-service/internal/application/service"
	datas "github.com/ikaiguang/go-srv-services/app/user-service/internal/infra/data"
	"github.com/ikaiguang/go-srv-services/app/user-service/internal/setup"
)

// RegisterUserRoutes 注册路由
func RegisterUserRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	// 数据库
	dbConn, err := engineHandler.GetPostgresGormDB()
	if err != nil {
		return err
	}
	redisCC, err := engineHandler.GetRedisClient()
	if err != nil {
		return err
	}
	authTokenRepo := engineHandler.GetAuthTokenRepo(redisCC)

	// user
	userRepo := datas.NewUserRepo(dbConn)
	userRegEmailRepo := datas.NewUserRegEmailRepo(dbConn)

	// oauth 授权
	userAuthSrv := services.NewUserAuthService(
		authTokenRepo,
		userRepo,
		userRegEmailRepo,
	)
	stdlog.Println("|*** 注册路由：UserAuth")
	userservicev1.RegisterSrvUserAuthHTTPServer(hs, userAuthSrv)
	userservicev1.RegisterSrvUserAuthServer(gs, userAuthSrv)

	return err
}
