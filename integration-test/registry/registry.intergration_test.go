package testregistry

import (
	"context"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	"testing"

	apppkg "github.com/ikaiguang/go-srv-services/pkg/app"
)

// make run service=ping-service
// go test -v -count=1 ./integration-test/registry -test.run=Test_RegistryDiscovery
func Test_RegistryDiscovery(t *testing.T) {
	consulCli, err := engineHandler.GetConsulClient()
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)

	// 引擎模块
	appID := apppkg.ID(engineHandler.AppConfig())
	endpoint := "discovery:///" + appID

	// new grpc client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	gClient := pingservicev1.NewSrvPingClient(conn)

	// new http client
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithEndpoint(endpoint),
		http.WithDiscovery(r),
		// 解析
		http.WithResponseDecoder(apppkg.ResponseDecoder),
	)
	if err != nil {
		logutil.Fatal(err)
	}
	defer func() { _ = hConn.Close() }()
	hClient := pingservicev1.NewSrvPingHTTPClient(hConn)

	//for {
	//	time.Sleep(time.Second)
	//	callGRPC(gClient)
	//	callHTTP(hClient)
	//}
	callGRPC(gClient)
	callHTTP(hClient)
}

func callGRPC(client pingservicev1.SrvPingClient) {
	reply, err := client.Ping(context.Background(), &pingv1.PingReq{Message: "grpc"})
	if err != nil {
		log.Fatal(err)
	}
	logutil.Infof("[grpc] SayHello %+v\n", reply)
}

func callHTTP(client pingservicev1.SrvPingHTTPClient) {
	reply, err := client.Ping(context.Background(), &pingv1.PingReq{Message: "http"})
	if err != nil {
		log.Fatal(err)
	}
	logutil.Printf("[http] SayHello %s\n", reply.Message)
}
