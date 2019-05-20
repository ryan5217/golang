package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/service/wx-program"
	"time"
)

var WPApi = NewWxProgramApi()
var liveIdVar = 9163

func NewWxProgramApi() *WxProgram {
	return &WxProgram{}
}

type WxProgram struct {
	Base
}

func (wp *WxProgram) Notice(c *gin.Context) {
	liveId := wp.GetInt(c, "live_id")
	if liveId == 0 {
		liveId = liveIdVar
	}
	noticeType := wp.GetInt(c, "notice_type", true)
	openid := wp.GetString(c, "openid", true)
	formId := wp.GetString(c, "form_id", true)
	if c.GetBool(Verify) {
		return
	}
	n := wx_program.NewWxProgramNotice(openid, formId, liveId)
	n.SendNotice(noticeType)
	wp.ServerJSONSuccess(c, nil)
}

func (wp *WxProgram) NoticeTime(c *gin.Context) {
	liveId := wp.GetInt(c, "live_id")
	if liveId == 0 {
		liveId = liveIdVar
	}
	openid := wp.GetString(c, "openid")
	formId := wp.GetString(c, "form_id")
	if c.GetBool(Verify) {
		return
	}
	n := wx_program.NewWxProgramNotice(openid, formId, liveId)
	l := n.Consume()
	wp.ServerJSONSuccess(c, l)
}

func (wp *WxProgram) GetLive(c *gin.Context) {
	liveId := wp.GetInt(c, "live_id")
	if liveId == 0 {
		liveId = liveIdVar
	}
	res, _ := wx_program.GetLiveDesc(liveId)
	start := res.Data.Starttime
	m := make(map[string]interface{})
	m["is_start"] = false
	m["live_id"] = liveIdVar
	m["data"] = res.Data
	if start-300 < time.Now().Unix() {
		m["is_start"] = true
	}

	wp.ServerJSONSuccess(c, m)
}


// 获取 access token
func (wp *WxProgram) GetAccessToken(c *gin.Context) {
	w := wx_program.NewWechatApi(setting.GinLogger(c))
	s, err := w.GetAccessToken()
	if err != nil {
		wp.ServerJSONError(c, err.Error(), error_code.FAIL)
		return
	}
	wp.ServerJSONSuccess(c, s)
}

// 获取新闻列表
func (wp *WxProgram) GetMaterialList(c *gin.Context) {
	w := wx_program.NewWechatApi(setting.GinLogger(c))
	offset := wp.GetInt(c, "page")
	count := wp.GetInt(c, "limit")
	res, err := w.GetMaterialList(offset, count)
	if err != nil {
		wp.ServerJSONError(c, err.Error(), error_code.FAIL)
	}
	wp.ServerJSONSuccess(c, res)
}

// 获取新闻列表
func (wp *WxProgram) DelMaterial(c *gin.Context) {
	w := wx_program.NewWechatApi(setting.GinLogger(c))
	id := wp.GetInt(c, "id", true)
	if c.GetBool(Verify) {
		return
	}
	res, err := w.DelMaterial(id)
	if err != nil {
		wp.ServerJSONError(c, err.Error(), error_code.FAIL)
	}
	wp.ServerJSONSuccess(c, res)
}

func (wp *WxProgram) InitWechatMaterial(c *gin.Context) {
	w := wx_program.NewWechatApi(setting.GinLogger(c))
	if err := w.InitWechatMaterial(); err != nil {
		wp.ServerJSONOther(c, nil, error_code.FAIL, err.Error())
		return
	}
	wp.ServerJSONSuccess(c, nil)
}
