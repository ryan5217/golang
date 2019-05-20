package collection

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/model/xorm-model"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type PriCollectionQuestion struct {
	Id        int64     `json:"id" xorm:"not null pk autoincr INT(11)"`
	Question  string    `json:"question" xorm:"default '' comment('问题') unique VARCHAR(500)"`
	ProjectId int64     `json:"project_id" xorm:"default 0 comment('项目ID') INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"<-"`
	UpdatedAt time.Time `json:"updated_at" xorm:"<-"`
}

type CollectionQuestionModel struct {
	xorm_model.ModelBase
	con  *xorm.Session
	one  PriCollectionQuestion
	list []PriCollectionQuestion
}

func NewCollectionQuestionModelModel() *CollectionQuestionModel {
	g := CollectionQuestionModel{}
	g.Engine = utils.GaodunPrimaryDb
	g.NewModelSession()
	g.con = g.Engine.NewSession()
	return &g
}

func (g *CollectionQuestionModel) Add(param *PriCollectionQuestion) (int64, error) {
	utils.ChangeStruct2OtherStruct(param, &g.one)
	_, err := g.Insert(&g.one)
	if err != nil {
		setting.Logger.Errorf("sql_Error" + err.Error())
	}
	return g.one.Id, err
}

func (g *CollectionQuestionModel) GetById(ID []int64) []PriCollectionQuestion {
	error := g.In("id", ID).Find(&g.list)
	if error != nil {
		setting.Logger.Errorf("sql_Error" + error.Error())
	}
	return g.list
}

func (g *CollectionQuestionModel) GetByQuestion(question string) PriCollectionQuestion {
	g.Where("question = ?", question).Get(&g.one)
	return g.one
}

func (g *CollectionQuestionModel) GetList(param *PriCollectionQuestion) []PriCollectionQuestion {
	engine := g.con
	if param.ProjectId > 0 {
		engine.Where("project_id = ?", param.ProjectId)
	}
	if param.Question != "" {
		engine.Where("question = ?", param.Question)
	}
	errList := engine.Find(&g.list)
	if errList != nil {
		setting.Logger.Errorf("sql_Error" + errList.Error())
	}
	return g.list
}
