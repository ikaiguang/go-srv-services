package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"
	stdhttp "net/http"

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
	handler := NewWsHandler(wssSrv)

	router := mux.NewRouter()
	router.HandleFunc("/ws/v1/websocket", handler.TestWebsocket)

	stdlog.Println("|*** 注册路由：Websocket")
	hs.Handle("/ws/v1/websocket", router)

	return
}

// WsHandler ...
type WsHandler struct {
	wsSrv *srvs.WebsocketSrv
}

// NewWsHandler ...
func NewWsHandler(wsSrv *srvs.WebsocketSrv) *WsHandler {
	return &WsHandler{
		wsSrv: wsSrv,
	}
}

// TestWebsocket ...
func (s *WsHandler) TestWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != stdhttp.MethodGet {
		err := errorutil.InternalServer(errorv1.ERROR_METHOD_NOT_ALLOWED.String(), "ERROR_METHOD_NOT_ALLOWED")
		w.WriteHeader(stdhttp.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err := s.wsSrv.Wss(r.Context(), w, r)
	if err != nil {
		w.WriteHeader(stdhttp.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	return
}
