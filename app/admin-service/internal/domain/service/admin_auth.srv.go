package srvs

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
	tokenutil "github.com/ikaiguang/go-srv-kit/kratos/token"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	adminerrorv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/errors"
	adminv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/resources"
	assemblers "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/assembler"
	entities "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/entity"
	repos "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/repo"
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
		log:               log.NewHelper(logger),
		adminRepo:         adminRepo,
		adminRegEmailRepo: adminRegEmailRepo,
	}
}

// CheckLoginEmail 邮箱是否存在
func (s *AdminAuthSrv) CheckLoginEmail(ctx context.Context, email string) (*entities.AdminRegEmail, error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.adminRegEmailRepo.QueryOneByAdminEmail(ctx, email)
	if err != nil {
		reason := errorv1.ERROR_DB.String()
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
		reason := errorv1.ERROR_DB.String()
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
func (s *AdminAuthSrv) SignToken(ctx context.Context, adminModel *entities.Admin) (out *adminv1.LoginResp, err error) {
	// 管理员信息
	adminInfo := &adminv1.Info{
		Id:            adminModel.Id,
		AdminUuid:     adminModel.AdminUuid,
		AdminNickname: adminModel.AdminNickname,
		AdminAvatar:   adminModel.AdminAvatar,
		AdminGender:   assemblers.ToAdminGenderEnum(adminModel.AdminGender),
		AdminStatus:   assemblers.ToAdminStatusEnum(adminModel.AdminStatus),
	}
	anyData, err := anypb.New(adminInfo)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}

	// token claims
	authClaims := &authutil.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: authutil.DefaultExpireTime(),
		},
		Payload: &authv1.Payload{
			Uid: adminModel.AdminUuid,
			Tt:  authv1.TokenTypeEnum_ADMIN,
			Lp:  authv1.PlatformEnum_WEB,
			Lt:  authv1.LimitTypeEnum_ONLY_ONE,
			St:  timestamppb.New(time.Now()),
		},
	}

	// 密码
	authCache := &authv1.Auth{
		Data:    anyData,
		Payload: authClaims.Payload,
		Secret:  adminModel.PasswordHash,
	}

	// 存储缓存
	err = s.authTokenRepo.SaveCacheData(ctx, authClaims, authCache)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}

	// generate token
	signingSecret := s.authTokenRepo.SigningSecret(ctx, authClaims.Payload.Tt, adminModel.PasswordHash)
	signedString, err := s.authTokenRepo.SignedToken(authClaims, signingSecret)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}

	// 响应
	out = &adminv1.LoginResp{
		AdminInfo: adminInfo,
		Token:     signedString,
	}
	return out, err
}
