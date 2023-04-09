package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	curlutil "github.com/ikaiguang/go-srv-kit/kit/curl"
	clientutil "github.com/ikaiguang/go-srv-kit/kratos/client"
	pkgerrors "github.com/pkg/errors"
	stdgrpc "google.golang.org/grpc"

	apppkg "github.com/ikaiguang/go-srv-services/pkg/app"
)

const (
	defaultTimeout = curlutil.DefaultTimeout
)

// NewDefaultMiddlewares 中间件
func NewDefaultMiddlewares(logger log.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(),
		metadata.Client(),
		tracing.Client(),
		apppkg.ClientLog(logger),
	}
}

// NewGRCPConnection grpc 链接
func NewGRCPConnection(logger log.Logger, otherOpts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	var opts = []grpc.ClientOption{
		grpc.WithTimeout(defaultTimeout),
	}

	// 中间件
	opts = append(opts, grpc.WithMiddleware(NewDefaultMiddlewares(logger)...))
	// 其他
	opts = append(opts, otherOpts...)

	// grpc 链接
	conn, err := grpc.DialInsecure(
		context.Background(),
		opts...,
	)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	return conn, nil
}

// NewHTTPConnection http 链接
func NewHTTPConnection(logger log.Logger, otherOpts ...http.ClientOption) (*http.Client, error) {
	var opts = []http.ClientOption{
		http.WithTimeout(defaultTimeout),
		http.WithResponseDecoder(apppkg.ResponseDecoder),
	}

	// 中间件
	opts = append(opts, http.WithMiddleware(NewDefaultMiddlewares(logger)...))
	// 其他
	opts = append(opts, otherOpts...)

	// http 链接
	conn, err := clientutil.NewHTTPClient(context.Background(), opts...)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	return conn, nil
}
