package banner

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
)

var (
	limit = 20
)

type GdTikuSystemConstant struct {
	SearchParam `xorm:"extends"`
	Id          int
	Thevalue    string
	Regdate     int64
	Modifydate  int64
	Comment     string
}

type SearchParam struct {
	Thekey         string
	Type           int32
	ForceUpdateCol []string `xorm:"-" json:"-"`
}

type TikuSystemConstantModel struct {
	*xorm.Engine
	s *xorm.Session
	p *SearchParam
}

func NewTikuSystemConstantModel(p *SearchParam) *TikuSystemConstantModel {
	return &TikuSystemConstantModel{Engine: utils.GaodunDb, p: p}
}

// 获取列表 按条件获取
func (b *TikuSystemConstantModel) GetKey() (GdTikuSystemConstant, error) {
	rowsSlicePtr := GdTikuSystemConstant{}

	_, err := b.WhereParam().Get(&rowsSlicePtr)

	return rowsSlicePtr, err
}

// 查询条件
func (b *TikuSystemConstantModel) WhereParam() *xorm.Session {
	if b.s == nil {
		b.s = b.NewSession()
		defer b.s.Close()
	}

	if b.p.Thekey != "" {
		b.s = b.s.Where("thekey=?", b.p.Thekey)
	}

	if b.p.Type > 0 {
		b.s = b.s.Where("type=?", b.p.Type)
	}
	return b.s
}
