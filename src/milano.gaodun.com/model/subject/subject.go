package subject

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type GdSubject struct {
	Id   int64
	Name string
}
type SubjectModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewSubjectModel() *SubjectModel {
	return &SubjectModel{Engine: utils.GaodunDb2}
}
func (g *SubjectModel) Get(id int64) (*GdSubject, error) {
	srm := GdSubject{}
	_, err := g.Id(id).Get(&srm)
	return &srm, err
}
