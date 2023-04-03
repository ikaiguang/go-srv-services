package services

import (
	"context"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/xid"

	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	testdatav1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/resources"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"

	services "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/service"
)

// testdata .
type testdata struct {
	testdataservicev1.UnimplementedSrvTestdataServer

	log    *log.Helper
	wssSrv *services.WebsocketSrv
}

// NewTestdataService .
func NewTestdataService(logger log.Logger, wssSrv *services.WebsocketSrv) testdataservicev1.SrvTestdataServer {
	return &testdata{
		log:    log.NewHelper(logger),
		wssSrv: wssSrv,
	}
}

// Websocket websocket
func (s *testdata) Websocket(ctx context.Context, in *testdatav1.TestReq) (resp *testdatav1.TestResp, err error) {
	// http
	httpContext, isHTTPContext := contextutil.MatchHTTPContext(ctx)
	if isHTTPContext {
		if httpContext.Request() == nil || httpContext.Request().Method != stdhttp.MethodGet {
			err = errorutil.InternalServer(errorv1.ERROR_METHOD_NOT_ALLOWED.String(), "ERROR_METHOD_NOT_ALLOWED")
			s.log.WithContext(ctx).Error(err)
			return resp, err
		}

		// ws
		err = s.wss(httpContext, httpContext.Response(), httpContext.Request())
		if err != nil {
			return resp, err
		}
		resp = &testdatav1.TestResp{
			Message: xid.New().String(),
		}
		err = errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "unimplemented")
		return resp, err
	}

	err = errorutil.NotImplemented(errorv1.ERROR_NOT_IMPLEMENTED.String(), "unimplemented")
	return &testdatav1.TestResp{}, err
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

// wss ws
func (s *testdata) wss(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
	// 升级连接
	cc, err := websocketutil.UpgradeConn(w, r, w.Header())
	if err != nil {
		err = errorutil.InternalServer(errorv1.ERROR_CONNECTION.String(), "upgrade conn failed", err)
		s.log.WithContext(ctx).Error(err)
		return
	}
	defer func() { _ = cc.Close() }()

	// 处理消息
	err = s.wssSrv.ProcessWssMsg(ctx, cc)
	if err != nil {
		return err
	}
	return err
}
