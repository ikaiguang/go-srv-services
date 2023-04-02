# go-srv-services

基于 [go-srv-kit](http://github.com/ikaiguang/go-srv-kit) 实现的微服务

## 运行服务

```shell
# 例子：运行 admin-service
make run service=admin-service
# 或
go run ./app/admin-service/cmd/main/... -conf=./app/admin-service/configs
```

## 数据库迁移

```shell
# 例子：迁移 admin-service 数据库
make migrate service=admin-service
# 或
go run ./app/admin-service/cmd/migration/... -conf=./app/admin-service/configs
```

## 编译PB协议文件

```shell
# 例子：编译 admin-service
make api service=admin-service
# 或
go run ./cmd/proto/... -path=api/admin-service
```

## 编译配置文件

```shell
# 例子：编译 admin-service 配置文件
make config service=admin-service
# 或
go run ./cmd/proto/... -path=./app/admin-service/internal/conf
```
