package study_rights

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type TkStudyRightsGradation struct {
	Id          int64
	ProjectId   int64
	SubjectId   int64
	Sort        int64
	Name        string
	CreatedTime string
	UpdatedTime string
}
type StudyRightsGradationModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewStudyRightsGradationModel() *StudyRightsGradationModel {
	return &StudyRightsGradationModel{Engine: utils.NewtikuDb}
}
func (g *StudyRightsGradationModel) Get(id int64) (*TkStudyRightsGradation, error) {
	srm := TkStudyRightsGradation{}
	_, err := g.Id(id).Get(&srm)
	return &srm, err
}
func (g *StudyRightsGradationModel) List(param SearchParam) (*[]TkStudyRightsGradation, error) {
	srg := []TkStudyRightsGradation{}
	err := g.Where("project_id=?", param.ProjectId).
		Where("subject_id=?", param.SubjectId).
		Asc("sort").Find(&srg)
	return &srg, err
}
