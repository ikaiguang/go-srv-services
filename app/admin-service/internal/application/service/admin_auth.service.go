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

	adminerrorv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/errors"
	adminv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/resources"
	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/services"
	assemblers "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/assembler"
	repos "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/repo"
)

// adminAuth ...
type adminAuth struct {
	adminservicev1.UnimplementedSrvAdminAuthServer

	authTokenRepo tokenutil.AuthTokenRepo

	adminRepo         repos.AdminRepo
	adminRegEmailRepo repos.AdminRegEmailRepo
}

// NewAdminAuthService ...
func NewAdminAuthService(
	authTokenRepo tokenutil.AuthTokenRepo,
	adminRepo repos.AdminRepo,
	adminRegEmailRepo repos.AdminRegEmailRepo,
) adminservicev1.SrvAdminAuthServer {
	return &adminAuth{
		authTokenRepo:     authTokenRepo,
		adminRepo:         adminRepo,
		adminRegEmailRepo: adminRegEmailRepo,
	}
}

// LoginByEmail Email登录
func (s *adminAuth) LoginByEmail(ctx context.Context, in *adminv1.LoginByEmailReq) (out *adminv1.LoginResp, err error) {
	// 注册邮箱
	regEmailModel, isNotFound, err := s.adminRegEmailRepo.QueryOneByAdminEmail(ctx, in.Email)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}
	if isNotFound {
		reason := adminerrorv1.ERROR_ADMIN_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// admin
	adminModel, isNotFound, err := s.adminRepo.QueryOneByAdminUuid(ctx, regEmailModel.AdminUuid)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return out, err
	}
	if isNotFound {
		reason := adminerrorv1.ERROR_ADMIN_NOT_EXIST.String()
		message := "用户不存在"
		err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

	// 验证密码
	err = passwordutil.Compare(adminModel.PasswordHash, in.Password)
	if err != nil {
		reason := adminerrorv1.ERROR_ADMIN_PASSWORD_INCORRECT.String()
		message := "密码不正确"
		err = errorutil.BadRequest(reason, message)
		//err = errorutil.BadRequest(reason, message, err)
		return out, err
	}

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

// Ping ping pong
func (s *adminAuth) Ping(ctx context.Context, in *adminv1.PingReq) (out *adminv1.PingResp, err error) {
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

	out = &adminv1.PingResp{
		Message: in.Message,
	}
	return out, err
}
