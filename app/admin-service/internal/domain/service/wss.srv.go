package srvs

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	websocketutil "github.com/ikaiguang/go-srv-kit/kratos/websocket"
)

// WebsocketSrv ...
type WebsocketSrv struct {
	log *log.Helper
}

// NewWebsocketSrv ...
func NewWebsocketSrv(logger log.Logger) *WebsocketSrv {
	return &WebsocketSrv{
		log: log.NewHelper(logger),
	}
}

// WsMessage ws
type WsMessage struct {
	Type    int
	Content []byte
}

func (s *WebsocketSrv) ProcessWssMsg(ctx context.Context, cc *websocket.Conn) error {
	// 读消息
	for {
		messageType, messageBytes, wsErr := cc.ReadMessage()
		if wsErr != nil {
			if websocketutil.IsCloseError(wsErr) {
				s.log.Infow("websocket close", wsErr.Error())
				break
			}
			err := errorutil.InternalServer(errorv1.ERROR_CONNECTION.String(), "ws读取信息失败", wsErr)
			s.log.Error(err)
			return err
		}

		// 消息
		wsMessage := &WsMessage{
			Type:    messageType,
			Content: messageBytes,
		}
		//messageChan <- wsMessage

		// 处理
		needCloseConn, err := s.processWsMessage(ctx, wsMessage)
		if err != nil {
			err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "ws处理信息失败", err)
			s.log.Error(err)
			return err
		}

		// 响应
		err = cc.WriteMessage(messageType, messageBytes)
		if err != nil {
			if websocketutil.IsCloseError(wsErr) {
				s.log.Infow("websocket close", wsErr.Error())
				break
			}
			err = errorutil.InternalServer(errorv1.ERROR_INTERNAL_SERVER.String(), "JSON编码错误", err)
			s.log.Error(err)
			return err
		}

		// 关闭
		if needCloseConn {
			return err
		}
	}
	return nil
}

// processWsMessage 处理ws-message
func (s *WebsocketSrv) processWsMessage(ctx context.Context, wsMessage *WsMessage) (needCloseConn bool, err error) {
	s.log.Infow("ws-message type", wsMessage.Type)
	s.log.Infow("ws-message message", string(wsMessage.Content))
	if string(wsMessage.Content) == "close" {
		needCloseConn = true
	}
	return needCloseConn, err
}