package adverts

import (
	"fmt"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type AdvertsApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type AdvertsParam struct {
	ProjectId   int64
	YxStudentId int64
}

func NewAdvertsApi(h *utils.HttpClient) *AdvertsApi {
	var i = AdvertsApi{}
	i.HttpClient = h
	return &i
}

func (g *AdvertsApi) GetAdvertsList(paIds string, projectId int64,clearCache int64) (*simpleJson.Json, int) {
	redisClient := utils.RedisHandle
	key := fmt.Sprintf("GetAdvertsList_%s_%d" ,paIds,projectId)
	respStr := redisClient.GetData(key).(string)
	if respStr == "" || clearCache == 1 {
		g.Uri = conf.SPARTA_DOMAIN + fmt.Sprintf("/v1/advert/getlist?pa_ids=%s&project_id=%d", paIds, projectId)
		r, err := req.Get(g.Uri)
		if err != nil {
			setting.Logger.Infof("GetAdvertsList_%s", err.Error())
			return &simpleJson.Json{}, error_code.ADVERTSAPIERR
		}
		respStr = r.String()
		redisClient.SetData(key, respStr, time.Second*86400)
	}
	sj, err := simpleJson.NewJson([]byte(respStr))
	if err != nil {
		setting.Logger.Infof("GetAdvertsList_%s", respStr)
		setting.Logger.Infof("GetAdvertsList_%s", err.Error())
		return &simpleJson.Json{}, error_code.ADVERTSAPIERR
	}
	resp := sj.Get("result")
	if len(resp.MustArray()) == 0 {
		setting.Logger.Infof("GetAdvertsListEmpty")
		return &simpleJson.Json{}, error_code.ADVERTSDATAEMPTY
	}
	return resp, error_code.SUCCESSSTATUS
}
