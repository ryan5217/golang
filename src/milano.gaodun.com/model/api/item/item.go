package item

import (
	"encoding/json"
	"errors"
	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api"
	"milano.gaodun.com/pkg/utils"
)

type ItemInfo struct {
	Title    string `json:"title"`
	Option   string `json:"option"`
	Partnum  string `json:"partnum"`
	Type     string `json:"type"`
	Answer   string `json:"answer"`
	Analysis string `json:"analysis"`
	ItemId   int64  `json:"item_id"`
	Flag     string `json:"flag"`
}

type ItemResApi struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	Result struct {
		RequestItem ItemInfo `json:"request_item"`
	} `json:"result"`
}

type ItemApi struct {
	*api.BaseApi
}

func NewItemApi(h *utils.HttpClient) *ItemApi {
	var i = ItemApi{}
	i.BaseApi = api.NewBaseApi(h)
	i.Uri = conf.BASE_DOMAIN + "/tiku/item"
	i.Key = conf.BASE_TIKU_KEY

	return &i
}

// 获取题目信息
func (i *ItemApi) GetOne(pkId int64) (*ItemInfo, error) {
	r, err := i.HttpClient.Get(i.Uri+"/"+fmt.Sprint(pkId), i.WhereParam, i.ApiHeader, i.Key)
	var res ItemResApi
	fmt.Println(r)
	if err == nil {
		err := json.Unmarshal(r.Bytes(), &res)
		return &res.Result.RequestItem, err
	}

	if res.Status != "0" {
		return &res.Result.RequestItem, errors.New(res.Info)
	}

	return &res.Result.RequestItem, err
}

func (i *ItemApi) SetBaseApi(b *api.BaseApi) *ItemApi {
	i.BaseApi = b

	return i
}
