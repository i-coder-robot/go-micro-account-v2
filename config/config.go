package config

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	source := consul.NewSource(
		//config center address
		consul.WithAddress(host + ":" + strconv.FormatInt(port, 10)),
		//默认 /micro/config
		consul.WithPrefix(prefix),
		//是否移除前缀 true:可以不带前缀，直接得到对应的配置
		consul.StripPrefix(true),
	)
	newConfig, err := config.NewConfig()
	if err != nil {
		return newConfig,err
	}
	err = config.Load(source)
	return newConfig,err
}
