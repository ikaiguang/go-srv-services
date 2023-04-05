package apputil

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	stdhttp "net/http"
	"strings"

	"google.golang.org/grpc/codes"

	commonv1 "github.com/ikaiguang/go-srv-services/api/common/v1"
)

const (
	OK = 0

	baseContentType = "application"
)

var (
	_ = http.DefaultRequestDecoder
	_ = http.DefaultErrorEncoder
	_ = http.DefaultResponseEncoder
	_ = http.DefaultResponseDecoder
)

// ID 程序ID
// @result = app.Name + app.Env + app.Branch + app.Version
func ID(appConfig *confv1.App) string {
	identifier := appConfig.Name
	identifier += ":" + ParseEnv(appConfig.Env).String()
	if appConfig.EnvBranch != "" {
		branchString := strings.Replace(appConfig.EnvBranch, " ", ":", -1)
		identifier += ":" + branchString
	}
	if appConfig.Version != "" {
		identifier += ":" + appConfig.Version
	}
	return identifier
}

// ParseEnv ...
func ParseEnv(appEnv string) (envEnum commonv1.EnvEnum_Env) {
	envInt32, ok := commonv1.EnvEnum_Env_value[strings.ToUpper(appEnv)]
	if ok {
		envEnum = commonv1.EnvEnum_Env(envInt32)
	}
	if envEnum == commonv1.EnvEnum_UNKNOWN {
		envEnum = commonv1.EnvEnum_PRODUCTION
		return envEnum
	}
	return envEnum
}

// IsDebugMode ...
func IsDebugMode(appEnv commonv1.EnvEnum_Env) bool {
	switch appEnv {
	case commonv1.EnvEnum_DEVELOP, commonv1.EnvEnum_TESTING:
		return true
	default:
		return false
	}
}

// IsSuccessCode 成功的响应码
func IsSuccessCode(code int32) bool {
	if code == OK {
		return true
	}
	return IsSuccessHTTPCode(int(code))
}

// IsSuccessHTTPCode 成功的HTTP响应吗
func IsSuccessHTTPCode(code int) bool {
	if code >= stdhttp.StatusOK && code < stdhttp.StatusMultipleChoices {
		return true
	}
	return false
}

// IsSuccessGRPCCode 成功的GRPC响应吗
func IsSuccessGRPCCode(code uint32) bool {
	return codes.Code(code) == codes.OK
}

// ToResponseError 转换为错误
func ToResponseError(response ResponseInterface) *errors.Error {
	return &errors.Error{
		Status: errors.Status{
			Code:     response.GetCode(),
			Reason:   response.GetReason(),
			Message:  response.GetMessage(),
			Metadata: response.GetMetadata(),
		},
	}
}

// HTTPError 转换为错误
func HTTPError(code int, message string) *errors.Error {
	return &errors.Error{
		Status: errors.Status{
			Code:    int32(code),
			Reason:  commonv1.ERROR_REQUEST_FAILED.String(),
			Message: message,
			Metadata: map[string]string{
				"status": stdhttp.StatusText(code),
			},
		},
	}
}
