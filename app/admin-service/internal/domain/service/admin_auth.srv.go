package srvs

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	adminerrorv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/errors"
	adminv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/resources"
	commonv1 "github.com/ikaiguang/go-srv-services/api/common/v1"
	entities "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/entity"
	repos "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/repo"
	authutil "github.com/ikaiguang/go-srv-services/business-util/auth"
	tokenutil "github.com/ikaiguang/go-srv-services/business-util/token"
)

// AdminAuthSrv ...
type AdminAuthSrv struct {
	authTokenRepo tokenutil.AuthTokenRepo

	log               *log.Helper
	adminRepo         repos.AdminRepo
	adminRegEmailRepo repos.AdminRegEmailRepo
}

// NewAdminAuthSrv ...
func NewAdminAuthSrv(
	authTokenRepo tokenutil.AuthTokenRepo,
	logger log.Logger,
	adminRepo repos.AdminRepo,
	adminRegEmailRepo repos.AdminRegEmailRepo,
) *AdminAuthSrv {
	return &AdminAuthSrv{
		authTokenRepo:     authTokenRepo,
		log:               log.NewHelper(log.With(logger, "module", "admin/domain/service/admin_auth")),
		adminRepo:         adminRepo,
		adminRegEmailRepo: adminRegEmailRepo,
	}
}

// CheckLoginEmail 邮箱是否存在
func (s *AdminAuthSrv) CheckLoginEmail(ctx context.Context, email string) (*entities.AdminRegEmail, error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.adminRegEmailRepo.QueryOneByAdminEmail(ctx, email)
	if err != nil {
		reason := commonv1.ERROR_DB.String()
		message := "数据库错误"
		err = errorutil.InternalServer(reason, message, err)
		return nil, err
	}
	if isNotFound {
		reason := adminerrorv1.ERROR_ADMIN_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return nil, err
	}
	return regEmailModel, err
}

// CheckLoginUserByUUID 用户是否存在
func (s *AdminAuthSrv) CheckLoginUserByUUID(ctx context.Context, adminUuid string) (*entities.Admin, error) {
	// admin
	adminModel, isNotFound, err := s.adminRepo.QueryOneByAdminUuid(ctx, adminUuid)
	if err != nil {
		reason := commonv1.ERROR_DB.String()
		message := "数据库错误"
		err = errorutil.InternalServer(reason, message, err)
		return nil, err
	}
	if isNotFound {
		reason := adminerrorv1.ERROR_ADMIN_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return nil, err
	}
	return adminModel, nil
}

// ComparePassword 对比密码
func (s *AdminAuthSrv) ComparePassword(hashPassword, password string) error {
	// 验证密码
	err := passwordutil.Compare(hashPassword, password)
	if err != nil {
		reason := adminerrorv1.ERROR_ADMIN_PASSWORD_INCORRECT.String()
		message := "密码不正确"
		err = errorutil.BadRequest(reason, message)
		return err
	}
	return nil
}

// SignToken token
func (s *AdminAuthSrv) SignToken(ctx context.Context, adminModel *entities.Admin, adminInfo *adminv1.Info) (signedString string, err error) {
	// 管理员信息
	anyData, err := anypb.New(adminInfo)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return signedString, err
	}

	// token claims
	authClaims := &authutil.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: authutil.DefaultExpireTime(),
		},
		Payload: &commonv1.Payload{
			Uid: adminInfo.AdminUuid,
			Tt:  commonv1.TokenTypeEnum_ADMIN,
			Lp:  commonv1.PlatformEnum_WEB,
			Lt:  commonv1.LimitTypeEnum_ONLY_ONE,
			St:  timestamppb.New(time.Now()),
		},
	}

	// 密码
	authCache := &commonv1.Auth{
		Data:    anyData,
		Payload: authClaims.Payload,
		Secret:  adminModel.PasswordHash,
	}

	// 存储缓存
	err = s.authTokenRepo.SaveCacheData(ctx, authClaims, authCache)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return signedString, err
	}

	// generate token
	signingSecret := s.authTokenRepo.SigningSecret(ctx, authClaims.Payload.Tt, adminModel.PasswordHash)
	signedString, err = s.authTokenRepo.SignedToken(authClaims, signingSecret)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return signedString, err
	}
	return signedString, err
}
