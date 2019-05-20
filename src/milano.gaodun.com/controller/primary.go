package controller

import (
	"github.com/gin-gonic/gin"
	priM "milano.gaodun.com/model/api/primary"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	priService "milano.gaodun.com/service/primary"
)

var PrimaryApi = NewPrimaryApi()

func NewPrimaryApi() *Primary {
	return &Primary{}
}

type Primary struct {
	Base
}

func (p Primary) GetGlive(c *gin.Context) {
	projectId := p.GetInt64(c, "project_id")
	source := p.GetString(c, "source")
	pri := priService.NewGliveService(setting.GinLogger(c))
	data, _ := pri.GetGlive(projectId, source)
	if data.GliveId > 0 {
		p.ServerJSONSuccess(c, &data)
	} else {
		p.ServerJSONSuccess(c, []priService.GliveView{})
	}
}
func (p Primary) StageSet(c *gin.Context) {
	user_stage := priM.PriUserStage{}
	user_stage.Id = p.GetInt64(c, "id")
	user_stage.Locked = p.GetInt(c, "locked")
	user_stage.StudentId = p.GetInt64(c, "student_id")
	user_stage.Stage = p.GetInt(c, "stage")
	if user_stage.StudentId <= 0 {
		p.ServerJSONError(c, "", error_code.MUSTFIELD)
		return
	}
	us_service := priService.NewUserStageService(setting.GinLogger(c))
	data, _ := us_service.Save(&user_stage)
	p.ServerJSONSuccess(c, &data)
	return
}
func (p Primary) GetStageSet(c *gin.Context) {
	studentId := p.GetInt64(c, "student_id")
	if studentId <= 0 {
		p.ServerJSONError(c, "", error_code.MUSTFIELD)
		return
	}
	us_service := priService.NewUserStageService(setting.GinLogger(c))
	data, _ := us_service.GetStage(studentId)
	if data.Id > 0 {
		p.ServerJSONSuccess(c, &data)
	} else {
		p.ServerJSONSuccess(c, "")
	}

	return
}

//获取笔记接口
func (p Primary) GetNoteList(c *gin.Context) {
	param := priM.SearchParams{}
	param.StudentId = p.GetInt64(c, "uid", true)
	param.CourseId = p.GetInt64(c, "course_id", true)
	param.ResourceId = p.GetInt64(c,"resource_id")
	param.Page = p.GetInt64(c, "page")
	param.Limit = p.GetInt64(c, "limit")
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Limit <= 0 {
		param.Limit = 15
	}
	noteService := priService.NewNoteService(setting.GinLogger(c))
	nt, err := noteService.GetNoteList(param)
	if err != nil {
		p.ServerJSONError(c, err, error_code.SYSERR)
		return
	}
	p.ServerJSONSuccess(c, nt)
}
//新增笔记接口
func (p Primary) Add(c *gin.Context) {
	param := priM.NoteData{}
	param.Student_id = p.GetInt64(c, "uid", true)
	param.Course_id = p.GetInt64(c, "course_id", true)
	param.Content = p.GetString(c, "content",true)
	param.Live_seconds = p.GetInt64(c, "live_seconds")
	param.Is_public = p.GetInt64(c, "is_public")
	param.Resource_id = p.GetInt64(c, "resource_id",true)
	param.Resource_type = p.GetString(c, "resource_type",true)
	param.Origin = p.GetString(c, "origin",true)
	param.Course_syllabus_item_id = p.GetInt64(c,"course_syllabus_item_id",true)
	noteService := priService.NewNoteService(setting.GinLogger(c))
	nt, err := noteService.AddNote(&param)
	if err != nil {
		p.ServerJSONError(c, err, error_code.SYSERR)
		return
	}
	if nt.Status != 0 {
		p.ServerJSONError(c, nt.Message, error_code.HTTPERR)
		return
	}
	p.ServerJSONSuccess(c, nt.Result)
}
//修改笔记接口
func (p Primary) Edit(c *gin.Context) {
	param := priM.Note{}
	param.Content = p.GetString(c, "content",true)
	param.IsPublic = p.GetInt64(c, "is_public")
	param.SourceId = p.GetInt64(c, "resource_id",true)
	param.SourceType = p.GetString(c, "resource_type",true)
	param.Uid = p.GetInt64(c, "uid",true)
	param.Id = p.GetInt64(c,"id",true)
	noteService := priService.NewNoteService(setting.GinLogger(c))
	nt, err := noteService.EditNote(&param)
	if err != nil {
		p.ServerJSONError(c, err, error_code.SYSERR)
		return
	}
	if nt.Status != 0 {
		p.ServerJSONError(c, nt.Message, error_code.HTTPERR)
		return
	}
	p.ServerJSONSuccess(c, nt.Result)
}