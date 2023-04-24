package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	snowflakeservicev1 "github.com/ikaiguang/go-srv-services/api/snowflake-service/v1/services"
	assemblers "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/application/assembler"
	services "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/application/service"
	srvs "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/domain/service"
	datas "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/infra/data"
	"github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/setup"
	snowflakeutil "github.com/ikaiguang/go-srv-services/business-kit/snowflake"
	"github.com/patrickmn/go-cache"
	stdlog "log"
	"time"
)

// RegisterSnowflakeRoutes 注册路由
func RegisterSnowflakeRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	stdlog.Println("|*** 注册路由：SnowflakeNodeID")

	// 数据库
	dbConn, err := engineHandler.GetMySQLGormDB()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	// config
	nodeIDConfig := engineHandler.SnowflakeNodeIDConfig()

	// cache
	var locker snowflakeutil.Locker
	cacheHandler := cache.New(5*time.Minute, 10*time.Minute)
	locker = snowflakeutil.NewLockerFromCache(cacheHandler)
	if nodeIDConfig.EnableRedisLocker {
		redisCC, err := engineHandler.GetRedisClient()
		if err != nil {
			logutil.Fatal(err)
			return
		}
		locker = snowflakeutil.NewLockerFromRedis(redisCC)
	}

	// repos
	snowflakeWorkerRepo := datas.NewSnowflakeNodeIDRepo(dbConn)

	snowflakeOptions := srvs.Options{
		MaxNodeID:    nodeIDConfig.MaxNodeId,
		IdleDuration: nodeIDConfig.IdleDuration.AsDuration(),
	}
	snowflakeSrv := srvs.NewSnowflakeSrv(snowflakeOptions, locker, snowflakeWorkerRepo)

	// 服务
	assembler := assemblers.NewAssembler()
	srv := services.NewSnowflakeNodeID(assembler, snowflakeSrv)

	snowflakeservicev1.RegisterSrvSnowflakeNodeIDHTTPServer(hs, srv)
	snowflakeservicev1.RegisterSrvSnowflakeNodeIDServer(gs, srv)
}
