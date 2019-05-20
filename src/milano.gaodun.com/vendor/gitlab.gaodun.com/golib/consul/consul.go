//Package consul package
package consul

import (
	"net/http"
	"os"
	_ "reflect"
	"strings"
	"sync"
	"time"

	"github.com/armon/consul-api"
	"github.com/bernos/go-retry"
	"gitlab.gaodun.com/golib/graylog"
)

var p sync.RWMutex

// AddGrayLog 日志
func AddGrayLog(info string) {
	m := make(map[string]interface{})
	p.Lock()
	m["item"] = "consul"
	p.Unlock()
	graylog.GdLog(info, m)
}

//ConsulConfig 获取 consul 地址
func ConsulConfig() *consulapi.Config {
	consulAddr, err := GetEnv()
	if err != nil {
		AddGrayLog("consule first " + err.Error())
	}
	return &consulapi.Config{
		Address:    consulAddr,
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	}
}

// RetryConsul 重试
func RetryConsul(kv *consulapi.KV) func() (interface{}, error) {
	return func() (interface{}, error) {
		pair, _, err := kv.List("gaodun/config_center/", nil)
		if err != nil {
			AddGrayLog("consule first " + err.Error())
			return nil, err
		}
		return pair, nil
	}
}

// 获取consul数据 (key: value)
func getConsulVal(consulConfig *consulapi.Config) (info map[string]string, err error) {
	info = make(map[string]string)
	client, err := consulapi.NewClient(consulConfig)

	if err != nil {
		AddGrayLog("consule first " + err.Error())
		return info, err
	}

	kv := client.KV()
	r := retry.Retry(RetryConsul(kv),
		retry.MaxRetries(3),
		retry.BaseDelay(time.Millisecond*1000))
	pair, err := r()
	if err != nil {
		return nil, err
	}

	if err != nil {
		AddGrayLog("consule first " + err.Error())
		return info, err
	}
	pa := pair.(consulapi.KVPairs)
	// 遍历key, value
	for _, item := range pa {
		keyList := strings.Split(item.Key, "/")
		keyNum := len(keyList)
		itemKey := keyList[keyNum-1]
		info[string(itemKey)] = string(item.Value)
	}

	return info, nil
}

// GetConf 获取配置项
// return map[string]string
func GetConf(projectName string) (minfo map[string]string, err error) {
	return getConsulVal(ConsulConfig())
}

//GetEnv 获取 consul address
func GetEnv() (string, error) {

	consuls := map[string]string{
		"dev":        "dev.consul.gaodunwangxiao.com",
		"test":       "t.consul.gaodunwangxiao.com",
		"prepare":    "pre.consul.gaodunwangxiao.com",
		"production": "pro.consul.gaodunwangxiao.com",
	}

	env := os.Getenv("SYSTEM_ENV")
	if env == "" {
		return consuls["dev"], nil
	}

	return consuls[env], nil
}

//GetEnvPath 获取 缓存目录环境
func GetEnvPath() string {

	consuls := map[string]string{
		"dev":        "dev-",
		"test":       "t-",
		"prepare":    "pre-",
		"production": "",
	}

	env := os.Getenv("SYSTEM_ENV")
	if env == "" {
		return consuls["dev"]
	}
	return consuls[env]
}
