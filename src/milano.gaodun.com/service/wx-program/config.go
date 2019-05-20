package wx_program

import (
	"time"
	"fmt"
)

const (
	TimeNotice  = 1
	InstantNotice  = 2
	PagePath = "pages/activity/activityDetail/activityDetail?liveId="
)


type DataValue struct {
	TopColor string `json:"top_color"`
	Keyword1Val string `json:"keyword1_val"`
	Keyword2Val string `json:"keyword2_val"`
	Keyword3Val string `json:"keyword3_val"`
	Openid string `json:"openid"`
	FormId string `json:"form_id"`
	Page string `json:"page"`
	EmphasisKeyword string `json:"emphasis_keyword"`
}

func NewTimeNoticeData(openid, formId string, liveId int) *DataValue {
	data := DataValue{}
	liveData, _ := GetLiveDesc(liveId)
	data.Openid = openid
	data.FormId = formId
	data.Page = PagePath + fmt.Sprintf("%d", liveId)
	data.Keyword1Val = liveData.Data.Title
	data.Keyword2Val = UnixToForm(liveData.Data.Starttime)
	data.Keyword3Val = "您预约的直播即将开始，点击卡片进入直播间。"

	return &data
}

func NewInstantNoticeData(openid, formId string, liveId int) *DataValue {
	data := DataValue{}
	liveData, _ := GetLiveDesc(liveId)
	data.Openid = openid
	data.FormId = formId
	//data.Page = PagePath + fmt.Sprintf("%d", liveId)
	data.Keyword1Val = liveData.Data.Title
	data.Keyword2Val = UnixToForm(liveData.Data.Starttime)
	data.Keyword3Val = "恭喜您预约成功，别忘了按时来上课哦!"

	return &data
}

func UnixToForm(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

