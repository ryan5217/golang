package activity

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

func NewActivityModel(p *ActivityParam) *ActivityModel {
	return &ActivityModel{Engine: utils.GaodunPrimaryDb, p: p}
}

type ActivityModel struct {
	*xorm.Engine
	s *xorm.Session
	p *ActivityParam
}

func (a *ActivityModel) List() ([]PriActivity, error) {
	rowsSlicer := []PriActivity{}
	err := a.WhereParam().Find(&rowsSlicer)

	return rowsSlicer, err
}

func (a ActivityModel) Add(m *PriActivity) (int64, error) {
	return a.Cols(a.p.ForceUpdateCol...).Insert(m)
}

func (a ActivityModel) IsExist() (bool, error) {
	return a.WhereParam().Exist(&PriActivity{})
}

// search condition
func (a *ActivityModel) WhereParam() *xorm.Session {

	if a.s == nil {
		a.s = a.NewSession()
		defer a.s.Close()
	}
	if a.p.Id > 0 {
		a.s = a.s.Where("id=?", a.p.Id)
	}

	if a.p.StudentId > 0 {
		a.s = a.s.Where("student_id=?", a.p.StudentId)
	}

	if len(a.p.ActName) > 0 {
		a.s = a.s.Where("act_name=?", a.p.ActName)
	}

	if a.p.ActType > 0 {
		a.s = a.s.Where("act_type=?", a.p.ActType)
	}

	if a.p.ProjectId > 0 {
		a.s = a.s.Where("project_id=?", a.p.ProjectId)
	}

	if a.p.SubjectId > 0 {
		a.s = a.s.Where("subject_id=?", a.p.SubjectId)
	}

	return a.s
}
