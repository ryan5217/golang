package wx_program

import (
	"encoding/json"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type WxProgramNotice struct {
	Openid string
	FormId string
	LiveId int
}

const (
	OpenidKey = "milano_openid_key_"
	FormidKey = "milano_formid_key_"
)

func NewWxProgramNotice(openid, formId string, liveId int) *WxProgramNotice {
	return &WxProgramNotice{Openid: openid, FormId: formId, LiveId: liveId}
}

// 发送小程序通知
func (wp *WxProgramNotice) SendNotice(noticeType int) error {
	var data = NewInstantNoticeData(wp.Openid, wp.FormId, wp.LiveId)
	if noticeType == TimeNotice {
		wp.Collect() // 入队列
		return nil
	}
	s, _ := json.Marshal(&data)
	param := req.Param{}
	param["which_wx"] = "jj_cybklm"
	param["moban_id"] = noticeType
	param["keys"] = "keyword1,keyword2,keyword3"
	param["data_value"] = []interface{}{string(s)}
	r, err := req.Post(conf.TORY_DOMAIN+"/wxmb/SendMoreBatch", param)
	setting.Logger.Info("wxprogram_res_" + r.String() + "_param_" + string(s))
	return err
}

// 发送小程序定时通知
func (wp *WxProgramNotice) sendNotice() error {
	data := NewTimeNoticeData(wp.Openid, wp.FormId, wp.LiveId)
	s, _ := json.Marshal(&data)
	param := req.Param{}
	param["which_wx"] = "jj_cybklm"
	param["moban_id"] = TimeNotice
	param["keys"] = "keyword1,keyword2,keyword3"
	param["data_value"] = []interface{}{string(s)}
	r, err := req.Post(conf.TORY_DOMAIN+"/wxmb/SendMoreBatch", param)
	setting.Logger.Info("wxprogram_res_" + r.String() + "_param_" + string(s))
	return err
}

// collect form id & openid to Redis list
func (wp *WxProgramNotice) Collect() {
	r := utils.RedisHandle.RedisClientHandle
	cmd := r.LPush(OpenidKey, wp.Openid)
	r.Expire(OpenidKey, 24*7*time.Hour) // 7 天过期时间
	if cmd.Err() != nil {
		setting.Logger.Info("wxprogram_redis_lpush_" + cmd.Err().Error())
	}
	statusC := r.Set(FormidKey+wp.Openid, wp.FormId, 24*7*time.Hour)
	if statusC.Err() != nil {
		setting.Logger.Info("wxprogram_redis_set_" + statusC.Err().Error())
	}
	return
}

// 发送上课 5 分钟通知
func (wp *WxProgramNotice) Consume() int64 {
	r := utils.RedisHandle.RedisClientHandle
	go func() {
		cmdStr := r.LPop(OpenidKey)
		for cmdStr.Val() != "" {
			openidVal := cmdStr.Val()
			formId := r.Get(FormidKey + openidVal).Val()
			wp.Openid = cmdStr.Val()
			wp.FormId = formId
			wp.sendNotice()
			cmdStr = r.LPop(OpenidKey)
		}
	}()

	return r.LLen(OpenidKey).Val()
}
