package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"

	services "github.com/ikaiguang/go-srv-services/app/user-service/internal/application/service"
	srvs "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/service"
	"github.com/ikaiguang/go-srv-services/app/user-service/internal/setup"
)

// RegisterWssRoutes 注册路由
// ref https://github.com/go-kratos/examples/blob/main/ws/main.go
func RegisterWssRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	wssSrv := srvs.NewWebsocketSrv(logger)
	handler := services.NewWebsocketService(wssSrv)

	router := mux.NewRouter()
	router.HandleFunc("/ws/v1/websocket", handler.TestWebsocket)

	stdlog.Println("|*** 注册路由：Websocket")
	hs.Handle("/ws/v1/websocket", router)

	return
}
