package banner

import (
	"github.com/apex/log"
	"milano.gaodun.com/model/exchange"
	"milano.gaodun.com/model/goods"
	"milano.gaodun.com/model/invitaion"
)

type GoodsServiceInterface interface {
	Add(gds *goods.PriGoods) (int64, error)
	Edit(gds *goods.PriGoods) (int64, error)
	List(gds *[]goods.PriGoods) error
	GetList(gds *goods.GoodsListRes) error
	FindAll(gds *[]goods.PriGoods) error
}

type GoodsService struct {
	GoodsM        *goods.GoodsModel
	InviteM       *invitaion.InvitationCodeModel
	ExchangeCodeM *exchange.ExchangeCodeModel
	logger        *log.Entry
	Page          int
	Limit         int
}

func NewGoodsService(logger *log.Entry) *GoodsService {
	return &GoodsService{
		GoodsM:        goods.NewGoodsModel(),
		InviteM:       invitaion.NewInvitationCodeModel(),
		ExchangeCodeM: exchange.NewExchangeCodeModel(),
		logger:        logger,
	}
}

func (g *GoodsService) Add(gds *goods.PriGoods) (int64, error) {
	row, err := g.GoodsM.Add(gds)
	if err != nil {
		g.logger.Error(err.Error())
		return row, err
	}
	g.GoodsM.Id(gds.Id).Get(gds)
	return row, err
}
func (g *GoodsService) Edit(gds *goods.PriGoods) (int64, error) {
	row, err := g.GoodsM.Edit(gds)
	if err != nil {
		g.logger.Error(err.Error())
	}
	g.GoodsM.Id(gds.Id).Get(gds)
	return row, err
}

// 查询所有列表
func (g *GoodsService) FindAll(gds *[]goods.PriGoods) error {
	g.GoodsM.Page(g.Limit, g.Page)
	err := g.GoodsM.FindAll(gds)
	if err != nil {
		g.logger.Error(err.Error())
	}
	return err
}

func (g *GoodsService) List(gds *[]goods.PriGoods) error {
	err := g.GoodsM.List(gds)
	if err != nil {
		g.logger.Error(err.Error())
	}
	return err
}
func (g *GoodsService) GetList(gds *goods.GoodsListRes) error {
	//获取商品列表
	err := g.GoodsM.List(&gds.GoodsList)
	if err != nil {
		g.logger.Error(err.Error())
		return err
	}
	//获取兑换列表
	epm := exchange.ExchangeParam{StudentId: gds.StudentId, ExchangeCode: ""}
	ecd, err := g.ExchangeCodeM.List(&epm)
	if err != nil {
		g.logger.Error(err.Error())
		return err
	}
	for _, v := range ecd {
		for key, val := range gds.GoodsList {
			if val.Id == v.GoodsId {
				gds.GoodsList[key].ExchangeTimes++
			}
		}
	}
	//获取邀请好友次数
	invit, err := g.InviteM.GetCode(gds.StudentId)
	if err != nil {
		g.logger.Error(err.Error())
	}
	gds.InvitedCount = invit.InvitedCount - invit.UsedInvitTimes
	return err
}
