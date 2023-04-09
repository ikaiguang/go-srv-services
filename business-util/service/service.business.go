package serviceutil

import (
	stderrors "errors"
	"github.com/go-kratos/kratos/v2/transport"
	iputil "github.com/ikaiguang/go-srv-kit/kit/ip"
	"sync"

	registrypkg "github.com/ikaiguang/go-srv-services/pkg/registry"
)

// ServiceName ...
type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

// ServiceInstance ...
type ServiceInstance struct {
	RegistryType registrypkg.RegistryType
	HTTPPort     string
	GRPCPort     string
}

// GetServiceInstance 获取服务实例
func GetServiceInstance(serviceName ServiceName) (serviceInstance *ServiceInstance, isExist bool) {
	serviceInstanceMutex.RLock()
	defer serviceInstanceMutex.RUnlock()

	serviceInstance, isExist = serviceInstances[serviceName]

	return serviceInstance, isExist
}

// GenLocalEndpoint ...
// 例子： 127.0.0.1:11101
func GenLocalEndpoint(serviceName ServiceName, serviceKind transport.Kind) (string, error) {
	serviceInstance, isExist := GetServiceInstance(serviceName)
	if !isExist {
		err := stderrors.New("service instance not found, serviceName: " + serviceName.String())
		return "", err
	}

	var (
		endpoint = iputil.LocalIP() + ":"
	)
	switch serviceKind {
	case transport.KindGRPC:
		endpoint += serviceInstance.GRPCPort
	default:
		endpoint += serviceInstance.HTTPPort
	}
	return endpoint, nil
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
