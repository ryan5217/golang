package youzan

import (
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"strings"
	"time"
)

const (
	Api_Uri = "/api/oauthentry" // uri：固定值
)

var ProjectMap = getProjectFunc()

type ResponseError struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubCode int    `json:"sub_code"`
		SubData string `json:"sub_data"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response, omitempty"`
}

type OrderResponse struct {
	ResponseError
	Response struct {
		FullOrderInfo struct {
			PayInfo   PayInfo   `json:"pay_info"`
			BuyerInfo BuyerInfo `json:"buyer_info"`
			Orders    []Orders  `json:"orders"`
			OrderInfo OrderInfo `json:"order_info"`
		} `json:"full_order_info"`
	} `json:"response"`
}

// 订单明细结构体
type Orders struct {
	OuterItemId   string `json:"outer_item_id"`  // 商品码
	Title         string `json:"title"`          // 商品名
	Payment       string `json:"payment"`        // 价格
	ItemId        int    `json:"item_id"`        // 商品 id
	BuyerMessages string `json:"buyer_messages"` // 开课帐号
}

// 交易明细详情
type OrderInfo struct {
	Tid     string `json:"tid"`      // 订单号
	Created string `json:"created"`  // 创建时间
	PayTime string `json:"pay_time"` //支付时间
}

// 支付流水
type PayInfo struct {
	Transaction []string `json:"transaction"`
	Payment     string   `json:"payment"`
}

// 买家信息
type BuyerInfo struct {
	OuterUserId  string `json:"outer_user_id"`
	BuyerPhone   string `json:"buyer_phone"`
	FansNickname string `json:"fans_nickname"`
}

/*
{
"error_response": {
"code": 40010,
"msg": "参数 token 无效",
"sub_code": 10000,
"sub_data": "",
"sub_msg": "token无效，该token不存在或已过期"
}
}
*/

// 项目名对应表
func getProjectFunc() map[string]int {
	m := map[string]int{}
	m["初级职称"] = 14
	m["证券从业"] = 9
	m["基金从业"] = 38
	m["CPA"] = 8
	return m
}

// 单个商品结构体
type GoodsResponse struct {
	ResponseError
	Response struct {
		Item struct {
			ItemTags []struct { // 分组
				Name string `json:"name"` // 项目名称
			} `json:"item_tags"`
		} `json:"item"`
	} `json:"response"`
}

// base
type YouzanApi struct {
	method  string
	token   string
	timeOut time.Duration
}

func (y *YouzanApi) param() map[string]interface{} {
	param := req.Param{}
	param["access_token"] = y.token
	param["method"] = y.method
	return param
}

// url
func (y *YouzanApi) url(method, appVersion string) string {
	methodArray := strings.Split(method, ".")
	method = "/" + appVersion + "/" + methodArray[len(methodArray)-1]
	method = strings.Join(methodArray[0:len(methodArray)-1], ".") + method
	return conf.YouZanHost + Api_Uri + "/" + method
}

func (y *YouzanApi) post(url string, param map[string]interface{}) (*req.Resp, error) {
	req.SetTimeout(y.timeOut)
	return req.Post(url, req.Param(param))
}
