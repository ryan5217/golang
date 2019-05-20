package collection

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/model/xorm-model"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type PriCollectionType struct {
	Id         int64     `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name       string    `json:"name" xorm:"default '' comment('名称') VARCHAR(50)"`
	PartnerKey string    `json:"partner_key" xorm:"default '' comment('关键字用于类型分类和业务分类') unique VARCHAR(50)"`
	Desc       string    `json:"desc" xorm:"default '' comment('用途说明') VARCHAR(200)"`
	CreatedAt  time.Time `json:"created_at" xorm:"<-"`
	UpdatedAt  time.Time `json:"updated_at" xorm:"<-"`
}

type CollectionTypeModel struct {
	xorm_model.ModelBase
	con  *xorm.Session
	one  PriCollectionType
	list []PriCollectionType
}

func NewCollectionTypeModel() *CollectionTypeModel {
	g := CollectionTypeModel{}
	g.Engine = utils.GaodunPrimaryDb
	g.NewModelSession()
	g.con = g.Engine.NewSession()
	return &g
}

func (g *CollectionTypeModel) GetByPartnerKey(partnerKey string) PriCollectionType {
	g.Where("partner_key = ?", partnerKey).Get(&g.one)
	return g.one
}

func (g *CollectionTypeModel) Add(param *PriCollectionType) (int64, error) {
	row, err := g.InsertOne(param)
	if err != nil {
		setting.Logger.Errorf("sql_Error" + err.Error())
	}
	return row, err
}

func (g *CollectionTypeModel) Edit(param *PriCollectionType) (int64, error) {
	row, err := g.con.ID(param.Id).Update(param)
	if err != nil {
		setting.Logger.Errorf("sql_Error" + err.Error())
	}
	return row, err
}

func (g *CollectionTypeModel) GetList(param *PriCollectionType) []PriCollectionType {
	engine := g.con
	if param.Name != "" {
		engine.Where("name like ? ", param.Name+"%")
	}
	if param.PartnerKey != "" {
		engine.Where("partner_key like ?", param.PartnerKey+"%")
	}
	error := engine.Find(&g.list)
	if error != nil {
		setting.Logger.Errorf("sql_Error" + error.Error())
	}
	return g.list
}
