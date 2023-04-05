package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"

	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/services"
	assemblers "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/assembler"
	services "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/service"
	srvs "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/service"
	datas "github.com/ikaiguang/go-srv-services/app/admin-service/internal/infra/data"
	"github.com/ikaiguang/go-srv-services/app/admin-service/internal/setup"
)

// RegisterAdminRoutes 注册路由
func RegisterAdminRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
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

	// assembler
	assembler := assemblers.NewAssembler()

	// admin
	adminRepo := datas.NewAdminRepo(dbConn)
	adminRegEmailRepo := datas.NewAdminRegEmailRepo(dbConn)

	// srv
	authSrv := srvs.NewAdminAuthSrv(
		authTokenRepo,
		logger,
		adminRepo,
		adminRegEmailRepo,
	)

	// oauth 授权
	adminAuthSrv := services.NewAdminAuthService(
		logger,
		assembler,
		authSrv,
	)
	stdlog.Println("|*** 注册路由：AdminAuth")
	adminservicev1.RegisterSrvAdminAuthHTTPServer(hs, adminAuthSrv)
	adminservicev1.RegisterSrvAdminAuthServer(gs, adminAuthSrv)

	return
}
