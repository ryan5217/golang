package message

import (
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type TkSystemMessageRelation struct {
	Uid       int64
	MessageId int64
	IsRead    int
	Isdel int
}

type SystemMessageRelationModel struct {
	*xorm.Engine
	s *xorm.Session
}

func NewSystemMessageRelationModel() *SystemMessageRelationModel {
	return &SystemMessageRelationModel{Engine: utils.TikuDbo}
}

func (g *SystemMessageRelationModel) GetSystemMessageRelations(uid int64) ([]TkSystemMessageRelation, error) {
	messages := []TkSystemMessageRelation{}
	var err error
	err = g.Where("uid=?", uid).Find(&messages)
	if err != nil {
		setting.Logger.Infof("GetSystemMessageRelations%d_%s", uid, err.Error())
	}
	return messages, err
}
func (g *SystemMessageRelationModel) Add(ts *TkSystemMessageRelation) (int64, error) {
	row, err := g.InsertOne(ts)
	if err != nil {
		setting.Logger.Infof("SystemMessageRelationModel_Add_%s", err.Error())
	}
	return row, err
}
func (g *SystemMessageRelationModel) DeleteRelation(msgId int64,uid int64) (int64, error) {
	var row int64
	var err error
	sm := TkSystemMessageRelation{Isdel:1}
	if msgId >0 {
		row, err = g.Where("uid=?",uid).Where("message_id=?",msgId).Update(sm)
	} else {
		row, err = g.Where("uid=?",uid).Update(sm)
	}

	if err != nil {
		setting.Logger.Infof("SystemMessageRelationModel_delete_%s", err.Error())
	}
	return row, err
}
func (g *SystemMessageRelationModel) ReadMessage(msgId int64,uid int64) (int64, error) {
	var row int64
	var err error
	sm := TkSystemMessageRelation{IsRead:1}
	if msgId >0 {
		row, err = g.Where("uid=?",uid).Where("message_id=?",msgId).Update(sm)
	} else {
		row, err = g.Where("uid=?",uid).Update(sm)
	}

	if err != nil {
		setting.Logger.Infof("SystemMessageRelationModel_ReadMessage_%s", err.Error())
	}
	return row, err
}