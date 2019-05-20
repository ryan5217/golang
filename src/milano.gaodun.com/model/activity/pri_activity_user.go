package activity

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type PriActivityUser struct {
	Number    string
	NumType   int
	ProjectId int64
	Uid       int64
	Source    string
}

func NewActivityUserModel() *ActivityUserModel {
	return &ActivityUserModel{Engine: utils.GaodunPrimaryDb}
}

type ActivityUserModel struct {
	*xorm.Engine
	s *xorm.Session
}

func (a ActivityUserModel) Add(m *PriActivityUser) (int64, error) {
	return a.Insert(m)
}
func (a ActivityUserModel) Edit(m *PriActivityUser) (int64, error) {
	num, err := a.Where("uid=?", m.Uid).Where("num_type=?", m.NumType).Update(m)
	if err != nil {
		setting.Logger.Infof("ActivityUserModel_Edit_%d"+err.Error(), m.Uid)
	}
	return num, err
}
func (a ActivityUserModel) GetAll(uid int64) ([]PriActivityUser, error) {
	pus := []PriActivityUser{}
	err := a.Where("uid=?", uid).Find(&pus)
	if err != nil {
		setting.Logger.Infof("ActivityUserModel_GetAll_%d"+err.Error(), uid)
	}
	return pus, err
}
