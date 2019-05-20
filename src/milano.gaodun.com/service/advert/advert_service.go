package advert

import (
	"github.com/apex/log"
	simpleJson "github.com/bitly/go-simplejson"
	adu "milano.gaodun.com/model/advert_user"
	"milano.gaodun.com/model/api/adverts"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
)

type AdvertService struct {
	aa     *adverts.AdvertsApi
	adm    *adu.AdvertUserRecordModel
	logger *log.Entry
}

func NewAdvertService(logger *log.Entry) *AdvertService {
	h := utils.HttpHandle
	var g AdvertService
	g.logger = logger
	g.aa = adverts.NewAdvertsApi(h)
	g.adm = adu.NewAdvertUserRecordModel()
	return &g
}

func (g *AdvertService) GetFloatAdvertList(uid int64, projectId int64, paIds string) (*simpleJson.Json, int) {
	result := simpleJson.Json{}
	adtList, code := g.aa.GetAdvertsList(paIds, projectId, 1)
	if code != error_code.SUCCESSSTATUS {
		return &result, code
	}
	recordList, code := g.adm.GetAdvertUserRecords(uid)
	if code != error_code.SUCCESSSTATUS {
		return &result, code
	}
	mapList := map[int64]adu.TkTikuAdvertUserRecord{}
	for _, v := range recordList {
		mapList[v.AdvertId] = v
	}
	for i := len(adtList.MustArray()) - 1; i >= 0; i-- {
		advertId, _ := adtList.GetIndex(i).Get("id").Int64()
		_, ok := mapList[advertId]
		if !ok {
			result = *adtList.GetIndex(i)
		}
	}
	return &result, error_code.SUCCESSSTATUS
}

func (g *AdvertService) AddAdvertRecord(uid int64, advertId int64) (int64, error) {

	adu := adu.TkTikuAdvertUserRecord{AdvertId: advertId, Uid: uid}
	exist, _ := g.adm.IsExist(&adu)
	if !exist {
		return g.adm.Add(&adu)
	} else {
		return 0, nil
	}
}
