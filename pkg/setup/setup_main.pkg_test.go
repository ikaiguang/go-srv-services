package setuppkg

import (
	"os"
	"testing"

	"github.com/ikaiguang/go-srv-services/testdata"
)

// confPath 配置目录
const confPath = "./../../testdata/configs"

func TestMain(m *testing.M) {
	_configFilepath = testdata.ConfigPath()

	// 初始化必要逻辑
	os.Exit(m.Run())
}
