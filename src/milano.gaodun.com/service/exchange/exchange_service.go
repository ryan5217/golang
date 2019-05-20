package exchange

import (
	"github.com/apex/log"
	//"milano.gaodun.com/model/invitaion"
	//"encoding/json"
	"fmt"
	"milano.gaodun.com/model/exchange"
	"milano.gaodun.com/model/goods"
	"milano.gaodun.com/model/invitaion"
	"milano.gaodun.com/pkg/utils"
	"milano.gaodun.com/pkg/error-code"
	"time"
)

type ExchangeServiceInterface interface {
	GetCode(studentId int64, goodsId int64) *exchange.PriExchangeCode
	GetList(param *exchange.ExchangeParam) []exchange.PriExchangeCode
	Edit(gab *exchange.PriExchangeCode) (int64, error)
}

type ExchangeService struct {
	ExchangeCodeM *exchange.ExchangeCodeModel
	GoodsM        *goods.GoodsModel
	InviteCodeM   *invitaion.InvitationCodeModel
	logger        *log.Entry
}

func NewExchangeService(logger *log.Entry) ExchangeServiceInterface {
	return &ExchangeService{
		ExchangeCodeM: exchange.NewExchangeCodeModel(),
		GoodsM:        goods.NewGoodsModel(),
		InviteCodeM:   invitaion.NewInvitationCodeModel(),
		logger:        logger}
}
func (i *ExchangeService) GetList(param *exchange.ExchangeParam) []exchange.PriExchangeCode {
	ecList, err := i.ExchangeCodeM.List(param)

	if err != nil {
		i.logger.Error(err.Error())
	}
	goodsIdList := []int64{}
	for _, v := range ecList {
		goodsIdList = append(goodsIdList, v.GoodsId)
	}
	goodsList := []goods.PriGoods{}
	err = i.GoodsM.In("id", goodsIdList).Find(&goodsList)
	for k, e := range ecList {
		for _, v := range goodsList {
			if v.Id == e.GoodsId {
				ecList[k].GoodsName = v.GoodsName
				ecList[k].PictureUrl = v.PictureUrl
			}
		}

	}
	return ecList
}
func (i *ExchangeService) Edit(gab *exchange.PriExchangeCode) (int64, error) {
	row, err := i.ExchangeCodeM.Edit(gab)
	if err != nil {
		i.logger.Error(err.Error())
	}
	i.ExchangeCodeM.Id(gab.Id).Get(gab)
	return row, err
}
// api 数据返回
func (i *ExchangeService) GetCode(studentId int64, goodsId int64) *exchange.PriExchangeCode {
	ec := exchange.PriExchangeCode{}
	gds := goods.PriGoods{}
	info := error_code.GetInfo()
	i.GoodsM.Id(goodsId).Get(&gds)
	if gds.Id == 0 {
		ec.StatusCode = error_code.GOODSNOTEXIST
		ec.Message = info[int(ec.StatusCode)]
		return &ec
	}
	if gds.IsDelete == 2 {
		ec.StatusCode = error_code.GOODSDELETE
		ec.Message = info[int(ec.StatusCode)]
		return &ec
	}
	start := utils.StrToUnix(gds.StartDate)
	end := utils.StrToUnix(gds.EndDate)
	now := time.Now().Unix()
	if now < start {
		ec.StatusCode = error_code.GOODSNOTSTART
		ec.Message = info[int(ec.StatusCode)]+":"+gds.StartDate+"-"+gds.EndDate
		return &ec
	}
	if now > end {
		ec.StatusCode = error_code.GOODSEXPIRED
		ec.Message = info[int(ec.StatusCode)]+":"+gds.StartDate+"-"+gds.EndDate
		return &ec
	}
	num, _ := i.ExchangeCodeM.Where("student_id=?", studentId).Where("goods_id=?", goodsId).Count(&ec)
	if int32(num) >= gds.CostLimit {
		ec.StatusCode = error_code.GOODSEXCHANGELIMIT
		ec.Message = info[int(ec.StatusCode)]
		return &ec
	}
	if gds.CostType == 1 {
		inv, err := i.InviteCodeM.GetCode(studentId)
		if err != nil {
			i.logger.Error(err.Error())
			ec.StatusCode = error_code.DBERR
			ec.Message = info[int(ec.StatusCode)]
			return &ec
		}
		leftTimes := inv.InvitedCount - inv.UsedInvitTimes
		if leftTimes < gds.ExchangeCost {
			ec.StatusCode = error_code.REMAINLIMIT
			ec.Message = fmt.Sprintf("该奖品需要邀请%d位好友,你的条件未满足", gds.ExchangeCost)
			return &ec
		}
		ec.ExchangeCode = utils.GetRandomString(16)
		ec.StudentId = studentId
		ec.GoodsId = goodsId
		if(gds.GoodsType == 1){
			i.ExchangeCodeM.SaveActivity(studentId)
			ec.Status = 2
		} else{
			ec.Status = 1
		}
		s := i.ExchangeCodeM.NewSession()
		defer s.Close()
		s.Begin()
		ec.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
		s.InsertOne(&ec)
		_, err = i.InviteCodeM.Where("student_id=?", studentId).Incr("used_invit_times", gds.ExchangeCost).Update(&invitaion.PriInvitationCode{})
		if err != nil {
			i.logger.Error(err.Error())
			ec.StatusCode = 2
			ec.Message = err.Error()
			s.Rollback()
			return &ec
		}
		s.Commit()
		ec.GoodsName = gds.GoodsName
		ec.StatusCode = 0
		ec.Message = info[error_code.SUCCESSSTATUS]
		return &ec
	} else {
		ec.StatusCode = error_code.EXCHANGETYPELIMIT
		ec.Message = info[int(ec.StatusCode)]
		return &ec
	}

}
