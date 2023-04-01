package services

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	authv1 "github.com/ikaiguang/go-srv-kit/api/auth/v1"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	passwordutil "github.com/ikaiguang/go-srv-kit/kit/password"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
	tokenutil "github.com/ikaiguang/go-srv-kit/kratos/token"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	usererrorv1 "github.com/ikaiguang/go-srv-services/api/user/v1/errors"
	userv1 "github.com/ikaiguang/go-srv-services/api/user/v1/resources"
	userservicev1 "github.com/ikaiguang/go-srv-services/api/user/v1/services"
	"github.com/ikaiguang/go-srv-services/app/user-service/internal/application/assembler"
	repos "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/repo"
)

// userAuth ...
type userAuth struct {
	userservicev1.UnimplementedSrvUserAuthServer

	authTokenRepo tokenutil.AuthTokenRepo

	userRepo         repos.UserRepo
	userRegEmailRepo repos.UserRegEmailRepo
}

// NewUserAuthService ...
func NewUserAuthService(
	authTokenRepo tokenutil.AuthTokenRepo,
	userRepo repos.UserRepo,
	userRegEmailRepo repos.UserRegEmailRepo,
) userservicev1.SrvUserAuthServer {
	return &userAuth{
		authTokenRepo:    authTokenRepo,
		userRepo:         userRepo,
		userRegEmailRepo: userRegEmailRepo,
	}
}

// LoginByEmail Email登录
func (s *userAuth) LoginByEmail(ctx context.Context, in *userv1.LoginByEmailReq) (out *userv1.LoginResp, err error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.userRegEmailRepo.QueryOneByUserEmail(ctx, in.Email)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}
	if isNotFound {
		reason := usererrorv1.ERROR_USER_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// user
	userModel, isNotFound, err := s.userRepo.QueryOneByUserUuid(ctx, regEmailModel.UserUuid)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}
	if isNotFound {
		reason := usererrorv1.ERROR_USER_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// 验证密码
	err = passwordutil.Compare(userModel.PasswordHash, in.Password)
	if err != nil {
		reason := usererrorv1.ERROR_USER_PASSWORD_INCORRECT.String()
		message := "密码不正确"
		err = errorutil.BadRequest(reason, message)
		//err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// 管理员信息
	userInfo := &userv1.Info{
		Id:           userModel.Id,
		UserUuid:     userModel.UserUuid,
		UserNickname: userModel.UserNickname,
		UserAvatar:   userModel.UserAvatar,
		UserGender:   assemblers.ToUserGenderEnum(userModel.UserGender),
		UserStatus:   assemblers.ToUserStatusEnum(userModel.UserStatus),
	}
	anyData, err := anypb.New(userInfo)
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
		return out, err
	}

	// generate token
	signingSecret := s.authTokenRepo.SigningSecret(ctx, authClaims.Payload.Tt, userModel.PasswordHash)
	signedString, err := s.authTokenRepo.SignedToken(authClaims, signingSecret)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}

	// 响应
	out = &userv1.LoginResp{
		UserInfo: userInfo,
		Token:    signedString,
	}
	return out, err
}

// Ping ping pong
func (s *userAuth) Ping(ctx context.Context, in *userv1.PingReq) (out *userv1.PingResp, err error) {
	// 可以解析
	tokenClaims, tokenClaimsOK := authutil.FromJWTContext(ctx)
	logutil.InfowWithContext(ctx,
		"tokenClaimsOK", tokenClaimsOK,
		"tokenClaims", tokenClaims,
	)

	// 当：使用 authutil.Claims 时，可以解析
	authClaims, authClaimsOK := authutil.FromAuthContext(ctx)
	logutil.InfowWithContext(ctx,
		"authClaimsOK", authClaimsOK,
		"authClaims", authClaims,
	)

	// 当：使用 tokenutil.AuthTokenRepo 时，可以解析
	authInfo, authInfoOK := authutil.FromRedisContext(ctx)
	logutil.InfowWithContext(ctx,
		"authInfoOK", authInfoOK,
		"authInfo", authInfo,
	)

	out = &userv1.PingResp{
		Message: in.Message,
	}
	return out, err
}
