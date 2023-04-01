package schemas

import (
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	uuidutil "github.com/ikaiguang/go-srv-kit/kit/uuid"
	"gorm.io/gorm"
	"time"

	adminenumv1 "github.com/ikaiguang/go-srv-services/api/admin/v1/enums"
)

// Initialize 初始化
func (s *Admin) Initialize(dbConn *gorm.DB) (err error) {
	return s.initialize(dbConn)
}

// initialize 初始化
func (s *Admin) initialize(dbConn *gorm.DB) (err error) {
	// 检查是否存在
	var (
		adminEmail      = "admin@admin.admin"
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
		adminUuid       = uuidutil.New()
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
	return err
}
