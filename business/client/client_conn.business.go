package clientutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	curlutil "github.com/ikaiguang/go-srv-kit/kit/curl"
	clientutil "github.com/ikaiguang/go-srv-kit/kratos/client"
	pkgerrors "github.com/pkg/errors"
	stdgrpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	commonv1 "github.com/ikaiguang/go-srv-services/api/common/v1"
	serviceutil "github.com/ikaiguang/go-srv-services/business/service"
	apppkg "github.com/ikaiguang/go-srv-services/pkg/app"
	registrypkg "github.com/ikaiguang/go-srv-services/pkg/registry"
	setuppkg "github.com/ikaiguang/go-srv-services/pkg/setup"
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
		//apppkg.ClientLog(logger),
	}
}

// NewGRPCConnection grpc 链接
func NewGRPCConnection(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName, otherOpts ...grpc.ClientOption) (*stdgrpc.ClientConn, error) {
	mLogger, _, err := engineHandler.LoggerMiddleware()
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var opts = []grpc.ClientOption{
		grpc.WithTimeout(defaultTimeout),
	}

	// 服务端点
	endpointOpts, err := GetGRPCEndpoint(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	opts = append(opts, endpointOpts...)

	// 中间件
	opts = append(opts, grpc.WithMiddleware(NewDefaultMiddlewares(mLogger)...))
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
func NewHTTPConnection(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName, otherOpts ...http.ClientOption) (*http.Client, error) {
	mLogger, _, err := engineHandler.LoggerMiddleware()
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var opts = []http.ClientOption{
		http.WithTimeout(defaultTimeout),
		http.WithResponseDecoder(apppkg.ResponseDecoder),
	}

	// 服务端点
	endpointOpts, err := GetHTTPEndpoint(engineHandler, serviceName)
	if err != nil {
		return nil, err
	}
	opts = append(opts, endpointOpts...)

	// 中间件
	opts = append(opts, http.WithMiddleware(NewDefaultMiddlewares(mLogger)...))
	// 其他
	opts = append(opts, otherOpts...)

	// http 链接
	conn, err := clientutil.NewHTTPClient(context.Background(), opts...)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	return conn, nil
}

// GetHTTPEndpoint 获取服务端点
func GetHTTPEndpoint(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) ([]http.ClientOption, error) {
	var opts []http.ClientOption
	registryType := engineHandler.RegistryType()
	switch registryType {
	case registrypkg.RegistryTypeConsul:
		consulClient, err := engineHandler.GetConsulClient()
		if err != nil {
			return nil, err
		}
		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return nil, err
		}
		// 端点 "discovery:///helloworld"
		appConfig := proto.Clone(engineHandler.AppConfig()).(*commonv1.App)
		appConfig.Name = serviceName.String()
		endpoint := "discovery:///" + apppkg.ID(appConfig)
		opts = append(opts,
			http.WithEndpoint(endpoint),
			http.WithDiscovery(r),
		)
	default:
		endpoint, err := serviceutil.GenLocalEndpoint(serviceName, transport.KindHTTP)
		if err != nil {
			return nil, err
		}
		opts = append(opts, http.WithEndpoint(endpoint))
	}
	return opts, nil
}

// GetGRPCEndpoint 获取服务端点
func GetGRPCEndpoint(engineHandler setuppkg.Engine, serviceName serviceutil.ServiceName) ([]grpc.ClientOption, error) {
	var opts []grpc.ClientOption
	registryType := engineHandler.RegistryType()
	switch registryType {
	case registrypkg.RegistryTypeConsul:
		consulClient, err := engineHandler.GetConsulClient()
		if err != nil {
			return nil, err
		}
		r, err := registrypkg.NewConsulRegistry(consulClient)
		if err != nil {
			return nil, err
		}
		// 端点 "discovery:///helloworld"
		appConfig := proto.Clone(engineHandler.AppConfig()).(*commonv1.App)
		appConfig.Name = serviceName.String()
		endpoint := "discovery:///" + apppkg.ID(appConfig)
		opts = append(opts,
			grpc.WithEndpoint(endpoint),
			grpc.WithDiscovery(r),
		)
	default:
		endpoint, err := serviceutil.GenLocalEndpoint(serviceName, transport.KindGRPC)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithEndpoint(endpoint))
	}
	return opts, nil
}
