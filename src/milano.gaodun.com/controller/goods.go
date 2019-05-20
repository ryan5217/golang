package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	gModle "milano.gaodun.com/model/goods"
	"milano.gaodun.com/pkg/setting"
	gService "milano.gaodun.com/service/goods"
	"strconv"
)

var GoodsApi = NewGoodsApi()

func NewGoodsApi() *Goods {
	return &Goods{}
}

type Goods struct {
	Base
}

func (i Goods) Add(c *gin.Context) {
	gds := gModle.PriGoods{}
	gds.GoodsName = i.GetString(c, "goods_name")
	gds.PictureUrl = i.GetString(c,"picture_url")
	gds.ExchangeCost = i.GetInt64(c, "exchange_cost")
	gds.CostType = i.GetInt32(c, "cost_type")
	gds.GoodsType = i.GetInt32(c, "goods_type")
	gds.CostLimit = i.GetInt32(c, "cost_limit")
	gds.ProjectId = i.GetInt64(c, "project_id")
	if gds.CostType == 0 {
		gds.CostType = 1
	}
	gds.IsDelete = 1
	gds.StartDate = i.GetString(c, "start_date")
	gds.EndDate = i.GetString(c, "end_date")
	bServer := gService.NewGoodsService(setting.GinLogger(c))
	bServer.Add(&gds)
	i.ServerJSONSuccess(c, gds)
}
func (i Goods) Edit(c *gin.Context) {
	gds := gModle.PriGoods{}
	gds.GoodsName = i.GetString(c, "goods_name")
	gds.ExchangeCost = i.GetInt64(c, "exchange_cost")
	gds.CostType = i.GetInt32(c, "cost_type")
	gds.IsDelete = i.GetInt32(c, "is_delete")
	gds.StartDate = i.GetString(c, "start_date")
	gds.EndDate = i.GetString(c, "end_date")
	gds.GoodsType = i.GetInt32(c, "goods_type")
	gds.CostLimit = i.GetInt32(c, "cost_limit")
	gds.ProjectId = i.GetInt64(c, "project_id")
	gds.PictureUrl = i.GetString(c,"picture_url")
	gds.ForceUpdateCol = i.PostMustCols(c)
	gds.Id, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	bServer := gService.NewGoodsService(setting.GinLogger(c))
	bServer.Edit(&gds)
	i.ServerJSONSuccess(c, gds)
}
func (i Goods) List(c *gin.Context) {
	gds := []gModle.PriGoods{}
	bServer := gService.NewGoodsService(setting.GinLogger(c))
	bServer.List(&gds)
	i.ServerJSONSuccess(c, gds)
}
func (i Goods) GetList(c *gin.Context) {
	gds := gModle.GoodsListRes{}
	gds.StudentId = i.GetInt64(c, "student_id",true)
	if c.GetBool(Verify) {
		return
	}
	bServer := gService.NewGoodsService(setting.GinLogger(c))
	bServer.GetList(&gds)

	i.ServerJSONSuccess(c, gds)
}
// 查询所有
func (i Goods) FindAll(c *gin.Context) {
	gds := []gModle.PriGoods{}
	bServer := gService.NewGoodsService(setting.GinLogger(c))
	bServer.Page = i.GetInt(c, "page")
	bServer.Limit = i.GetInt(c, "limit")
	bServer.FindAll(&gds)
	i.ServerJSONSuccess(c, gds)
}