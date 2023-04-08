package serviceutil

// ServerName 服务名称
type (
	ServerName string
	ServerPort int
)

// Value ...
func (s ServerName) Value() string {
	return string(s)
}

// Value ...
func (s ServerPort) Value() int {
	return int(s)
}

const (
	AdminService         ServerName = "admin-service"
	AdminServiceHTTPPort ServerPort = 11101
	AdminServiceGRPCPort ServerPort = 11102

	UserService         ServerName = "user-service"
	UserServiceHTTPPort ServerPort = 11201
	UserServiceGRPCPort ServerPort = 11202
)
