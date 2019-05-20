package members_student

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type GdMembersStudent struct {
	Id       int64
	MemberId int64
	Nickname string
	Phone    string
}
type MembersStudentModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewMembersStudentModel() *MembersStudentModel {
	return &MembersStudentModel{Engine: utils.GaodunDb2}
}
func (g *MembersStudentModel) Get(id int64) (*GdMembersStudent, error) {
	srm := GdMembersStudent{}
	_, err := g.Id(id).Get(&srm)
	if err != nil {
		setting.Logger.Infof("MembersStudentModel_%s", err.Error())
	}
	return &srm, err
}
