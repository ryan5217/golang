package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/setting"
	//"milano.gaodun.com/service/invitation"
	"milano.gaodun.com/service/exchange"
	EM "milano.gaodun.com/model/exchange"
	"strconv"
)

var ExchangeApi = NewExchangeApi()

func NewExchangeApi() *Exchange {
	return &Exchange{}
}

type Exchange struct {
	Base
}

func (e Exchange) GetCode(c *gin.Context) {
	studentId := e.GetInt64(c, "student_id")
	goodsId := e.GetInt64(c, "goods_id")
	bServer := exchange.NewExchangeService(setting.GinLogger(c))
	r := bServer.GetCode(studentId, goodsId)
	if(r.StatusCode == 0){
		e.ServerJSONSuccess(c, r)
	} else {
		e.ServerJSONError(c,r,int(r.StatusCode))
	}
}
func (e Exchange) Modify(c *gin.Context) {
	ech := EM.PriExchangeCode{}
	Id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ech.Id =Id
	ech.StudentId = e.GetInt64(c, "student_id")
	ech.ExchangeCode = e.GetString(c,"exchange_code")
	ech.GoodsId = e.GetInt64(c, "goods_id")
	ech.Status = e.GetInt32(c,"status")
	ech.ForceUpdateCol = e.PostMustCols(c)
	bServer := exchange.NewExchangeService(setting.GinLogger(c))
	bServer.Edit(&ech)
	e.ServerJSONSuccess(c, ech)
}
func (e Exchange) GetList(c *gin.Context) {
	emp := EM.ExchangeParam{}
	emp.StudentId = e.GetInt64(c, "student_id")
	emp.ExchangeCode = e.GetString(c, "exchange_code")
	bServer := exchange.NewExchangeService(setting.GinLogger(c))
	r := bServer.GetList(&emp)
	e.ServerJSONSuccess(c, r)
}
