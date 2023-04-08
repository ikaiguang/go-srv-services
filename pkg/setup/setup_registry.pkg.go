package setuppkg

import registrypkg "github.com/ikaiguang/go-srv-services/pkg/registry"

// SetRegistryType 设置 服务注册类型
func (s *engines) SetRegistryType(rt registrypkg.RegistryType) {
	s.registryType = rt
}

// RegistryType 服务注册类型
func (s *engines) RegistryType() registrypkg.RegistryType {
	return s.registryType
}
