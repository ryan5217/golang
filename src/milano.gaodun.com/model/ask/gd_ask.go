package ask

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)


type AskModel struct {
	*xorm.Engine
	s *xorm.Session
}
type GdAsk struct {
	Id int64
	Status        int64
}
func NewAskModel() *AskModel {
	return &AskModel{Engine: utils.GaodunDb3}
}

func (g *AskModel) Edit(a *GdAsk) (int64, error) {
	row, err := g.Where("id=?",a.Id).Update(a)
	return row, err
}
