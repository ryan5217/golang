package activity

import (
	"fmt"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
)

type ActivityApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

func NewActivityApi(h *utils.HttpClient) *ActivityApi {
	var p = ActivityApi{}
	p.HttpClient = h
	return &p
}

func (g *ActivityApi) GetSeckillActivityDetail(seckillId int64) (*simpleJson.Json, int) {
	g.Uri = fmt.Sprintf(conf.SPARTA_DOMAIN+"/v1/seckill/one?seckill_id=", seckillId)
	r, err := req.Get(g.Uri)
	if err != nil {
		return &simpleJson.Json{}, error_code.SECKILLAPIERR
	}
	res, err := simpleJson.NewJson(r.Bytes())
	if err != nil {
		return &simpleJson.Json{}, error_code.SECKILLAPIERR
	}
	return res, error_code.SUCCESSSTATUS
}
func (g *ActivityApi) GetGroupActivityDetail(groupId int64, uid int64) (*simpleJson.Json, int) {
	g.Uri = fmt.Sprintf(conf.SPARTA_DOMAIN + "/v1/group-buying/detail")
	header := req.Param{}
	header["student_id"] = uid
	header["group_id"] = groupId
	r, err := req.Post(g.Uri, header)
	if err != nil {
		return &simpleJson.Json{}, error_code.SECKILLAPIERR
	}
	res, err := simpleJson.NewJson(r.Bytes())
	if err != nil {
		return &simpleJson.Json{}, error_code.SECKILLAPIERR
	}
	return res, error_code.SUCCESSSTATUS
}
