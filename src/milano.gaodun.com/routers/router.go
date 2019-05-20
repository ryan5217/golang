package routers

import (
	"milano.gaodun.com/middleware/logs"
	"milano.gaodun.com/middleware/panicHandle"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "milano.gaodun.com/docs"

	"milano.gaodun.com/conf"
	"milano.gaodun.com/controller"
	"milano.gaodun.com/pkg/utils"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	gin.SetMode(conf.RunMode)

	r.Use(cors.Default())

	r.Use(logs.Logger())
	r.Use(gin.Recovery())
	r.Use(panicHandle.CatchError())
	// primary
	r.Any("/", Root)
	r.GET("/status", utils.GetStatus)

	// v1 app 接口使用
	v1 := r.Group("/v1")

	// banner
	banner := v1.Group("/banner")
	banner.GET("/getList", controller.BannerApi.GetList)
	banner.POST("/modify/:id", controller.BannerApi.Modify)
	banner.POST("/add", controller.BannerApi.Add)
	banner.DELETE("/delete/:id", controller.BannerApi.Delete)

	studyRights := v1.Group("/right")
	studyRights.GET("/list", controller.StudyRightsApi.GetList)
	studyRights.GET("/one", controller.StudyRightsApi.GetOne)
	studyRights.GET("/subject/list", controller.StudyRightsApi.GetSubjectList)
	// activity
	activity := v1.Group("/activity")
	// home
	home := v1.Group("/primary")
	home.GET("/note/list", controller.PrimaryApi.GetNoteList)
	home.POST("/note/add", controller.PrimaryApi.Add)
	home.POST("/note/edit", controller.PrimaryApi.Edit)
	home.GET("/glive", controller.PrimaryApi.GetGlive)
	home.POST("/user/stage", controller.PrimaryApi.StageSet)   //学员阶段解锁设置
	home.GET("/user/stage", controller.PrimaryApi.GetStageSet) //获取学员阶段解锁设置
	// junior home
	juniorHome := v1.Group("/junior")
	juniorHome.GET("/home", controller.JuniorHomeApi.GetHome)
	juniorHome.GET("/salesMan", controller.JuniorHomeApi.GetSalesMan)
	juniorHome.GET("/home/study", controller.JuniorHomeApi.GetHomeStudyInfo)
	juniorHome.GET("/column/goods/list", controller.CourseApi.GetColumnGoodsByGoodsIds)
	juniorHome.GET("/marketing", controller.JuniorHomeApi.GetMarketingInfo) //获取营销模块
	// message home
	message := v1.Group("/message")
	message.POST("/modify", controller.MessageApi.Modify)
	message.GET("/system/list", controller.MessageApi.GetSystemMessages)
	message.GET("/list", controller.MessageApi.List)
	message.POST("/system/delete", controller.MessageApi.Delete)
	message.POST("/notification/delete", controller.MessageApi.DeleteNotification)
	message.POST("/notification/read", controller.MessageApi.ReadMessage)
	message.GET("/notification/num", controller.MessageApi.GetNotifyMessageNum)
	// advert
	advert := v1.Group("/advert")
	advert.GET("/list", controller.AdvertApi.GetAdvertList)
	advert.POST("/add/record", controller.AdvertApi.AddAdvertRecord)
	// exercise
	exercise := v1.Group("/exercise")
	exercise.POST("/add/record", controller.ExerciseRecordApi.AddRecord)
	exercise.GET("/get/record", controller.ExerciseRecordApi.GetRecord)
	// home
	ask := v1.Group("/ask")
	ask.GET("/list", controller.AskApi.List)
	ask.GET("/detail", controller.AskApi.AskDetail)
	ask.POST("/answer/ask", controller.AskApi.AnswerAsk)
	// course
	course := v1.Group("/course")
	course.GET("/syllabus", controller.CourseApi.GetCourseSyllabus)
	course.GET("/list", controller.CourseApi.GetCourseListByGoodsId)
	course.GET("/detailList", controller.CourseApi.GetCourseList)
	course.POST("/syllabus/study", controller.CourseApi.CourseSyllabusStudyStatusRecord)
	course.GET("/buy/list", controller.CourseApi.GetGoodsCourseList)

	activity.GET("/getList", controller.ActivityApi.GetList)         // 查询学员分享状态
	activity.POST("/add", controller.ActivityApi.Add)                // 记录状态
	activity.POST("/add/user", controller.ActivityApi.AddUser)       // 记录状态联系方式
	activity.POST("/add/contact", controller.ActivityApi.AddContact) // 记录状态联系方式
	activity.GET("/get/contact", controller.ActivityApi.GetContact)  // 记录状态联系方式
	// app h5 使用
	app := r.Group("/app")
	h5 := app.Group("/h5")
	h5.GET("/get", controller.H5Api.GetApp) // 分享

	// gd info
	gd := v1.Group("/gdInfo")
	gd.GET("/get", controller.GdInfApi.Get)
	gd.GET("/getP", controller.GdInfApi.GetInfo)
	// invitation info
	inv := v1.Group("/invitation")
	inv.GET("/getUuid", controller.InvitationApi.GetUuid)
	inv.GET("/getCode", controller.InvitationApi.GetCode)
	inv.POST("/addInvitation", controller.VerifyF.VerifyFlag, controller.InvitationApi.AddInvitation)
	inv.GET("/invitationList", controller.InvitationApi.InvitationList)
	inv.GET("/getPhoneCode", controller.VerifyF.VerifyFlag, controller.InvitationApi.GetPhoneCode)
	goods := v1.Group("/goods")
	goods.POST("/add", controller.GoodsApi.Add)
	goods.POST("/modify/:id", controller.GoodsApi.Edit)
	goods.GET("/list", controller.GoodsApi.List)
	goods.GET("/getList", controller.GoodsApi.GetList)
	goods.GET("/findAll", controller.GoodsApi.FindAll)
	// invitation info
	exc := v1.Group("/exchange")
	exc.POST("/getCode", controller.ExchangeApi.GetCode)
	exc.POST("/modify/:id", controller.ExchangeApi.Modify)
	exc.GET("/getList", controller.ExchangeApi.GetList)
	qu := v1.Group("/questionnaire")
	qu.POST("/add", controller.QuestionnaireApi.Add)
	qu.POST("/edit/:id", controller.QuestionnaireApi.Edit)

	// wx program
	pro := v1.Group("wx")
	pro.POST("notice", controller.WPApi.Notice)
	pro.GET("notice_time", controller.WPApi.NoticeTime)
	pro.GET("live", controller.WPApi.GetLive)
	pro.GET("access_token", controller.WPApi.GetAccessToken)
	pro.GET("batchget_material", controller.WPApi.GetMaterialList)
	pro.GET("init_material", controller.WPApi.InitWechatMaterial)
	pro.GET("del_material", controller.WPApi.DelMaterial)

	//用户反馈
	collection := v1.Group("/collection")
	collection.GET("/getCollection", controller.CollectionApi.GetCollection)
	collection.POST("/addCollection", controller.CollectionApi.AddCollection)
	collection.GET("/getInfoList", controller.CollectionApi.InfoList)
	collection.GET("/getTypeList", controller.CollectionApi.TypeList)
	collection.GET("/getQuestionList", controller.CollectionApi.QuestionList)
	collection.GET("/getInfo", controller.CollectionApi.GetInfoAndQuestion)
	collection.POST("/addType", controller.CollectionApi.AddType)
	collection.POST("/editType", controller.CollectionApi.EditType)

	//salesman
	salesman := v1.Group("/salesman")
	salesman.GET("/getList", controller.SalseApi.GetList)

	// youzhan
	youzhan := v1.Group("/youzan")
	youzhan.POST("/accept", controller.YouzanApi.Accept)
	youzhan.GET("/push_crm", controller.YouzanApi.PushCrm)

	//获取消息推送
	msg := v1.Group("/msg")
	msg.GET("send",controller.SendMessageApi.SendMsg)

	// tiku const
	tiku := v1.Group("/tiku")
	tiku.GET("/const", controller.TikuBaseApi.TkConst)

	return r
}
