package study_rights

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type TkStudyRightsModules struct {
	Id          int64
	ProjectId   int64
	SubjectId   int64
	GradationId int64
	Name        string
	PaperType   int64
	Sort        int64
	CreatedTime string
	UpdatedTime string
}
type StudyRightsModulesModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewStudyRightsModulesModel() *StudyRightsModulesModel {
	return &StudyRightsModulesModel{Engine: utils.NewtikuDb}
}
func (g *StudyRightsModulesModel) Get(id int64) (*TkStudyRightsModules, error) {
	srm := TkStudyRightsModules{}
	_, err := g.Id(id).Get(&srm)
	return &srm, err
}
func (g *StudyRightsModulesModel) List(param SearchParam) (*[]TkStudyRightsModules, error) {
	srm := []TkStudyRightsModules{}
	err := g.Where("project_id=?", param.ProjectId).
		Where("subject_id=?", param.SubjectId).
		Asc("sort").Find(&srm)
	return &srm, err
}
