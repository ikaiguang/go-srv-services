package setuputil

import (
	"github.com/ikaiguang/go-srv-services/business-kit/registry"
)

// SetRegistryType 设置 服务注册类型
func (s *engines) SetRegistryType(rt registryutil.RegistryType) {
	s.registryType = rt
}

// RegistryType 服务注册类型
func (s *engines) RegistryType() registryutil.RegistryType {
	return s.registryType
}
