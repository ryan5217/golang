package tag

import (
	"fmt"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type TagApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}
type TagResApi struct {
	Code    int64
	Message string
	Result  map[string]map[string]TagClass
}
type TagFidResApi struct {
	Code    int64
	Message string
	Result  map[string][]Tag
}
type TagClass struct {
	Id       int64
	TagId    int64
	FidId    int64
	FidType  int64
	IsDelete int64
	Tags     []Tag
}
type Tag struct {
	Id       int64
	Name     string
	EName    string
	Summary  string
	Hot      int64
	ClassId  int64
	IsDelete int64
	ParentId int64
}

func NewTagApi(h *utils.HttpClient) *TagApi {
	var p = TagApi{}
	p.HttpClient = h
	return &p
}
func (g *TagApi) GetTagListByFid(fidType int64, fidId int64) (*TagFidResApi, error) {
	var res TagFidResApi
	g.Uri = conf.PROMETHEUS_DOMAIN + fmt.Sprintf("/v1/tag/tag-fid?fid_type=%d&fid_id=%d", fidType, fidId)
	r, err := req.Get(g.Uri)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
func (g *TagApi) GetTagList(fidType int64, fidIds string) (*TagResApi, error) {
	var res TagResApi
	g.Uri = conf.PROMETHEUS_DOMAIN + fmt.Sprintf("/v1/tag/tag-fids?fid_type=%d&fid_ids=%s", fidType, fidIds)
	r, err := req.Get(g.Uri)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
