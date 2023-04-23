package schemas

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	"gorm.io/gorm"
	"time"

	adminenumv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/enums"
)

// InitializeAdmin 初始化
// admin@admin.admin : Admin.123456
func (s *Admin) InitializeAdmin(dbConn *gorm.DB) migrationutil.MigrationInterface {
	var (
		adminEmail          = "admin@admin.admin"
		adminUuid           = adminEmail
		migrationVersion    = migrationutil.Version
		migrationIdentifier = migrationVersion + ":" + s.TableName() + ":initializeAdmin:" + adminEmail
	)
	migrationUp := func() error {
		return s.initializeAdminUp(dbConn, adminUuid, adminEmail)
	}
	migrationDown := func() error {
		return s.initializeAdminDown(dbConn, adminUuid, adminEmail)
	}
	return migrationutil.NewAnyMigrator(
		migrationVersion,
		migrationIdentifier,
		migrationUp,
		migrationDown,
	)
}

// initializeAdminDown 初始化
func (s *Admin) initializeAdminDown(dbConn *gorm.DB, adminUuid, adminEmail string) (err error) {
	tx := dbConn.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback().Error
		}
	}()

	err = tx.Table(AdminRegEmailSchema.TableName()).
		Where("admin_email = ?", adminEmail).
		Delete(AdminRegEmail{}).Error
	if err != nil {
		return err
	}
	err = tx.Table(AdminSchema.TableName()).
		Where("admin_uuid = ?", adminUuid).
		Delete(Admin{}).Error
	if err != nil {
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return err
}

// initializeAdminUp 初始化
func (s *Admin) initializeAdminUp(dbConn *gorm.DB, adminUuid, adminEmail string) (err error) {
	// 检查是否存在
	var (
		existEmailModel = &AdminRegEmail{}
	)
	err = dbConn.Table(AdminRegEmailSchema.TableName()).
		Where("admin_email = ?", adminEmail).
		First(existEmailModel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		} else {
			return err
		}
	}
	if existEmailModel.Id > 0 {
		return err
	}

	// 新增
	var (
		now             = time.Now()
		activeEndTime   = now.Add(time.Hour * 24 * 365 * 10)
		adminNickname   = "admin"
		passwordHash, _ = passwordutil.Encrypt("Admin.123456")
		registerFrom    = "EMAIL"
	)
	regEmailModel := &AdminRegEmail{
		AdminUuid:   adminUuid,
		AdminEmail:  adminEmail,
		CreatedTime: now,
		UpdatedTime: now,
	}
	adminModel := &Admin{
		AdminUuid:       adminUuid,
		CreatedTime:     now,
		UpdatedTime:     now,
		AdminEmail:      regEmailModel.AdminEmail,
		AdminNickname:   adminNickname,
		AdminAvatar:     "",
		AdminGender:     adminenumv1.AdminGenderEnum_SECRET.String(),
		AdminStatus:     adminenumv1.AdminStatusEnum_ENABLE.String(),
		ActiveBeginTime: &now,
		ActiveEndTime:   &activeEndTime,
		PasswordHash:    string(passwordHash),
		RegisterFrom:    registerFrom,
	}

	// 新增用户
	tx := dbConn.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback().Error
		}
	}()

	if err = tx.Create(regEmailModel).Error; err != nil {
		return err
	}
	if err = tx.Create(adminModel).Error; err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return err
}
