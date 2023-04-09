package services

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	logutil "github.com/ikaiguang/go-srv-kit/log"

	userv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/resources"
	userservicev1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/services"
	assemblers "github.com/ikaiguang/go-srv-services/app/user-service/internal/application/assembler"
	srvs "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/service"
	authutil "github.com/ikaiguang/go-srv-services/business-util/auth"
)

// userAuth ...
type userAuth struct {
	userservicev1.UnimplementedSrvUserAuthServer

	log       *log.Helper
	assembler *assemblers.Assembler
	authSrv   *srvs.UserAuthSrv
}

// NewUserAuthService ...
func NewUserAuthService(
	logger log.Logger,
	assembler *assemblers.Assembler,
	authSrv *srvs.UserAuthSrv,
) userservicev1.SrvUserAuthServer {
	return &userAuth{
		log:       log.NewHelper(log.With(logger, "module", "user/application/service/user_auth")),
		assembler: assembler,
		authSrv:   authSrv,
	}
}

// LoginByEmail Email登录
func (s *userAuth) LoginByEmail(ctx context.Context, in *userv1.LoginByEmailReq) (out *userv1.LoginResp, err error) {
	// 注册邮箱
	regEmailModel, err := s.authSrv.CheckLoginEmail(ctx, in.Email)
	if err != nil {
		return out, err
	}

	// user
	userModel, err := s.authSrv.CheckLoginUserByUUID(ctx, regEmailModel.UserUuid)
	if err != nil {
		return out, err
	}

	// 验证密码
	err = s.authSrv.ComparePassword(userModel.PasswordHash, in.Password)
	if err != nil {
		return out, err
	}

	// 签证
	authInfo := s.assembler.AuthInfo(userModel)
	signedString, err := s.authSrv.SignToken(ctx, userModel, authInfo)
	if err != nil {
		return out, err
	}

	out = s.assembler.LoginResp(authInfo, signedString)
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
