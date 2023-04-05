package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	clientutil "github.com/ikaiguang/go-srv-kit/kratos/client"
)

// NewUserHTTPClient ...
func NewUserHTTPClient(name ServerName) {
	_ = clientutil.NewHTTPClient
	_ = http.NewClient
}

// NewUserRPCClient ...
func NewUserRPCClient(name ServerName) {
	_ = clientutil.NewGRPCClient
	_ = grpc.Dial
}
