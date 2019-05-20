package ask

import (
	"fmt"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type TreeListApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type TreeListResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewTreeListApi(h *utils.HttpClient) *TreeListApi {
	var p = TreeListApi{}
	p.HttpClient = h
	return &p
}

func (g *TreeListApi) GetTreeListBySign(classSign string, clearCache int64) (string, int) {
	var code = error_code.SUCCESSSTATUS
	redisClient := utils.RedisHandle
	treeLisKey := "TreeList" + classSign
	respStr := redisClient.GetData(treeLisKey).(string)
	if respStr == "" || clearCache == 1 {
		g.Uri = fmt.Sprintf(conf.SPARTA_DOMAIN+"/v1/list/index?class_sign=%s&clear_cache=%d", classSign, clearCache)
		r, err := req.Get(g.Uri)
		if err != nil {
			setting.Logger.Infof("GetTreeListBySign_%s", err.Error())
			code = error_code.TREELISTAPI
		}
		redisClient.SetData(treeLisKey, r.String(), time.Second*86400)
		respStr = r.String()
	}
	return respStr, code
}
