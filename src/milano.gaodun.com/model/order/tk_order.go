package order

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
)

type TkOrder struct {
	Id         int64
	Modifydate int64
}

type OrderModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewOrderModel() *OrderModel {
	return &OrderModel{Engine: utils.NewtikuDb}
}

func (g *OrderModel) GetOrderByIds(orderIds []int64) (*[]TkOrder, int) {
	orders := []TkOrder{}
	err := g.In("id", orderIds).OrderBy("modifydate desc").Find(&orders)
	if err != nil {
		return &orders, error_code.TKORDERMODELERR
	}
	return &orders, error_code.SUCCESSSTATUS
}
