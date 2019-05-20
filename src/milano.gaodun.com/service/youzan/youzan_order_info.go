package youzan

import (
	"fmt"
	"time"
)

var (
	order_method = "youzan.trade.get" // 获取订单详情：固定值
	appVersion   = "4.0.0"
)

type YouzanOrder struct {
	YouzanApi
}

// 初始化
func NewYouzanOrder(token string) *YouzanOrder {
	y := YouzanOrder{}
	y.method = order_method
	y.token = token
	y.timeOut = time.Second * 20
	return &y
}

// 获取订单信息
func (y *YouzanOrder) GetOrder(tid string) (*OrderResponse, error) {
	orderRes := OrderResponse{}
	param := y.param()
	param["tid"] = tid
	uri := y.url(y.method, appVersion)
	res, err := y.post(uri, param)
	if err != nil {
		return &orderRes, err
	}
	if err := res.ToJSON(&orderRes); err != nil {
		return &orderRes, err
	}
	if orderRes.ErrorResponse.Code != 0 {
		return &orderRes, fmt.Errorf(orderRes.ErrorResponse.Msg)
	}
	return &orderRes, nil
}
