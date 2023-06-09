package testping

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"
	"testing"
	"time"

	testdata "github.com/ikaiguang/go-srv-services/testdata"
)

// make run service=ping-service
// go test -v -count=1 ./integration-test/ping -test.run=TestGRPC_Ping_Hello
func TestGRPC_Ping_Hello(t *testing.T) {
	var (
		ctx  = context.Background()
		opts = []grpc.ClientOption{
			grpc.WithEndpoint(testdata.GRPCAddr()),
			grpc.WithTimeout(time.Second * 60),
		}
	)

	// 链接
	grpcConn, err := testdata.NewGRPCInsecureConn(ctx, opts...)
	require.Nil(t, err)
	defer func() {
		_ = grpcConn.Close()
	}()

	// 客户端
	cc := pingservicev1.NewSrvPingClient(grpcConn)

	// ping
	pingReq := &pingv1.PingReq{
		Message: "hello",
	}
	pingResp, err := cc.Ping(ctx, pingReq)
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			t.Error("grpc response error :", err)
			t.FailNow()
		}
		t.Logf("grpc response error code : %d\n", s.Code())
		t.Logf("grpc response error msg : %v\n", s.Message())
		t.Logf("grpc response error detail : %v\n", s.Details())
		t.FailNow()
	}

	t.Logf("grpc response body lenght : %d\n", len(pingResp.String()))
	if testdata.EnablePrintResult() {
		t.Logf("http response body content : %v\n", pingResp.String())
	}
}
