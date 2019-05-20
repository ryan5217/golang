package primary

import (
	"fmt"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type CourseApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}
type CourseListResp struct {
	PackageName string
	SaasCourses []CourseInfo
	IsListen    string
}
type CourseInfo struct {
	StudentsNum  int64
	CourseId     int64
	CourseName   string
	CourseType   int64
	SaasCourseId int64
	AllChapter   int64
	Progress     float32
	LiveList     interface{}
}
type CourseResp struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Result  []StudyProgressInfo
}
type StudyProgressInfo struct {
	AllChapter          int64   `json:"allChapter"`
	FinishedChapter     int64   `json:"finishedChapter"`
	CourseId            string  `json:"courseId"`
	CurrentStudySubject string  `json:"currentStudySubject"`
	CourseType          string  `json:"courseType"`
	Progress            float32 `json:"progress"`
	LastLeaningTime     int64   `json:"lastLeaningTime"`
	CurrentStudyUrl     string  `json:"currentStudyUrl"`
}
type StudyProgressParam struct {
	CourseId int64
	Uid      int64
	CsItemId int64
	Status   int64
}

func NewCourseApi(h *utils.HttpClient) *CourseApi {
	var p = CourseApi{}
	p.HttpClient = h
	return &p
}

//获取课程下的阶段大纲列表
func (g *CourseApi) GetCourseById(CourseId int64) (string, error) {
	g.Uri = conf.COURSE_SERVICE_DOMAIN + fmt.Sprintf("/course/%d", CourseId)
	rep, err := req.Get(g.Uri)
	if err != nil {
		setting.Logger.Infof("CourseApi_GetCourseById_%s", err.Error())
	}
	return rep.String(), err
}
func (g *CourseApi) GetCourseStudyProgress(uid int64, param req.Param) (*CourseResp, error) {
	res :=CourseResp{Result:[]StudyProgressInfo{}}
	g.Uri = conf.STUDYAPI_DOMAIN + fmt.Sprintf("/student/%d/course/progress/saas", uid)
	r, err := req.Get(g.Uri, param)
	if err != nil {
		setting.Logger.Infof("CourseApi_GetCourseStudyProgress_%s", err.Error())
		return &res,err
	}
	err = r.ToJSON(&res)
	return &res, err
}
func (g *CourseApi) GetSyllabusStudyStatus(uid int64, CourseId int64, syllabusId string) (string, error) {
	g.Uri = conf.STUDY_SERVICE_DOMAIN + fmt.Sprintf("/student/%d/course/%d/syllabus/%s/tree", uid, CourseId, syllabusId)
	rep, err := req.Get(g.Uri)
	if err != nil {
		setting.Logger.Infof("CourseApi_GetSyllabusStudyStatus_%s", err.Error())
	}
	return rep.String(), err
}
func (g *CourseApi) RecordSyllabusStudyStatus(param StudyProgressParam) (string, error) {
	requestParam := req.QueryParam{
		"cs_item_id": param.CsItemId,
		"status":     param.Status,
	}
	g.Uri = conf.STUDY_SERVICE_DOMAIN + fmt.Sprintf("/student/%d/course/%d/progress", param.Uid, param.CourseId)
	rep, err := req.Post(g.Uri, requestParam)
	if err != nil {
		setting.Logger.Infof("CourseApi_RecordSyllabusStudyStatus_%s", err.Error())
	}
	return rep.String(), err
}
