package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

//设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀，不设置默认前缀/micro/config
		consul.WithPrefix(prefix),
		//不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	//配置初始化
	cfg, err := config.NewConfig()
	if err != nil {
		return cfg, err
	}
	err = cfg.Load(consulSource)
	return cfg, err
}
