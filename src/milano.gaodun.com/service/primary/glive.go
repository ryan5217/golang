package primary

import (
	"fmt"
	"github.com/apex/log"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api/primary"
	sub "milano.gaodun.com/model/sub_glive"
	"milano.gaodun.com/pkg/utils"
	"strconv"
	"time"
)

type GliveService struct {
	ga       *primary.GliveApi
	subGlive *sub.GdStudentSubGliveModel
	logger   *log.Entry
	gm       *primary.Glive
}
type GliveView struct {
	Title           string    `json:"title"`
	GliveId         int64     `json:"glive_id"`
	StartTimestamp  string    `json:"start_timestamp"`
	EndTimestamp    string    `json:"end_timestamp"`
	S_time          string    `json:"s_time"`
	E_time          string    `json:"e_time"`
	Show_url        string    `json:"show_url"`
	Back_url        string    `json:"back_url"`
	Teacher_info    Teacher   `json:"teacher_info"`
	Issub           string    `json:"issub"`
	Share_info      ShareInfo `json:"share_info"`
	Glive_sub_count string    `json:"glive_sub_count"`
}

type Teacher struct {
	Name string `json:"name"`
}
type ShareInfo struct {
	Url     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewGliveService(logger *log.Entry) *GliveService {
	h := utils.HttpHandle
	var g GliveService
	g.logger = logger
	g.ga = primary.NewGliveApi(h)
	g.subGlive = sub.NewGdStudentSubGliveModel()
	return &g
}

func (g *GliveService) GetGlive(projectId int64, source string) (*GliveView, error) {
	gra, err := g.ga.GetGlive(projectId)
	gm := gra.Glist
	glv := GliveView{}
	gls := []GliveView{}
	local, _ := time.LoadLocation("Asia/Chongqing")
	now := time.Now().Format("2006-01-02")
	for _, glive := range *gm {
		glv.Title = glive.Name
		glv.GliveId = glive.Id
		glv.StartTimestamp = fmt.Sprintf("%d", glive.StartTime/1000)
		glv.EndTimestamp = fmt.Sprintf("%d", glive.StartTime/1000+glive.Duration)
		glv.S_time = time.Unix(glive.StartTime/1000, 0).In(local).Format("15:04")
		glv.E_time = time.Unix(glive.StartTime/1000+glive.Duration, 0).In(local).Format("15:04")
		glv.Show_url = conf.GLIVE_GAODUN_COM + "/i.html?meeting=" + fmt.Sprintf("%d", glive.Id)
		if len(glive.Replay) > 0 {
			replay := glive.Replay[0]
			if replay != "" && replay[0:4] != "http" {
				glv.Back_url = fmt.Sprintf("http://video1.cdn.gaodun.com/pub/%s/SD.mp4", replay)
			}
		}
		if len(glive.Teacher) > 0 {
			glv.Teacher_info.Name = gra.Teacher[fmt.Sprintf("%d", glive.Teacher[0])].Name
		} else {
			glv.Teacher_info.Name = "高顿名师"
		}
		glv.Share_info.Url = glv.Show_url + "&viewername=&appid=" + source
		glv.Share_info.Title = "我在看【" + glv.Title + "】直播，帮助很大"
		glv.Share_info.Content = "提分点讲解精准到位"
		glv.Issub = "0"
		glv.Glive_sub_count = "1"
		start := time.Unix(glive.StartTime/1000, 0).Format("2006-01-02")
		if start == now {
			gls = append(gls, glv)
		}
	}
	gls = g.SortGliveListDesc(gls)
	nowUnixt := time.Now().Unix()
	glv = GliveView{}
	for _, v := range gls {
		start, _ := strconv.ParseInt(v.StartTimestamp, 10, 64)
		end, _ := strconv.ParseInt(v.EndTimestamp, 10, 64)
		if nowUnixt < start || (nowUnixt > start && nowUnixt < end) {
			glv = v
		}
	}
	subCount, err := g.subGlive.GetCount(glv.GliveId, projectId)
	glv.Glive_sub_count = fmt.Sprintf("%d", subCount)
	if err != nil {
		g.logger.Error(err.Error())
	}
	return &glv, err
}
func (g *GliveService) SortGliveListDesc(arr []GliveView) []GliveView {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			ii, _ := strconv.ParseInt(arr[i].StartTimestamp, 10, 64)
			jj, _ := strconv.ParseInt(arr[j].StartTimestamp, 10, 64)
			if ii < jj {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}
