package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	"github.com/ikaiguang/go-srv-services/business-kit/app"
	stdlog "log"

	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"

	"github.com/ikaiguang/go-srv-services/app/user-service/internal/setup"
)

// RegisterRootRoutes 注册路由
func RegisterRootRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &pingv1.PingResp{
			Message: "Hello World!",
		}
		err := apputil.ResponseEncoder(w, r, data)
		if err != nil {
			apputil.ErrorEncoder(w, r, err)
		}
	})

	stdlog.Println("|*** 注册路由：Root(/)")
	hs.Handle("/", router)

	return
}
