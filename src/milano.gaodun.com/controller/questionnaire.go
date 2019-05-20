package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	QM "milano.gaodun.com/model/questionnaire"
	"milano.gaodun.com/pkg/setting"
	QS "milano.gaodun.com/service/questionnaire"
	"strconv"
)

var QuestionnaireApi = NewQuestionnaireApi()

func NewQuestionnaireApi() *Questionnaire {
	return &Questionnaire{}
}

type Questionnaire struct {
	Base
}

func (q Questionnaire) Add(c *gin.Context) {
	qu := QM.PriQuestionnaire{}
	qu.ItemId = q.GetString(c, "item_id")
	qu.ItemName = q.GetString(c, "item_name")
	qu.TypeId = q.GetInt64(c, "type_id")
	qu.Choice = q.GetString(c, "choice")
	qu.Evaluate = q.GetString(c, "evaluate")
	qu.Comment = q.GetString(c, "comment")
	bServer := QS.NewQuestionnaireService(setting.GinLogger(c))
	bServer.AddQuestionnaire(&qu)
	q.ServerJSONSuccess(c, qu)
}
func (q Questionnaire) Edit(c *gin.Context) {
	qu := QM.PriQuestionnaire{}
	qu.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	qu.ItemId = q.GetString(c, "item_id")
	qu.ItemName = q.GetString(c, "item_name")
	qu.TypeId = q.GetInt64(c, "type_id")
	qu.Choice = q.GetString(c, "choice")
	qu.Evaluate = q.GetString(c, "evaluate")
	qu.Comment = q.GetString(c, "comment")
	bServer := QS.NewQuestionnaireService(setting.GinLogger(c))
	qu.ForceUpdateCol = q.PostMustCols(c)
	bServer.Edit(&qu)
	q.ServerJSONSuccess(c, qu)
}
