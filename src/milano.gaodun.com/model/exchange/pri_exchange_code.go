package exchange

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
)

type PriExchangeCode struct {
	Id             int64
	ExchangeParam `xorm:"extends"`
	GoodsId        int64
	GoodsName      string `xorm:"-"`
	PictureUrl     string `xorm:"-"`
	Status         int32
	CreatedTime    string
	StatusCode     int32    `xorm:"-"`
	Message        string   `xorm:"-"`
	ForceUpdateCol []string `xorm:"-" json:"-"`
}
type ExchangeParam struct {
	ExchangeCode   string
	StudentId      int64
}
type ExchangeCodeModel struct {
	*xorm.Engine
	s *xorm.Session
}
type Activity struct {
	Code    int64       `json:"code"`
	Message int64       `json:"message"`
	Result    interface{} `json:"result"`
}
func NewExchangeCodeModel() *ExchangeCodeModel {
	return &ExchangeCodeModel{Engine: utils.GaodunPrimaryDb}
}

func (e *ExchangeCodeModel) Add(ec *PriExchangeCode) (int64, error) {
	row, err := e.InsertOne(ec)
	return row, err
}
func (e *ExchangeCodeModel) List(ep *ExchangeParam) ([]PriExchangeCode, error) {
	ecd := []PriExchangeCode{}
	err := e.WhereParam(ep).Desc("created_time").Limit(100,0).Find(&ecd)
	return ecd, err
}
func (e *ExchangeCodeModel) WhereParam(ep *ExchangeParam) *xorm.Session {
	if e.s == nil {
		e.s = e.NewSession()
		defer e.s.Close()
	}
	if ep.StudentId > 0 {
		e.s = e.s.Where("student_id=?", ep.StudentId)
	}

	if ep.ExchangeCode != "" {
		e.s = e.s.Where("exchange_code=?", ep.ExchangeCode)
	}


	return e.s
}
func (b *ExchangeCodeModel) Edit(gab *PriExchangeCode) (int64, error) {
	return b.Id(gab.Id).Update(gab)
}
func (b *ExchangeCodeModel) SaveActivity(studentId int64) (*Activity, error) {
	var res Activity
	//ip := req.
	param := req.Param{
		"subject_id":     210,
		"project_id":     38,
		"act_type": 1,
		"act_name":  2018,
		"student_id": studentId,
	}
	r, err := req.Post(conf.MILANO_DOMAIN+"/v1/activity/add", req.Header{}, param)
	//fmt.Println(r)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}