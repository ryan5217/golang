package primary

import (
	"fmt"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin/json"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"strings"
	"time"
)

type GoodsApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

func NewGoodsApi(h *utils.HttpClient) *GoodsApi {
	var p = GoodsApi{}
	p.HttpClient = h
	return &p
}

type GoodsDetail struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Result  map[string]interface{}
}
type GoodsBuyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Result  map[string]GoodsBuy
}
type GoodsBuy struct {
	BuyStatus bool   `json:"buy_status"`
	GoodsId   int64  `json:"goods_id"`
	GoodsName string `json:"goods_name"`
}
type GoodsOrder struct {
	Uid       int64
	StudentId int64
	ProjectId int64
	GoodsId   int64
}

//获取商品信息
func (g *GoodsApi) GetGoodsById(GoodsId int64, uid int64, clearCache int64, orderType string) (string, error) {
	var respStr = ""
	redisClient := utils.RedisHandle
	if orderType == "order" {
		key := fmt.Sprintf("GetGoodsById_%d", GoodsId) + "new"
		respStr = redisClient.GetData(key).(string)
		if respStr == "" || clearCache == 1 {
			g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/column/goodsdetailnew?goods_id=%d&type=order&gd_student_id=%d", GoodsId, uid)
			rep, err := req.Get(g.Uri)
			if err != nil {
				setting.Logger.Infof("GetGoodsById_%s", err.Error())
				return respStr, err
			}
			respStr = rep.String()
			redisClient.SetData(key, respStr, time.Hour*1)
		}
	} else {
		key := fmt.Sprintf("GetGoodsById_%d", GoodsId)
		respStr = redisClient.GetData(key).(string)
		if respStr == "" || clearCache == 1 {
			g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/column/goodsdetail?goods_id=%d", GoodsId)
			rep, err := req.Get(g.Uri)
			if err != nil {
				setting.Logger.Infof("GetGoodsById_%s", err.Error())
				return respStr, err
			}
			respStr = rep.String()
			redisClient.SetData(key, respStr, time.Second*86400)
		}
	}
	return respStr, nil
}

//获取商品订单信息
func (g *GoodsApi) GetGoodsOrderList(uid int64, projectId int64, isAudit int64) (*simpleJson.Json, int) {
	g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/order/goodslist?uid=%d&project_id=%d&status=0&is_auditing=%d&order_type=-1", uid, projectId, isAudit)
	rep, err := req.Get(g.Uri)
	if err != nil {
		setting.Logger.Infof("GetGoodsOrderList_uid_%d_projectId_%d", uid, projectId)
		setting.Logger.Infof("GetGoodsOrderList_%s", err.Error())
		return &simpleJson.Json{}, error_code.GOODSORDERLISTAPIERR
	}
	resp, err := simpleJson.NewJson(rep.Bytes())
	if err != nil {
		setting.Logger.Infof("GetGoodsOrderList_%s", err.Error())
		return &simpleJson.Json{}, error_code.GOODSORDERLISTAPIERR
	}
	code, _ := resp.Get("code").Int64()
	if code != 11999999 {
		setting.Logger.Infof("GetGoodsOrderList_%s", rep.String())
		return &simpleJson.Json{}, error_code.GOODSORDERLISTAPIERR
	}
	return resp.Get("result"), error_code.SUCCESSSTATUS
}

//商品是否购买接口
func (g *GoodsApi) GetGoodsBuyList(goodsIds []string, uid int64) (map[string]GoodsBuy, int) {
	code := error_code.SUCCESSSTATUS
	gbr := GoodsBuyResp{}
	param := req.Param{}
	param["goods_id"] = strings.Join(goodsIds, ",")
	param["uid"] = uid
	g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/order/isbuy")
	rep, err := req.Post(g.Uri, param)
	if err != nil {
		setting.Logger.Infof("GetGoodsBuyList_uid_%d_goodsIds_%s", uid, param["goods_id"])
		setting.Logger.Infof("GetGoodsBuyList_%s", err.Error())
		code = error_code.GOODSBUYAPI
		return gbr.Result, code
	}
	rep.ToJSON(&gbr)
	return gbr.Result, code
}

//专栏商品接口
func (g *GoodsApi) GetColumnGoods(columnId int64, projectId int64, clearCache int64) (*simpleJson.Json, int) {
	redisClient := utils.RedisHandle
	key := fmt.Sprintf("GetColumnGoods_%d_%d", columnId, projectId)
	respStr := redisClient.GetData(key).(string)
	if respStr == "" || clearCache == 1 {
		g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/column/goodslist?column_id=%d&project_id=%d", columnId, projectId)
		rep, err := req.Get(g.Uri)
		if err != nil {
			setting.Logger.Infof("GetColumnGoods_%d_%d", columnId, projectId)
			setting.Logger.Infof("GetColumnGoods_%s", err.Error())
			return &simpleJson.Json{}, error_code.GOODSDETAILAPIERR
		}
		respStr = rep.String()
		redisClient.SetData(key, respStr, time.Second*86400)
	}
	res, err := simpleJson.NewJson([]byte(respStr))
	if err != nil {
		setting.Logger.Infof("GetColumnGoods_json_err_%s", respStr)
		setting.Logger.Infof("GetColumnGoods_%s", err.Error())
		return &simpleJson.Json{}, error_code.GOODSDETAILAPIERR
	}
	return res.Get("result"), error_code.SUCCESSSTATUS
}

//商品详情
func (g *GoodsApi) GetGoodsDetail(goodsIds string, clearCache int64) (*simpleJson.Json, int) {
	code := error_code.SUCCESSSTATUS
	redisClient := utils.RedisHandle
	key := fmt.Sprintf("GetGoodsDetail_%s", goodsIds)
	respStr := redisClient.GetData(key).(string)
	if respStr == "" || clearCache == 1 {
		param := req.Param{}
		param["goods_id"] = goodsIds
		g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/column/goodsdetail")
		rep, err := req.Get(g.Uri, param)
		if err != nil {
			setting.Logger.Infof("GetGoodsDetailGoodsIds_%s", goodsIds)
			setting.Logger.Infof("GetGoodsDetail_%s", err.Error())
			return &simpleJson.Json{}, error_code.GOODSDETAILAPIERR
		}
		respStr = rep.String()
		redisClient.SetData(key, respStr, time.Second*86400)
	}
	res, _ := simpleJson.NewJson([]byte(respStr))
	return res, code
}

//购买免费课程
func (g *GoodsApi) GoodsBuy(goodsOrder GoodsOrder) (int, int) {
	code := error_code.SUCCESSSTATUS
	param := req.Param{}
	param["student_id"] = goodsOrder.StudentId
	param["uid"] = goodsOrder.Uid
	param["goods_id"] = goodsOrder.GoodsId
	param["project_id"] = goodsOrder.ProjectId
	g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/order/create")
	rep, err := req.Post(g.Uri, param)
	if err != nil {
		param, _ := json.Marshal(param)
		setting.Logger.Infof("GoodsBuyParam_%s", string(param[:]))
		setting.Logger.Infof("GoodsBuy_%s", err.Error())
		code = error_code.GOODSORDERAPIERR
		return 0, code
	}
	js, _ := simpleJson.NewJson([]byte(rep.String()))
	orderId, _ := js.Get("result").Get("order_id").Int()
	if orderId == 0 {
		setting.Logger.Infof("GoodsBuyResult_%s", rep.String())
		code = error_code.GOODSORDERCREATEERR
		return 0, code
	}
	return orderId, code
}
