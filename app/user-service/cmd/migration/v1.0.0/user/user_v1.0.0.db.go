package dbv1_0_0_user

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	"gorm.io/gorm"

	schemas "github.com/ikaiguang/go-srv-services/app/user-service/internal/infra/schema"
)

// Upgrade .
func Upgrade(dbConn *gorm.DB, migrateRepo migrationutil.MigrateRepo) (err error) {
	upgradeHandler := NewMigrateHandler(dbConn, migrateRepo)

	// 创建表 user
	err = upgradeHandler.CreateTableUser()
	if err != nil {
		return err
	}
	// 创建表 user
	err = upgradeHandler.CreateTableUserRegEmail()
	if err != nil {
		return err
	}
	// 创建表 user
	err = upgradeHandler.CreateTableUserRegMobile()
	if err != nil {
		return err
	}
	// 创建表 user
	err = upgradeHandler.CreateTableUserRegUsername()
	if err != nil {
		return err
	}

	// 初始化用户
	err = upgradeHandler.InitializeUser()
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

// CreateTableUser ...
func (s *migrate) CreateTableUser() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.UserSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.UserSchema)
}

// CreateTableUserRegEmail ...
func (s *migrate) CreateTableUserRegEmail() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.UserRegEmailSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.UserRegEmailSchema)
}

// CreateTableUserRegMobile ...
func (s *migrate) CreateTableUserRegMobile() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.UserRegMobileSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.UserRegMobileSchema)
}

// CreateTableUserRegUsername ...
func (s *migrate) CreateTableUserRegUsername() (err error) {
	if s.dbConn.Migrator().HasTable(schemas.UserRegUsernameSchema) {
		return err
	}
	return s.dbConn.Migrator().CreateTable(schemas.UserRegUsernameSchema)
}

// InitializeUser 初始化管理员
func (s *migrate) InitializeUser() (err error) {
	var (
		serverVersion     = "v1.0.0"
		migrateIdentifier = serverVersion + "user:InitializeUser"
	)
	// 已进行数据库迁移
	dataModel, _, err := s.migrateRepo.QueryOneByIdentifier(migrateIdentifier)
	if err != nil {
		return
	}
	if dataModel.Id > 0 {
		return
	}
	// 记录数据库迁移
	defer func() {
		if err == nil {
			err = s.migrateRepo.CreateDefaultRecord(serverVersion, migrateIdentifier)
		}
	}()

	err = schemas.UserSchema.Initialize(s.dbConn)
	if err != nil {
		return err
	}
	return
}
