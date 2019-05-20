package sales

import (
	"github.com/apex/log"
	"milano.gaodun.com/model/sales"
)

type SalesServiceInterface interface {
	GetSalesManList()
}

type SalesService struct {
	SalesMan *sales.TkYxSalesmanModel
	logger   *log.Entry
}

func NewSalesService(logger *log.Entry) *SalesService {
	return &SalesService{
		SalesMan: sales.NewCollectionInfoModel(),
		logger:   logger,
	}
}

//获取列表
func (service SalesService) GetList(param *sales.TkYxSalesman) []sales.TkYxSalesman {
	var info sales.TkYxSalesman
	var infoList []sales.TkYxSalesman
	info.ProjectId = param.ProjectId
	infoList = service.SalesMan.FindAll(&info)
	return infoList
}
