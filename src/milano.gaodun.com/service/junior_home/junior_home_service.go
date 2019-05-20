package banner

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/imroc/req"
	"math"
	"milano.gaodun.com/conf"
	atApi "milano.gaodun.com/model/api/activity"
	adApi "milano.gaodun.com/model/api/adverts"
	icApi "milano.gaodun.com/model/api/item_category"
	"milano.gaodun.com/model/api/primary"
	gdApi "milano.gaodun.com/model/api/primary"
	smApi "milano.gaodun.com/model/api/sales_man"
	tlApi "milano.gaodun.com/model/api/tree_list"
	cm "milano.gaodun.com/model/course"
	ddiM "milano.gaodun.com/model/doone_data_item"
	msM "milano.gaodun.com/model/members_student"
	examM "milano.gaodun.com/model/project_exam_time"
	yucM "milano.gaodun.com/model/user_change"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
	"strconv"
	"strings"
	"time"
)

type JuniorHomeServiceInterface interface {
	GetHome(param *Param) (*JuniorHomeResp, int)
	GetSalesMan(uid int64, clearCache int64, isVip bool, needVerifyVip bool) (*SalesMan, int)
	GetHomeStudyInfo(param *Param) (*ExperienceResp, int)
	GetTreeListMap(key string, clearCache int64) (map[string]*simpleJson.Json, int)
	GetMarketingInfo(param *MarketingParam) (*LiveInfo, int)
}

type JuniorHomeService struct {
	ica    *icApi.ItemCategoryApi
	act    *atApi.ActivityApi
	ddi    *ddiM.DooneDataItemModel
	tla    *tlApi.TreeListApi
	ada    *adApi.AdvertsApi
	gda    *gdApi.GoodsApi
	yuc    *yucM.UserChangeModel
	toc    *primary.TocApi
	cm     *cm.CourseModel
	capi   *primary.CourseApi
	exm    *examM.ProjectExamTimeModel
	smapi  *smApi.SalesManApi
	msm    *msM.MembersStudentModel
	logger *log.Entry
}

func NewJuniorHomeService(logger *log.Entry) JuniorHomeServiceInterface {
	h := utils.HttpHandle
	return &JuniorHomeService{
		ica:    icApi.NewItemCategoryApi(h),
		act:    atApi.NewActivityApi(h),
		tla:    tlApi.NewTreeListApi(h),
		ada:    adApi.NewAdvertsApi(h),
		gda:    gdApi.NewGoodsApi(h),
		yuc:    yucM.NewUserChangeModel(),
		ddi:    ddiM.NewDooneDataItemModel(),
		toc:    primary.NewTocApi(h),
		cm:     cm.NewCourseModel(),
		capi:   primary.NewCourseApi(h),
		exm:    examM.NewProjectExamTimeModel(),
		smapi:  smApi.NewSalesManApi(h),
		msm:    msM.NewMembersStudentModel(),
		logger: logger,
	}
}

type JuniorHomeResp struct {
	Message  string        `json:"message"`
	ExamDate string        `json:"exam_date"`
	LeftDays int64         `json:"left_days"`
	ExamInfo string        `json:"exam_info"`
	List     []interface{} `json:"list"`
}
type JuniorHome struct {
	StudyInfo StudyInfo `json:"study_info"`
	LiveInfo  LiveInfo  `json:"live_info"`
	Chapter   Chapter   `json:"chapter"`
	Advert    Advert    `json:"advert"`
}
type Advert struct {
	HaveData   bool   `json:"have_data"`
	Type       string `json:"type"`
	ModuleName string `json:"module_name"`
	Route      string `json:"route"`
	Img        string `json:"img"`
}
type Chapter struct {
	HaveData   bool        `json:"have_data"`
	Type       string      `json:"type"`
	ModuleName string      `json:"module_name"`
	MoreInfo   string      `json:"more_info"`
	Route      string      `json:"route"`
	Data       interface{} `json:"data"`
}
type LiveInfo struct {
	HaveData   bool     `json:"have_data"`
	IsBuy      bool     `json:"is_buy"`
	Type       string   `json:"type"`
	ModuleName string   `json:"module_name"`
	Route      string   `json:"route"`
	More       string   `json:"more"`
	Activity   Activity `json:"activity"`
}
type BookInfo struct {
	HaveData   bool   `json:"have_data"`
	Type       string `json:"type"`
	ModuleName string `json:"module_name"`
	Goods      Goods  `json:"book"`
}
type Activity struct {
	Id              int64  `json:"id"`
	Type            string `json:"type"`
	TypeImg         string `json:"type_img"`
	Name            string `json:"name"`
	Img             string `json:"img"`
	Avatar          string `json:"avatar"`
	StudentNum      int64  `json:"student_num"`
	StartTime       int64  `json:"start_time"`
	EndTime         int64  `json:"end_time"`
	ActivityEndTime int64  `json:"activity_end_time"`
	Gif             string `json:"gif"`
	ButtonInfo      string `json:"button_info"`
	ButtonColor     string `json:"button_color"`
	Route           string `json:"route"`
}
type StudyInfo struct {
	HaveData   bool     `json:"have_data"`
	IsFree     bool     `json:"is_free"`
	ModuleName string   `json:"module_name"`
	MoreInfo   string   `json:"more_info"`
	MoreRoute  string   `json:"route"`
	Type       string   `json:"type"`
	GoodsList  []Goods  `json:"goods_list"`
	SalesMan   SalesMan `json:"sales_man"`
}
type SalesMan struct {
	Id            int64    `json:"id"`
	Avatar        string   `json:"avatar"`
	Name          string   `json:"name"`
	QRCode        string   `json:"qr_code"`
	MsMsgCode     string   `json:"ms_msg_code"`
	IsVip         bool     `json:"is_vip"`
	Info          string   `json:"info"`
	JoinInfo      string   `json:"join_info"`
	ButtonTitle   string   `json:"button_title"`
	Route         string   `json:"route"`
	StudentNum    int64    `json:"student_num"`
	GroupInfoList []string `json:"group_info_list"`
}
type Goods struct {
	GoodsId string `json:"goods_id"`
	Img     string `json:"img"`
	Route   string `json:"route"`
}
type Param struct {
	IsAudit     int64
	ClearCache  int64
	ProjectId   int64
	SubjectId   int64
	StudentId   int64
	YxStudentId int64
	Source      int64
	Version     string
}
type ExperienceResp struct {
	Cover           string        `json:"cover"`
	Title           string        `json:"title"`
	Banner          string        `json:"banner"`
	GoodsCourseList []GoodsCourse `json:"goods_course_list"`
}
type GoodsCourse struct {
	GoodsId       int64    `json:"goods_id"`
	IsBuy         bool     `json:"is_buy"`
	Info          string   `json:"info"`
	Img           string   `json:"img"`
	Name          string   `json:"name"`
	StudentNum    int64    `json:"student_num"`
	VCourseId     int64    `json:"v_course_id"`
	CourseList    []Course `json:"course_list"`
	AllChapterNum int64    `json:"all_chapter_num"`
	Remark        string   `json:"remark"`
	Route         string   `json:"route"`
}
type Course struct {
	CourseId         int64            `json:"course_id"`
	Name             string           `json:"name"`
	AllChapterNum    int64            `json:"all_chapter_num"`
	CourseProgress   float32          `json:"course_progress"`
	SyllabusItemList [][]SyllabusItem `json:"syllabus_item_list"`
}
type SyllabusItem struct {
	Id         int64       `json:"id"`
	Name       string      `json:"name"`
	Depth      string      `json:"depth"`
	ParentId   string      `json:"parent_id"`
	ResourceId string      `json:"resource_id"`
	Resource   interface{} `json:"resource"`
	Progress   int64       `json:"progress"`
	Children   interface{} `json:"children"`
}

type MarketingParam struct {
	ClearCache int64
	StudentId  int64
	ClassKey   string //分类key
	PublicKey  string //公开课key
}

func (g *JuniorHomeService) GetTreeListMap(key string, clearCache int64) (map[string]*simpleJson.Json, int) {
	result := error_code.SUCCESSSTATUS
	//获取首页课程设置
	g.logger.Infof("GetTreeListMap_start")
	tla, result := g.tla.GetTreeListBySign(key, clearCache)
	g.logger.Infof("GetTreeListMap_end")
	if result != error_code.SUCCESSSTATUS {
		return map[string]*simpleJson.Json{}, result
	}
	tlJson, _ := simpleJson.NewJson([]byte(tla))
	list := tlJson.Get("result").Get("list")
	listJson := map[string]*simpleJson.Json{}
	for i := 0; i < len(list.MustArray()); i++ {
		key, _ := list.GetIndex(i).Get("sign").String()
		listJson[key] = list.GetIndex(i)
	}
	return listJson, result
}
func (g *JuniorHomeService) GetHomeStudyInfo(param *Param) (*ExperienceResp, int) {
	result := error_code.SUCCESSSTATUS
	resp := ExperienceResp{}
	g.logger.Infof("GetHomeStudyInfo_start%d", param.StudentId)
	//获取首页课程设置
	if result != error_code.SUCCESSSTATUS {
		return &resp, result
	}
	listJson, result := g.GetTreeListMap("junior_study", param.ClearCache)
	if result != error_code.SUCCESSSTATUS {
		return &resp, result
	}
	_, ok := listJson["free_course"]
	_, ok1 := listJson["light_course"]
	g.logger.Infof("GetHomeStudyInfo_GetGoodsCourse_start%d", param.StudentId)
	if ok && ok1 {
		//根据商品id获取v课程信息
		freeGoods, code := g.GetGoodsCourse(*param, listJson["free_course"], 1)
		g.logger.Infof("GetHomeStudyInfo_GetGoodsCourse_freeGoods_end_%d", param.StudentId)
		if code != error_code.SUCCESSSTATUS {
			return &resp, code
		}
		lightGoods, code := g.GetGoodsCourse(*param, listJson["light_course"], 0)
		g.logger.Infof("GetHomeStudyInfo_GetGoodsCourse_lightGoods_end_%d", param.StudentId)
		if code != error_code.SUCCESSSTATUS {
			return &resp, code
		}
		paidGoods, code := g.GetGoodsCourse(*param, listJson["paid_course"], 0)
		g.logger.Infof("GetHomeStudyInfo_GetGoodsCourse_paidGoods_end_%d", param.StudentId)
		if code != error_code.SUCCESSSTATUS {
			return &resp, code
		}
		goodsIds := []string{}
		goodsIds = append(goodsIds, fmt.Sprintf("%d", freeGoods.GoodsId))
		goodsIds = append(goodsIds, fmt.Sprintf("%d", lightGoods.GoodsId))
		goodsIds = append(goodsIds, fmt.Sprintf("%d", paidGoods.GoodsId))
		//获取商品是否购买信息
		goodsBuyList, code := g.gda.GetGoodsBuyList(goodsIds, param.StudentId)
		g.logger.Infof("GetHomeStudyInfo_GetGoodsBuyList_end_%d", param.StudentId)
		if code != error_code.SUCCESSSTATUS {
			return &resp, code
		}
		fg, ok := goodsBuyList[fmt.Sprintf("%d", freeGoods.GoodsId)]
		if ok {
			freeGoods.IsBuy = fg.BuyStatus
		}
		lg, ok := goodsBuyList[fmt.Sprintf("%d", lightGoods.GoodsId)]
		if ok {
			lightGoods.IsBuy = lg.BuyStatus
		}
		pg, ok := goodsBuyList[fmt.Sprintf("%d", paidGoods.GoodsId)]
		if ok {
			paidGoods.IsBuy = pg.BuyStatus
		}
		if utils.VersionCompare("5.0.2", param.Version) {
			freeGoods.Route = fmt.Sprintf("tkofkjzc://study/course/list?goods_id=%d", freeGoods.GoodsId)
			lightGoods.Route = fmt.Sprintf("tkofkjzc://study/course/list?goods_id=%d", lightGoods.GoodsId)
			paidGoods.Route = fmt.Sprintf("tkofkjzc://study/course/list?goods_id=%d", paidGoods.GoodsId)
		} else {
			freeGoods.Route = fmt.Sprintf("tkofkjzc://common/h5/index?url="+conf.MTIKU_DOMAIN+"/pj/ke-list?goods_id=%d", freeGoods.GoodsId)
			lightGoods.Route = fmt.Sprintf("tkofkjzc://common/h5/index?url="+conf.MTIKU_DOMAIN+"/pj/ke-list?goods_id=%d", lightGoods.GoodsId)
			paidGoods.Route = fmt.Sprintf("tkofkjzc://common/h5/index?url="+conf.MTIKU_DOMAIN+"/pj/ke-list?goods_id=%d", paidGoods.GoodsId)
		}
		resp.GoodsCourseList = append(resp.GoodsCourseList, freeGoods)
		if param.IsAudit != 1 {
			resp.GoodsCourseList = append(resp.GoodsCourseList, lightGoods)
			resp.GoodsCourseList = append(resp.GoodsCourseList, paidGoods)
		}
		experience, ok := listJson["experience_course"]
		resp.Title = "零基础口碑班体验课"
		if ok {
			resp.Title, _ = experience.Get("name").String()
			resp.Cover, _ = experience.Get("data").GetIndex(0).Get("desc").String()
			resp.Banner, _ = experience.Get("data").GetIndex(1).Get("desc").String()
		}
		if resp.Cover == "" {
			resp.Cover = "http://simg01.gaodunwangxiao.com/uploadfiles/tmp/upload/201903/08/1e3aa_20190308101044.png"
		}
		if resp.Banner == "" {
			resp.Banner = "http://simg01.gaodunwangxiao.com/uploadfiles/tmp/upload/201903/12/b0586_20190312145825.png"
		}
	} else {
		return &resp, error_code.GOODSCONFIGERR
	}
	return &resp, result
}
func (g *JuniorHomeService) GetGoodsCourse(param Param, json *simpleJson.Json, needSyllabus int64) (GoodsCourse, int) {
	goodsCourseBytes, _ := json.Get("data").GetIndex(0).Get("json_detail").Bytes()
	goodsCourseJson, _ := simpleJson.NewJson(goodsCourseBytes)
	goodsIdStr, _ := goodsCourseJson.Get("goods_id").String()
	//获取商品详情字符串
	g.logger.Infof("GetGoodsCourse_GetGoodsById_start_%d_goodsIdStr_%s", param.StudentId, goodsIdStr)
	goodsDetailStr, err1 := g.gda.GetGoodsById(utils.String2Int64(goodsIdStr), 0, param.ClearCache, "")
	g.logger.Infof("GetGoodsCourse_GetGoodsById_end_%d_goodsIdStr_%s", param.StudentId, goodsIdStr)
	if err1 != nil {
		return GoodsCourse{}, error_code.GOODSDETAILAPIERR
	}
	//字符串转json
	goodsDetailJson, _ := simpleJson.NewJson([]byte(goodsDetailStr))
	//获取v课程id
	goodsVcourseId, _ := goodsDetailJson.Get("result").Get("info").Get("product_info").GetIndex(0).Get("vid").Int64()
	//获取v课程信息
	goodsVcourse, _ := g.cm.Get(goodsVcourseId)
	g.logger.Infof("GetGoodsCourse_GetGoodsVcourse_end_goodsVcourseId_%d", goodsVcourseId)
	//组装课程列表
	goodsCourse := GoodsCourse{AllChapterNum: 0}
	goodsCourse.VCourseId = goodsVcourseId
	goodsCourse.Name, _ = goodsDetailJson.Get("result").Get("info").Get("title").String()
	buyNum, _ := goodsDetailJson.Get("result").Get("info").Get("buy_num").Int64()
	buyNumFalse, _ := goodsDetailJson.Get("result").Get("info").Get("buy_num_false").Int64()
	goodsCourse.StudentNum = buyNum + buyNumFalse
	goodsCourse.GoodsId = utils.String2Int64(goodsIdStr)
	goodsCourse.Img, _ = goodsDetailJson.Get("result").Get("info").Get("horizontal_img").String()
	goodsCourse.Remark, _ = goodsDetailJson.Get("result").Get("info").Get("remark").String()
	remarkJson, err := simpleJson.NewJson([]byte(goodsCourse.Remark))
	if err != nil {
		g.logger.Infof("GetGoodsCourse_remarkJson_err_%s", err.Error())
	}
	goodsCourse.Info, _ = remarkJson.Get("experience_info").String()

	var vCourseList []cm.GdCourse
	courseList := []Course{}
	if goodsVcourse.Isbig > 0 {
		courseIds := goodsVcourse.Courses
		g.logger.Infof("GetGoodsCourse_GetCourseByIds_start_%s", courseIds)
		vList, _ := g.cm.GetCourseByIds(courseIds)
		g.logger.Infof("GetGoodsCourse_GetCourseByIds_end_%s", courseIds)
		vCourseList = *vList
	} else {
		vCourseList = append(vCourseList, *goodsVcourse)
	}
	for _, v := range vCourseList {
		course := Course{}
		course.CourseId = v.RelationCourse
		course.Name = v.Name
		courseList = append(courseList, course)
	}
	courseListResp := map[string]Course{}
	if needSyllabus != 1 {
		goodsCourse.CourseList = courseList
		return goodsCourse, error_code.SUCCESSSTATUS
	}
	courseSyllabus := map[string][]*simpleJson.Json{}
	client := utils.RedisHandle
	for _, v := range vCourseList {
		//获取大纲阶段字符串
		var freeGradationStr string
		gradationKey := fmt.Sprintf("SaasCourseSyllabusCourseId_%d", v.RelationCourse)
		g.logger.Infof("GetGoodsCourse_GetCache_start_%s", gradationKey)
		freeGradationStr = client.GetData(gradationKey).(string)
		g.logger.Infof("GetGoodsCourse_GetCache_end_%s", freeGradationStr)
		if freeGradationStr == "" || param.ClearCache == 1 {
			g.logger.Infof("GetGoodsCourse_GetSyllabusList_start_saasCourseId_%d", v.RelationCourse)
			freeGradationStr, err1 = g.toc.GetSyllabusList(v.RelationCourse)
			g.logger.Infof("GetGoodsCourse_GetSyllabusList_end_saasCourseId_%d", v.RelationCourse)
			if err1 != nil {
				return GoodsCourse{}, error_code.TOCAPIERR
			}
			client.SetData(gradationKey, freeGradationStr, time.Second*86400)
		}
		gradationJson, _ := simpleJson.NewJson([]byte(freeGradationStr))
		//课程大纲信息转map
		for i := 0; i < len(gradationJson.MustArray()); i++ {
			syllabusId, _ := gradationJson.GetIndex(i).Get("syllabus_id").String()
			courseId, _ := gradationJson.GetIndex(i).Get("course_id").String()
			var syllabusStr string
			syllabusStatusKey := fmt.Sprintf("SaasCourseSyllabusStudy_StudentId_%d_CourseId_%d_SyllabusId_%s", param.StudentId, v.RelationCourse, syllabusId)
			g.logger.Infof("GetGoodsCourse_GetCache_start_%s", syllabusStatusKey)
			syllabusStr = client.GetData(syllabusStatusKey).(string)
			g.logger.Infof("GetGoodsCourse_GetCache_end_%s", syllabusStr)
			if syllabusStr == "" || param.ClearCache == 1 {
				g.logger.Infof("GetGoodsCourse_GetSyllabusStudyStatus_start_%d_%s_%s", param.StudentId, courseId, syllabusId)
				syllabusStr, _ = g.capi.GetSyllabusStudyStatus(param.StudentId, utils.String2Int64(courseId), syllabusId)
				g.logger.Infof("GetGoodsCourse_GetSyllabusStudyStatus_end_%s", syllabusStr)
				client.SetData(syllabusStatusKey, syllabusStr, time.Second*1200)
			}
			syllabusJson, _ := simpleJson.NewJson([]byte(syllabusStr))
			syllabus := syllabusJson.Get("result")
			courseSyllabus[courseId] = append(courseSyllabus[courseId], syllabus)
		}
	}
	//获取课程学习进度
	reqParam := req.Param{}
	for i := 0; i < len(vCourseList); i++ {
		reqParam[fmt.Sprintf("course_id_list[%d]", i)] = vCourseList[i].RelationCourse
	}
	g.logger.Infof("GetGoodsCourse_GetCourseStudyProgress_start_%d", param.StudentId)
	progress, err := g.capi.GetCourseStudyProgress(param.StudentId, reqParam)
	g.logger.Infof("GetGoodsCourse_GetCourseStudyProgress_end")
	if err != nil {
		return GoodsCourse{}, error_code.COURSEPROGRESSAPIERR
	}
	progressMap := map[int64]primary.StudyProgressInfo{}
	for _, v := range courseList {
		progressMap[v.CourseId] = primary.StudyProgressInfo{CourseId: fmt.Sprintf("%d", v.CourseId)}
	}
	for _, v := range progress.Result {
		progressMap[utils.String2Int64(v.CourseId)] = v
	}
	//组装课程进度，大纲信息
	courseMap := make(map[string][][]SyllabusItem)
	for courseId, val := range courseSyllabus {
		for _, v := range val {
			syllabusItemList, code := g.GetSyllabusByTree(v, 1)
			courseMap[courseId] = append(courseMap[courseId], syllabusItemList)
			if code != error_code.SUCCESSSTATUS {
				return GoodsCourse{}, code
			}
		}
	}
	for _, c := range courseList {
		v, _ := progressMap[c.CourseId]
		if c.CourseId == utils.String2Int64(v.CourseId) {
			c.AllChapterNum = v.AllChapter
			c.CourseProgress = v.Progress
			_, ok := courseMap[v.CourseId]
			if ok {
				for _, v := range courseMap[v.CourseId] {
					c.SyllabusItemList = append(c.SyllabusItemList, v)
				}
			}
			courseListResp[v.CourseId] = c
		}
	}
	for _, v := range courseListResp {
		goodsCourse.CourseList = append(goodsCourse.CourseList, v)
		goodsCourse.AllChapterNum += v.AllChapterNum
	}
	return goodsCourse, error_code.SUCCESSSTATUS
}
func (g *JuniorHomeService) GetSyllabusByTree(tree *simpleJson.Json, depth int64) ([]SyllabusItem, int) {
	res := []SyllabusItem{}
	for i := 0; i < len(tree.MustArray()); i++ {
		item := SyllabusItem{}
		one := tree.GetIndex(i)
		item.Id, _ = one.Get("id").Int64()
		item.Name, _ = one.Get("name").String()
		item.Depth = fmt.Sprintf("%d", depth)
		item.ParentId, _ = one.Get("parent_id").String()
		item.ResourceId, _ = one.Get("resource_id").String()
		item.Progress, _ = one.Get("progress").Int64()
		item.Resource = one.Get("resource")

		children := tree.GetIndex(i).Get("children")
		if len(children.MustArray()) > 0 {
			dep := depth + 1
			childList, code := g.GetSyllabusByTree(children, dep)
			if code != error_code.SUCCESSSTATUS {
				return res, code
			}
			if depth == 1 {
				item.Children = childList
				res = append(res, item)
			} else {
				res = append(res, item)
				for _, v := range childList {
					v.Depth = "2"
					res = append(res, v)
				}
			}
		} else {
			res = append(res, item)
		}
	}
	return res, error_code.SUCCESSSTATUS
}
func (g *JuniorHomeService) GetHome(param *Param) (*JuniorHomeResp, int) {
	result := error_code.SUCCESSSTATUS
	res := JuniorHome{}
	resp := JuniorHomeResp{}
	//获取首页课程设置
	g.logger.Infof("JuniorGetHome_start_%d", param.StudentId)
	listJson, result := g.GetTreeListMap("junior_study", param.ClearCache)
	g.logger.Infof("JuniorGetHome_GetTreeListBySign_end_%d", param.StudentId)
	if result != error_code.SUCCESSSTATUS {
		return &resp, result
	}
	var code int
	//获取首页分章做真题知识点
	_, ok := listJson["chapter_study"]
	_, ok1 := listJson["tiku_more"]
	if ok {
		res.Chapter.Data, code = g.GetKnowledge(param, listJson["chapter_study"].Get("data"))
		g.logger.Infof("JuniorGetHome_GetKnowledge_end_%d", param.StudentId)
		res.Chapter.HaveData = true
		res.Chapter.Type = "Chapter"
		res.Chapter.ModuleName, _ = listJson["chapter_study"].Get("name").String()
		if ok1 {
			res.Chapter.MoreInfo, _ = listJson["tiku_more"].Get("name").String()
		} else {
			res.Chapter.MoreInfo = "更多题库模块"
		}
		res.Chapter.Route = "tkofkjzc://study/home/index?page=tiku"
		if code != error_code.SUCCESSSTATUS {
			res.Chapter.HaveData = false
			resp.Message += error_code.INFO[code] + ";"
		}
	}
	//查询营销学生id
	yuc, code := g.yuc.Get(param.StudentId)
	g.logger.Infof("JuniorGetHome_UserChangeModel_end_%d", param.StudentId)
	if code != error_code.SUCCESSSTATUS {
		return &resp, code
	}
	param.YxStudentId = yuc.YxStudentId
	//获取首页课程学习
	_, ok = listJson["free_course"]
	_, ok1 = listJson["light_course"]
	_, ok2 := listJson["paid_course"]
	_, ok3 := listJson["lecture_note"]
	_, ok4 := listJson["tiku"]
	_, ok5 := listJson["sales_man"]
	_, ok6 := listJson["study_zone"]
	_, ok7 := listJson["study_more"]
	if ok && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 {
		g.logger.Infof("JuniorGetHome_GetGoodsList_start_%d", param.StudentId)
		res.StudyInfo, code = g.GetGoodsList(param, listJson)
		g.logger.Infof("JuniorGetHome_GetGoodsList_end_%d", param.StudentId)
		res.StudyInfo.HaveData = true
		res.StudyInfo.Type = "StudyInfo"
		res.StudyInfo.ModuleName, _ = listJson["study_zone"].Get("name").String()
		if ok7 {
			res.StudyInfo.MoreInfo, _ = listJson["study_more"].Get("name").String()
		} else {
			res.StudyInfo.MoreInfo = "查看全部"
		}
		res.StudyInfo.MoreRoute = "tkofkjzc://study/home/index?page=task"
		if code != error_code.SUCCESSSTATUS {
			res.StudyInfo.HaveData = false
			resp.Message += error_code.INFO[code] + ";"
		}
	}
	//获取首页公开课
	_, ok = listJson["public_course"]
	if ok && !utils.VersionCompare("5.0.2", param.Version) {
		g.logger.Infof("JuniorGetHome_GetLiveInfo_start_%d", param.StudentId)
		res.LiveInfo, code = g.GetLiveInfo(param, listJson["public_course"])
		g.logger.Infof("JuniorGetHome_GetLiveInfo_start_%d", param.StudentId)
		res.LiveInfo.HaveData = true
		res.LiveInfo.Type = "PublicCourse"
		if code != error_code.SUCCESSSTATUS {
			res.LiveInfo.HaveData = false
			resp.Message += error_code.INFO[code] + ";"
		}
	} else if utils.VersionCompare("5.0.2", param.Version) {
		res.LiveInfo.Type = "LiveInfo"
		lc, ok := listJson["live_course"]
		res.LiveInfo.ModuleName, _ = listJson["public_course"].Get("name").String()
		if ok {
			remarkByte, _ := lc.Get("remark").Bytes()
			remarkJson, _ := simpleJson.NewJson(remarkByte)
			res.LiveInfo.Route, _ = remarkJson.Get("route_url").String()
			res.LiveInfo.More, _ = remarkJson.Get("call_info").String()
			haveData, _ := remarkJson.Get("have_data").String()
			if haveData != ""{
				res.LiveInfo.HaveData = true
			}
		}
	}
	g.logger.Infof("JuniorGetHome_GetExamDateInfo_start_%d", param.StudentId)
	exam, code := g.exm.Get(param.ProjectId)
	g.logger.Infof("JuniorGetHome_GetExamDateInfo_end_%d", param.StudentId)
	if code != error_code.SUCCESSSTATUS {
		return &resp, code
	}
	if exam.ExamTime != 0 {
		resp.ExamDate = time.Unix(exam.ExamTime, 0).Format("2006-01-02")
		theTime, _ := time.Parse("2006-01-02", resp.ExamDate)
		unix := theTime.Unix()
		nowUnix := time.Now().Unix()
		if unix > nowUnix {
			resp.LeftDays = (unix-nowUnix)/86400 + 1
			resp.ExamInfo = fmt.Sprintf("距离考试仅剩%d天，坚持就是胜利！", resp.LeftDays)
		}
	}
	//获取学习计划广告位
	if !utils.VersionCompare("5.0.2", param.Version) {
		g.logger.Infof("JuniorGetHome_GetAdvertsList_start_%d", param.StudentId)
		advert, code := g.ada.GetAdvertsList("31-", param.ProjectId, param.ClearCache)
		g.logger.Infof("JuniorGetHome_GetAdvertsList_start_%d", param.StudentId)
		res.Advert.Type = "Advert"
		res.Advert.ModuleName = "初级会计学习计划"
		if code != error_code.SUCCESSSTATUS {
			res.Advert.HaveData = false
			resp.Message += "获取学习计划广告位失败;"
		} else {
			index := len(advert.MustArray()) - 1
			res.Advert.HaveData = true
			url, _ := advert.GetIndex(index).Get("url").String()
			res.Advert.Route = "tkofkjzc://common/h5/index?url=" + url
			res.Advert.Img, _ = advert.GetIndex(index).Get("pic_url").String()
		}
		if param.IsAudit != 1 {
			resp.List = append(resp.List, res.Advert)
		}
	}
	//排序结果
	resp.List = append(resp.List, res.StudyInfo)
	if param.IsAudit != 1 {
		resp.List = append(resp.List, res.LiveInfo)
	}
	resp.List = append(resp.List, res.Chapter)
	if param.IsAudit == 1 {
		recommendBook, ok := listJson["recommend_book"]
		if ok {
			data := recommendBook.Get("data")
			if len(data.MustArray()) > 0 {
				for i := 0; i < len(data.MustArray()); i++ {
					subjectId, _ := data.GetIndex(i).Get("service_id").Int64()
					if subjectId == param.SubjectId {
						jsonDetailByte, _ := data.GetIndex(i).Get("json_detail").Bytes()
						goods := Goods{}
						json.Unmarshal(jsonDetailByte, &goods)
						BookInfo := BookInfo{}
						if goods.GoodsId != "" {
							BookInfo.HaveData = true
						}
						BookInfo.ModuleName, _ = recommendBook.Get("name").String()
						BookInfo.Type = "Book"
						BookInfo.Goods = goods
						resp.List = append(resp.List, BookInfo)
					}
				}
			}
		}
	}
	g.logger.Infof("JuniorGetHome_end_%d", param.StudentId)
	return &resp, result
}
func (g *JuniorHomeService) GetLiveInfo(param *Param, json *simpleJson.Json) (LiveInfo, int) {
	Live := LiveInfo{}
	var gdJson *simpleJson.Json
	goodsId := ""
	code := error_code.SUCCESSSTATUS
	data := json.Get("data")
	indexMax := len(data.MustArray()) - 1
	//indexMax := 0
	if indexMax < 0 {
		Live.HaveData = false
		code = error_code.PUBLICCOURSEERR
		return Live, code
	}
	activityInfoJson, _ := json.Get("data").GetIndex(indexMax).Get("json_detail").String()
	activityJsonInfo, _ := simpleJson.NewJson([]byte(activityInfoJson))
	activityInfo := json.Get("data").GetIndex(indexMax)
	serviceType, _ := data.GetIndex(indexMax).Get("service_type").String()
	serviceId, _ := data.GetIndex(indexMax).Get("service_id").Int64()
	price := "0"
	Live.ModuleName, _ = json.Get("name").String()
	routeInfo, _ := activityJsonInfo.Get("Route").String()
	Live.Activity.Route = fmt.Sprintf(routeInfo, serviceType, serviceId)
	g.logger.Infof("JuniorGetHome_GetGoodsDetail_start_%d", param.StudentId)
	if serviceType == "seckill" {
		Live.Activity.Type = "seckill"
		Live.Activity.ActivityEndTime, _ = activityInfo.Get("data").Get("end_date").Int64()
		Live.Activity.Name, _ = activityInfo.Get("data").Get("title").String()
		Live.Activity.TypeImg, _ = activityJsonInfo.Get("SeckillImg").String()
		price, _ = activityInfo.Get("data").Get("seckill_price").String()
		gid, _ := activityInfo.Get("data").Get("goods_id").Int64()
		goodsId = fmt.Sprintf("%d", gid)
		gdJson, code = g.gda.GetGoodsDetail(goodsId, param.ClearCache)
		if code != error_code.SUCCESSSTATUS {
			return Live, code
		}
	} else if serviceType == "group" {
		Live.Activity.Type = "group"
		Live.Activity.ActivityEndTime, _ = activityInfo.Get("data").Get("detail").Get("end_time").Int64()
		Live.Activity.Name, _ = activityInfo.Get("data").Get("detail").Get("title").String()
		Live.Activity.TypeImg, _ = activityJsonInfo.Get("GroupImg").String()
		price, _ = activityInfo.Get("data").Get("detail").Get("group_price").String()
		gid, _ := activityInfo.Get("data").Get("detail").Get("goods_id").Int64()
		goodsId = fmt.Sprintf("%d", gid)
		gdJson, code = g.gda.GetGoodsDetail(goodsId, param.ClearCache)
		if code != error_code.SUCCESSSTATUS {
			return Live, code
		}
	} else {
		//获取免费类型的直播公开课
		Live.Activity.Type = "free"
		gdJson, code = g.gda.GetGoodsDetail(strconv.FormatInt(serviceId, 10), param.ClearCache)
		if code != error_code.SUCCESSSTATUS {
			return Live, code
		}
		gid, _ := gdJson.Get("result").Get("info").Get("id").Int64()
		if gid == 0 {
			return Live, error_code.PUBLICCOURSEERR
		}
		goodsId = fmt.Sprintf("%d", gid)
		Live.Activity.ActivityEndTime, _ = gdJson.Get("result").Get("info").Get("end_time").Int64()
		Live.Activity.Name, _ = gdJson.Get("result").Get("info").Get("title").String()

	}
	g.logger.Infof("JuniorGetHome_GetGoodsDetail_end_%d", param.StudentId)
	buyNum, _ := gdJson.Get("result").Get("info").Get("buy_num").Int64()
	buyNumFalse, _ := gdJson.Get("result").Get("info").Get("buy_num_false").Int64()
	Live.Activity.StartTime, _ = gdJson.Get("result").Get("info").Get("start_time").Int64()
	Live.Activity.EndTime, _ = gdJson.Get("result").Get("info").Get("end_time").Int64()
	Live.Activity.StudentNum = buyNum + buyNumFalse
	if code != error_code.SUCCESSSTATUS {
		return Live, code
	}
	now := time.Now().Unix()
	Live.Activity.Id, _ = activityJsonInfo.Get("Id").Int64()
	Live.Activity.Avatar, _ = activityJsonInfo.Get("Avatar").String()
	Live.Activity.Img, _ = activityJsonInfo.Get("Img").String()

	Live.Activity.Gif, _ = activityJsonInfo.Get("Gif").String()
	Live.Activity.ButtonInfo, _ = activityJsonInfo.Get("ButtonInfo").String()
	if Live.Activity.ButtonInfo == "" {
		Live.Activity.ButtonInfo = price + "元限时体验"
	}
	if now < Live.Activity.EndTime && now > Live.Activity.StartTime {
		Live.Activity.Gif, _ = activityJsonInfo.Get("DynamicGif").String()
		Live.Activity.ButtonInfo = "正在直播中"
	}
	if now > Live.Activity.ActivityEndTime {
		Live.Activity.ButtonInfo = "活动已结束，查看详情"
		Live.Activity.ButtonColor = "A0A2BA"
	} else {
		Live.Activity.ButtonColor = "FD7762"
	}
	Live.HaveData = true
	return Live, code
}
func (g *JuniorHomeService) GetGoodsList(param *Param, jsonMap map[string]*simpleJson.Json) (StudyInfo, int) {
	code := error_code.SUCCESSSTATUS
	goodsIds := []string{}
	goodsMap := map[string][]Goods{}
	goodsResp := StudyInfo{IsFree: true, GoodsList: []Goods{}, SalesMan: SalesMan{}}
	freeList := jsonMap["free_course"].Get("data")
	lightList := jsonMap["light_course"].Get("data")
	paidList := jsonMap["paid_course"].Get("data")
	lectureNoteList := jsonMap["lecture_note"].Get("data")
	tikuList := jsonMap["tiku"].Get("data")
	goodsBuy := Goods{}
	lectureNote := Goods{}
	tiku := Goods{}
	for i := len(freeList.MustArray()) - 1; i >= 0; i-- {
		jsonStr, _ := freeList.GetIndex(i).Get("json_detail").String()
		goods := Goods{}
		json.Unmarshal([]byte(jsonStr), &goods)
		goodsIds = append(goodsIds, goods.GoodsId)
		if (goods != Goods{}) {
			goodsBuy = goods
			goodsMap["free_course"] = append(goodsMap["free_course"], goods)
		}
	}
	for i := 0; i < len(lightList.MustArray()); i++ {
		jsonStr, _ := lightList.GetIndex(i).Get("json_detail").String()
		goods := Goods{}
		json.Unmarshal([]byte(jsonStr), &goods)
		goodsIds = append(goodsIds, goods.GoodsId)
		if (goods != Goods{}) {
			goodsMap["light_course"] = append(goodsMap["light_course"], goods)
		}
	}
	paidGoodsMap := map[string]Goods{}
	for i := 0; i < len(paidList.MustArray()); i++ {
		jsonStr, _ := paidList.GetIndex(i).Get("json_detail").String()
		goods := Goods{}
		json.Unmarshal([]byte(jsonStr), &goods)
		goodsIds = append(goodsIds, goods.GoodsId)
		if (goods != Goods{}) {
			if utils.VersionCompare("5.0.2", param.Version) {
				goods.Route = "tkofkjzc://study/course/list?goods_id=" + goods.GoodsId
			}
			paidGoodsMap[goods.GoodsId] = goods
			goodsMap["paid_course"] = append(goodsMap["paid_course"], goods)
		}
	}
	for i := 0; i < len(lectureNoteList.MustArray()); i++ {
		subjectId, _ := lectureNoteList.GetIndex(i).Get("service_id").Int64()
		if param.SubjectId == subjectId {
			jsonStr, _ := lectureNoteList.GetIndex(i).Get("json_detail").String()
			goods := Goods{}
			json.Unmarshal([]byte(jsonStr), &goods)
			goodsIds = append(goodsIds, goods.GoodsId)
			if (goods != Goods{}) {
				lectureNote = goods
				goodsMap["lecture_note"] = append(goodsMap["lecture_note"], goods)
			}
		}
	}
	for i := 0; i < len(tikuList.MustArray()); i++ {
		jsonStr, _ := tikuList.GetIndex(i).Get("json_detail").String()
		subjectId, _ := tikuList.GetIndex(i).Get("service_id").Int64()
		goods := Goods{}
		json.Unmarshal([]byte(jsonStr), &goods)
		if (goods != Goods{} && subjectId == param.SubjectId) {
			tiku = goods
			goodsMap["tiku"] = append(goodsMap["tiku"], goods)
		}
	}
	if len(goodsMap["free_course"]) == 0 ||
		len(goodsMap["light_course"]) == 0 ||
		len(goodsMap["paid_course"]) == 0 ||
		len(goodsMap["lecture_note"]) == 0 ||
		len(goodsMap["tiku"]) == 0 {
		code = error_code.GOODSCONFIGERR
		return goodsResp, code
	}
	g.logger.Infof("JuniorGetHome_GetGoodsBuyList_start_%d_%s", param.StudentId, strings.Join(goodsIds, "_"))
	redisClient := utils.RedisHandle
	key := fmt.Sprintf("GetGoodsBuyList%d_%d_%s", param.StudentId, strings.Join(goodsIds, "_"))
	//respStr := redisClient.GetData(key).(string)
	gbM := map[string]gdApi.GoodsBuy{}
	gbM, code = g.gda.GetGoodsBuyList(goodsIds, param.StudentId)
	//if respStr == "" || param.ClearCache == 1 {
	//	gbM, code = g.gda.GetGoodsBuyList(goodsIds, param.StudentId)
	//	if code != error_code.SUCCESSSTATUS {
	//		return goodsResp, code
	//	}
	//	cacheByte, err := json.Marshal(gbM)
	//	if err != nil {
	//		g.logger.Infof("GetGoodsBuyList_%s", err.Error())
	//		code = error_code.ITEMCATEGORYAPI
	//		return goodsResp, error_code.GOODSBUYAPI
	//	}
	//	respStr = string(cacheByte)
	//	g.logger.Infof("JuniorGetHome_GetGoodsBuyList_from_api_%d_%s", param.StudentId, respStr)
	//	redisClient.SetData(key, respStr, time.Second*600)
	//} else {
	//	g.logger.Infof("JuniorGetHome_GetGoodsBuyList_from_cache_%d_%s", param.StudentId, respStr)
	//	json.Unmarshal([]byte(respStr), &gbM)
	//}
	g.logger.Infof("JuniorGetHome_GetGoodsBuyList_end_%d", param.StudentId)

	//已购买付费课程
	for _, v := range goodsMap["paid_course"] {
		if gbM[v.GoodsId].BuyStatus == true {
			goodsResp.IsFree = false
			goods, ok := paidGoodsMap[v.GoodsId]
			if ok {
				goodsBuy = goods
			}
		}
	}
	salesMan, code := g.GetSalesMan(param.StudentId, param.ClearCache, !goodsResp.IsFree, false)
	g.logger.Infof("JuniorGetHome_GetSalesMan_end_%d", param.StudentId)
	if code != error_code.SUCCESSSTATUS {
		return goodsResp, code
	}
	goodsResp.SalesMan = *salesMan
	//尚未购买免费课程，自动购买
	g.logger.Infof("JuniorGetHome_BuyFreeCourse_start_%d_%d_", param.StudentId, param.YxStudentId)
	for _, v := range goodsMap["free_course"] {
		if gbM[v.GoodsId].BuyStatus == false {
			gd := gdApi.GoodsOrder{
				Uid:       param.StudentId,
				StudentId: param.YxStudentId,
				ProjectId: param.ProjectId,
				GoodsId:   utils.String2Int64(v.GoodsId),
			}
			g.logger.Infof("JuniorGetHome_BuyFreeCourse_%d_%s", param.StudentId, v.GoodsId)
			redisClient.Expire(key, 0)
			_, code := g.gda.GoodsBuy(gd)
			if code != error_code.SUCCESSSTATUS {
				return goodsResp, code
			}
		}
	}
	g.logger.Infof("JuniorGetHome_BuyFreeCourse_end_%d", param.StudentId)
	if param.IsAudit == 1 {
		_, ok := goodsMap["light_course"]
		if ok {
			goodsBuy.Route = goodsMap["light_course"][0].Route
		} else {
			goodsBuy.Route = goodsBuy.Route + "/1"
		}
	}
	goodsResp.GoodsList = append(goodsResp.GoodsList, goodsBuy)
	goodsResp.GoodsList = append(goodsResp.GoodsList, lectureNote)
	goodsResp.GoodsList = append(goodsResp.GoodsList, tiku)
	return goodsResp, code
}
func (g *JuniorHomeService) GetKnowledge(param *Param, json *simpleJson.Json) (*simpleJson.Json, int) {
	picid := ""
	var icp *simpleJson.Json
	icids := []string{}
	code := error_code.SUCCESSSTATUS
	for i := 0; i < len(json.MustArray()); i++ {
		id, _ := json.GetIndex(i).Get("service_id").Int64()
		if param.SubjectId == id {
			picid, _ = json.GetIndex(i).Get("desc").String()
		}
	}
	if picid == "" {
		code = error_code.KNOWLEDGEIDEMPTY
		return icp, code
	}
	g.logger.Infof("JuniorGetHome_GetItemCategoryList_start_%d", param.StudentId)
	res, code := g.ica.GetItemCategoryList(&icApi.ItemCategoryParam{ProjectId: param.ProjectId, SubjectId: param.SubjectId, Picid: picid}, param.ClearCache)
	g.logger.Infof("JuniorGetHome_GetItemCategoryList_end_%d", param.StudentId)
	if code != error_code.SUCCESSSTATUS {
		return icp, code
	}
	icp = res.Result.Get("result")
	for i := 0; i < len(icp.MustArray()); i++ {
		icid, _ := icp.GetIndex(i).Get("icid").String()
		icids = append(icids, icid)
	}
	g.logger.Infof("JuniorGetHome_DooneDataItem_start_%d", param.StudentId)
	countMap, err := g.ddi.GetCount(param.StudentId, icids)
	g.logger.Infof("JuniorGetHome_DooneDataItem_end_%d", param.StudentId)
	if err != nil {
		return icp, error_code.DOONEITEMMODELERR
	}
	for i := 0; i < len(icp.MustArray()); i++ {
		icid, _ := icp.GetIndex(i).Get("icid").String()
		numString, _ := icp.GetIndex(i).Get("num").String()
		num := utils.String2Float64(numString)
		_, ok := countMap[icid]
		if ok {
			var rate float64
			if num == 0 {
				rate = 0
			} else {
				rate = math.Floor(float64(countMap[icid]*100)/num + 0.5)
			}
			icp.GetIndex(i).Set("alread_do_num", countMap[icid])
			icp.GetIndex(i).Set("already_rate", countMap[icid])
			icp.GetIndex(i).Set("alread_do_rate", rate)
		} else {
			icp.GetIndex(i).Set("alread_do_num", 0)
			icp.GetIndex(i).Set("already_rate", 0)
			icp.GetIndex(i).Set("alread_do_rate", 0)
		}

	}
	return icp, code
}
func (g *JuniorHomeService) GetSalesMan(uid int64, clearCache int64, isVip bool, needVerifyVip bool) (*SalesMan, int) {
	result := error_code.SUCCESSSTATUS
	resp := SalesMan{}
	g.logger.Infof("GetSalesMan_start")
	jsonMap, result := g.GetTreeListMap("junior_study", clearCache)
	g.logger.Infof("GetSalesMan_GetTreeListMap_end")
	if result != error_code.SUCCESSSTATUS {
		return &resp, result
	}
	//为了提高性能，首页接口vip校验单独处理
	if needVerifyVip {
		paid, ok := jsonMap["paid_course"]
		if ok {
			pdlist := paid.Get("data")
			goodsIdsList := []string{}
			for i := 0; i < len(pdlist.MustArray()); i++ {
				courseBytes, _ := pdlist.GetIndex(i).Get("json_detail").Bytes()
				goodsJson, _ := simpleJson.NewJson(courseBytes)
				goodsId, _ := goodsJson.Get("goods_id").String()
				goodsIdsList = append(goodsIdsList, goodsId)
			}
			if len(goodsIdsList) > 0 {
				goodsBuy, code := g.gda.GetGoodsBuyList(goodsIdsList, uid)
				if code != error_code.SUCCESSSTATUS {
					return &resp, result
				}
				for _, v := range goodsBuy {
					if v.BuyStatus == true {
						isVip = true
					}
				}
			}
		}
	}
	_, ok := jsonMap["sales_man"]
	_, ok2 := jsonMap["group_student_num"]
	_, ok3 := jsonMap["salesman_vip"]
	if ok && ok2 {
		//查询营销学生id
		yuc, code := g.yuc.Get(uid)
		g.logger.Infof("GetSalesMan_GetYxStudentId_end_%d", yuc.YxStudentId)
		if code != error_code.SUCCESSSTATUS {
			return &resp, code
		}
		salesMan, code := g.smapi.GetSalesManList(smApi.SalesManParam{ProjectId: 14, YxStudentId: yuc.YxStudentId}, clearCache)
		g.logger.Infof("GetSalesMan_GetSalesManList_end")
		if code != error_code.SUCCESSSTATUS {
			return &resp, code
		}
		salesManList := jsonMap["sales_man"].Get("data")
		salesManMap := map[int64]*simpleJson.Json{}
		for i := 0; i < len(salesManList.MustArray()); i++ {
			jsonByte, _ := salesManList.GetIndex(i).Get("json_detail").Bytes()
			smJson, _ := simpleJson.NewJson(jsonByte)
			smId, _ := smJson.Get("id").Int64()
			salesManMap[smId] = smJson

		}
		studentInfo, _ := g.msm.Get(uid)
		g.logger.Infof("GetSalesMan_GdMembersStudent_end_")
		salesManId, _ := salesMan.Get("id").Int64()
		sm, ok := salesManMap[salesManId]
		if ok {
			resp.Avatar, _ = sm.Get("avatar").String()
			resp.Info, _ = sm.Get("info").String()
			resp.Route, _ = sm.Get("route").String()
		} else {
			resp.Avatar = "http://simg01.gaodunwangxiao.com/uploadimgs/tmp/upload/201902/21/083d2_20190221131325.png"
			resp.Info = "针对性学习规划，获取学习礼包"
			resp.Route = "tkofkjzc://common/h5/index?isLogin=1&url=http://muses.gaodun.com/pj/add-teacher?isLogin=1"
		}

		resp.Id, _ = salesMan.Get("id").Int64()
		resp.QRCode, _ = salesMan.Get("qcode").String()
		resp.Name, _ = salesMan.Get("name").String()
		resp.JoinInfo = fmt.Sprintf("HI,%s同学你好!\n我是你的伴学师，%s老师。体验效果如何？现邀请你进入班级共同学习~", studentInfo.Nickname, resp.Name)
		resp.MsMsgCode, _ = salesMan.Get("weixin_code").String()
		resp.ButtonTitle = "进班级群"
		if resp.QRCode == "" {
			resp.QRCode = "https://simg01.gaodunwangxiao.com/uploadimgs/goada/201903/12/ff494_20190312112418.png"
			resp.Name = "高顿会计小助"
			resp.MsMsgCode = "GD18321406726"
		}
		groupStudentNum := jsonMap["group_student_num"].Get("data")
		for i := 0; i < len(groupStudentNum.MustArray()); i++ {
			resp.StudentNum, _ = groupStudentNum.GetIndex(i).Get("service_id").Int64()
		}
		groupInfo := jsonMap["group_info"].Get("data")
		for i := 0; i < len(groupInfo.MustArray()); i++ {
			info, _ := groupInfo.GetIndex(i).Get("desc").String()
			resp.GroupInfoList = append(resp.GroupInfoList, info)
		}
	}
	resp.IsVip = isVip
	if ok3 && isVip {
		salesManVipList := jsonMap["salesman_vip"].Get("data")
		num := len(salesManVipList.MustArray())
		if num > 0 {
			index := int(uid) % num
			salesManVipBytes, _ := salesManVipList.GetIndex(index).Get("json_detail").Bytes()
			salesManVipJson, _ := simpleJson.NewJson(salesManVipBytes)
			resp.Name, _ = salesManVipList.GetIndex(index).Get("name").String()
			resp.Avatar, _ = salesManVipJson.Get("avatar").String()
			resp.Route, _ = salesManVipJson.Get("route").String()
			resp.QRCode, _ = salesManVipJson.Get("qr_code").String()
			resp.MsMsgCode, _ = salesManVipJson.Get("ms_msg_code").String()
		}
	}
	return &resp, result
}

/**
获取营销模块
*/
func (g *JuniorHomeService) GetMarketingInfo(param *MarketingParam) (*LiveInfo, int) {
	result := error_code.SUCCESSSTATUS
	LiveInfo := LiveInfo{}
	Param := Param{}
	//获取配置数据列表
	listJson, result := g.GetTreeListMap(param.ClassKey, param.ClearCache)
	if result != error_code.SUCCESSSTATUS {
		return &LiveInfo, result
	}
	var code int
	//获取首页公开课
	_, ok := listJson[param.PublicKey]
	if ok {
		Param.StudentId = param.StudentId
		Param.ClearCache = param.ClearCache
		g.logger.Infof("JuniorGetHome_GetLiveInfo_start_%d", param.StudentId)
		LiveInfo, code = g.GetLiveInfo(&Param, listJson[param.PublicKey])
		g.logger.Infof("JuniorGetHome_GetLiveInfo_start_%d", param.StudentId)
		LiveInfo.HaveData = true
		LiveInfo.Type = "PublicCourse"
		if code != error_code.SUCCESSSTATUS {
			LiveInfo.HaveData = false
			return &LiveInfo, code
		}
	}
	return &LiveInfo, result
}
