package consul

import (
	"gitlab.gaodun.com/golib/consul"
)

// GdConfig 配置文件
var GdConsul = InitConfig()

// InitConfig 配置文件
func InitConfig() map[string]string {
	config, err := consul.GetConf("")
	if err != nil {
		panic(err)
	}
	return config
}
