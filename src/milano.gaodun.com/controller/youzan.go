package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/service/crm_order_push"
	"milano.gaodun.com/pkg/error-code"
)

// 有赞订单通知
type Youzan struct {
	Base
}

var YouzanApi = NewYouzanApi()

func NewYouzanApi() *Youzan {
	return &Youzan{}
}

// 有赞消息通知回调接口
func (y Youzan) Accept(c *gin.Context) {
	// 接收有赞订单号
	body := struct {
		Id string `json:"id"`
	}{}
	l := setting.GinLogger(c)
	if err := c.ShouldBindJSON(&body); err != nil {
		l.Info("error_youzan_push_" + err.Error())
	} else {
		l.Info("tid_" + body.Id)
		// 异步处理， 快速响应 youzan 回调
		go func() {
			crm := crm_order_push.NewCrmOrderPush()
			crm.SetLog(setting.GinLogger(c))
			crm.PushOrder(body.Id)
		}()
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"}) // 有赞专用响应信息
}

// 手动推 tid to crm order
func (y *Youzan) PushCrm(c *gin.Context) {
	tid := y.GetString(c, "tid")
	crm := crm_order_push.NewCrmOrderPush()
	crm.SetLog(setting.GinLogger(c))
	if err := crm.PushOrder(tid); err != nil {
		y.ServerJSONOther(c, nil, error_code.FAIL, err.Error())
		return
	}
	y.ServerJSONSuccess(c, nil)
}
