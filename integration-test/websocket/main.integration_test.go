package testwebsocket

import (
	stdlog "log"
	"os"
	"testing"

	"github.com/ikaiguang/go-srv-services/pkg/setup"
)

var (
	engineHandler setuppkg.Engine
)

func TestMain(m *testing.M) {
	var err error

	//configPath := "../app/ping-service/configs"
	configPath := "./../../app/ping-service/configs/develop"
	engineHandler, err = setuppkg.New(setuppkg.WithConfigPath(configPath))
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	defer func() { _ = engineHandler.Close() }()

	code := m.Run()

	// 初始化必要逻辑
	os.Exit(code)
}
