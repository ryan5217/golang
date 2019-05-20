package controller

import (
	"github.com/gin-gonic/gin"
	bModle "milano.gaodun.com/model/banner"
	"milano.gaodun.com/pkg/setting"
	bService "milano.gaodun.com/service/banner"
	"strconv"
)

var BannerApi = NewBannerApi()

func NewBannerApi() *Banner {
	return &Banner{}
}

type Banner struct {
	Base
}

func (b Banner) GetList(c *gin.Context) {
	projectId := b.GetInt64(c, "project_id")
	subjectId := b.GetInt64(c, "subject_id")
	showMode := b.GetInt32(c, "show_mode")
	source := b.GetInt32(c, "source")

	param := bModle.BannerParam{ProjectId: projectId, ShowMode: showMode, SubjectId: subjectId, Source: source}

	bServer := bService.NewBannerService(&param, setting.GinLogger(c))
	r := bServer.List()
	b.ServerJSONSuccess(c, r)
}
func (b Banner) Modify(c *gin.Context) {

	//var ForceUpdateColumn map[string]bool
	Id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	BannerParam := bModle.BannerParam{Id: Id}

	gab := bModle.PriAppBanner{}
	gab.Id = Id
	gab.Title = b.GetString(c, "title")
	gab.ShowMode = b.GetInt32(c, "show_mode")
	gab.Source = b.GetInt32(c, "source")
	gab.ProjectId = b.GetInt64(c, "project_id")
	gab.SubjectId = b.GetInt64(c, "subject_id")
	gab.PictureUrl = b.GetString(c, "picture_url")
	gab.SkipUrl = b.GetString(c, "skip_url")
	gab.ShowType = b.GetInt32(c, "show_type")
	gab.ShowTableId = b.GetInt64(c, "show_table_id")
	gab.ShowTableDict = b.GetString(c, "show_table_dict")
	gab.Sort = b.GetInt32(c, "sort")
	gab.Status = b.GetInt32(c, "status")
	gab.StartDate = b.GetString(c, "start_date")
	gab.EndDate = b.GetString(c, "end_date")
	gab.Remark = b.GetString(c, "remark")
	gab.ForceUpdateCol = b.PostMustCols(c)
	bServer := bService.NewBannerService(&BannerParam, setting.GinLogger(c))
	bServer.Edit(&gab)
	b.ServerJSONSuccess(c, gab)
}
func (b Banner) Add(c *gin.Context) {
	gab := bModle.PriAppBanner{}
	gab.Title = b.GetString(c, "title")
	gab.ShowMode = b.GetInt32(c, "show_mode")
	gab.Source = b.GetInt32(c, "source")
	gab.ProjectId = b.GetInt64(c, "project_id")
	gab.SubjectId = b.GetInt64(c, "subject_id")
	gab.PictureUrl = b.GetString(c, "picture_url")
	gab.SkipUrl = b.GetString(c, "skip_url")
	gab.ShowType = b.GetInt32(c, "show_type")
	gab.ShowTableId = b.GetInt64(c, "show_table_id")
	gab.ShowTableDict = b.GetString(c, "show_table_dict")
	gab.Sort = b.GetInt32(c, "sort")
	gab.Status = b.GetInt32(c, "status")
	gab.StartDate = b.GetString(c, "start_date")
	gab.EndDate = b.GetString(c, "end_date")
	gab.Remark = b.GetString(c, "remark")
	gab.ForceUpdateCol = b.PostMustCols(c)
	bServer := bService.NewBannerService(&bModle.BannerParam{}, setting.GinLogger(c))
	bServer.Add(&gab)
	b.ServerJSONSuccess(c, gab)
}
func (b Banner) Delete(c *gin.Context) {

	//var ForceUpdateColumn map[string]bool
	Id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	BannerParam := bModle.BannerParam{Id: Id}
	bServer := bService.NewBannerService(&BannerParam, setting.GinLogger(c))
	r := bServer.Delete()
	b.ServerJSONSuccess(c, r)
}
