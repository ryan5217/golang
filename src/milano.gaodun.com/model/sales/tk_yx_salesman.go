package sales

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/model/xorm-model"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type TkYxSalesman struct {
	Id         int64  `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name       string `json:"name" xorm:"not null default '' comment('销售名') VARCHAR(10)"`
	WeixinCode string `json:"weixin_code" xorm:"not null default '' comment('微信code') VARCHAR(50)"`
	Qcode      string `json:"qcode" xorm:"not null default '' comment('微信二维码') VARCHAR(200)"`
	ProjectId  int64  `json:"project_id" xorm:"not null default 0 comment('项目id') INT(11)"`
	Status     int    `json:"status" xorm:"not null default 1 comment('1启用0禁用') TINYINT(4)"`
	Regdate    int    `json:"regdate" xorm:"not null INT(10)"`
	Modifydate int    `json:"modifydate" xorm:"not null INT(10)"`
	Account    string `json:"account" xorm:"comment('登录账号') VARCHAR(50)"`
	CrmId      int    `json:"crm_id" xorm:"not null comment('crm对应的id') INT(11)"`
	OpenId     string `json:"open_id" xorm:"not null comment('微信open_id') VARCHAR(255)"`
	/*	SetNum     int    `json:"set_num" xorm:"not null default 0 comment('设定数量') INT(11)"`
		AddNum     int    `json:"add_num" xorm:"not null comment('添加数量') INT(11)"`
		TotalNum   int    `json:"total_num" xorm:"not null comment('累计添加数') INT(11)"`*/
}

type TkYxSalesmanModel struct {
	xorm_model.ModelBase
	con  *xorm.Session
	list []TkYxSalesman
}

func NewCollectionInfoModel() *TkYxSalesmanModel {
	g := TkYxSalesmanModel{}
	g.Engine = utils.NewtikuDb
	g.NewModelSession()
	g.con = g.Engine.NewSession()
	return &g
}

func (g *TkYxSalesmanModel) FindAll(param *TkYxSalesman) []TkYxSalesman {
	engine := g.con
	logger := setting.Logger
	var list []TkYxSalesman
	engine = engine.Table("tk_yx_salesman")
	if param.ProjectId > 0 {
		engine.Where("project_id = ?", param.ProjectId)
	}
	error := engine.Find(&list)
	if error != nil {
		logger.Info("sql_Error" + error.Error())
	}
	logger.Info("SELECT_RESULT" + fmt.Sprintf("%v", list))
	return list
}
