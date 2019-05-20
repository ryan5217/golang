package conf

import (
	"log"
	"milano.gaodun.com/pkg/consul"
)

var (

	// base key
	BASE_TIKU_KEY   = "gd_tk_app_kjzc"
	BASE_Source_KEY = "gd_tk_app_nosource"
	BaseKey         = map[string]string{
		BASE_TIKU_KEY:   "gdtkappkjzc11lei32xiao232peng4180",
		BASE_Source_KEY: "gdtkappnosourcelei32xiao237peng413",
	}

	BaseCoinKey = map[string]string{
		"app_id":     "gd_gaodun_coin",
		"app_secret": "0123456789abcdefghdsfww2lzarstuvwxyz",
	}
	// base domain
	BASE_DOMAIN = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "base.gaodun.com"
	// base domain
	SSO_DOMAIN            = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "sso.gaodun.com"
	MTIKU_DOMAIN          = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "mtiku.gaodun.com"
	MILANO_DOMAIN         = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "milano.gaodun.com"
	SPARTA_DOMAIN         = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "sparta.gaodun.com"
	STUDY_SERVICE_DOMAIN  = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "study-service.gaodun.com"
	TOC_SERVICE_DOMAIN    = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "toc-service.gaodun.com"
	COURSE_SERVICE_DOMAIN = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "course-service.gaodun.com"
	STUDYAPI_DOMAIN       = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "studyapi.gaodun.com"
	PROMETHEUS_DOMAIN     = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "prometheus.gaodun.com"
	UPLOAD_DOMAIN         = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "upload.gaodun.com"
	SIMG_DOMAIN           = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "simg01.gaodunwangxiao.com"
	BAIYI_DOMAIN          = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "baiyiapi.gaodun.com"
	MUSES_DOMAIN          = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "muses.gaodun.com"
	NOTICE_DOMAIN         = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "notice.gaodunwangxiao.com"
	NOTICE_DOMAIN_ONLINE  = "http://notice.gaodunwangxiao.com"
	RESOURCE_DOMAIN       = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "resource-service.gaodun.com"
	//MTIKU_DOMAIN = "http://" + "t-" + "mtiku.gaodun.com"
	// gd live  lachesis
	//LACHESIS_DOMAIN = "http://" + consul.GdConsul["PUBLIC_DEPLOY_ENV_SIMG"] + "lachesis.gaodun.com"
	LACHESIS_DOMAIN = "http://lachesis.gaodun.com"
	// 微信小程序通知
	// http://troy.gaodun.com/wxmb/SendMoreBatch
	TORY_DOMAIN = "http://troy.gaodun.com"
	// gd info domain
	GD_INFO_DOMAIN = "https://gadmin.gaodun.com"
)

const (
	YouZanHost = "https://open.youzan.com"
)

func GetAppSecret(key string) string {
	if v, ok := BaseKey[key]; ok {
		return v
	}
	log.Print("key is not find")

	return ""
}
