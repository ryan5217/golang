package conf

import (
	"time"
)

var (
	RunMode string = "release"

	HTTPPort     int           = 6000
	ReadTimeout  time.Duration = 60 * time.Second
	WriteTimeout time.Duration = 60 * time.Second

	// gray log
	LOG_UDP           = "udp://g02.graylog.gaodunwangxiao.com:5504"
	IS_SHOW_CONSOLE   = false
	LOG_FIELDS        = "item:milano"
	LOG_FIELDS_MYSQL  = "item:milano_mysql"
	LOG_LEVEL         = -1
	GLIVE_LIST_DOMAIN = "http://gaodun-uploadfiles.oss-cn-hangzhou.aliyuncs.com"
	GLIVE_GAODUN_COM  = "http://glive.gaodun.com"
	APP_ID            = 180603

	// wechat config 杨浦主体 --- 微信配置
	WxGrantType = "client_credential" // 获取access_token填写client_credential
	WxAppid     = "wx33b5aabafc6e073c"
	WxSecret    = "8bfa8a3e18f890dd3e1508b9a95ebcc4"

	// upload host
	Upload_Host = "http://upload.gaodun.com"
	SImg_Host =  "http://simg01.gaodunwangxiao.com"
)

const (
	YouZanClientId = "9083efb5b0199cd682" //有赞云控制台的应用client_id
	YouZanClientSecret = "58160f8521d09c44f6422cb1d5c976e1" //有赞云控制台的应用client_secret
	YouZanKtdId = "13320582" //有赞云商家 id
)
