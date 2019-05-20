package message

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	"time"
)

type TkSystemMessage struct {
	Id             int64
	Title          string
	Url            string
	Uid            int64
	ProjectId      int64
	UpdatedTime    string
	CreatedTime    string
	Isdel          int      `json:"-"`
	ForceUpdateCol []string `xorm:"-" json:"-"`
}

type SystemMessageModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewSystemMessageModel() *SystemMessageModel {
	return &SystemMessageModel{
		Engine: utils.TikuDbo,
		s:      utils.TikuDbo.NewSession(),
	}
}

type Param struct {
	Uid       int64
	ProjectId int64
	Type      int
	Page      int
	Limit     int
}

func (g *SystemMessageModel) GetSystemMessages(param Param) ([]TkSystemMessage, error) {
	messages := []TkSystemMessage{}
	engine := g.s
	var err error
	if param.Uid > 0 {
		engine.In("uid", []int64{param.Uid, 0})
	}
	if param.Limit > 0 {
		err = engine.Where("project_id=?", param.ProjectId).Where("isdel=?", 0).OrderBy("id desc").Limit(param.Limit, (param.Page-1)*param.Limit).Find(&messages)
	} else {
		err = engine.Where("project_id=?", param.ProjectId).Where("isdel=?", 0).OrderBy("id desc").Find(&messages)
	}
	if err != nil {
		setting.Logger.Infof("SystemMessageModel_GetMessages_%d_%s", param.ProjectId, err.Error())
	}
	return messages, err
}
func (g *SystemMessageModel) GetSystemMessageCount(param Param) (int, error) {
	message := TkSystemMessage{}
	engine := g.s
	var err error
	var num int64
	if param.Uid > 0 {
		engine.In("uid", []int64{param.Uid, 0})
	}
	num, err = engine.Where("project_id=?", param.ProjectId).Where("isdel=?", 0).Count(&message)
	if err != nil {
		setting.Logger.Infof("SystemMessageModel_GetSystemMessageCount_%d_%s", param.ProjectId, err.Error())
	}
	return int(num), err
}
func (g *SystemMessageModel) Modify(ts *TkSystemMessage) (int64, error) {
	var row int64
	var err error
	if ts.Id > 0 {
		ts.UpdatedTime = time.Now().String()
		row, err = g.Id(ts.Id).Cols(ts.ForceUpdateCol...).Update(ts)

	} else {
		var cstSh, _ = time.LoadLocation("Asia/Shanghai")
		ts.UpdatedTime = time.Now().In(cstSh).Format("2006-01-02 15:04:05")
		row, err = g.Omit("created_time").Omit("updated_time").InsertOne(ts)
	}
	if err != nil {
		setting.Logger.Infof("SystemMessageModel_AddMessage_%s", err.Error())
	}
	return row, err
}
func (g *SystemMessageModel) DeleteMessage(id int64, isdel int) (int64, error) {
	var row int64
	var err error
	sm := TkSystemMessage{Isdel: isdel}
	row, err = g.Id(id).Cols("isdel").Update(sm)
	if err != nil {
		setting.Logger.Infof("SystemMessageModel_delete_%s", err.Error())
	}
	return row, err
}
