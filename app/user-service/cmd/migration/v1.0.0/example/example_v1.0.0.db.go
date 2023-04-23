package dbv1_0_0_example

import (
	"context"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"

	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"

	schemas "github.com/ikaiguang/go-srv-services/app/user-service/internal/infra/schema"
)

// Migrate 数据库迁移
type Migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationutil.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) *Migrate {
	return &Migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// Upgrade ...
func (s *Migrate) Upgrade(ctx context.Context) error {
	var (
		mr       migrationutil.MigrationInterface
		migrator = s.dbConn.WithContext(ctx).Migrator()
	)

	// 创建表
	mr = schemas.ExampleSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		return pkgerrors.WithStack(err)
	}
	return nil
}
