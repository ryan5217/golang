package collection

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/model/xorm-model"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"milano.gaodun.com/request-params"
	"time"
)

type PriCollectionInfo struct {
	Id            int64     `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uid           int64     `json:"uid" xorm:"default 0 comment('gaodun user id') INT(11)"`
	QuestionId    int64     `json:"question_id" xorm:"default 0 comment('问题 id') INT(11)"`
	Context       string    `json:"context" xorm:"default '' comment('用户回答内容') VARCHAR(255)"`
	ContextExtend string    `json:"context_extend" xorm:"default '' comment('扩展信息：图片地址, 分号分逗') VARCHAR(500)"`
	PartnerId     int64     `json:"partner_id" xorm:"default 0 comment('第三方业务 id') index(idx_partner) INT(11)"`
	PartnerKey    string    `json:"partner_key" xorm:"default '' comment('第三方业务关键字(partner_key 创建)') index(idx_partner) VARCHAR(50)"`
	PartnerName   string    `json:"partner_name" xorm:"default '' comment('第三方业务名称(partner_key 创建)') index(idx_partner) VARCHAR(255)"`
	ProjectId     int64     `json:"project_id" xorm:"default 0 comment('项目 id') INT(11)"`
	Source        string    `json:"source" xorm:"default '' comment('来源 android,ios,web,other') VARCHAR(255)"`
	CreatedAt     time.Time `json:"created_at" xorm:"<-"`
	UpdatedAt     time.Time `json:"updated_at" xorm:"<-"`
}

type Collection struct {
	Question      string    `json:"question"`
	QuestionId    int64     `json:"question_id"`
	ProjectId     int64     `json:"project_id"`
	Uid           int64     `json:"uid" `
	KeyName       string    `json:"key_name"`
	PartnerKey    string    `json:"partner_key"`
	PartnerName   string    `json:"partner_name"`
	Context       string    `json:"context"`
	PartnerId     int64     `json:"partner_id"`
	ContextExtend string    `json:"context_extend"`
	Source        string    `json:"source"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CollectionList struct {
	Question      string    `json:"question"`
	QuestionId    int64     `json:"question_id"`
	ProjectId     int64     `json:"project_id"`
	Uid           int64     `json:"uid" `
	PartnerKey    string    `json:"partner_key"`
	KeyName       string    `json:"key_name"`
	PartnerName   string    `json:"partner_name"`
	Context       string    `json:"context"`
	PartnerId     int64     `json:"partner_id"`
	ContextExtend string    `json:"context_extend"`
	Source        string    `json:"source"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	params.Page
}

type ReportList struct {
	Context string  `json:"context"`
	Num     float64 `json:"num"`
	Rate    string  `json:"rate"`
}

type CollectionInfoModel struct {
	xorm_model.ModelBase
	con  *xorm.Session
	list []PriCollectionInfo
}

func NewCollectionInfoModel() *CollectionInfoModel {
	g := CollectionInfoModel{}
	g.Engine = utils.GaodunPrimaryDb
	g.NewModelSession()
	g.con = g.Engine.NewSession()
	return &g
}

func (g *CollectionInfoModel) Add(param *PriCollectionInfo) (int64, error) {
	row, err := g.InsertOne(param)
	if err != nil {
		setting.Logger.Errorf("sql_Error" + err.Error())
	}
	return row, err
}

func (g *CollectionInfoModel) FindAll(param *[]PriCollectionInfo) error {
	err := g.GetSession().Find(param)
	if err != nil {
		setting.Logger.Errorf("sql_Error" + err.Error())
	}
	return err
}

func (g *CollectionInfoModel) GetInfoAndQuestion(param *CollectionList) ([]Collection, int64) {
	engine := g.con
	var list []Collection
	param.HandlePage()
	engine = engine.Table("pri_collection_info").Join("LEFT", "pri_collection_question", "pri_collection_info.question_id = pri_collection_question.id")
	if param.Question != "" {
		engine.Where("pri_collection_question.question like ?", param.Question+"%")
	}
	if param.QuestionId > 0 {
		engine.Where("pri_collection_info.question_id = ?", param.QuestionId)
	}
	if param.ProjectId > 0 {
		engine.Where("pri_collection_info.project_id = ?", param.ProjectId)
	}
	if param.Source != "" {
		engine.Where("pri_collection_info.source = ?", param.Source)
	}
	if param.PartnerKey != "" {
		engine.Where("pri_collection_info.partner_key = ?", param.PartnerKey)
	}
	if param.PartnerName != "" {
		engine.Where("pri_collection_info.partner_name = ?", param.PartnerName+"%")
	}
	if param.Uid > 0 {
		engine.Where("pri_collection_info.uid = ?", param.Uid)
	}
	if param.PartnerId > 0 {
		engine.Where("pri_collection_info.partner_id = ?", param.PartnerId)
	}
	engine = engine.Limit(param.Offset, (param.PageNum-1)*param.Offset)
	total, err := engine.Clone().Count()
	if err != nil {
		setting.Logger.Errorf("sql_Error" + err.Error())
	}
	errList := engine.Find(&list)
	if errList != nil {
		setting.Logger.Errorf("sql_Error" + errList.Error())
	}
	return list, total
}

/**
获取报表
*/
func (g *CollectionInfoModel) GetInfoReport(questionId int64, partnerKey string, limit int64, top int64) []map[string]string {
	engine := g.con
	var list []map[string]string
	var errList error
	if limit > 1000 || limit < 1 {
		limit = 100
	}
	if top <= 0 {
		top = 10
	}
	if questionId > 0 && partnerKey != "" {
		list, errList = engine.SQL("SELECT  info.context,COUNT(info.id) AS num FROM (SELECT id,context,question_id,partner_key FROM `pri_collection_info` WHERE question_id = ? AND partner_key = ? ORDER BY id DESC LIMIT ?) AS info GROUP BY info.context ORDER BY num DESC LIMIT ?", questionId, partnerKey, limit, top).QueryString()
	}
	if errList != nil {
		setting.Logger.Errorf("sql_Error" + errList.Error())
	}
	return list
}
