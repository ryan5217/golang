package primary

import (
	"github.com/imroc/req"

	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type NoteApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}
type NoteResApi struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Result  []Note `json:"result"`
}
type NoteAddResApi struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
type Note struct {
	Id                   int64  `json:"id"`
	Content              string `json:"content"`
	LiveSeconds          int64  `json:"live_seconds"`
	IsPublic             int64  `json:"is_public"`
	SourceId             int64  `json:"source_id"`
	ResourceItemId       int64  `json:"resource_item_id"`
	SourceType           string `json:"source_type"`
	UpdatedAt            string `json:"updated_at"`
	Uid                  int64
	Origin               string `json:"origin"`
	CourseSyllabusItemId int64  `json:"course_syllabus_item_id"`
	CourseId             int64  `json:"course_id"`
	CourseName           string `json:"course_name"`
	SourceName           string `json:"source_name"`
}
type SearchParams struct {
	StudentId  int64
	CourseId   int64
	ResourceId int64
	Page       int64
	Limit      int64
}
type NoteData struct {
	Content                 string
	Live_seconds            int64
	Is_public               int64
	Resource_id             int64
	Resource_item_id        int64
	Resource_type           string
	Updated_at              string
	Student_id              int64
	Origin                  string
	Course_syllabus_item_id int64
	Course_id               int64
}

func NewNoteApi(h *utils.HttpClient) *NoteApi {
	var p = NoteApi{}
	p.HttpClient = h
	return &p
}

func (g *NoteApi) GetNoteList(param SearchParams) (*NoteResApi, error) {
	var res NoteResApi
	g.Uri = fmt.Sprintf(conf.STUDY_SERVICE_DOMAIN+"/note/list?student_id=%d&course_id=%d&limit=%d&page=%d", param.StudentId, param.CourseId, param.Limit, param.Page)
	if(param.ResourceId > 0){
		g.Uri += fmt.Sprintf("&source_id=%d",param.ResourceId)
	}
	r, err := req.Get(g.Uri)

	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
func (g *NoteApi) AddNote(param *NoteData) (*NoteAddResApi, error) {
	var res NoteAddResApi
	g.Uri = conf.STUDY_SERVICE_DOMAIN + "/note"
	query_param := req.QueryParam{
		"content":                 param.Content,
		"live_seconds":            param.Live_seconds,
		"is_public":               param.Is_public,
		"resource_id":             param.Resource_id,
		"resource_type":           param.Resource_type,
		"origin":                  param.Origin,
		"course_syllabus_item_id": param.Course_syllabus_item_id,
		"course_id":               param.Course_id,
		"student_id":              param.Student_id,
	}
	r, err := req.Post(g.Uri, query_param)

	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
func (g *NoteApi) EditNote(param *Note) (*NoteAddResApi, error) {
	var res NoteAddResApi
	g.Uri = fmt.Sprintf(conf.STUDY_SERVICE_DOMAIN+"/note/%d", param.Id)
	query_param := req.QueryParam{
		"content":                 param.Content,
		"live_seconds":            param.LiveSeconds,
		"is_public":               param.IsPublic,
		"resource_id":             param.SourceId,
		"resource_type":           param.SourceType,
		"student_id":              param.Uid,
		"course_id":               1,
		"origin":                  "app",
		"course_syllabus_item_id": 1,
	}
	r, err := req.Put(g.Uri, query_param)

	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
