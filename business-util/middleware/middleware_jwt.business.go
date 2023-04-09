package middlewareutil

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwt "github.com/golang-jwt/jwt/v4"

	errorutil "github.com/ikaiguang/go-srv-kit/error"

	authutil "github.com/ikaiguang/go-srv-services/business-util/auth"
	setuppkg "github.com/ikaiguang/go-srv-services/pkg/setup"
)

// NewWhiteListMatcher 路由白名单
func NewWhiteListMatcher(whiteList map[string]struct{}) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		//if tr, ok := contextutil.MatchHTTPServerContext(ctx); ok {
		//	if _, ok := whiteList[tr.Request().URL.Path]; ok {
		//		return false
		//	}
		//}

		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewJWTMiddleware jwt中间
func NewJWTMiddleware(engineHandler setuppkg.Engine, whiteList map[string]struct{}) (m middleware.Middleware, err error) {
	// redis
	redisCC, err := engineHandler.GetRedisClient()
	if err != nil {
		return m, errorutil.WithStack(err)
	}
	authTokenRepo := engineHandler.GetAuthTokenRepo(redisCC)

	m = selector.Server(
		authutil.Server(
			authTokenRepo.JWTKeyFunc(),
			authutil.WithSigningMethod(authTokenRepo.JWTSigningMethod()),
			authutil.WithClaims(func() jwt.Claims { return &authutil.Claims{} }),
			authutil.WithValidator(authTokenRepo.ValidateFunc()),
		),
	).
		Match(NewWhiteListMatcher(whiteList)).
		Build()

	return m, err
}
