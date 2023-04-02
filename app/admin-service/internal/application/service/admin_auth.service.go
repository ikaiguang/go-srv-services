package services

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	authutil "github.com/ikaiguang/go-srv-kit/kratos/auth"
	logutil "github.com/ikaiguang/go-srv-kit/log"

	adminv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/resources"
	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/services"
	srvs "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/service"
)

// adminAuth ...
type adminAuth struct {
	adminservicev1.UnimplementedSrvAdminAuthServer

	log     *log.Helper
	authSrv *srvs.AdminAuthSrv
}

// NewAdminAuthService ...
func NewAdminAuthService(
	logger log.Logger,
	authSrv *srvs.AdminAuthSrv,
) adminservicev1.SrvAdminAuthServer {
	return &adminAuth{
		log:     log.NewHelper(logger),
		authSrv: authSrv,
	}
}

// LoginByEmail Email登录
func (s *adminAuth) LoginByEmail(ctx context.Context, in *adminv1.LoginByEmailReq) (out *adminv1.LoginResp, err error) {
	// 注册邮箱
	regEmailModel, err := s.authSrv.CheckLoginEmail(ctx, in.Email)
	if err != nil {
		return out, err
	}

	// admin
	adminModel, err := s.authSrv.CheckLoginUserByUUID(ctx, regEmailModel.AdminUuid)
	if err != nil {
		return out, err
	}

	// 验证密码
	err = s.authSrv.ComparePassword(adminModel.PasswordHash, in.Password)
	if err != nil {
		return out, err
	}

	return s.authSrv.SignToken(ctx, adminModel)
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
