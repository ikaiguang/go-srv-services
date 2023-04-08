package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	consulutil "github.com/ikaiguang/go-srv-kit/data/consul"
	filepathutil "github.com/ikaiguang/go-srv-kit/kit/filepath"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
	"os"
	"path/filepath"

	commonv1 "github.com/ikaiguang/go-srv-services/api/common/v1"
	apppkg "github.com/ikaiguang/go-srv-services/pkg/app"
)

var (
	_conf string
)

func init() {
	flag.StringVar(&_conf, "conf", "", "config path, eg: -conf ./configs")
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	var err error
	defer func() {
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
	}()

	// 配置
	bootConfig, err := loadingConfig()
	if err != nil {
		return
	}

	// consul
	consulHandler, err := NewConsulConfig(bootConfig)
	if err != nil {
		return
	}

	// 开始配置
	stdlog.Println("|==================== 更新配置到Consul 开始 ====================|")
	defer stdlog.Println("|==================== 更新配置到Consul 结束 ====================|")
	err = consulHandler.StoreToConsul()
	if err != nil {
		return
	}
}

// loadingConfig 加载配置
func loadingConfig() (*commonv1.Bootstrap, error) {
	handler := config.New(config.WithSource(file.NewSource(_conf)))
	err := handler.Load()
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	var conf = &commonv1.Bootstrap{}
	err = handler.Scan(conf)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	// App配置
	if conf.App == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : app")
		return nil, err
	}

	// 服务配置
	if conf.Server == nil {
		err = pkgerrors.New("[请配置服务再启动] config key : server")
		return nil, err
	}
	return conf, nil
}

// ConsulConfig ...
type ConsulConfig struct {
	config *commonv1.Bootstrap
	cc     *api.Client
	path   string
}

// NewConsulConfig 初始化
func NewConsulConfig(config *commonv1.Bootstrap) (*ConsulConfig, error) {
	if config.Base == nil || config.Base.Consul == nil {
		err := pkgerrors.New("请先配置Consul配置再试")
		return nil, err
	}
	cc, err := consulutil.NewClient(config.Base.Consul)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	absPath, err := filepath.Abs(_conf)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	return &ConsulConfig{
		config: config,
		cc:     cc,
		path:   absPath,
	}, nil
}

// StoreToConsul 存到consul
func (s *ConsulConfig) StoreToConsul() error {
	configDataM, err := s.ReadConfigFiles()
	if err != nil {
		return err
	}
	ctx := context.Background()
	opt := &api.WriteOptions{}
	opt = opt.WithContext(ctx)
	for key := range configDataM {
		kv := &api.KVPair{
			Key:   key,
			Value: configDataM[key],
		}
		_, err := s.cc.KV().Put(kv, opt)
		if err != nil {
			return pkgerrors.WithStack(err)
		}
	}
	return nil
}

// ReadConfigFiles 读取文件
func (s *ConsulConfig) ReadConfigFiles() (map[string][]byte, error) {
	fs, err := filepathutil.ReadDir(s.path)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	consulPath := apppkg.ConfigPath(s.config.App)
	configDataM := make(map[string][]byte)
	for i := range fs {
		if fs[i].IsDir() {
			continue
		}
		destPath := filepath.Join(consulPath, fs[i].Name())
		filePath := filepath.Join(s.path, fs[i].Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, pkgerrors.WithStack(err)
		}
		configDataM[destPath] = content
		//fmt.Println(destPath, len(content))
	}
	return configDataM, nil
}
