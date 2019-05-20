package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/service/tiku_constant"
	"milano.gaodun.com/pkg/error-code"
)

// 题库相关
type TikuBase struct {
	Base
}

var TikuBaseApi = NewTikuBaseApi()

func NewTikuBaseApi() *TikuBase {
	return &TikuBase{}
}

// 获取常量
func (y *TikuBase) TkConst(c *gin.Context) {
	key := y.GetString(c, "key")
	val, err := tiku_constant.TkConstList.GetKey(key)
	if err != nil {
		setting.GinLogger(c).Warn("tiku_const_" + err.Error())
		y.ServerJSONError(c, val, error_code.NOT_FIND)
		return
	}
	y.ServerJSONSuccess(c, val)
}
