package study_rights

import (
	"fmt"
	"github.com/apex/log"
	"milano.gaodun.com/model/api/goods"
	"milano.gaodun.com/model/api/tag"
	msM "milano.gaodun.com/model/members_student"
	rightM "milano.gaodun.com/model/study_rights"
	"milano.gaodun.com/pkg/utils"
	"strings"
)

var PaperTypeList = map[int64]string{
	1: "大纲试卷",
	2: "知识点",
	3: "每日一练",
	4: "智能组卷",
	5: "真题/专享",
}
var StatusList = map[int64]string{
	1: "开通",
	2: "禁用",
	3: "免费",
	4: "内容制作中",
}
var TagFidType = int64(119)

type StudyRightsServiceInterface interface {
	GetList(rightM.SearchParam) (*StudyRightsRes, error)
	GetOne(id int64) (*StudyRightsModuleRes, error)
	GetSubjectList(projectIdList []string) (map[int64][]int64, error)
}
type StudyRightsRes struct {
	GradationList []StudyRightsGradationRes
}
type StudyRightsGradationRes struct {
	Id         int64
	Name       string
	Sort       int64
	ModuleList []StudyRightsModuleRes
}
type StudyRightsModuleRes struct {
	Id          int64
	Sort        int64
	PaperType   Info
	Description string
	GoodsIds    string
	CourseIds   string
	ResourceId  int64
	Status      Info
	Gradation   Info
	Module      Info
	Tags        []tag.Tag
	Url         string
}
type Info struct {
	Id   int64
	Name string
}
type StudyRightsService struct {
	RightsMa       *rightM.StudyRightsManageModel
	RightsGra      *rightM.StudyRightsGradationModel
	RightsMo       *rightM.StudyRightsModulesModel
	MemberStudentM *msM.MembersStudentModel
	TagApi         *tag.TagApi
	GoodsApi       *goods.GoodsApi
	logger         *log.Entry
}

func NewStudyRightsService(logger *log.Entry) StudyRightsServiceInterface {
	h := utils.HttpHandle
	return &StudyRightsService{
		RightsMa:       rightM.NewStudyRightsManageModel(),
		RightsGra:      rightM.NewStudyRightsGradationModel(),
		RightsMo:       rightM.NewStudyRightsModulesModel(),
		MemberStudentM: msM.NewMembersStudentModel(),
		TagApi:         tag.NewTagApi(h),
		GoodsApi:       goods.NewGoodsApi(h),
		logger:         logger,
	}
}

//获取一个权限设置
func (g *StudyRightsService) GetOne(id int64) (*StudyRightsModuleRes, error) {

	rightMa, err := g.RightsMa.Get(id)
	if (err != nil || *rightMa == rightM.TkStudyRightsManage{}) {
		return &StudyRightsModuleRes{}, err
	}
	gradation, err := g.RightsGra.Get(rightMa.GradationId)
	module, err := g.RightsMo.Get(rightMa.ModuleId)
	if err != nil {
		return &StudyRightsModuleRes{}, err
	}
	tags, _ := g.TagApi.GetTagListByFid(TagFidType, id)
	sr := StudyRightsModuleRes{
		Id:   rightMa.Id,
		Sort: module.Sort,
		PaperType: Info{
			Id:   rightMa.PaperType,
			Name: PaperTypeList[rightMa.PaperType],
		},
		Description: rightMa.Description,
		GoodsIds:    rightMa.GoodsIds,
		ResourceId:  rightMa.ResourceId,
		Status: Info{
			Id:   rightMa.Status,
			Name: StatusList[rightMa.Status],
		},
		Gradation: Info{
			Id:   gradation.Id,
			Name: gradation.Name,
		},
		Module: Info{
			Id:   module.Id,
			Name: module.Name,
		},
		Tags: tags.Result["tag_list"],
	}
	if sr.Tags == nil {
		sr.Tags = []tag.Tag{}
	}
	return &sr, err
}
func (g *StudyRightsService) GetList(param rightM.SearchParam) (*StudyRightsRes, error) {
	sr := StudyRightsRes{GradationList: []StudyRightsGradationRes{}}
	//获取阶段列表
	gra, err := g.RightsGra.List(param)
	//获取模块列表
	module, err := g.RightsMo.List(param)
	//获取权限设置列表
	manage, err := g.RightsMa.List(param)
	gradationExist := map[int64]bool{}
	moduleList := map[int64]rightM.TkStudyRightsManage{}
	idStrings := ""
	goodsString := ""
	for _, m := range *manage {
		gradationExist[m.GradationId] = true
		moduleList[m.ModuleId] = m
		idStrings += fmt.Sprintf("%d,", m.Id)
		goodsString += m.GoodsIds + ","
	}
	idStrings = idStrings[0 : len(idStrings)-1]
	goodsString = goodsString[0 : len(goodsString)-1]
	goodsBuyList := map[string]goods.GoodsBuy{}
	if param.Uid > 0 {
		ms, _ := g.MemberStudentM.Get(param.Uid)
		goodBuyListResp, _ := g.GoodsApi.GetGoodsBuyList(ms.MemberId, goodsString)
		goodsBuyList = goodBuyListResp.Result
	}
	tagRes, _ := g.TagApi.GetTagList(TagFidType, idStrings)
	//根据阶段列表，模块列表把权限设置列表进行排序
	for _, g := range *gra {
		if gradationExist[g.Id] == true {
			graOne := StudyRightsGradationRes{
				Id:         g.Id,
				Name:       g.Name,
				Sort:       g.Sort,
				ModuleList: []StudyRightsModuleRes{},
			}
			for _, mo := range *module {
				if (moduleList[mo.Id] != rightM.TkStudyRightsManage{} && moduleList[mo.Id].GradationId == g.Id) {
					moduleOne := StudyRightsModuleRes{
						Id: moduleList[mo.Id].Id,
						Gradation: Info{
							Id:   graOne.Id,
							Name: graOne.Name,
						},
						Module:      Info{Id: mo.Id, Name: mo.Name},
						Sort:        mo.Sort,
						PaperType:   Info{Id: moduleList[mo.Id].PaperType, Name: PaperTypeList[moduleList[mo.Id].PaperType]},
						Description: moduleList[mo.Id].Description,
						GoodsIds:    moduleList[mo.Id].GoodsIds,
						CourseIds:    moduleList[mo.Id].CourseIds,
						ResourceId:  moduleList[mo.Id].ResourceId,
						Status:      Info{Id: moduleList[mo.Id].Status, Name: StatusList[moduleList[mo.Id].Status]},
						Tags:        tagRes.Result["tag_list"][fmt.Sprintf("%d", moduleList[mo.Id].Id)].Tags,
					}
					if tagRes.Result["tag_list"][fmt.Sprintf("%d", moduleList[mo.Id].Id)].Tags == nil {
						moduleOne.Tags = []tag.Tag{}
					}
					//免费小图片
					if moduleList[mo.Id].Status == 3 {
						moduleOne.Url = "http://simg01.gaodunwangxiao.com/uploadimgs/tmp/upload/201811/30/86d18_20181130135805.png"
					}
					//必刷小图片
					if moduleList[mo.Id].Status == 1 {
						goodsL := strings.Split(moduleList[mo.Id].GoodsIds, ",")
						for _, v := range goodsL {
							if goodsBuyList[v].BuyStatus {
								moduleOne.Url = "http://simg01.gaodunwangxiao.com/uploadimgs/tmp/upload/201812/05/8ab2f_20181205155157.png"
							}
						}
					}
					graOne.ModuleList = append(graOne.ModuleList, moduleOne)
				}
			}
			sr.GradationList = append(sr.GradationList, graOne)
		}
	}
	return &sr, err
}
func (g *StudyRightsService) GetSubjectList(projectIdList []string) (map[int64][]int64, error) {
	return g.RightsMa.SubjectList(projectIdList)
}
