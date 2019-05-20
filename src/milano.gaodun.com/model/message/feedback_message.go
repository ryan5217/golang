package message

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type TkFeedback struct {
	Id             int64
	Uid            int64
	Content        string
	Pic            string
	Url            string
	Status         int
	RetContent     string
	IsRead         int
	ProjectId      int64
	Isdel          int
	UpdatedTime    string
	CreatedTime    string
	ForceUpdateCol []string `xorm:"-" json:"-"`
}

type FeedbackMessageModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewFeedbackMessageModel() *FeedbackMessageModel {
	return &FeedbackMessageModel{Engine: utils.TikuDbo}
}

func (g *FeedbackMessageModel) GetFeedbackMessages(param Param) ([]TkFeedback, error) {
	messages := []TkFeedback{}
	var err error
	err = g.Where("project_id=?", param.ProjectId).
		Where("isdel=?", 0).
		Where("uid=?", param.Uid).
		Where("status=?", 2).
		OrderBy("id desc").
		Limit(param.Limit, (param.Page-1)*param.Limit).
		Find(&messages)
	if err != nil {
		setting.Logger.Infof("FeedbackMessageModel_GetMessages_%d_%s", param.ProjectId, err.Error())
	}
	return messages, err
}
func (g *FeedbackMessageModel) GetFeedbackMessageCount(param Param) (int, error) {
	message := TkFeedback{}
	var err error
	var num int64
	num, err = g.Where("project_id=?", param.ProjectId).
		Where("isdel=?", 0).
		Where("uid=?", param.Uid).
		Where("status=?", 2).
		Count(&message)
	if err != nil {
		setting.Logger.Infof("FeedbackMessageModel_GetFeedbackMessageCount_%d_%s", param.ProjectId, err.Error())
	}
	return int(num), err
}
func (g *FeedbackMessageModel) GetUnreadFeedbackMessageCount(param Param) (int, error) {
	message := TkFeedback{}
	var err error
	var num int64
	num, err = g.Where("project_id=?", param.ProjectId).
		Where("isdel=?", 0).
		Where("uid=?", param.Uid).
		Where("status=?", 2).
		Where("is_read=?", 0).
		Count(&message)
	if err != nil {
		setting.Logger.Infof("FeedbackMessageModel_GetUnreadFeedbackMessageCount_%d_%s", param.ProjectId, err.Error())
	}
	return int(num), err
}
func (g *FeedbackMessageModel) Modify(ts *TkFeedback) (int64, error) {
	var row int64
	var err error
	if ts.Id > 0 {
		ts.UpdatedTime = time.Now().String()
		row, err = g.Id(ts.Id).Cols(ts.ForceUpdateCol...).Omit("created_time").Omit("updated_time").Update(ts)

	} else {
		var cstSh, _ = time.LoadLocation("Asia/Shanghai")
		ts.UpdatedTime = time.Now().In(cstSh).Format("2006-01-02 15:04:05")
		row, err = g.Omit("created_time").Omit("updated_time").InsertOne(ts)
	}
	if err != nil {
		setting.Logger.Infof("FeedbackMessageModel_AddMessage_%s", err.Error())
	}
	return row, err
}
func (g *FeedbackMessageModel) DeleteMessage(id int64, uid int64) (int64, error) {
	var row int64
	var err error
	sm := TkFeedback{Isdel: 1}
	if id == -1 {
		row, err = g.Where("status=?", 2).Where("uid=?", uid).Update(sm)
	} else {
		row, err = g.Id(id).Where("status=?", 2).Where("uid=?", uid).Update(sm)
	}

	if err != nil {
		setting.Logger.Infof("FeedbackMessageModel_delete_%s", err.Error())
	}
	return row, err
}
func (g *FeedbackMessageModel) ReadMessage(id int64, uid int64) (int64, error) {
	var row int64
	var err error
	sm := TkFeedback{IsRead: 1}
	if id == -1 {
		row, err = g.Where("status=?", 2).Where("uid=?", uid).Update(sm)
	} else {
		row, err = g.Id(id).Where("status=?", 2).Where("uid=?", uid).Update(sm)
	}

	if err != nil {
		setting.Logger.Infof("FeedbackMessageModel_ReadMessage_%s", err.Error())
	}
	return row, err
}
