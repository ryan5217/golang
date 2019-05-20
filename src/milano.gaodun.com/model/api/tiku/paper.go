package tiku

import (
	"encoding/json"
	"errors"
	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api"
	"milano.gaodun.com/pkg/utils"
)

type PaperResApi struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	Result Paper  `json:"result"`
}

// base tiku 返回数据结构
type Paper struct {
	Id        string     `json:"id"`
	Title     string     `json:"title"`
	Etype     string     `json:"etype"`
	SubjectId string     `json:"subject_id"`
	Year      string     `json:"year"`
	Score     string     `json:"score"`
	ItemData  []ItemData `json:"item_data"`
}

type ItemData struct {
	Id       string  `json:"id"`
	Title    string  `json:"title"`
	Type     string  `json:"type"`
	Itemsnum string  `json:"itemsnum"`
	Items    []Items `json:"items"`
}

type Items struct {
	Id       string `json:"id"`
	Scoreses string `json:"scoreses"`
}

type PaperApi struct {
	*api.BaseApi
}

func NewPaperApi(h *utils.HttpClient) *PaperApi {
	var p = PaperApi{}
	p.BaseApi = api.NewBaseApi(h)
	p.Uri = conf.BASE_DOMAIN + "/tiku/paper"
	p.Key = conf.BASE_TIKU_KEY

	return &p
}

// 获取试卷
func (p *PaperApi) GetPaperOne(pkId int64) (*Paper, error) {
	r, err := p.HttpClient.Get(p.Uri+"/"+fmt.Sprint(pkId), p.WhereParam, p.ApiHeader, p.Key)
	var res PaperResApi
	if err == nil {
		err := json.Unmarshal(r.Bytes(), &res)
		return &res.Result, err
	}

	if res.Status != "0" {
		return &res.Result, errors.New(res.Info)
	}

	return &res.Result, err
}

func (p *PaperApi) SetBaseApi(b *api.BaseApi) *PaperApi {
	p.BaseApi = b

	return p
}
