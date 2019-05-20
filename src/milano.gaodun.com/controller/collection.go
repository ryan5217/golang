package controller

import (
	"github.com/gin-gonic/gin"
	form "milano.gaodun.com/model/collection"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	service "milano.gaodun.com/service/collection"
)

//用户反馈
var CollectionApi = NewCollectionApi()

func NewCollectionApi() *Collection {
	return &Collection{}
}

type Collection struct {
	Base
}

/**
添加问题调查
*/
func (C Collection) AddCollection(c *gin.Context) {
	AddCollection := form.Collection{}
	question := C.GetString(c, "question", true)
	context := C.GetString(c, "context", true)
	projectId := C.GetInt64(c, "project_id")
	uId := C.GetInt64(c, "student_id", true)
	partnerKey := C.GetString(c, "partner_key", true)
	partnerName := C.GetString(c, "partner_name")
	contextExtend := C.GetString(c, "context_extend")
	partnerId := C.GetInt64(c, "partner_id")
	Source := C.GetString(c, "source_from")
	if c.GetBool(Verify) {
		return
	}
	service := service.NewCollectionService(setting.GinLogger(c))
	AddCollection.PartnerId = partnerId
	AddCollection.Question = question
	AddCollection.ProjectId = projectId
	AddCollection.Uid = uId
	AddCollection.PartnerKey = partnerKey
	AddCollection.PartnerName = partnerName
	AddCollection.Context = context
	AddCollection.ContextExtend = contextExtend
	AddCollection.Source = Source
	id, code := service.AddCollection(&AddCollection)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	C.ServerJSONSuccess(c, id)
}

/**
获取是否已经提交
*/
func (C Collection) GetCollection(c *gin.Context) {
	selectCollection := form.Collection{}
	question := C.GetString(c, "question", true)
	projectId := C.GetInt64(c, "project_id")
	uId := C.GetInt64(c, "student_id", true)
	partnerKey := C.GetString(c, "partner_key", true)
	partnerId := C.GetInt64(c, "partner_id")
	source := C.GetString(c, "source_from")
	if c.GetBool(Verify) {
		return
	}
	service := service.NewCollectionService(setting.GinLogger(c))
	selectCollection.PartnerId = partnerId
	selectCollection.Question = question
	selectCollection.ProjectId = projectId
	selectCollection.Uid = uId
	selectCollection.PartnerKey = partnerKey
	selectCollection.Source = source
	list, code := service.GetCollection(&selectCollection)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	if list.Id > 0 {
		C.ServerJSONSuccess(c, list)
	} else {
		C.ServerJSONError(c, nil, error_code.NODATA)
	}
}

/**
获取列表
*/
func (C Collection) InfoList(c *gin.Context) {
	selectCollection := form.PriCollectionInfo{}
	questionId := C.GetInt64(c, "question_id")
	projectId := C.GetInt64(c, "project_id")
	uId := C.GetInt64(c, "student_id")
	partnerKey := C.GetString(c, "partner_key")
	partnerId := C.GetInt64(c, "partner_id")
	limit := C.GetInt(c, "limit")
	page := C.GetInt(c, "page")
	if c.GetBool(Verify) {
		return
	}
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	service := service.NewCollectionService(setting.GinLogger(c))
	selectCollection.PartnerId = partnerId
	selectCollection.QuestionId = questionId
	selectCollection.ProjectId = projectId
	selectCollection.Uid = uId
	selectCollection.PartnerKey = partnerKey
	list, code := service.InfoList(&selectCollection, limit, page)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	if len(list) > 0 {
		C.ServerJSONSuccess(c, list)
	} else {
		C.ServerJSONError(c, nil, error_code.NODATA)
	}
}

/**
获取列表
*/
func (C Collection) TypeList(c *gin.Context) {
	selectCollection := form.PriCollectionType{}
	name := C.GetString(c, "name")
	partnerKey := C.GetString(c, "partner_key")
	if c.GetBool(Verify) {
		return
	}
	service := service.NewCollectionService(setting.GinLogger(c))
	selectCollection.PartnerKey = partnerKey
	selectCollection.Name = name
	list, code := service.TypeList(&selectCollection)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	if len(list) > 0 {
		C.ServerJSONSuccess(c, list)
	} else {
		C.ServerJSONError(c, nil, error_code.NODATA)
	}
}

/**
获取列表
*/
func (C Collection) QuestionList(c *gin.Context) {
	selectCollection := form.PriCollectionQuestion{}
	service := service.NewCollectionService(setting.GinLogger(c))
	list, code := service.QuestionList(&selectCollection)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	if len(list) > 0 {
		C.ServerJSONSuccess(c, list)
	} else {
		C.ServerJSONError(c, nil, error_code.NODATA)
	}
}

/**
获取回答列表
*/
func (C Collection) GetInfoAndQuestion(c *gin.Context) {
	typeCollection := form.PriCollectionType{}
	selectCollection := form.CollectionList{}
	question := C.GetString(c, "question")
	questionId := C.GetInt64(c, "question_id")
	projectId := C.GetInt64(c, "project_id")
	uId := C.GetInt64(c, "student_id")
	partnerKey := C.GetString(c, "partner_key")
	partnerId := C.GetInt64(c, "partner_id")
	source := C.GetString(c, "source_from")
	pageNum := C.GetInt(c, "page_num")
	offset := C.GetInt(c, "offset")
	num := C.GetInt64(c, "num")
	top := C.GetInt64(c, "top")
	if c.GetBool(Verify) {
		return
	}
	service := service.NewCollectionService(setting.GinLogger(c))
	selectCollection.PartnerId = partnerId
	selectCollection.Question = question
	selectCollection.QuestionId = questionId
	selectCollection.ProjectId = projectId
	selectCollection.Uid = uId
	selectCollection.PartnerKey = partnerKey
	selectCollection.Source = source
	selectCollection.PageNum = pageNum
	selectCollection.Offset = offset
	list, total := service.InfoM.GetInfoAndQuestion(&selectCollection)
	typeList, code := service.TypeList(&typeCollection)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	typeArr := make(map[string]string)
	for _, v := range typeList {
		typeArr[v.PartnerKey] = v.Name
	}
	if len(list) > 0 {
		report := service.GetInfoReport(questionId, partnerKey, num, top)
		for k, v := range list {
			list[k].KeyName = typeArr[v.PartnerKey]
		}
		C.ServerJSONSuccess(c, map[string]interface{}{"list": list, "total": total, "report": report})
	} else {
		C.ServerJSONError(c, nil, error_code.NODATA)
	}
}

/**
添加类型
*/
func (C Collection) AddType(c *gin.Context) {
	var Addtype form.PriCollectionType
	name := C.GetString(c, "name", true)
	partnerKey := C.GetString(c, "partner_key", true)
	desc := C.GetString(c, "desc")
	if c.GetBool(Verify) {
		return
	}
	Addtype.Name = name
	Addtype.PartnerKey = partnerKey
	Addtype.Desc = desc
	service := service.NewCollectionService(setting.GinLogger(c))
	id, code := service.AddType(&Addtype)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	C.ServerJSONSuccess(c, id)
}

func (C Collection) EditType(c *gin.Context) {
	var Addtype form.PriCollectionType
	id := C.GetInt64(c, "id")
	name := C.GetString(c, "name", true)
	desc := C.GetString(c, "desc")
	if c.GetBool(Verify) {
		return
	}
	Addtype.Id = id
	Addtype.Name = name
	Addtype.Desc = desc
	service := service.NewCollectionService(setting.GinLogger(c))
	id, code := service.EditType(&Addtype)
	if code > 0 {
		C.ServerJSONError(c, nil, code)
		return
	}
	C.ServerJSONSuccess(c, id)
}
