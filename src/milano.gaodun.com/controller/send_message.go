package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/service/message"
)

var SendMessageApi = NewSendMessage()

type SendMessage struct {
	Base
}

func NewSendMessage() *SendMessage {
	return &SendMessage{}
}

/**
发送短信和消息通知
证券从业使用
*/
func (s *SendMessage) SendMsg(c *gin.Context) {
	send := message.Send{}
	send.Uid = s.GetInt64(c, "uid", true)
	send.ProjectId = s.GetInt64(c, "project_id", true)
	if c.GetBool(Verify) {
		return
	}
	SendMessageService := message.NewSendMessageService(setting.GinLogger(c))
	SendMessageService.Send(&send)
	s.ServerJSONSuccess(c, 0)
}
