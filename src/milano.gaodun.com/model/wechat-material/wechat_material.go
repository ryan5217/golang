package wechat_material

import (
	"milano.gaodun.com/model/xorm-model"
	"milano.gaodun.com/pkg/utils"
)

type PriWechatMaterial struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	MediaId      string `json:"media_id"`
	Digest       string `json:"digest"`
	ThumbMediaId string `json:"thumb_media_id"`
	ThumbUrl     string `json:"thumb_url"`
	ThumbPicUrl  string `json:"thumb_pic_url"`
	Url          string `json:"url"`
	Weight       int    `json:"-"`
	UpdateTime   int    `json:"update_time"`
	CreatedAt    string `json:"created_at" xorm:"<-"`
}

type WechatMaterialModel struct {
	xorm_model.ModelBase
}

func NewWechatMaterialModel() *WechatMaterialModel {
	w := new(WechatMaterialModel)
	w.Engine = utils.GaodunPrimaryDb
	w.NewModelSession()
	return w
}
