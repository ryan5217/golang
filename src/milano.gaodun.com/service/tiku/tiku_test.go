package tiku

import (
	"fmt"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api/item"
	"milano.gaodun.com/model/api/tiku"
	"milano.gaodun.com/pkg/utils"
	"testing"
)

func TestGetPaper(t *testing.T) {
	param := map[string]interface{}{"combine": "inIds"}
	//param["condition"] = `{"id":"234,4442,5522,3542"}`
	param["is_need_item"] = "y"
	param["is_need_page"] = "n"
	//param["student_id"] = 1694698

	res, err := utils.HttpHandle.Get("http://t-base.gaodun.com/tiku/paper/3542", param, map[string]string{}, conf.BASE_TIKU_KEY)
	//res, err := utils.HttpHandle.Get("http://www.baidu.com", param, map[string]string{}, conf.BASE_TIKU_KEY)
	fmt.Println(string(res.Bytes()), err)
}

func TestGetPaperOne(t *testing.T) {
	paperApi := tiku.NewPaperApi(utils.HttpHandle)
	r, err := paperApi.SetBaseApi(paperApi.SetParam("is_need_item", "y")).GetPaperOne(3542)
	fmt.Println(r, err)

	itemApi := item.NewItemApi(utils.HttpHandle)
	re, err := itemApi.SetBaseApi(itemApi.SetParam("is_need_all", 1)).GetOne(50120)
	fmt.Println(re, err)

}

func TestPaperService_GetPaper(t *testing.T) {
	p := NewPaperService(utils.HttpHandle, nil)
	p.GetPaper(3542)
}
