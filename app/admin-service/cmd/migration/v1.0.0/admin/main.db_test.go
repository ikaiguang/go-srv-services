package dbv1_0_0_admin

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	"github.com/ikaiguang/go-srv-services/business-kit/setup"
	stdlog "log"
	"os"
	"testing"

	"github.com/ikaiguang/go-srv-services/app/admin-service/internal/setup"
)

// upHandler handler
var upHandler *Migrate

func TestMain(m *testing.M) {
	// 初始化
	configPath := "./../../../../configs"
	if err := setup.Init(setuputil.WithConfigPath(configPath)); err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	//defer func() { _ = setup.Close() }()

	// 模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	//defer func() { _ = setup.Close() }()

	// 数据库链接
	dbConn, err := engineHandler.GetMySQLGormDB()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// migrateHandler 迁移手柄
	migrateRepo := migrationutil.NewMigrateRepo(dbConn)

	// handler
	upHandler = NewMigrateHandler(dbConn, migrateRepo)

	os.Exit(m.Run())
}
