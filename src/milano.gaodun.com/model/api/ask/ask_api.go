package ask

import (
	"github.com/imroc/req"

	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type AskApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type Ask struct {
	Id                   string   `json:"id"`
	CourseId             string   `json:"course_id"`
	CourseName           string   `json:"course_name"`
	ResourceId           string   `json:"resource_id"`
	ResourceName         string   `json:"resource_name"`
	StudentId            string   `json:"student_id"`
	NickName             string   `json:"nick_name"`
	Content              string   `json:"content"`
	SourceType           string   `json:"source_type"`
	CourseSyllabusItemId string   `json:"course_syllabus_item_id"`
	Picture              string   `json:"picture"`
	VideoTime            string   `json:"video_time"`
	Reply                Reply    `json:"answer"`
	Images               []string `json:"images"`
	CreatedAt            string   `json:"created_at"`
}
type Reply struct {
	Id        string   `json:"id"`
	Content   string   `json:"content"`
	StudentId string   `json:"student_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Images    []string `json:"images"`
}
type AskParam struct {
	Uid        int64
	CourseId   int64
	SourceType string
	Page       int64
	Limit      int64
}

func NewAskApi(h *utils.HttpClient) *AskApi {
	var p = AskApi{}
	p.HttpClient = h
	return &p
}

func (g *AskApi) GetAskList(param *AskParam) (string, error) {
	g.Uri = fmt.Sprintf(conf.BASE_DOMAIN+"/care/home/asklist?student_id=%d&course_id=%d&source_type=saas&teacher_answer=1&page=%d&ps=%d&order=is_top desc,regdate desc",
		param.Uid, param.CourseId, param.Page, param.Limit)
	r, err := req.Get(g.Uri)
	return r.String(), err
}
func (g *AskApi) GetAskDetail(askId int64) (string, error) {
	g.Uri = fmt.Sprintf(conf.BASE_DOMAIN+"/care/Home/QuestionRest/getOneAskInfo?ask_id=%d&json_result=1", askId)
	r, err := req.Get(g.Uri)
	return r.String(), err
}
