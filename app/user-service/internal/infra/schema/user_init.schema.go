package schemas

import (
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	uuidutil "github.com/ikaiguang/go-srv-kit/kit/uuid"
	"gorm.io/gorm"
	"time"

	userenumv1 "github.com/ikaiguang/go-srv-services/api/user/v1/enums"
)

// Initialize 初始化
func (s *User) Initialize(dbConn *gorm.DB) (err error) {
	return s.initialize(dbConn)
}

// initialize 初始化
func (s *User) initialize(dbConn *gorm.DB) (err error) {
	// 检查是否存在
	var (
		userEmail       = "user@user.user"
		existEmailModel = &UserRegEmail{}
	)
	err = dbConn.Table(UserRegEmailSchema.TableName()).
		Where("user_email = ?", userEmail).
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
		userUuid        = uuidutil.New()
		userNickname    = "user"
		passwordHash, _ = passwordutil.Encrypt("User.123456")
		registerFrom    = "EMAIL"
	)
	regEmailModel := &UserRegEmail{
		UserUuid:    userUuid,
		UserEmail:   userEmail,
		CreatedTime: now,
		UpdatedTime: now,
	}
	userModel := &User{
		UserUuid:        userUuid,
		CreatedTime:     now,
		UpdatedTime:     now,
		UserEmail:       regEmailModel.UserEmail,
		UserNickname:    userNickname,
		UserAvatar:      "",
		UserGender:      userenumv1.UserGenderEnum_SECRET.String(),
		UserStatus:      userenumv1.UserStatusEnum_ENABLE.String(),
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
	if err = tx.Create(userModel).Error; err != nil {
		return err
	}

	err = tx.Commit().Error
	return err
}
