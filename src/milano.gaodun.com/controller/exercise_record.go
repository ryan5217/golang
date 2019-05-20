package controller

import (
	"github.com/gin-gonic/gin"
	er "milano.gaodun.com/model/exercise_record"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	es "milano.gaodun.com/service/exercise_record"
)

var ExerciseRecordApi = NewExerciseRecordApi()

func NewExerciseRecordApi() *ExerciseRecord {
	return &ExerciseRecord{}
}

type ExerciseRecord struct {
	Base
}

func (i ExerciseRecord) AddRecord(c *gin.Context) {
	param := er.TkExerciseRecord{}
	param.Uid = i.GetInt64(c, "uid", true)
	param.ResourceId = i.GetInt64(c, "resource_id", true)
	param.Type = i.GetInt64(c, "type", true)
	param.ChapterName = i.GetString(c, "chapter_name", true)
	param.ModuleId = i.GetInt64(c, "module_id")
	if c.GetBool(Verify) {
		return
	}
	server := es.NewExerciseRecordService(setting.GinLogger(c))
	code := server.AddRecord(&param)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, "添加成功")
	}
	return
}
func (i ExerciseRecord) GetRecord(c *gin.Context) {
	uid := i.GetInt64(c, "uid", true)
	if c.GetBool(Verify) {
		return
	}
	server := es.NewExerciseRecordService(setting.GinLogger(c))
	res, code := server.GetRecord(uid)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, res)
	}
	return
}
