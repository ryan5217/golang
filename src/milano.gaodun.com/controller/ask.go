package controller

import (
	"github.com/gin-gonic/gin"
	aApi "milano.gaodun.com/model/api/ask"
	askM "milano.gaodun.com/model/ask"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	aService "milano.gaodun.com/service/ask"
)

var AskApi = NewAskApi()

func NewAskApi() *Ask {
	return &Ask{}
}

type Ask struct {
	Base
}

func (i Ask) List(c *gin.Context) {
	askParam := aApi.AskParam{}
	askParam.Uid = i.GetInt64(c, "uid", true)
	askParam.CourseId = i.GetInt64(c, "course_id", true)
	askParam.SourceType = "saas"
	askParam.Page = i.GetInt64(c, "page")
	askParam.Limit = i.GetInt64(c, "limit")
	if askParam.Page == 0 || askParam.Limit == 0 {
		askParam.Page = 1
		askParam.Limit = 15
	}
	bServer := aService.NewAskService(setting.GinLogger(c))
	bServer.GetAskList(&askParam)
	askList, err := bServer.GetAskList(&askParam)
	if err != nil {
		i.ServerJSONError(c, err, error_code.SYSERR)
	} else {
		i.ServerJSONSuccess(c, askList)
	}
}
func (i Ask) AskDetail(c *gin.Context) {
	askId := i.GetInt64(c, "ask_id", true)
	bServer := aService.NewAskService(setting.GinLogger(c))
	res, err := bServer.GetAskDetail(askId)

	if err != nil {
		i.ServerJSONError(c, err, error_code.SYSERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}

}
func (i Ask) AnswerAsk(c *gin.Context) {
	aParam := askM.GdAskQuestions{}
	aParam.Content = i.GetString(c, "content", true)
	aParam.StudentId = i.GetInt64(c, "uid", true)
	aParam.AskId = i.GetInt64(c, "ask_id", true)
	aParam.Pid = i.GetInt64(c, "answer_id", true)
	aParam.FileUrl = i.GetString(c,"file_url")
	bServer := aService.NewAskService(setting.GinLogger(c))
	res, err := bServer.AnswerAsk(&aParam)
	if err != nil {
		i.ServerJSONError(c, err, error_code.SYSERR)
	} else {
		i.ServerJSONSuccess(c, res)
	}

}
