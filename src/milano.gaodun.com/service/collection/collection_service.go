package collection

import (
	"fmt"
	"github.com/apex/log"
	"milano.gaodun.com/model/collection"
	"milano.gaodun.com/pkg/error-code"
	"strconv"
)

type CollectionServiceInterface interface {
	AddCollection(param *collection.Collection) (int64, int)
	GetCollection(param *collection.Collection) (collection.PriCollectionInfo, int)
	InfoList(param *collection.PriCollectionInfo, limit int, page int) ([]collection.PriCollectionInfo, int)
	TypeList(param *collection.PriCollectionType) ([]collection.PriCollectionType, int)
	GetInfoAndQuestion(param *collection.CollectionList) ([]collection.Collection, int64)
	AddType(param *collection.PriCollectionType) (int64, int)
	EditType(param *collection.PriCollectionType) (int64, int)
	GetInfoReport(questionId int64, partnerKey string, limit int64) []collection.ReportList
}

type CollectionService struct {
	InfoM     *collection.CollectionInfoModel
	TypeM     *collection.CollectionTypeModel
	QuestionM *collection.CollectionQuestionModel
	logger    *log.Entry
}

func NewCollectionService(logger *log.Entry) *CollectionService {
	return &CollectionService{
		InfoM:     collection.NewCollectionInfoModel(),
		TypeM:     collection.NewCollectionTypeModel(),
		QuestionM: collection.NewCollectionQuestionModelModel(),
		logger:    logger,
	}
}

//添加问卷调查
func (service CollectionService) AddCollection(param *collection.Collection) (int64, int) {
	var questionId int64
	var questionAdd collection.PriCollectionQuestion
	var info collection.PriCollectionInfo
	var infoList []collection.PriCollectionInfo
	result := service.QuestionM.GetByQuestion(param.Question)
	if result.Id == 0 {
		questionAdd.Question = param.Question
		questionAdd.ProjectId = param.ProjectId
		id, _ := service.QuestionM.Add(&questionAdd) //添加一个问题
		if id > 0 {
			questionId = id
		} else {
			return 0, error_code.DBERR
		}
	} else {
		questionId = result.Id
	}
	//反馈类型
	typeResult := service.TypeM.GetByPartnerKey(param.PartnerKey)
	if typeResult.Id == 0 {
		return 0, error_code.COLLECTION_TYPE_NOT_EXITS
	}
	info.Uid = param.Uid
	info.QuestionId = questionId
	info.ProjectId = param.ProjectId
	info.PartnerKey = param.PartnerKey
	info.PartnerName = param.PartnerName
	info.PartnerId = param.PartnerId
	info.Source = param.Source
	service.InfoM.WhereParams(info)
	service.InfoM.FindAll(&infoList)
	if len(infoList) > 0 {
		return 0, error_code.DATA_ADD_REPEAD
	}
	info.Context = param.Context
	info.ContextExtend = param.ContextExtend
	id, err := service.InfoM.Add(&info)
	if err != nil {
		return 0, error_code.DBERR
	}
	return id, 0
}

//查询问卷
func (service CollectionService) GetCollection(param *collection.Collection) (collection.PriCollectionInfo, int) {
	//反馈问题查询或添加
	var questionId int64
	var questionAdd collection.PriCollectionQuestion
	var info collection.PriCollectionInfo
	var infoList []collection.PriCollectionInfo
	result := service.QuestionM.GetByQuestion(param.Question)
	if result.Id == 0 {
		questionAdd.Question = param.Question
		questionAdd.ProjectId = param.ProjectId
		id, _ := service.QuestionM.Add(&questionAdd) //添加一个问题
		if id > 0 {
			questionId = id
		} else {
			return info, error_code.DBERR
		}
	} else {
		questionId = result.Id
	}
	//反馈类型
	typeResult := service.TypeM.GetByPartnerKey(param.PartnerKey)
	if typeResult.Id == 0 {
		return info, error_code.COLLECTION_TYPE_NOT_EXITS
	}
	info.Uid = param.Uid
	info.QuestionId = questionId
	info.ProjectId = param.ProjectId
	info.PartnerKey = param.PartnerKey
	info.PartnerId = param.PartnerId
	info.Source = param.Source
	service.InfoM.WhereParams(info)
	err := service.InfoM.FindAll(&infoList)
	if err != nil {
		return info, error_code.DBERR
	}
	if len(infoList) > 0 {
		return infoList[0], 0
	}
	return collection.PriCollectionInfo{}, 0
}

//获取列表
func (service CollectionService) InfoList(param *collection.PriCollectionInfo, limit int, page int) ([]collection.PriCollectionInfo, int) {
	var info collection.PriCollectionInfo
	var infoList []collection.PriCollectionInfo
	info.Uid = param.Uid
	info.QuestionId = param.QuestionId
	info.ProjectId = param.ProjectId
	info.PartnerKey = param.PartnerKey
	info.PartnerId = param.PartnerId
	service.InfoM.WhereParams(info)
	service.InfoM.Page(limit, page)
	err := service.InfoM.FindAll(&infoList)
	if err != nil {
		return infoList, error_code.DBERR
	}
	return infoList, 0
}

//获取列表
func (service CollectionService) TypeList(param *collection.PriCollectionType) ([]collection.PriCollectionType, int) {
	var info collection.PriCollectionType
	info.PartnerKey = param.PartnerKey
	info.Name = param.Name
	list := service.TypeM.GetList(&info)
	return list, 0
}

//问题列表
func (service CollectionService) QuestionList(param *collection.PriCollectionQuestion) ([]collection.PriCollectionQuestion, int) {
	var info collection.PriCollectionQuestion
	list := service.QuestionM.GetList(&info)
	return list, 0
}

//获取获取反馈信息
func (service CollectionService) GetInfoAndQuestion(param *collection.CollectionList) ([]collection.Collection, int64) {
	list, total := service.InfoM.GetInfoAndQuestion(param)
	return list, total
}

//添加反馈类型
func (service CollectionService) AddType(param *collection.PriCollectionType) (int64, int) {
	detail := service.TypeM.GetByPartnerKey(param.PartnerKey)
	if detail.Id > 0 {
		return 0, error_code.DATA_ADD_REPEAD
	}
	id, err := service.TypeM.Add(param)
	if err != nil {
		return 0, error_code.DBERR
	}
	return id, 0
}

//编辑反馈类型
func (service CollectionService) EditType(param *collection.PriCollectionType) (int64, int) {
	id, err := service.TypeM.Edit(param)
	if err != nil {
		return 0, error_code.DBERR
	}
	return id, 0
}

/**
获取报表
*/
func (service CollectionService) GetInfoReport(questionId int64, partnerKey string, limit int64, top int64) []collection.ReportList {
	var result []collection.ReportList
	var line collection.ReportList
	var total float64
	list := service.InfoM.GetInfoReport(questionId, partnerKey, limit, top)
	for _, v := range list {
		line.Context = v["context"]
		num, _ := strconv.ParseFloat(v["num"], 64)
		line.Num = num
		total += num
		result = append(result, line)
	}

	for k, v := range result {
		if total > 0 {
			rate := v.Num / total
			result[k].Rate = fmt.Sprintf("%.2f", rate) + "%"
		} else {
			result[k].Rate = "0.0%"
		}
	}
	return result
}
