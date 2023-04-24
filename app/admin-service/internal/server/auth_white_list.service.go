package servers

import (
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	testdataservicev1 "github.com/ikaiguang/go-srv-kit/api/testdata/v1/services"
)

// getAuthWhiteList 验证白名单
func getAuthWhiteList() map[string]struct{} {
	// 白名单
	whiteList := make(map[string]struct{})

	// 测试
	whiteList[pingservicev1.OperationSrvPingPing] = struct{}{}
	whiteList[testdataservicev1.OperationSrvTestdataWebsocket] = struct{}{}

	return whiteList
}
