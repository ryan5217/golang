package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	priM "milano.gaodun.com/model/api/primary"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	priService "milano.gaodun.com/service/primary"
	"time"
)

var CourseApi = NewCourseApi()

func NewCourseApi() *Course {
	return &Course{}
}

type Course struct {
	Base
}

//获取课程大纲
func (p Course) GetCourseSyllabus(c *gin.Context) {
	courseId := p.GetInt64(c, "course_id", true)
	uid := p.GetInt64(c, "uid", true)
	if c.GetBool(Verify) {
		return
	}
	courseService := priService.NewCourseService(setting.GinLogger(c))
	nt, err := courseService.GetCourseSyllabus(courseId, uid)
	if err != nil {
		p.ServerJSONError(c, err, error_code.SYSERR)
		return
	}
	p.ServerJSONSuccess(c, nt)
}

//获取课程大纲
func (p Course) GetCourseListByGoodsId(c *gin.Context) {
	goodsId := p.GetInt64(c, "goods_id", true)
	uid := p.GetInt64(c, "uid", true)
	orderType := p.GetString(c, "type")
	clearCache := p.GetInt64(c, "clear_cache")
	if c.GetBool(Verify) {
		return
	}
	courseService := priService.NewCourseService(setting.GinLogger(c))
	nt, err := courseService.GetCourseListByGoodsId(goodsId, uid, clearCache, orderType)
	if err != nil {
		p.ServerJSONError(c, err.Error(), error_code.SYSERR)
		return
	}
	p.ServerJSONSuccess(c, nt)
}

//获取课程大纲
func (p Course) CourseSyllabusStudyStatusRecord(c *gin.Context) {
	param := priM.StudyProgressParam{}
	param.CourseId = p.GetInt64(c, "course_id", true)
	param.Uid = p.GetInt64(c, "uid", true)
	param.CsItemId = p.GetInt64(c, "cs_item_id", true)
	param.Status = p.GetInt64(c, "status", true)
	if c.GetBool(Verify) {
		return
	}
	courseService := priService.NewCourseService(setting.GinLogger(c))
	nt, err := courseService.CourseSyllabusStudyStatusRecord(param)
	if err != nil {
		p.ServerJSONError(c, err, error_code.SYSERR)
		return
	}
	p.ServerJSONSuccess(c, nt)
}
func (p Course) GetGoodsCourseList(c *gin.Context) {
	uid := p.GetInt64(c, "uid", true)
	isAudit := p.GetInt64(c, "is_sh")
	if c.GetBool(Verify) {
		return
	}
	bServer := priService.NewCourseService(setting.GinLogger(c))
	homeResp, code := bServer.GetGoodsCourseList(uid, isAudit)
	if code != error_code.SUCCESSSTATUS {
		p.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		p.ServerJSONSuccess(c, homeResp)
	}
}
func (p Course) GetColumnGoodsByGoodsIds(c *gin.Context) {
	columnId := p.GetInt64(c, "column_id", true)
	projectId := p.GetInt64(c, "project_id", true)
	goodsStr := p.GetString(c, "goods_ids", true)
	clearCache := p.GetInt64(c, "clear_cache")
	if c.GetBool(Verify) {
		return
	}
	bServer := priService.NewCourseService(setting.GinLogger(c))
	homeResp, code := bServer.GetColumnGoods(columnId, projectId, goodsStr, clearCache)
	if code != error_code.SUCCESSSTATUS {
		p.ServerJSONError(c, error_code.INFO[code], code)
	} else {
		p.ServerJSONSuccess(c, homeResp)
	}
}

//商品详情获取课程列表
func (p Course) GetCourseList(c *gin.Context) {
	resultList := priM.CourseListResp{}
	redis := utils.RedisHandle
	goodsId := p.GetInt64(c, "goods_id", true)
	clearCache := p.GetInt64(c, "clear_cache")
	if c.GetBool(Verify) {
		return
	}
	key := fmt.Sprintf("NEW_COURSES_LIST_%d", goodsId)
	if clearCache > 0 {
		redis.Expire(key, time.Second*1)
	}
	if list := redis.GetData(key); list != "" {
		json.Unmarshal([]byte(list.(string)), &resultList)
		p.ServerJSONSuccess(c, resultList)
	} else {
		courseService := priService.NewCourseService(setting.GinLogger(c))
		nt, err := courseService.GetCourseList(goodsId, 0, clearCache, "")
		data, _ := json.Marshal(nt)
		redis.SetData(key, data, time.Hour*12)
		if err != nil {
			p.ServerJSONError(c, err.Error(), error_code.SYSERR)
			return
		}
		p.ServerJSONSuccess(c, nt)
	}
	return
}
