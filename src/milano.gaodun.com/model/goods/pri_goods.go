package goods

import (
	"milano.gaodun.com/pkg/utils"
	"time"
	"milano.gaodun.com/model/xorm-model"
)

type PriGoods struct {
	Id             int64
	GoodsName      string
	GoodsType      int32
	PictureUrl     string
	CostType       int32
	CostLimit      int32
	ProjectId      int64
	ExchangeCost   int64
	StartDate      string
	EndDate        string
	IsDelete       int32
	ExchangeTimes  int64    `xorm:"-"`
	ForceUpdateCol []string `xorm:"-" json:"-"`
}

type GoodsListRes struct {
	StudentId    int64
	InvitedCount int64
	GoodsList    []PriGoods
}
type GoodsModel struct {
	xorm_model.ModelBase
}

func NewGoodsModel() *GoodsModel {
	g := GoodsModel{}
	g.Engine = utils.GaodunPrimaryDb
	g.NewModelSession()
	return &g
}

func (g *GoodsModel) Add(pg *PriGoods) (int64, error) {
	row, err := g.InsertOne(pg)
	return row, err
}
func (g *GoodsModel) Edit(pg *PriGoods) (int64, error) {
	return g.Where("id=?", pg.Id).Cols(pg.ForceUpdateCol...).Update(pg)
}
func (g *GoodsModel) List(pg *[]PriGoods) error {
	dateNow := time.Now().Format("2006-01-02")
	err := g.Where("start_date<?", dateNow).
		Where("end_date>?", dateNow).
		Where("is_delete=?", 1).
		Asc("exchange_cost").Find(pg)
	return err
}

func (g *GoodsModel) FindAll(pg *[]PriGoods) error {
	err := g.GetSession().Find(pg)
	return err
}