package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/service/gd-info"
)

var GdInfApi = NewGdInfoApi()

func NewGdInfoApi() *GdInfo {
	return &GdInfo{}
}

type GdInfo struct {
	Base
}

func (g GdInfo) Get(c *gin.Context) {
	limit := g.GetInt64(c, "limit")
	page := g.GetInt64(c, "page")
	infoType := g.GetString(c, "info_type", true)
	if c.GetBool(Verify) {
		return
	}
	gs := gd_info.NewGdInfoService(setting.GinLogger(c))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 15
	}

	res, _ := gs.Get(infoType, page, limit)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(200, fmt.Sprintf(`{"code":0, "message":"请求成功", "result":%s}`, res))
}

func (g GdInfo) GetInfo(c *gin.Context) {
	fileUrl := g.GetString(c, "file_url", true)
	if c.GetBool(Verify) {
		return
	}
	gs := gd_info.NewGdInfoService(setting.GinLogger(c))

	res, _ := gs.GetInfo(fileUrl)
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(200, fmt.Sprintf(`{"code":0, "message":"请求成功", "result":%s}`, res))
}
