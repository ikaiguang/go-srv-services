package dbv1_0_0

import (
	"context"
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	"gorm.io/gorm"

	dbv1_0_0_admin "github.com/ikaiguang/go-srv-services/app/admin-service/cmd/migration/v1.0.0/admin"
	dbv1_0_0_example "github.com/ikaiguang/go-srv-services/app/admin-service/cmd/migration/v1.0.0/example"
)

const (
	Version = "v1.0.0"
)

// Upgrade update
func Upgrade(ctx context.Context, dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	//var (
	//	serverVersion     = Version
	//	migrateIdentifier = serverVersion + ":migrate"
	//)
	//// 已进行数据库迁移
	//dataModel, _, err := migrateRepo.QueryRecord(ctx, migrateIdentifier)
	//if err != nil {
	//	return
	//}
	//if dataModel.Id > 0 {
	//	return
	//}
	//// 记录数据库迁移
	//defer func() {
	//	if err == nil {
	//		err = migrateRepo.CreateRecord(ctx, serverVersion, migrateIdentifier)
	//	}
	//}()

	// example
	err = dbv1_0_0_example.NewMigrateHandler(dbConn, migrateRepo).Upgrade(ctx)
	if err != nil {
		return err
	}

	// admin
	err = dbv1_0_0_admin.NewMigrateHandler(dbConn, migrateRepo).Upgrade(ctx)
	if err != nil {
		return err
	}

	return err
}
