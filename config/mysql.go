package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
)

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	Database string `json:"database"`
	Port int64 `json:"port"`
}

func GetMysql4Consul(config config.Config,path ...string) *MysqlConfig {
	c:= &MysqlConfig{}
	err := config.Get(path...).Scan(c)
	if err != nil {
		fmt.Println(err)
	}
	return c

}
