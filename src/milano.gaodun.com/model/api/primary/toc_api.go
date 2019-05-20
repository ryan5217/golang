package primary

import (
	"fmt"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type TocApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

func NewTocApi(h *utils.HttpClient) *TocApi {
	var p = TocApi{}
	p.HttpClient = h
	return &p
}

//获取课程下的阶段大纲列表
func (g *TocApi) GetSyllabusList(CourseId int64) (string, error) {
	g.Uri = conf.TOC_SERVICE_DOMAIN + fmt.Sprintf("/course/%d/syllabus/items", CourseId)
	rep, err := req.Get(g.Uri)
	return rep.String(), err
}
