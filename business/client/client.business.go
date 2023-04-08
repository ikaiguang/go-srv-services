package clientutil

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	clientutil "github.com/ikaiguang/go-srv-kit/kratos/client"

	serviceutil "github.com/ikaiguang/go-srv-services/business/service"
)

// NewUserHTTPClient ...
func NewUserHTTPClient(name serviceutil.ServerName) {
	_ = clientutil.NewHTTPClient
	_ = http.NewClient
}

// NewUserRPCClient ...
func NewUserRPCClient(name serviceutil.ServerName) {
	_ = clientutil.NewGRPCClient
	_ = grpc.Dial
}
