package schemas

import (
	migrationutil "github.com/ikaiguang/go-srv-kit/data/migration"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	"gorm.io/gorm"
	"time"

	userenumv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/enums"
)

// InitializeUser 初始化
// user@user.user : User.123456
func (s *User) InitializeUser(dbConn *gorm.DB) migrationutil.MigrationInterface {
	var (
		userEmail           = "user@user.user"
		userUuid            = userEmail
		migrationVersion    = migrationutil.Version
		migrationIdentifier = migrationVersion + ":" + s.TableName() + ":initializeUser:" + userEmail
	)
	migrationUp := func() error {
		return s.initializeUserUp(dbConn, userUuid, userEmail)
	}
	migrationDown := func() error {
		return s.initializeUserDown(dbConn, userUuid, userEmail)
	}
	return migrationutil.NewAnyMigrator(
		migrationVersion,
		migrationIdentifier,
		migrationUp,
		migrationDown,
	)
}

// initializeUserDown 初始化
func (s *User) initializeUserDown(dbConn *gorm.DB, userUuid, userEmail string) (err error) {
	tx := dbConn.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback().Error
		}
	}()

	err = tx.Table(UserRegEmailSchema.TableName()).
		Where("user_email = ?", userEmail).
		Delete(UserRegEmail{}).Error
	if err != nil {
		return err
	}
	err = tx.Table(UserSchema.TableName()).
		Where("user_uuid = ?", userUuid).
		Delete(User{}).Error
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return err
}

// initializeUserUp 初始化
func (s *User) initializeUserUp(dbConn *gorm.DB, userUuid, userEmail string) (err error) {
	// 检查是否存在
	var (
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
	if err != nil {
		return err
	}
	return err
}
