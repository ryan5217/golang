package ask

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

type SalesManApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type SalesManParam struct {
	ProjectId   int64
	YxStudentId int64
}

func NewSalesManApi(h *utils.HttpClient) *SalesManApi {
	var i = SalesManApi{}
	i.HttpClient = h
	return &i
}

func (g *SalesManApi) GetSalesManList(param SalesManParam, clearCache int64) (*simpleJson.Json, int) {
	var code = error_code.SUCCESSSTATUS
	redisClient := utils.RedisHandle
	key := fmt.Sprintf("GetSalesManList_%d_%d", param.ProjectId, param.YxStudentId)
	respStr := redisClient.GetData(key).(string)
	if respStr == "" || clearCache == 1 {
		apiHeader := req.Param{}
		apiHeader["project_id"] = param.ProjectId
		apiHeader["student_id"] = param.YxStudentId
		g.Uri = conf.SPARTA_DOMAIN + "/v1/wxuser/my-salesman"
		r, err := req.Post(g.Uri, apiHeader)
		if err != nil {
			setting.Logger.Infof("GetSalesManList_%s", err.Error())
			code = error_code.SALESMANAPIERR
		}
		respStr = r.String()
		redisClient.SetData(key, respStr, time.Second*86400)
	}
	sj, _ := simpleJson.NewJson([]byte(respStr))
	return sj.Get("result"), code
}
