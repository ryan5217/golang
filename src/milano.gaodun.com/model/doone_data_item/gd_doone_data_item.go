package doone_data_item

import (
	"github.com/go-xorm/builder"
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type DooneDataItemModel struct {
	*xorm.Engine
	s *xorm.Session
}
type GdDooneDataItem struct {
	Icid string
	Num  int64
}

func NewDooneDataItemModel() *DooneDataItemModel {
	return &DooneDataItemModel{Engine: utils.GaoduntikuDb}
}

func (g *DooneDataItemModel) GetCount(uid int64, icids []string) (map[string]int64, error) {
	var cou []GdDooneDataItem
	res := map[string]int64{}
	for _, v := range icids {
		res[v] = 0
	}
	eg := g.Select("icid,COUNT(1) AS num ").Where("student_id=?", uid).Where(builder.In("icid", icids)).GroupBy("icid")
	err := eg.Find(&cou)
	if err != nil {
		setting.Logger.Infof("GdDooneDataItem_%s", err.Error())
		return res, err
	}
	for _, v := range cou {
		res[v.Icid] = v.Num
	}
	return res, err
}
