package dbv1_0_0

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	dbv1_0_0_snowflake "github.com/ikaiguang/go-srv-services/app/snowflake-service/cmd/migration/v1.0.0/snowflake"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
)

// Upgrade update
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	//var (
	//	serverVersion     = "v1.0.0"
	//	migrateIdentifier = serverVersion + ":migrate"
	//)
	//// 已进行数据库迁移
	//dataModel, _, err := migrateRepo.QueryOneByIdentifier(migrateIdentifier)
	//if err != nil {
	//	return
	//}
	//if dataModel.Id > 0 {
	//	return
	//}
	//// 记录数据库迁移
	//defer func() {
	//	if err == nil {
	//		err = migrateRepo.CreateDefaultRecord(serverVersion, migrateIdentifier)
	//	}
	//}()

	// snowflake
	err = dbv1_0_0_snowflake.Upgrade(dbConn, migrateRepo)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	return err
}
