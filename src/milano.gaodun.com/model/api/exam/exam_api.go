package primary

import (
	"fmt"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
	simpleJson "github.com/bitly/go-simplejson"
	"milano.gaodun.com/pkg/error-code"
)

type ExamApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

func NewExamApi(h *utils.HttpClient) *ExamApi {
	var p = ExamApi{}
	p.HttpClient = h
	return &p
}

func (g *ExamApi) GetExamInfo(subjectId int64) (*simpleJson.Json, int) {
	g.Uri = conf.STUDYAPI_DOMAIN + fmt.Sprintf("/api/exam-dates?subject_id=%d", subjectId)
	rep, err := req.Get(g.Uri)
	if err != nil {
		return &simpleJson.Json{},error_code.EXAMAPIERR
	}
	res,err := simpleJson.NewJson(rep.Bytes())
	if err != nil {
		return &simpleJson.Json{},error_code.EXAMAPIERR
	}
	status,_ := res.Get("http_status").Int64()
	if status == 404 {
		return &simpleJson.Json{},error_code.EXAMDATEEMPTY
	}
	return res, error_code.SUCCESSSTATUS
}