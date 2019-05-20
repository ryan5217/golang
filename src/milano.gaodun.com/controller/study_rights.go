package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	srModel "milano.gaodun.com/model/study_rights"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	srService "milano.gaodun.com/service/study_rights"
	"strings"
)

var StudyRightsApi = NewStudyRightsApi()

func NewStudyRightsApi() *StudyRights {
	return &StudyRights{}
}

type StudyRights struct {
	Base
}

func (i StudyRights) GetList(c *gin.Context) {
	param := srModel.SearchParam{}
	param.ProjectId = i.GetInt64(c, "project_id", true)
	param.SubjectId = i.GetInt64(c, "subject_id", true)
	param.Uid = i.GetInt64(c, "uid")
	bServer := srService.NewStudyRightsService(setting.GinLogger(c))
	srRes, err := bServer.GetList(param)
	if err != nil {
		i.ServerJSONError(c, err, error_code.SYSERR)
	} else {
		i.ServerJSONSuccess(c, srRes)
	}
}
func (i StudyRights) GetOne(c *gin.Context) {
	id := i.GetInt64(c, "id", true)
	bServer := srService.NewStudyRightsService(setting.GinLogger(c))
	srRes, err := bServer.GetOne(id)
	if err != nil {
		i.ServerJSONError(c, err, error_code.SYSERR)
	} else {
		i.ServerJSONSuccess(c, srRes)
	}
}
func (i StudyRights) GetSubjectList(c *gin.Context) {
	projectIds := i.GetString(c,"project_id",true)
	ProjectIdList := strings.Split(projectIds,",")
	bServer := srService.NewStudyRightsService(setting.GinLogger(c))
	srRes, err := bServer.GetSubjectList(ProjectIdList)
	if err != nil {
		i.ServerJSONError(c, err, error_code.SYSERR)
	} else {
		i.ServerJSONSuccess(c, srRes)
	}
}
