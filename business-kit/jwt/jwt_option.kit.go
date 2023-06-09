// Package jwtutil 摘自kratos子项目
package jwtutil

import (
	"github.com/golang-jwt/jwt/v4"

	authutil "github.com/ikaiguang/go-srv-services/business-kit/auth"
)

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod jwt.SigningMethod
	claims        func() jwt.Claims
	tokenHeader   map[string]interface{}
	validator     authutil.ValidateFunc
}

// WithSigningMethod with signing method option.
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithClaims with customer claim
// If you use it in Server, f needs to return a new jwt.Claims object each time to avoid concurrent write problems
// If you use it in Client, f only needs to return a single object to provide performance
func WithClaims(f func() jwt.Claims) Option {
	return func(o *options) {
		o.claims = f
	}
}

// WithTokenHeader withe customer tokenHeader for client side
func WithTokenHeader(header map[string]interface{}) Option {
	return func(o *options) {
		o.tokenHeader = header
	}
}

// WithValidator 自定义验证
func WithValidator(validator authutil.ValidateFunc) Option {
	return func(o *options) {
		o.validator = validator
	}
}
