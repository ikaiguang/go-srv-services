package dbv1_0_0_org

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	"gorm.io/gorm"

	schemas "github.com/ikaiguang/go-srv-services/app/user-service/internal/infra/schema"
)

// Upgrade .
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	upgradeHandler := NewMigrateHandler(dbConn, migrateRepo)

	// CreateTableOrg 创建表 org
	err = upgradeHandler.CreateTableOrg()
	if err != nil {
		return err
	}
	return err
}

// migrate 数据库迁移
type migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationutil.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) *migrate {
	return &migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// CreateTableOrg ...
func (s *migrate) CreateTableOrg() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.OrgSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.OrgSchema)
}
