package controller

import (
	"github.com/gin-gonic/gin"
	mm "milano.gaodun.com/model/message"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	mService "milano.gaodun.com/service/message"
)

var MessageApi = NewMessageApi()

func NewMessageApi() *Message {
	return &Message{}
}

type Message struct {
	Base
}

func (i Message) Modify(c *gin.Context) {
	param := mm.TkSystemMessage{}
	param.ProjectId = i.GetInt64(c, "project_id", true)
	param.Id = i.GetInt64(c, "id")
	param.Title = i.GetString(c, "title", true)
	param.Url = i.GetString(c, "url", true)
	param.Uid = i.GetInt64(c, "uid") //指定用户推送
	param.ForceUpdateCol = i.PostMustCols(c)
	if c.GetBool(Verify) {
		return
	}
	bServer := mService.NewMessageService(setting.GinLogger(c))
	_, err := bServer.ModifySystemMessage(&param)
	if err != nil {
		i.ServerJSONError(c, "", error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, "")
	}
}
func (i Message) GetNotifyMessageNum(c *gin.Context) {
	param := mm.Param{}
	param.ProjectId = i.GetInt64(c, "project_id", true)
	param.Uid = i.GetInt64(c, "uid", true)
	if c.GetBool(Verify) {
		return
	}

	bServer := mService.NewMessageService(setting.GinLogger(c))
	res, err := bServer.GetNotifyNum(param)
	if err != nil {
		i.ServerJSONError(c, []int{}, error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}
}
func (i Message) GetSystemMessages(c *gin.Context) {
	param := mm.Param{}
	param.ProjectId = i.GetInt64(c, "project_id")
	param.Limit = i.GetInt(c, "limit")
	param.Page = i.GetInt(c, "page")
	if c.GetBool(Verify) {
		return
	}
	if param.Limit < 1 {
		param.Limit = 10
	}
	if param.Page < 1 {
		param.Page = 1
	}
	bServer := mService.NewMessageService(setting.GinLogger(c))
	res, err := bServer.GetSystemMessages(param)
	if err != nil {
		i.ServerJSONError(c, []int{}, error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}
}
func (i Message) Delete(c *gin.Context) {
	id := i.GetInt64(c, "id", true)
	isdel := i.GetInt(c, "is_del")
	if c.GetBool(Verify) {
		return
	}
	bServer := mService.NewMessageService(setting.GinLogger(c))
	res, err := bServer.DeleteSystemMessage(id, isdel)
	if err != nil {
		i.ServerJSONError(c, []int{}, error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}
}
func (i Message) DeleteNotification(c *gin.Context) {
	id := i.GetInt64(c, "id", true)
	uid := i.GetInt64(c, "uid", true)
	msgType := i.GetInt64(c, "type")
	if c.GetBool(Verify) {
		return
	}
	bServer := mService.NewMessageService(setting.GinLogger(c))
	res, err := bServer.DeleteNotifyMessage(id, uid, msgType)
	if err != nil {
		i.ServerJSONError(c, []int{}, error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}
}
func (i Message) ReadMessage(c *gin.Context) {
	id := i.GetInt64(c, "id", true)
	uid := i.GetInt64(c, "uid", true)
	msgType := i.GetInt64(c, "type")
	if c.GetBool(Verify) {
		return
	}
	bServer := mService.NewMessageService(setting.GinLogger(c))
	res, err := bServer.ReadMessage(id, uid, msgType)
	if err != nil {
		i.ServerJSONError(c, []int{}, error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}
}
func (i Message) List(c *gin.Context) {
	param := mm.Param{}
	param.ProjectId = i.GetInt64(c, "project_id", true)
	param.Uid = i.GetInt64(c, "uid", true)
	param.Limit = i.GetInt(c, "limit")
	param.Page = i.GetInt(c, "page")
	param.Type = i.GetInt(c, "type")
	if c.GetBool(Verify) {
		return
	}
	if param.Limit < 1 {
		param.Limit = 10
	}
	if param.Page < 1 {
		param.Page = 1
	}
	bServer := mService.NewMessageService(setting.GinLogger(c))
	res, err := bServer.List(param)
	if err != nil {
		i.ServerJSONError(c, []int{}, error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}
}
