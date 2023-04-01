package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdlog "log"

	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin/v1/services"
	services "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/service"
	datas "github.com/ikaiguang/go-srv-services/app/admin-service/internal/infra/data"
	"github.com/ikaiguang/go-srv-services/app/admin-service/internal/setup"
)

// RegisterAdminRoutes 注册路由
func RegisterAdminRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
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

	// admin
	adminRepo := datas.NewAdminRepo(dbConn)
	adminRegEmailRepo := datas.NewAdminRegEmailRepo(dbConn)

	// oauth 授权
	adminAuthSrv := services.NewAdminAuthService(
		authTokenRepo,
		adminRepo,
		adminRegEmailRepo,
	)
	stdlog.Println("|*** 注册路由：AdminAuth")
	adminservicev1.RegisterSrvAdminAuthHTTPServer(hs, adminAuthSrv)
	adminservicev1.RegisterSrvAdminAuthServer(gs, adminAuthSrv)

	return err
}
