package advert_user

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type TkTikuAdvertUserRecord struct {
	Uid         int64
	AdvertId    int64
	UpdatedTime string
	CreatedTime string
}

type AdvertUserRecordModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewAdvertUserRecordModel() *AdvertUserRecordModel {
	return &AdvertUserRecordModel{Engine: utils.NewtikuDb}
}

func (g *AdvertUserRecordModel) GetAdvertUserRecords(uid int64) ([]TkTikuAdvertUserRecord, int) {
	records := []TkTikuAdvertUserRecord{}
	err := g.Where("uid=?", uid).Find(&records)
	if err != nil {
		setting.Logger.Infof("AdvertUserRecordModel_GetMessages_%d_%s", uid, err.Error())
		return records, error_code.DBERR
	}
	return records, error_code.SUCCESSSTATUS
}
func (g *AdvertUserRecordModel) Add(ts *TkTikuAdvertUserRecord) (int64, error) {
	var row int64
	var err error
	row, err = g.Omit("created_time").Omit("updated_time").InsertOne(ts)
	if err != nil {
		setting.Logger.Infof("AdvertUserRecordModel_AddRecord_%s", err.Error())
	}
	return row, err
}
func (g *AdvertUserRecordModel) IsExist(ts *TkTikuAdvertUserRecord) (bool, error) {
	exist, err := g.Exist(ts)
	if err != nil {
		setting.Logger.Infof("AdvertUserRecordModel_FindOne_%s", err.Error())
	}
	return exist, err
}
