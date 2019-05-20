package wx_program

import (
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"fmt"
)

type LiveRes struct {
	Status int `json:"status"`
	Info string `json:"info"`
	Data LiveDesc `json:"data"`
} 

type LiveDesc struct {
	Title string `json:"title"`
	Starttime int64 `json:"Starttime"`
	Endtime int64 `json:"Endtime"`
}



// 获取直播详情，暂时放在这
func GetLiveDesc(liveId int) (*LiveRes, error) {
	resStruct := LiveRes{}
	res, err := req.Get(conf.LACHESIS_DOMAIN + "/api/zhibo/" + fmt.Sprintf("%d", liveId))
	if err != nil {
		return &resStruct, err
	}
	err = res.ToJSON(&resStruct)

	return &resStruct, err
}
