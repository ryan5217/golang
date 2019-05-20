package youzan

import (
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"fmt"
	"milano.gaodun.com/pkg/utils"
	"time"
)

var (
	token_uri = "/oauth/token"
)

const (
	TokenKey = "youzan_key_"
)

type YouzanTokenResp struct {
	Error            string `json:"error"` // 错误码 如果错误的话
	ErrorDescription string `json:"error_description"` // 错误码描述

	AccessToken string `json:"access_token"` // token
	ExpiresIn   int    `json:"expires_in"` // 有效期 7 天
	Scope       string `json:"scope"` // 相应权限
}

// youzan token api
type YouzanToken struct {
	clientId     string
	clientSecret string
	kdtId        string
}

// 默认初始化
func NewYouzanToken() *YouzanToken {
	y := new(YouzanToken)
	y.clientId = conf.YouZanClientId
	y.clientSecret = conf.YouZanClientSecret
	y.kdtId = conf.YouZanKtdId
	return y
}

// 重新初始化
func (y *YouzanToken) InitConf(clientId, clientSecret, kdtId string) {
	y.clientId = clientId
	y.clientSecret = clientSecret
	y.kdtId = kdtId
}

// token 存缓存中 6 天
func (y *YouzanToken) GetAccessToken() (string, error) {
	r := utils.RedisHandle.RedisClientHandle
	if token := r.Get(TokenKey).Val(); token != "" {
		return token, nil
	}

	param := req.Param{}
	param["grant_type"] = "silent"
	param["kdt_id"] = y.kdtId
	param["client_id"] = y.clientId
	param["client_secret"] = y.clientSecret
	h := req.Header{"user-agent": "X-YZ-Client 2.0.0 - PHP"}
	token, err := y.pareResponse(req.Post(conf.YouZanHost + token_uri, param, h))
	if err == nil && token != "" {
		r.Set(TokenKey, token, time.Hour * 24 * 6) // 6 天
	}

	return token, err
}

// 解析 res
func (y *YouzanToken) pareResponse(res *req.Resp, e error) (string, error) {
	r := YouzanTokenResp{}
	if e != nil {
		return "", e
	}
	if err := res.ToJSON(&r); err != nil {
		return "", err
	}
	if r.Error == "" {
		return r.AccessToken, nil
	}

	return r.AccessToken, fmt.Errorf(r.ErrorDescription)
}
