package ask

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

type GdAskQuestions struct {
	Id         int64
	Pid        int64
	AskId      int64
	StudentId  int64
	Content    string
	Regdate    int64
	FileUrl    string
	Modifydate int64
}

type AskQuestionsModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewAskQuestionsModel() *AskQuestionsModel {
	return &AskQuestionsModel{Engine: utils.GaodunDb3}
}

func (g *AskQuestionsModel) Add(pg *GdAskQuestions) (int64, error) {
	row, err := g.InsertOne(pg)
	return row, err
}
