package user_change

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type TkYxUserChange struct {
	Id          int64
	YxStudentId int64
	GdStudentId int64
	Phone       string
	CreatedAt   string
	UpdatedAt   string
}
type UserChangeModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewUserChangeModel() *UserChangeModel {
	return &UserChangeModel{Engine: utils.NewtikuDb}
}
func (g *UserChangeModel) Get(uid int64) (*TkYxUserChange, int) {
	srm := TkYxUserChange{}
	code := error_code.SUCCESSSTATUS
	_, err := g.Where("gd_student_id=?", uid).Get(&srm)
	if err != nil {
		setting.Logger.Infof("UserChangeModel_%s", err.Error())
		code = error_code.TKYXUSERCHANGEMODELDBERR
	}
	if srm.YxStudentId == 0 {
		code = error_code.TKSTUDENTIDERR
	}
	return &srm, code
}
