package dbv1_0_0_snowflake

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	schemas "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/infra/schema"
	"gorm.io/gorm"
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

// Upgrade .
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	upgradeHandler := NewMigrateHandler(dbConn, migrateRepo)

	// 创建表 example
	err = upgradeHandler.CreateTableSnowflakeNodeID()
	if err != nil {
		return err
	}
	return err
}

// CreateTableSnowflakeNodeID ...
func (s *Migrate) CreateTableSnowflakeNodeID() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.SnowflakeNodeIDSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.SnowflakeNodeIDSchema)
}
