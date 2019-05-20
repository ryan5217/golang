package crm_order_push

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api/user"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"milano.gaodun.com/service/youzan"
	"strings"
)

var OrderPath = "/OrderCenter/SyncOrder"

// 往 crm 推送 有赞 订单
type CrmOrderPush struct {
	l *log.Entry
}

func NewCrmOrderPush() *CrmOrderPush {
	c := new(CrmOrderPush)
	c.l = setting.GrayLog()
	return c
}

func (c *CrmOrderPush) SetLog(l *log.Entry) {
	c.l = l
}

// crm 订单推送 形参: 有赞订单 id
func (c *CrmOrderPush) PushOrder(tid string) error {
	yzClient := youzan.NewYouzanToken()
	token, err := yzClient.GetAccessToken()
	if err != nil {
		c.l.Info("error_youzan_token_" + err.Error())
		return err
	}
	orderRes, err := youzan.NewYouzanOrder(token).GetOrder(tid)
	if err != nil {
		c.l.Info("error_youzan_order_" + err.Error())
		return err
	}
	if len(orderRes.Response.FullOrderInfo.Orders) < 1 {
		c.l.Info("error_youzan_order_没有商品")
		return fmt.Errorf("error_youzan_order_没有商品")
	}
	goodsId := orderRes.Response.FullOrderInfo.Orders[0].ItemId
	goodsRes, err := youzan.NewYouzanGoods(token).GetGoodsInfo(goodsId)
	if err != nil {
		c.l.Info("error_youzan_goods_" + err.Error())
		return err
	}
	if len(goodsRes.Response.Item.ItemTags) < 1 {
		c.l.Info("error_youzan_order_没有项目或不需要推送的项目")
		return fmt.Errorf("error_youzan_order_没有项目或不需要推送的项目")
	}
	buyer := struct {
		Phone string `json:"手机号(开课凭证)"`
	}{}
	var trueName, phone string
	phone = orderRes.Response.FullOrderInfo.BuyerInfo.BuyerPhone
	if err := json.Unmarshal([]byte(orderRes.Response.FullOrderInfo.Orders[0].BuyerMessages), &buyer); err == nil && buyer.Phone != "" {
		phone = buyer.Phone
		trueName = "youzan" + phone + orderRes.Response.FullOrderInfo.BuyerInfo.FansNickname
	} else {
		c.l.Info("error_youzan_buyer_不是<手机号(开课凭证)>")
		trueName = "youzan" + orderRes.Response.FullOrderInfo.BuyerInfo.BuyerPhone + orderRes.Response.FullOrderInfo.BuyerInfo.FansNickname
	}
	userId, err := c.getUser(phone, trueName)
	if err != nil {
		return err
	}
	// 项目 id
	projectId := youzan.ProjectMap[goodsRes.Response.Item.ItemTags[0].Name]

	productList := []string{}
	for _, val := range orderRes.Response.FullOrderInfo.Orders {
		productList = append(productList, val.OuterItemId)
	}

	crmPost := newPostData()
	crmPost.ProductList = productList
	crmPost.StudentInfo.UcenterUid = userId
	crmPost.StudentInfo.TrueName = trueName
	crmPost.StudentInfo.Telphone = phone
	crmPost.StudentInfo.CertificateNo = phone
	crmPost.StudentInfo.StartTime = orderRes.Response.FullOrderInfo.OrderInfo.PayTime
	crmPost.StudentInfo.Create_Time = orderRes.Response.FullOrderInfo.OrderInfo.PayTime
	crmPost.OrderInfo.CourseType = projectId
	crmPost.OrderInfo.SellPrise = orderRes.Response.FullOrderInfo.PayInfo.Payment
	crmPost.OrderInfo.DiscountPrise = 0
	crmPost.ClueInfo.FollowRecords = phone + crmPost.ClueInfo.FollowRecords
	ser := Serial{}
	ser.OrderNo = tid
	ser.PaySerialNo = strings.Join(orderRes.Response.FullOrderInfo.PayInfo.Transaction, ",")
	ser.PayCount = orderRes.Response.FullOrderInfo.PayInfo.Payment
	ser.PayType = CrmPayType
	ser.PayTypeName = CrmPayTypeName
	ser.PayTime = orderRes.Response.FullOrderInfo.OrderInfo.PayTime
	crmPost.SerialList = append(crmPost.SerialList, ser)
	crmPost.VOrderNo = tid

	c.buildCrmOrder(crmPost)
	return nil
}

// 返回用户 id
func (c *CrmOrderPush) getUser(phone, name string) (int, error) {
	password := "123456"
	u := user.NewUserApi(utils.HttpHandle)
	res, err := u.RegisterAndFindByPhone(phone, "youzan-"+name, password)
	if err != nil {
		c.l.Info("error_youzan_sso_user_" + err.Error())
	}
	return res.Data.StudentId, err
}

//
func (c *CrmOrderPush) buildCrmOrder(crmPost PostData) {
	p, _ := json.Marshal(crmPost)
	l := c.l.WithField("crm_order", string(p))
	if res, err := req.Post(conf.BAIYI_DOMAIN+OrderPath, req.BodyJSON(crmPost)); err == nil {
		s, _ := res.ToString()
		l.Info(s)
	} else {
		l.Info("error_youzan_crm_request_" + err.Error())
	}
}
