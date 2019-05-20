package youzan

import (
	"fmt"
	"time"
)

var (
	goods_method = "youzan.item.get" // 获取订单详情：固定值
	appGoodsVersion = "3.0.0"
)

type YouzanGoods struct {
	YouzanApi
}

// 初始化
func NewYouzanGoods(token string) *YouzanGoods {
	y := YouzanGoods{}
	y.method = goods_method
	y.token = token
	y.timeOut = time.Second * 20

	return &y
}

func (y *YouzanGoods) GetGoodsInfo(itemId int) (*GoodsResponse, error) {
	goodsRes := GoodsResponse{}
	param := y.param()
	param["item_id"] = itemId
	uri := y.url(y.method, appGoodsVersion)
	res, err := y.post(uri, param)
	if err != nil {
		return &goodsRes, err
	}
	if err := res.ToJSON(&goodsRes); err != nil {
		return &goodsRes, err
	}
	if goodsRes.ErrorResponse.Code != 0 {
		return &goodsRes, fmt.Errorf(goodsRes.ErrorResponse.Msg)
	}
	return &goodsRes, nil
}
