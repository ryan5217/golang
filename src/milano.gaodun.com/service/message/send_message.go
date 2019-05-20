package message

import (
	"encoding/json"
	"github.com/apex/log"
	simpleJson "github.com/bitly/go-simplejson"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api/sms"
	"milano.gaodun.com/model/members_student"
	"milano.gaodun.com/model/message"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"milano.gaodun.com/service/tiku_constant"
)

type sendMessageService struct {
	logger *log.Entry
}

/**
短信的消息
*/
type SendSMS struct {
	Message   string `json:"message"`
	Phone     string `json:"phone"`
	Ip        string `json:"ip"`
	AppId     int    `json:"appId"`
	Type      string `json:"type"`
	SessionId string `json:"sessionId"`
	Uid       int64  `json:"uid"`
	TplCode   string `json:"tplCode"`
}

/**
模板ID
*/
type SMS155860947 struct {
	Course string `json:"course"`
	Time   string `json:"time"`
}

/**
系统消息通知
*/
type SendMessage struct {
	Title     string `json:"title"`
	Url       string `json:"url"`
	ProjectId int64  `json:"project_id"`
	Uid       int64  `json:"uid"`
}

/**
系统消息通知
*/
type Send struct {
	SessionId string `json:"session_id"`
	Uid       int64  `json:"uid"`
	ProjectId int64  `json:"project_id"`
	Ip        string `json:"ip"`
}

func NewSendMessageService(logger *log.Entry) *sendMessageService {
	return &sendMessageService{
		logger: logger,
	}
}

/**
发送消息
*/
func (s sendMessageService) Send(param *Send) int {
	sendSms := SendSMS{}
	Sms := SMS155860947{}
	sendMessage := SendMessage{}
	jsonData, error := s.GetKeyList("CONFIG_FOR_TEMP_ZQCY_ATY_APP")
	if error > 0 {
		return error
	}

	//为用户发送短信
	Sms.Course, _ = jsonData.Get("shareTitle").String()
	Sms.Time, _ = jsonData.Get("date_time").String()
	utils.ChangeStruct2OtherStruct(param, sendSms)
	//获取用户信息
	student := members_student.NewMembersStudentModel()
	studentInfo, _ := student.Get(param.Uid)
	if studentInfo.Phone == "" {
		return error_code.USER_PHONE_NOT_FOUND
	}
	sendSms.Phone = studentInfo.Phone
	sendSms.SessionId = studentInfo.Phone
	Message, _ := json.Marshal(Sms)
	sendSms.Message = string(Message)
	error = s.SendSMS(&sendSms)
	//为用户推送消息
	sendMessage.Uid = param.Uid
	sendMessage.ProjectId = param.ProjectId
	sendMessage.Url, _ = jsonData.Get("info_url").String()
	sendMessage.Title, _ = jsonData.Get("info_title").String()
	s.SendMessage(&sendMessage)
	if error > 0 {
		return error
	} else {
		return 0
	}
}

/**
发送短信消息
*/
func (s sendMessageService) SendSMS(param *SendSMS) int {
	sendApi := sms.SendApi{}
	_, code := sendApi.Send(conf.NOTICE_DOMAIN_ONLINE+"/api/v1/sendmessage", param.Phone, param.Message)
	if code > 0 {
		return code
	} else {
		return 0
	}
}

/**
发送系统消息
*/
func (s sendMessageService) SendMessage(param *SendMessage) int {
	systemMsg := message.TkSystemMessage{}
	systemMsg.ProjectId = param.ProjectId
	systemMsg.Url = param.Url
	systemMsg.Title = param.Title
	systemMsg.Uid = param.Uid
	id, _ := message.NewSystemMessageModel().Modify(&systemMsg)
	if id > 0 {
		return 0
	} else {
		return error_code.SYSTEM_INFO_SEND_FAILD
	}
}

/**
获取常量信息列表
*/
func (s sendMessageService) GetKeyList(key string) (*simpleJson.Json, int) {
	tikuContent := tiku_constant.NewTikuConstantList()
	data, err := tikuContent.GetKey(key)
	if err != nil {
		setting.Logger.Errorf("getSysTemContent" + err.Error())
	}
	list, error := simpleJson.NewJson([]byte(data))
	if error != nil {
		setting.Logger.Errorf("json_Error" + error.Error())
	}
	return list, 0
}
