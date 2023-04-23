package dbmigrate

import (
	"context"
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"

	dbv1_0_0 "github.com/ikaiguang/go-srv-services/app/admin-service/cmd/migration/v1.0.0"
	"github.com/ikaiguang/go-srv-services/app/ping-service/internal/setup"
)

// Run 运行迁移
func Run(opts ...Option) {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}

	// 初始化
	if err := setup.Init(); err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// 模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	// 关闭
	if opt.closeEngine {
		defer func() { _ = setup.Close() }()
	}

	// 数据库链接
	dbConn, err := engineHandler.GetMySQLGormDB()
	//dbConn, err := engineHandler.GetPostgresGormDB()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}

	// migrateHandler 迁移手柄
	var (
		ctx         = context.Background()
		migrateRepo = migrationutil.NewMigrateRepo(dbConn)
	)

	// 初始化迁移记录
	if err = migrateRepo.InitializeSchema(ctx); err != nil {
		logutil.Fatalw("migrateHandler.InitializeMigrationSchema failed", err)
	}

	// v1.0.0
	if err = dbv1_0_0.Upgrade(ctx, dbConn, migrateRepo); err != nil {
		logutil.Fatalw("dbv1_0_0.Upgrade failed", err)
	}
}
