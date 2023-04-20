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

# 服务端口号和全局错误码枚举值区间划分

### 划分目的

1. 端口号的划分是为了更好的管理服务,避免在本地运行时发生端口冲突.
2. Kratos中错误码推荐用英文表示,在用枚举定义全局错误码时,划分好各个服务枚举值区间可以有效减少发生冲突的概率

### 划分原则:

1. 服务端口号用5位数字表示,避免与常用组件端口号冲突,最高不超65535
2. 服务编号用3位数表示,server为奇数,service为偶数
3. 最后1位用于区分http和grpc,1表示http协议端口,2表示grpc协议端口
4. 错误码用7位数字表示,其中前3位与服务序号保持一致,后4位用于表示具体错误码的枚举值. (注意:
   错误码枚举值在errors.proto有唯一约束,但在业务和代码中无意义,不应对外暴露)

| 服务               | 服务编号 | http惯用端口号 | grpc惯用端口号 | 错误码枚举区间         |      
|:-----------------|:-----|:----------|:----------|:----------------|
| ping-service     | 110  | 11001     | 11002     | 1100001-1109999 |
| snowflake-server | 111  | 11101     | 11102     | 1110001-1119999 |
| admin-server     | 112  | 11201     | 11202     | 1120001-1129999 |
| user-service     | 113  | 11301     | 11302     | 1130001-1139999 |

