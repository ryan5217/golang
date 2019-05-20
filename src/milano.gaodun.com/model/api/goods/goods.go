package goods

import (
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type GoodsApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}
type GoodsResApi struct {
	Code    int64
	Message string
	Result  map[string]GoodsBuy
}
type GoodsBuy struct {
	GoodsId   int64  `json:"goods_id"`
	GoodsName string `json:"goods_name"`
	BuyStatus bool   `json:"buy_status"`
}

func NewGoodsApi(h *utils.HttpClient) *GoodsApi {
	var p = GoodsApi{}
	p.HttpClient = h
	return &p
}
func (g *GoodsApi) GetGoodsBuyList(uid int64, goods_ids string) (*GoodsResApi, error) {
	var res GoodsResApi
	reqParam := req.QueryParam{
		"uid":      uid,
		"goods_id": goods_ids,
	}
	g.Uri = conf.SPARTA_DOMAIN + "/v1/order/isbuy"
	r, err := req.Post(g.Uri, reqParam)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
