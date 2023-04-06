package setuppkg

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	consulutil "github.com/ikaiguang/go-srv-kit/data/consul"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"

	commonv1 "github.com/ikaiguang/go-srv-services/api/common/v1"
	apputil "github.com/ikaiguang/go-srv-services/business/app"
)

// newConfigWithConsul 初始化配置手柄
func newConfigWithConsul(setupOpts ...Option) (configImpl Config, consulClient *api.Client, err error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	// 启动选项
	setupOpt := &options{}
	for i := range setupOpts {
		setupOpts[i](setupOpt)
	}

	stdlog.Println("|==================== 初始化Consul配置中心 开始 ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== 初始化Consul配置中心 结束 ====================|")

	// 配置路径
	confPath := _configConsulPath
	if setupOpt.configPath != "" {
		confPath = setupOpt.configPath
	}
	stdlog.Println("|*** 加载：Consul初始化配置文件路径: ", confPath)
	configHandler := config.New(config.WithSource(
		file.NewSource(confPath),
	))

	// 加载配置
	if err = configHandler.Load(); err != nil {
		err = pkgerrors.WithStack(err)
		return configImpl, consulClient, err
	}

	// 读取配置文件
	cfg := &commonv1.Bootstrap{}
	if err = configHandler.Scan(cfg); err != nil {
		err = pkgerrors.WithStack(err)
		return configImpl, consulClient, err
	}

	// App配置
	if cfg.App == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : app")
		return configImpl, consulClient, err
	}

	// 服务配置
	if cfg.Base == nil || cfg.Base.Consul == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : base.consul")
		return configImpl, consulClient, err
	}

	// consul客户端
	stdlog.Println("|*** 加载：Consul客户端：for 配置中心")
	consulClient, err = consulutil.NewConsulClient(cfg.Base.Consul)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return configImpl, consulClient, err
	}

	// 配置source
	cs, err := consul.New(consulClient, consul.WithPath(apputil.ConfigPath(cfg.App)))
	if err != nil {
		err = pkgerrors.WithStack(err)
		return configImpl, consulClient, err
	}

	var opts []config.Option
	stdlog.Println("|*** 加载：Consul配置中心的配置: ...")
	opts = append(opts, config.WithSource(cs))

	// config impl
	configImpl, err = NewConfiguration(opts...)
	if err != nil {
		err = pkgerrors.WithStack(err)
		return configImpl, consulClient, err
	}
	return configImpl, consulClient, err
}
