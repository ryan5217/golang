package sms

import (
	"github.com/imroc/req"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"net"
)

type SendApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type SendResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (s SendApi) Send(url string, phone string, message string) (string, int) {
	p := req.Param{}
	p["phone"] = phone
	p["appId"] = "130444"
	p["type"] = "2"
	p["ip"] = s.ClientIP()
	p["tplCode"] = "SMS_155860947"
	p["sessionId"] = phone
	p["message"] = message
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
		"ORIGIN":       "gaodun.com",
	}
	r, err := req.Post(url, header, p)
	if err != nil {
		setting.Logger.Infof("SendSMS%s", err.Error())
		return "", error_code.TREELISTAPI
	}
	return r.String(), 0
}

func (s SendApi) ClientIP() string {
	ip := "127.0.0.1"
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
