package users

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type User struct {
	Name string
	Age  int
}

type UserFac struct {
	Table        User
	RowsSlicePtr []User
	engine       *xorm.Engine
}

func NewUserFac() {
	var u = UserFac{}
	u.engine = utils.GaodunDb

}

func (u UserFac) Find() User {
	u.Table.Name = "user"
	return u.Table
}
