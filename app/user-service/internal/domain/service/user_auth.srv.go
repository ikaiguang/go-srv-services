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
	usererrorv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/errors"
	userv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
	repos "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/repo"
)

// UserAuthSrv ...
type UserAuthSrv struct {
	authTokenRepo tokenutil.AuthTokenRepo
	log           *log.Helper

	userRepo         repos.UserRepo
	userRegEmailRepo repos.UserRegEmailRepo
}

func NewUserAuthSrv(
	authTokenRepo tokenutil.AuthTokenRepo,
	logger log.Logger,
	userRepo repos.UserRepo,
	userRegEmailRepo repos.UserRegEmailRepo,
) *UserAuthSrv {
	return &UserAuthSrv{
		authTokenRepo:    authTokenRepo,
		log:              log.NewHelper(log.With(logger, "module", "user/domain/service/user_auth")),
		userRepo:         userRepo,
		userRegEmailRepo: userRegEmailRepo,
	}
}

// CheckLoginEmail 邮箱是否存在
func (s *UserAuthSrv) CheckLoginEmail(ctx context.Context, email string) (*entities.UserRegEmail, error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.userRegEmailRepo.QueryOneByUserEmail(ctx, email)
	if err != nil {
		reason := errorv1.ERROR_DB.String()
		message := "数据库错误"
		err = errorutil.InternalServer(reason, message, err)
		return nil, err
	}
	if isNotFound {
		reason := usererrorv1.ERROR_USER_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return nil, err
	}
	return regEmailModel, nil
}

// CheckLoginUserByUUID 用户是否存在
func (s *UserAuthSrv) CheckLoginUserByUUID(ctx context.Context, userUuid string) (*entities.User, error) {
	userModel, isNotFound, err := s.userRepo.QueryOneByUserUuid(ctx, userUuid)
	if err != nil {
		reason := errorv1.ERROR_DB.String()
		message := "数据库错误"
		err = errorutil.InternalServer(reason, message, err)
		return nil, err
	}
	if isNotFound {
		reason := usererrorv1.ERROR_USER_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return nil, err
	}
	return userModel, nil
}

// ComparePassword 对比密码
func (s *UserAuthSrv) ComparePassword(hashPassword, password string) error {
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
func (s *UserAuthSrv) SignToken(ctx context.Context, userModel *entities.User, authInfo *userv1.Info) (signedString string, err error) {
	// 用户信息
	anyData, err := anypb.New(authInfo)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return signedString, err
	}

	// token claims
	authClaims := &authutil.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: authutil.DefaultExpireTime(),
		},
		Payload: &authv1.Payload{
			Uid: userModel.UserUuid,
			Tt:  authv1.TokenTypeEnum_API,
			Lp:  authv1.PlatformEnum_WEB,
			Lt:  authv1.LimitTypeEnum_ONLY_ONE,
			St:  timestamppb.New(time.Now()),
		},
	}

	// 密码
	authCache := &authv1.Auth{
		Data:    anyData,
		Payload: authClaims.Payload,
		Secret:  userModel.PasswordHash,
	}

	// 存储缓存
	err = s.authTokenRepo.SaveCacheData(ctx, authClaims, authCache)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return signedString, err
	}

	// generate token
	signingSecret := s.authTokenRepo.SigningSecret(ctx, authClaims.Payload.Tt, userModel.PasswordHash)
	signedString, err = s.authTokenRepo.SignedToken(authClaims, signingSecret)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return signedString, err
	}
	return signedString, err
}
