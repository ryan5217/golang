package banner

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
	"time"
)

var (
	limit = 20
)

type PriAppBanner struct {
	BannerParam        `xorm:"extends"`
	Title              string
	PictureUrl         string
	SkipUrl            string
	ShowType           int32
	ShowTableId        int64
	ShowTableDict      string                 `json:"-"`
	ShowTableDictJASON map[string]interface{} `json:"ShowTableDict" xorm:"-"`
	Sort               int32
	StartDate          string
	EndDate            string
	Status             int32
	Remark             string                 `json:"-"`
	RemarkJSON         map[string]interface{} `json:"Remark" xorm:"-"`
}

type BannerParam struct {
	Id             int64
	ShowMode       int32
	Source         int32
	ProjectId      int64
	SubjectId      int64
	Skip           `xorm:"-"`
	ForceUpdateCol []string `xorm:"-" json:"-"`
}

type BannerVO struct {
	Data  []PriAppBanner
	Count int64
	Page  int
}

type Skip struct {
	Page  int `json:"-"`
	Limit int `json:"-"`
}

type BannerModel struct {
	*xorm.Engine
	s *xorm.Session
	p *BannerParam
}

func NewBannerModel(p *BannerParam) *BannerModel {
	return &BannerModel{Engine: utils.GaodunPrimaryDb, p: p}
}

// 获取列表 按条件获取
func (b *BannerModel) List() ([]PriAppBanner, error) {
	rowsSlicePtr := []PriAppBanner{}

	dateNow := time.Now().Format("2006-01-02")
	err := b.WhereParam().Where("start_date<?", dateNow).
		Where("end_date>?", dateNow).
		Where("status=?", 1).
		Desc("sort").Find(&rowsSlicePtr)

	return rowsSlicePtr, err
}

func (b *BannerModel) Edit(gab *PriAppBanner) (int64, error) {
	return b.WhereParam().Cols(gab.ForceUpdateCol...).Update(gab)
}
func (b *BannerModel) Add(gab *PriAppBanner) (int64, error) {
	row, err := b.Cols(gab.ForceUpdateCol...).InsertOne(gab)
	return row, err
}
func (b *BannerModel) getPage() (int, int) {
	if b.p.Limit == 0 {
		b.p.Limit = limit
	}
	if b.p.Page > 0 {
		b.p.Page--
	}
	return b.p.Limit, b.p.Page * b.p.Limit
}

// 查询条件
func (b *BannerModel) WhereParam() *xorm.Session {
	if b.s == nil {
		b.s = b.NewSession()
		defer b.s.Close()
	}

	if b.p.Id > 0 {
		b.s = b.s.Where("id=?", b.p.Id)
	}

	if b.p.ProjectId > 0 {
		b.s = b.s.Where("project_id=?", b.p.ProjectId)
	}

	if b.p.SubjectId > 0 {
		b.s = b.s.Where("subject_id=?", b.p.SubjectId)
	}

	if b.p.Source > 0 {
		b.s = b.s.Where("source=?", b.p.Source)
	}

	if b.p.ShowMode > 0 {
		b.s = b.s.Where("show_mode=?", b.p.ShowMode)
	}

	return b.s
}
