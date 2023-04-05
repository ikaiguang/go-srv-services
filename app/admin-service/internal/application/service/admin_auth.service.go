package services

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	logutil "github.com/ikaiguang/go-srv-kit/log"

	adminv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/resources"
	adminservicev1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/services"
	assemblers "github.com/ikaiguang/go-srv-services/app/admin-service/internal/application/assembler"
	srvs "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/service"
	authutil "github.com/ikaiguang/go-srv-services/business/auth"
)

// adminAuth ...
type adminAuth struct {
	adminservicev1.UnimplementedSrvAdminAuthServer

	log       *log.Helper
	assembler *assemblers.Assembler
	authSrv   *srvs.AdminAuthSrv
}

// NewAdminAuthService ...
func NewAdminAuthService(
	logger log.Logger,
	assembler *assemblers.Assembler,
	authSrv *srvs.AdminAuthSrv,
) adminservicev1.SrvAdminAuthServer {
	return &adminAuth{
		log:       log.NewHelper(log.With(logger, "module", "admin/application/service/admin_auth")),
		assembler: assembler,
		authSrv:   authSrv,
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

	// 签证
	authInfo := s.assembler.AuthInfo(adminModel)
	signedString, err := s.authSrv.SignToken(ctx, adminModel, authInfo)
	if err != nil {
		return out, err
	}

	out = s.assembler.LoginResp(authInfo, signedString)
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
