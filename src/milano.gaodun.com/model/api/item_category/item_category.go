package ask

import (
	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
	simpleJson "github.com/bitly/go-simplejson"
	"milano.gaodun.com/pkg/setting"
	"time"
)

type ItemCategoryApi struct {
	*api.BaseApi
	Uri string
}

type ItemCategoryResApi struct {
	Status string
	Info   string
	Result *simpleJson.Json
}
type ItemCategoryParam struct {
	ProjectId int64
	SubjectId int64
	Picid     string
}

func NewItemCategoryApi(h *utils.HttpClient) *ItemCategoryApi {
	var i = ItemCategoryApi{}
	i.BaseApi = api.NewBaseApi(h)
	i.Key = conf.BASE_Source_KEY
	return &i
}

func (g *ItemCategoryApi) GetItemCategoryList(param *ItemCategoryParam,clearCache int64) (ItemCategoryResApi, int) {
	var code = error_code.SUCCESSSTATUS
	var res ItemCategoryResApi
	redisClient := utils.RedisHandle
	key := fmt.Sprintf("GetItemCategoryList_%d_%d_%s" ,param.ProjectId,param.SubjectId,param.Picid)
	respStr := redisClient.GetData(key).(string)
	if respStr == "" || clearCache == 1 {
		g.Uri = fmt.Sprintf(conf.BASE_DOMAIN+"/tiku/ItemCategory?project_id=%d&subject_id=%d&picid=%s&type=one&is_high_error=false",
			param.ProjectId, param.SubjectId, param.Picid)
		r, err := g.HttpClient.Get(g.Uri, g.WhereParam, g.ApiHeader, g.Key)
		if err != nil {
			setting.Logger.Infof("GetItemCategoryList_%s",err.Error())
			code = error_code.ITEMCATEGORYAPI
			return res, code
		}
		respStr = r.String()
		redisClient.SetData(key, respStr, time.Second*86400)
	}
	sj, _ := simpleJson.NewJson([]byte(respStr))
	res.Result = sj
	return res, code
}
