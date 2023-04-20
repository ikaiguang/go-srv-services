# 配置说明

每个服务的配置文件都位于该服务根目录，如`./app/xxx-service/configs/xxx.yaml`

为了更好的区分配置类型，特意拆分为`四个`配置

* config_app.yaml : 服务主配置，如端口、名称、日志。。。
* config_base.yaml : 服务的基础配置，如服务注册与发现、服务追踪、雪花算法。。。
* config_business.yaml : 服务的业务配置，如登录、验证码。。。
* config_data.yaml : 服务的数据配置，如数据库MySQL、Redis、MongoDB。。。

## 关于配置说明

* 配置目录：`./app/xxx-service/configs`，用于常规开发、测试
* 配置目录：`./app/xxx-service/configs/consul`，用于主要开发，使用配置中心的配置启动服务
    * 启动服务命令：`make dev service=xxx-service`
    * 或：`go run ./app/${service}/cmd/main/... -conf-consul=./app/${service}/configs/consul`
* 配置目录：`./app/xxx-service/configs/develop`，用于主要开发，配置开启服务注册发现、服务追踪。。。
    * 也会同步此配置到服务配置中心`Consul`
    * 同步命令：`make consul-config service=xxx-service`
    * 或：`go run ./cmd/consul-config/... -conf=./app/${service}/configs/develop`

## 关于服务端口

更新/修改服务配置`./app/xxx-service/configs/config_app.yaml`的端口，

请同步更新以下文件：

1. 同服务下各个配置，如：`./app/xxx-service/configs/develop/config_app.yaml`
2. 把更新后的服务同步到`Consul`配置中心，运行命令：`make consul-config service=xxx-service`
3. 更新业务代码端口，方便本地调试；代码文件：`./business-util/service/service.business.go:serviceInstances`
