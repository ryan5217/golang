package primary

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"milano.gaodun.com/conf"
	tsc "milano.gaodun.com/model/tiku_system_constant"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type GliveApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}
type GliveResApi struct {
	Name    string             `json:"name"`
	Teacher map[string]Teacher `json:"teacher"`
	Glist   *[]Glive           `json:"meeting"`
}
type Teacher struct {
	Name    string `json:"name"`
	Brief   string `json:"brief"`
	Avantar string `json:"avantar"`
	Weixin  string `json:"weixin"`
}

type Glive struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	StartTime int64    `json:"startTime"`
	EndTime   int64    `json:"endTime"`
	Duration  int64    `json:"duration"`
	Data      int64    `json:"data"`
	Teacher   []int64  `json:"teacher"`
	Replay    []string `json:"replay"`
}

func NewGliveApi(h *utils.HttpClient) *GliveApi {
	var p = GliveApi{}
	p.HttpClient = h
	return &p
}

func (g *GliveApi) GetGlive(projectId int64) (*GliveResApi, error) {
	var res GliveResApi
	var classid string
	var jsonData map[string]map[string]string
	client := utils.RedisHandle
	cache := client.GetData("TikuSystemConstant_GLIVE_CLASS_MEETING_INFO")
	if cache == "" {
		parm := tsc.SearchParam{Thekey: "GLIVE_CLASS_MEETING_INFO"}
		tscm := tsc.NewTikuSystemConstantModel(&parm)
		data, err := tscm.GetKey()
		if err != nil {
			return &res, err
		}
		theValue := data.Thevalue
		json.Unmarshal([]byte(theValue), &jsonData)
		prjs := fmt.Sprintf("%d", projectId)
		classid = jsonData[prjs]["classid"]
		client.SetData("TikuSystemConstant_GLIVE_CLASS_MEETING_INFO", theValue, time.Second*5)
	} else {
		classid = gjson.Parse(fmt.Sprintf("%s", cache)).Get(fmt.Sprintf("%d", projectId)).Get("classid").String()
	}
	g.Uri = conf.GLIVE_LIST_DOMAIN + "/class/" + classid + ".js"
	r, err := req.Get(g.Uri)

	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	}
	return &res, err
}
