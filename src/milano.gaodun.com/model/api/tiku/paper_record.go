package tiku

import (
	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
	simpleJson "github.com/bitly/go-simplejson"
	"milano.gaodun.com/pkg/setting"
)

type PaperRecordApi struct {
	*api.BaseApi
	Uri string
}

func NewPaperRecordApi(h *utils.HttpClient) *PaperRecordApi {
	var i = PaperRecordApi{}
	i.BaseApi = api.NewBaseApi(h)
	i.Key = conf.BASE_Source_KEY
	return &i
}

func (g *PaperRecordApi) GetPaperRecordList(uid int64, paperIds string) (*simpleJson.Json, int) {
	var code = error_code.SUCCESSSTATUS
	g.Uri = fmt.Sprintf(conf.BASE_DOMAIN+"/tiku/paper?is_need_page=n&student_id=%d&is_need_other_attribute=y&condition={\"id\":%s}&combine=inIds",
		uid,paperIds)
	r, err := g.HttpClient.Get(g.Uri, g.WhereParam, g.ApiHeader, g.Key)
	if err != nil {
		setting.Logger.Infof("PaperRecordApi_GetPaperRecordList_%s",err.Error())
		code = error_code.PAPERRECORDAPI
		return &simpleJson.Json{}, code
	}
	sj, err2 := simpleJson.NewJson(r.Bytes())
	if err2 != nil {
		setting.Logger.Infof("PaperRecordApi_GetPaperRecordList_%s",err.Error())
		code = error_code.PAPERRECORDAPI
		return &simpleJson.Json{}, code
	}
	return sj.Get("result"), code
}
func (g *PaperRecordApi) GetPaperStatus(uid int64, projectId int64,subjectId int64,pdid string) (*simpleJson.Json, int) {
	var code = error_code.SUCCESSSTATUS
	g.Uri = fmt.Sprintf(conf.BASE_DOMAIN+"/tiku/PaperDataStudentLog?combine=inPdIds&student_id=%d&condition={\"project_id\":%d,\"subject_id\":%d,\"paper_data_id\":%s}&type=-1",
		uid,projectId,subjectId,pdid)
	r, err := g.HttpClient.Get(g.Uri, g.WhereParam, g.ApiHeader, g.Key)
	if err != nil {
		setting.Logger.Infof("PaperRecordApi_GetPaperStatus_%s",err.Error())
		code = error_code.PAPERRECORDAPI
		return &simpleJson.Json{}, code
	}
	sj, err2 := simpleJson.NewJson(r.Bytes())
	if err2 != nil {
		setting.Logger.Infof("PaperRecordApi_GetPaperStatus_%s",err.Error())
		code = error_code.PAPERRECORDAPI
		return &simpleJson.Json{}, code
	}
	return sj.Get("result"), code
}