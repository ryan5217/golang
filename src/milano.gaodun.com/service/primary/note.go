package primary

import (
	"github.com/apex/log"
	simpleJson "github.com/bitly/go-simplejson"
	"milano.gaodun.com/model/api/primary"
	"milano.gaodun.com/pkg/utils"
	"strconv"
)

type NoteService struct {
	nt     *primary.NoteApi
	logger *log.Entry
	nv     *NoteView
	toc    *primary.TocApi
	course *primary.CourseApi
}
type NoteView struct {
	NoteList []primary.Note
}
type SyllabusItem struct {
	List map[string]Item
}
type Item struct {
	Id       string
	Name     string
	ParentId string
}

func NewNoteService(logger *log.Entry) *NoteService {
	h := utils.HttpHandle
	var g NoteService
	g.logger = logger
	g.nt = primary.NewNoteApi(h)
	g.toc = primary.NewTocApi(h)
	g.course = primary.NewCourseApi(h)
	g.nv = &NoteView{}
	return &g
}
func (g *NoteService) GetNoteList(params primary.SearchParams) (*NoteView, error) {
	//获取笔记原数据
	ntapi, err := g.nt.GetNoteList(params)
	if err != nil {
		return g.nv, err
	}
	//获取课程信息
	course, err := g.course.GetCourseById(params.CourseId)
	courseJson, err := simpleJson.NewJson([]byte(course))
	if err != nil {
		return g.nv, err
	}
	//获取阶段大纲列表
	gradation, err := g.toc.GetSyllabusList(params.CourseId)
	if err != nil {
		return g.nv, err
	}
	gra, err := simpleJson.NewJson([]byte(gradation))
	syllabusList := []simpleJson.Json{}
	for k, _ := range gra.MustArray() {
		for i := 0; i < len(gra.GetIndex(k).Get("syllabus").MustArray()); i++ {
			syllabusList = append(syllabusList, *gra.GetIndex(k).Get("syllabus").GetIndex(i))
		}
	}
	//获取大纲条目列表列表si用来获取条目名称
	si := SyllabusItem{make(map[string]Item)}
	g.GetSyllabusItemList(syllabusList, &si)
	if err != nil {
		return g.nv, err
	}

	for _, v := range ntapi.Result {
		//根据资源id获取大纲条目名称
		v.SourceName = si.List[si.List[strconv.Itoa(int(v.SourceId))].ParentId].Name
		v.CourseName, _ = courseJson.Get("result").Get("course_name").String()
		v.Uid = params.StudentId
		g.nv.NoteList = append(g.nv.NoteList, v)
	}

	return g.nv, err
}
func (g *NoteService) GetSyllabusItemList(syllabusList []simpleJson.Json, si *SyllabusItem) {
	for _, v := range syllabusList {
		key, err := v.Get("id").Int()
		if key != 0 && err == nil {
			it := Item{}
			it.Id, _ = v.Get("id").String()
			it.Name, _ = v.Get("name").String()
			it.ParentId, _ = v.Get("parent_id").String()
			si.List[strconv.Itoa(key)] = it
		}
		resourceId, err := v.Get("resource_id").String()
		if resourceId != "" && err == nil && resourceId != "0" {
			it := Item{}
			it.Id, _ = v.Get("id").String()
			it.Name, _ = v.Get("name").String()
			it.ParentId, _ = v.Get("parent_id").String()
			si.List[resourceId] = it
		}
		children := v.Get("children").MustArray()
		if len((children)) > 0 {
			s_list := []simpleJson.Json{}
			for k, _ := range children {
				s_list = append(s_list, *v.Get("children").GetIndex(k))
			}
			g.GetSyllabusItemList(s_list, si)
		}
	}
}
func (g *NoteService) AddNote(data *primary.NoteData)(*primary.NoteAddResApi,error) {
	return g.nt.AddNote(data)
}
func (g *NoteService) EditNote(data *primary.Note)(*primary.NoteAddResApi,error) {
	return g.nt.EditNote(data)
}