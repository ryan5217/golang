package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	jhService "milano.gaodun.com/service/junior_home"
)

var JuniorHomeApi = NewJuniorHomeApi()

func NewJuniorHomeApi() *JuniorHome {
	return &JuniorHome{}
}

type JuniorHome struct {
	Base
}

func (i JuniorHome) GetHome(c *gin.Context) {
	param := jhService.Param{}
	param.ProjectId = i.GetInt64(c, "project_id", true)
	param.SubjectId = i.GetInt64(c, "subject_id", true)
	param.StudentId = i.GetInt64(c, "uid")
	param.IsAudit = i.GetInt64(c, "is_sh")
	param.ClearCache = i.GetInt64(c, "clear_cache")
	param.Version = i.GetString(c, "version")
	param.Source = i.GetInt64(c, "source")
	if param.StudentId == 0 {
		i.ServerJSONError(c, &jhService.JuniorHomeResp{}, error_code.UIDEMPTYERR)
		return
	}
	if c.GetBool(Verify) {
		return
	}
	bServer := jhService.NewJuniorHomeService(setting.GinLogger(c))
	homeResp, code := bServer.GetHome(&param)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, homeResp)
	}
}
func (i JuniorHome) GetSalesMan(c *gin.Context) {
	uid := i.GetInt64(c, "uid", true)
	clearCache := i.GetInt64(c, "clear_cache")
	if c.GetBool(Verify) {
		return
	}
	bServer := jhService.NewJuniorHomeService(setting.GinLogger(c))
	homeResp, code := bServer.GetSalesMan(uid, clearCache, false, true)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, homeResp)
	}
}
func (i JuniorHome) GetHomeStudyInfo(c *gin.Context) {
	param := jhService.Param{}
	param.StudentId = i.GetInt64(c, "uid", true)
	param.IsAudit = i.GetInt64(c, "is_sh")
	param.ClearCache = i.GetInt64(c, "clear_cache")
	param.Version = i.GetString(c, "version")
	param.Source = i.GetInt64(c, "source")
	if c.GetBool(Verify) {
		return
	}
	bServer := jhService.NewJuniorHomeService(setting.GinLogger(c))
	homeResp, code := bServer.GetHomeStudyInfo(&param)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, homeResp)
	}
}

/**
获取营销模块
证券从业
*/
func (i JuniorHome) GetMarketingInfo(c *gin.Context) {
	param := jhService.MarketingParam{}
	param.StudentId = i.GetInt64(c, "uid")
	param.ClearCache = i.GetInt64(c, "clear_cache")
	param.ClassKey = i.GetString(c, "class_sign", true)
	param.PublicKey = i.GetString(c, "public_sign", true)
	if param.StudentId == 0 {
		i.ServerJSONError(c, &jhService.JuniorHomeResp{}, error_code.UIDEMPTYERR)
		return
	}
	if c.GetBool(Verify) {
		return
	}
	bServer := jhService.NewJuniorHomeService(setting.GinLogger(c))
	info, code := bServer.GetMarketingInfo(&param)
	if code != error_code.SUCCESSSTATUS {
		i.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		i.ServerJSONSuccess(c, info)
	}
}
