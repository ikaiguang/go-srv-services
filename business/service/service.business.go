package serviceutil

import (
	"sync"

	registrypkg "github.com/ikaiguang/go-srv-services/pkg/registry"
)

// ServiceName ...
type ServiceName string

// ServiceInstance ...
type ServiceInstance struct {
	RegistryType registrypkg.RegistryType
	HTTPPort     string
	GRPCPort     string
}

// GetServiceInstance 获取服务实例
func GetServiceInstance(serviceName ServiceName) (serviceInstance *ServiceInstance, isExist bool) {
	serviceInstanceMutex.RLock()
	defer serviceInstanceMutex.Unlock()

	serviceInstance, isExist = serviceInstances[serviceName]

	return serviceInstance, isExist
}

// GenLocalEndpoint ...
func GenLocalEndpoint(serviceName ServiceName) {

}

const (
	AdminService ServiceName = "admin-service" // 定义后记得添加 serviceInstances
	UserService  ServiceName = "user-service"  // 定义后记得添加 serviceInstances
)

var (
	serviceInstanceMutex = sync.RWMutex{}
	serviceInstances     = map[ServiceName]*ServiceInstance{
		AdminService: {
			HTTPPort: "11101",
			GRPCPort: "11102",
		},
		UserService: {
			HTTPPort: "11201",
			GRPCPort: "11202",
		},
	}
)
