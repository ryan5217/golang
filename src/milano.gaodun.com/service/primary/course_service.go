package primary

import (
	"fmt"
	"github.com/apex/log"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/api/primary"
	"milano.gaodun.com/model/api/resources"
	"milano.gaodun.com/model/course"
	msa "milano.gaodun.com/model/members_student_assign"
	"milano.gaodun.com/model/order"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	jhService "milano.gaodun.com/service/junior_home"
	"milano.gaodun.com/service/tiku_constant"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type CourseService struct {
	nt        *primary.CourseApi
	logger    *log.Entry
	toc       *primary.TocApi
	goodsApi  *primary.GoodsApi
	course    *primary.CourseApi
	cm        *course.CourseModel
	ord       *order.OrderModel
	assignM   *msa.MembersStudentAssignModel
	jhs       jhService.JuniorHomeServiceInterface
	waitGroup sync.WaitGroup
	lock      sync.Mutex
}

func NewCourseService(logger *log.Entry) *CourseService {
	h := utils.HttpHandle
	var g CourseService
	g.logger = logger
	g.nt = primary.NewCourseApi(h)
	g.toc = primary.NewTocApi(h)
	g.course = primary.NewCourseApi(h)
	g.goodsApi = primary.NewGoodsApi(h)
	g.cm = course.NewCourseModel()
	g.assignM = msa.NewMembersStudentAssignModel()
	g.jhs = jhService.NewJuniorHomeService(logger)
	g.ord = order.NewOrderModel()
	return &g
}

type RecommendGoods struct {
	GoodsId string
	Info    string
	*simpleJson.Json
}

func (g *CourseService) GetCourseSyllabus(courseId int64, uid int64) (*simpleJson.Json, error) {
	//获取阶段大纲列表
	gradation, err := g.toc.GetSyllabusList(courseId)
	courseIds := []int64{}
	//获取派课信息
	courseIds = append(courseIds, courseId)
	course, _ := g.cm.GetCourseBySaasCourseId(courseId)
	courseAssigns, _ := g.assignM.GetAssigns(uid, []int64{course.Id})
	gra, err := simpleJson.NewJson([]byte(gradation))
	if err != nil || len(courseAssigns) == 0 {
		return &simpleJson.Json{}, err
	}
	//获取大纲学习状态
	for i := 0; i < len(gra.MustArray()); i++ {
		syllabusId, _ := gra.GetIndex(i).Get("syllabus_id").String()
		ss, _ := g.course.GetSyllabusStudyStatus(uid, courseId, syllabusId)
		syllabus, _ := simpleJson.NewJson([]byte(ss))
		//添加是否试听字段
		syllbusChildren := syllabus.Get("result").GetIndex(0).Get("children")
		graChildren := gra.GetIndex(i).Get("syllabus").GetIndex(0).Get("children")
		for j := 0; j < len(syllbusChildren.MustArray()); j++ {
			for k := 0; k < len(graChildren.MustArray()); k++ {
				jid, _ := syllbusChildren.GetIndex(j).Get("id").Int()
				kid, _ := graChildren.GetIndex(k).Get("id").Int()
				if jid == kid {
					audition := graChildren.GetIndex(k).Get("audition").Interface()
					syllabus.Get("result").GetIndex(0).Get("children").GetIndex(j).Set("audition", audition)
				}
			}
		}
		gra.GetIndex(i).Set("syllabus", syllabus.Get("result"))
	}
	return gra, err
}

//获取课程列表
func (g *CourseService) GetCourseListByGoodsId(goodsId int64, uid int64, clearCache int64, orderType string) (*primary.CourseListResp, error) {
	courseListResp := primary.CourseListResp{SaasCourses: []primary.CourseInfo{}}
	saasCourses := []primary.CourseInfo{}
	//根据商品id获取v课程信息
	gdresp, err := g.goodsApi.GetGoodsById(goodsId, uid, clearCache, orderType)
	goods, err := simpleJson.NewJson([]byte(gdresp))
	vCourseId, err := goods.Get("result").Get("info").Get("product_info").GetIndex(0).Get("vid").Int64()
	vCourse, err := g.cm.Get(vCourseId)
	courseListResp.PackageName = vCourse.Name
	courseIds := []int64{}
	saasCourseIds := req.Param{}
	var courseAssigns = []msa.GdMembersStudentAssign{}
	if vCourse.Isbig == 0 && vCourse.Id > 0 {
		//单课程
		respOne := primary.CourseInfo{
			StudentsNum:  vCourse.Studentsnum,
			CourseId:     vCourseId,
			CourseName:   vCourse.Name,
			SaasCourseId: vCourse.RelationCourse,
			CourseType:   vCourse.Istasks,
		}
		saasCourses = append(saasCourses, respOne)
		courseIds = append(courseIds, vCourseId)
		saasCourseIds["course_id_list[0]"] = vCourse.RelationCourse
		courseAssigns, err = g.assignM.GetAssigns(uid, courseIds)
	} else if vCourse.Courses != "" {
		//大包
		courses, err := g.cm.GetCourseByIds(vCourse.Courses)
		if err != nil {
			return &courseListResp, err
		}
		var isAssign = map[int64]bool{}
		for _, v := range *courses {
			courseIds = append(courseIds, v.Id)
			isAssign[v.Id] = false
		}
		courseAssigns, err = g.assignM.GetAssigns(uid, courseIds)
		if err != nil {
			return &courseListResp, err
		}
		for _, w := range courseAssigns {
			isAssign[w.CourseId] = true
		}
		i := 0
		for _, x := range *courses {
			if isAssign[x.Id] {
				respOne := primary.CourseInfo{
					StudentsNum:  x.Studentsnum,
					CourseId:     x.Id,
					CourseName:   x.Name,
					SaasCourseId: x.RelationCourse,
					CourseType:   x.Istasks,
				}
				saasCourses = append(saasCourses, respOne)
				saasCourseIds["course_id_list["+strconv.Itoa(i)+"]"] = x.RelationCourse
				i++
			}
		}
	}
	if len(courseAssigns) == 0 {
		return &courseListResp, err
	}
	progressMap := map[string]primary.StudyProgressInfo{}
	courseProgress, err := g.course.GetCourseStudyProgress(uid, saasCourseIds)
	if len(courseProgress.Result) > 0 {
		for _, s := range courseProgress.Result {
			progressMap[s.CourseId] = s
		}
	}
	for _, c := range saasCourses {
		s, ok := progressMap[fmt.Sprintf("%d", c.SaasCourseId)]
		if ok && strconv.FormatInt(c.SaasCourseId, 10) == s.CourseId {
			c.Progress = s.Progress
			c.AllChapter = s.AllChapter
		}
		courseListResp.SaasCourses = append(courseListResp.SaasCourses, c)
	}
	//if err != nil || vCourseId == 0 {
	//	return &courseListResp, err
	//}

	return &courseListResp, nil
}
func (g *CourseService) CourseSyllabusStudyStatusRecord(param primary.StudyProgressParam) (*simpleJson.Json, error) {
	//记录学习进度
	record, err := g.course.RecordSyllabusStudyStatus(param)
	rec, err := simpleJson.NewJson([]byte(record))
	status, _ := rec.Get("status").Int64()
	if err != nil || status != 0 {
		return &simpleJson.Json{}, err
	}
	return rec.Get("result"), err
}
func (g *CourseService) GetGoodsCourseList(uid int64, isAudit int64) ([]*simpleJson.Json, int) {
	resp := []*simpleJson.Json{}
	g.logger.Infof("CourseService_GetGoodsCourseList_start_%d", uid)
	treeListMap, code := g.jhs.GetTreeListMap("junior_study", 0)
	if code != error_code.SUCCESSSTATUS {
		return resp, code
	}
	free, ok := treeListMap["free_course"]
	light, ok1 := treeListMap["light_course"]
	paid, ok2 := treeListMap["paid_course"]
	if !ok || !ok1 || !ok2 {
		return resp, error_code.GOODSCONFIGERR
	}
	g.logger.Infof("CourseService_GetTreeListMap_end_%d", uid)
	freeGoodsMap := g.GetGoodsId(free, "one")
	lightGoodsMap := g.GetGoodsId(light, "all")
	paidGoodsMap := g.GetGoodsId(paid, "all")
	paidExpCourse := g.GetGoodsId(paid, "one")
	paidGoodsExpMap := map[int64]bool{}
	for K, _ := range freeGoodsMap {
		paidGoodsExpMap[K] = true
	}
	for K, _ := range lightGoodsMap {
		paidGoodsExpMap[K] = true
	}
	for K, _ := range paidExpCourse {
		paidGoodsExpMap[K] = true
	}
	g.logger.Infof("CourseService_GetGoodsOrderList_start_%d", uid)
	gdList, code := g.goodsApi.GetGoodsOrderList(uid, 14, isAudit)
	g.logger.Infof("CourseService_GetGoodsOrderList_end_%d", uid)
	if code != error_code.SUCCESSSTATUS {
		return resp, code
	}
	orderGoodsMap := map[int64]*simpleJson.Json{}
	for i := 0; i < len(gdList.MustArray()); i++ {
		goodsId, _ := gdList.GetIndex(i).Get("goods_id").Int64()
		orderGoodsMap[goodsId] = gdList.GetIndex(i)
	}
	experience, ok := treeListMap["experience_course"]
	title := "零基础口碑班体验课"
	if ok {
		title, _ = experience.Get("name").String()
	}

	for k, _ := range orderGoodsMap {
		orderGoodsMap[k].Set("route", "")
		paidGoodsIsBuy, ok := paidExpCourse[k]
		if ok && paidGoodsIsBuy {
			for lightGoodsId, _ := range lightGoodsMap {
				delete(orderGoodsMap, lightGoodsId)
			}
			for freeGoodsId, _ := range freeGoodsMap {
				delete(orderGoodsMap, freeGoodsId)
			}
		}
		lightGoodsIsBuy, ok := lightGoodsMap[k]
		if ok && lightGoodsIsBuy {
			for freeGoodsId, _ := range freeGoodsMap {
				delete(orderGoodsMap, freeGoodsId)
			}
		}
	}
	freeJsonBytes, _ := free.Get("data").GetIndex(0).Get("json_detail").Bytes()
	freeGoodsJson, _ := simpleJson.NewJson(freeJsonBytes)
	route, _ := freeGoodsJson.Get("route").String()
	for freeGoodsId, _ := range freeGoodsMap {
		_, ok = orderGoodsMap[freeGoodsId]
		if ok {
			orderGoodsMap[freeGoodsId].Set("route", route)
			orderGoodsMap[freeGoodsId].Set("title", title)
		}
	}
	for lightGoodsId, _ := range lightGoodsMap {
		_, ok = orderGoodsMap[lightGoodsId]
		if ok {
			orderGoodsMap[lightGoodsId].Set("route", route)
			orderGoodsMap[lightGoodsId].Set("title", title)
		}
	}
	for paidGoodsId, _ := range paidExpCourse {
		_, ok = orderGoodsMap[paidGoodsId]
		if ok {
			orderGoodsMap[paidGoodsId].Set("title", title)
		}
	}
	orderIds := []int64{}
	for _, v := range orderGoodsMap {
		orderId, _ := v.Get("order_id").Int64()
		orderIds = append(orderIds, orderId)

	}
	//查询订单下单时间
	g.logger.Infof("CourseService_GetOrderByIds_start_%d", uid)
	orders, code := g.ord.GetOrderByIds(orderIds)
	g.logger.Infof("CourseService_GetOrderByIds_end_%d", uid)
	if code != error_code.SUCCESSSTATUS {
		return resp, code
	}
	//付费课程排在前面
	for _, v := range *orders {
		for _, g := range orderGoodsMap {
			oid, _ := g.Get("order_id").Int64()
			gid, _ := g.Get("goods_id").Int64()
			paidGoodsIsBuy, ok := paidGoodsMap[gid]
			if oid == v.Id && ok && paidGoodsIsBuy {
				g.Set("modify_date", v.Modifydate)
				resp = append(resp, g)
				delete(orderGoodsMap, gid)
			}
		}
	}
	//其他课程排在后面
	for _, v := range *orders {
		for _, g := range orderGoodsMap {
			oid, _ := g.Get("order_id").Int64()
			if oid == v.Id {
				g.Set("modify_date", v.Modifydate)
				resp = append(resp, g)
			}
		}
	}
	//0基础口碑班放在vip后面
	for i := 0; i < len(resp)-1; i++ {
		gid1, _ := resp[i].Get("goods_id").Int64()
		gid2, _ := resp[i+1].Get("goods_id").Int64()
		paidExp, ok := paidGoodsExpMap[gid1]
		paidGoodsIsBuy, ok2 := paidGoodsMap[gid2]
		if paidExp && ok && ok2 && paidGoodsIsBuy {
			tmp := resp[i]
			resp[i] = resp[i+1]
			resp[i+1] = tmp
		}
	}
	//0基础口碑班放在其他非vip课前面
	for i := len(resp) - 1; i > 0; i-- {
		gid1, _ := resp[i].Get("goods_id").Int64()
		gid2, _ := resp[i-1].Get("goods_id").Int64()
		paidExp, ok := paidGoodsExpMap[gid1]
		paidGoodsIsBuy, ok2 := paidGoodsMap[gid2]
		if paidExp && ok && (!ok2 || !paidGoodsIsBuy) {
			tmp := resp[i]
			resp[i] = resp[i-1]
			resp[i-1] = tmp
		}
	}
	g.logger.Infof("CourseService_GetGoodsCourseList_end_%d", uid)
	return resp, error_code.SUCCESSSTATUS
}
func (g *CourseService) GetGoodsId(json *simpleJson.Json, num string) map[int64]bool {
	result := map[int64]bool{}
	data := json.Get("data")

	if num == "one" {
		jsonDetailBytes, _ := data.GetIndex(0).Get("json_detail").Bytes()
		jsonDetail, _ := simpleJson.NewJson(jsonDetailBytes)
		goodsIdStr, _ := jsonDetail.Get("goods_id").String()
		result[utils.String2Int64(goodsIdStr)] = true
	} else {
		for i := 0; i < len(data.MustArray()); i++ {
			jsonDetailBytes, _ := data.GetIndex(i).Get("json_detail").Bytes()
			jsonDetail, _ := simpleJson.NewJson(jsonDetailBytes)
			goodsIdStr, _ := jsonDetail.Get("goods_id").String()
			result[utils.String2Int64(goodsIdStr)] = true
		}
	}
	return result
}
func (g *CourseService) GetColumnGoods(columnId int64, projectId int64, goodsStr string, clearCache int64) (*simpleJson.Json, int) {
	columnGoods, code := g.goodsApi.GetColumnGoods(columnId, projectId, 1)
	if code != error_code.SUCCESSSTATUS {
		return &simpleJson.Json{}, code
	}
	goodsIds := strings.Split(goodsStr, ",")
	goodsList := columnGoods.Get("info").Get("goods_list")
	goodsListMap := map[int64]*simpleJson.Json{}
	for i := 0; i < len(goodsList.MustArray()); i++ {
		goodsId, _ := goodsList.GetIndex(i).Get("goods_id").Int64()
		goodsListMap[goodsId] = goodsList.GetIndex(i)
	}
	respGoodsList := []*simpleJson.Json{}
	for _, v := range goodsIds {
		goods, ok := goodsListMap[utils.String2Int64(v)]
		if ok {
			respGoodsList = append(respGoodsList, goods)
		}
	}
	columnGoods.Get("info").Set("goods_list", respGoodsList)
	return columnGoods, error_code.SUCCESSSTATUS
}

func (g *CourseService) GetCourseList(goodsId int64, uid int64, clearCache int64, orderType string) (*primary.CourseListResp, error) {
	courseListResp := primary.CourseListResp{SaasCourses: []primary.CourseInfo{}}
	saasCourses := keInfo{}
	//根据商品id获取v课程信息
	defer func() {
		if p := recover(); p != nil {
			setting.Logger.Infof("SystemMessageModel_delete_%s", p)
		}
	}()
	gdresp, err := g.goodsApi.GetGoodsById(goodsId, uid, clearCache, orderType)
	goods, err := simpleJson.NewJson([]byte(gdresp))
	vCourseId, err := goods.Get("result").Get("info").Get("product_info").GetIndex(0).Get("vid").Int64()
	vCourse, err := g.cm.Get(vCourseId)
	courseListResp.PackageName = vCourse.Name
	courseListResp.IsListen = "0"
	//拉取直播
	tikuContent := tiku_constant.NewTikuConstantList()
	data, err := tikuContent.GetKey("ZQCY_LIVE_COURSE_LIST")
	list, error := simpleJson.NewJson([]byte(data))
	if error != nil {
		setting.Logger.Errorf("json_Error" + error.Error())
	}
	var liveList map[string]interface{}
	if list != nil {
		liveList = list.MustMap()
	}
	if vCourse.Isbig == 0 && vCourse.Id > 0 {
		//单课程
		respOne := primary.CourseInfo{
			StudentsNum:  vCourse.Studentsnum,
			CourseId:     vCourseId,
			CourseName:   vCourse.Name,
			SaasCourseId: vCourse.RelationCourse,
			CourseType:   vCourse.Istasks,
		}
		respOne.LiveList = []interface{}{}
		if vCourse.RelationCourse == 0 && len(liveList) > 0 {
			experience, ok := liveList[string(vCourseId)]
			if ok {
				respOne.LiveList = g.getLiveList(experience, &courseListResp)
			} else {
				respOne.LiveList = []interface{}{}
			}
		}
		saasCourses = append(saasCourses, respOne)
	} else if vCourse.Courses != "" {
		//大包
		courses, err := g.cm.GetCourseByIds(vCourse.Courses)
		if err != nil {
			return &courseListResp, err
		}
		i := 0
		for _, x := range *courses {
			respOne := primary.CourseInfo{
				StudentsNum:  x.Studentsnum,
				CourseId:     x.Id,
				CourseName:   x.Name,
				SaasCourseId: x.RelationCourse,
				CourseType:   x.Istasks,
			}
			respOne.LiveList = []interface{}{}
			if x.RelationCourse == 0 && len(liveList) > 0 {
				vid := fmt.Sprintf("%v", x.Id)
				experience, ok := liveList[vid]
				if ok {
					respOne.LiveList = g.getLiveList(experience, &courseListResp)
				}
			} else {
				respOne.LiveList = []interface{}{}
			}
			saasCourses = append(saasCourses, respOne)
			i++
		}
	}
	//排序
	sort.Sort(saasCourses)
	for _, c := range saasCourses {
		courseListResp.SaasCourses = append(courseListResp.SaasCourses, c)
	}
	return &courseListResp, err
}

func (g *CourseService) getLiveList(params interface{}, courseListResp *primary.CourseListResp) []interface{} {
	var result *simpleJson.Json
	var LiveList []interface{}
	experience := params.([]interface{})
	courseList := make(chan interface{}, len(experience))
	if len(experience) > 0 {
		for _, c := range experience {
			courseInfo := c.(map[string]interface{})
			id, ok := courseInfo["id"]
			if ok && id != "" {
				g.waitGroup.Add(1)
				go func(info map[string]interface{}) {
					g.lock.Lock()
					res := resources.Resource{}
					s, _ := res.Get(conf.RESOURCE_DOMAIN + "/resource/" + info["id"].(string))
					list, error := simpleJson.NewJson([]byte(s))
					if error != nil {
						setting.Logger.Errorf("json_Error" + error.Error())
					}
					result = list.Get("result")
					result.Set("name", info["name"])
					result.Set("listen", info["listen"])
					if info["listen"].(string) == "1" {
						courseListResp.IsListen = "1"
					}
					courseList <- result.MustMap()
					g.lock.Unlock()
					g.waitGroup.Done()
				}(courseInfo)
			}
		}
		g.waitGroup.Wait()
		for i := 0; i <= len(courseList); i++ {
			list := <-courseList
			LiveList = append(LiveList, list)
		}
	}
	for k, v := range experience {
		vId := v.(map[string]interface{})["id"].(string)
		for _, lv := range LiveList {
			lId := lv.(map[string]interface{})["id"]
			if vId == fmt.Sprintf("%v", lId) {
				experience[k] = lv
			}
		}
	}
	return experience
}

type keInfo []primary.CourseInfo

//排序
//Len()
func (s keInfo) Len() int {
	return len(s)
}

//Less()
func (s keInfo) Less(i, j int) bool {
	iLive := s[i].LiveList.([]interface{})
	jLive := s[j].LiveList.([]interface{})
	return len(iLive) > len(jLive)
}

//Swap()
func (s keInfo) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
