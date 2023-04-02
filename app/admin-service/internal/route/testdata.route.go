package routes

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"

	services "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/service"
	srvs "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/service"
)

// RegisterTestdataRoutes 注册路由
func RegisterTestdataRoutes(hs *http.Server, gs *grpc.Server, logger log.Logger) {

	wssSrv := srvs.NewWebsocketSrv(logger)
	testdata := services.NewTestdataService(logger, wssSrv)
	stdlog.Println("|*** 注册路由：NewTestdataService")
	testdataservicev1.RegisterSrvTestdataHTTPServer(hs, testdata)
	testdataservicev1.RegisterSrvTestdataServer(gs, testdata)
}
