package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	advertService "milano.gaodun.com/service/advert"
)

var AdvertApi = NewAdvertApi()

func NewAdvertApi() *Advert {
	return &Advert{}
}

type Advert struct {
	Base
}

func (i Advert) GetAdvertList(c *gin.Context) {
	uid := i.GetInt64(c, "uid", true)
	projectId := i.GetInt64(c, "project_id", true)
	paIds :=i.GetString(c,"pa_ids",true)
	if c.GetBool(Verify) {
		return
	}
	server := advertService.NewAdvertService(setting.GinLogger(c))
	resp, code := server.GetFloatAdvertList(uid,projectId,paIds)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, resp)
	}
}
func (i Advert) AddAdvertRecord(c *gin.Context) {
	uid := i.GetInt64(c, "uid", true)
	advertId := i.GetInt64(c, "advert_id", true)
	if c.GetBool(Verify) {
		return
	}
	server := advertService.NewAdvertService(setting.GinLogger(c))
	resp, err := server.AddAdvertRecord(uid,advertId)
	if err != nil {
		i.ServerJSONError(c, error_code.INFO[error_code.DBERR], error_code.DBERR)
	} else {
		i.ServerJSONSuccess(c, resp)
	}
}
