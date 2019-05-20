package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type H5 struct {
	Base
}

var H5Api = NewH5Api()

func NewH5Api() *H5 {
	return &H5{}
}

func (h H5) GetApp(c *gin.Context) {
	agent := c.GetHeader("User-Agent")
	agent = strings.ToUpper(agent)
	// iPhone
	if strings.Contains(agent, "IPHONE") {
		c.Redirect(http.StatusTemporaryRedirect, "https://itunes.apple.com/us/app/初级会计职称题库-2018会计师考试网校课堂/id858332938?mt=8&ign-mpt=uo=4")
	}
	// Android
	if strings.Contains(agent, "ANDROID") {
		c.Redirect(http.StatusTemporaryRedirect, "https://v.gaodun.com/Apps/client_iosazapk/iosazurl/kjzc")
	}
	// m gaodun
	c.Redirect(http.StatusTemporaryRedirect, "https://m.gaodun.com")
}
