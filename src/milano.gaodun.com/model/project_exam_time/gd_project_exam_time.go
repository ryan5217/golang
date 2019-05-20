package project_exam_time

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type GdProjectExamTime struct {
	Id          int64
	ProjectId   int64
	ExamTime    int64
	ExamEndTime int64
}
type ProjectExamTimeModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewProjectExamTimeModel() *ProjectExamTimeModel {
	return &ProjectExamTimeModel{Engine: utils.GaodunDb3}
}
func (g *ProjectExamTimeModel) Get(projectId int64) (*GdProjectExamTime, int) {
	srm := GdProjectExamTime{}
	_, err := g.Where("project_id=?", projectId).Where("isdel=0").Where("exam_time>?", time.Now().Unix()).OrderBy("exam_time asc").Get(&srm)
	if err != nil {
		setting.Logger.Infof("GdProjectExamTime_%s", err.Error())
		return &srm, error_code.EXAMMODELERR
	}
	return &srm, error_code.SUCCESSSTATUS
}
