package graylog

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"net"
	"os"
)

var (
	ProjectName     = "gaodun-test"
	Endpoint        = "cn-hangzhou.log.aliyuncs.com"
	LogStoreName    = "gaodun-logstore"
	AccessKeyID     = "LTAIOE8CPOZ6AFrx"
	AccessKeySecret = "q6QtSFklTJQGZ56om8Am25XsbQJYkV"
	Client          *GdAliLogger
	LocalIPV4       string
	SlsClient       *sls.Client
	SysEnv          = getEnv()
)

func init() {
	ProjectName = ProjectName
	AccessKeyID = AccessKeyID
	AccessKeySecret = AccessKeySecret
	Endpoint = Endpoint
	LogStoreName = LogStoreName
	SlsClient = &sls.Client{
		Endpoint:        Endpoint,
		AccessKeyID:     AccessKeyID,
		AccessKeySecret: AccessKeySecret,
	}
	i, _ := LocalIPv4s()
	if len(i) > 0 {
		LocalIPV4 = i[0]
	}
}

// 获取本地 ip4 地址
func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

//GetEnv 获取 consul address
func getEnv() string {
	env := os.Getenv("SYSTEM_ENV")

	if env == "production" {
		return ""
	}

	if env == "" {
		return "dev-"
	}

	return env + "-"
}
