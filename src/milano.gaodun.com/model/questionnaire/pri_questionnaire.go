package questionnaire

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type PriQuestionnaire struct {
	Id             int64
	ItemId         string
	ItemName       string
	TypeId         int64
	Choice         string
	Evaluate       string
	Comment        string
	ForceUpdateCol []string `xorm:"-" json:"-"`
}

type QuestionnaireModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewQuestionnaireModel() *QuestionnaireModel {
	return &QuestionnaireModel{Engine: utils.GaodunPrimaryDb}
}

func (e *QuestionnaireModel) Add(ec *PriQuestionnaire) (int64, error) {
	row, err := e.InsertOne(ec)
	return row, err
}
func (e *QuestionnaireModel) Edit(ec *PriQuestionnaire) (int64, error) {
	return e.Id(ec.Id).Cols(ec.ForceUpdateCol...).Update(ec)
}
