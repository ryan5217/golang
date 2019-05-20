package wx_program

// 微信素材文章

// api host domain
const (
	WechatHost = "https://api.weixin.qq.com"
)

// redis key
const (
	AccessTokenKey = "wechat_access_token"
)

// 错误消息
type WxResp struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

// token
type AccessToken struct {
	WxResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// 列表
type MaterialList struct {
	WxResp
	TotalCount int        `json:"total_count"`
	ItemCount  int        `json:"item_count"`
	Item       []ItemList `json:"item"`
}

// item list
type ItemList struct {
	MediaId    string  `json:"media_id"`
	Content    Content `json:"content"`
	UpdateTime int     `json:"update_time"`
}

// content
type Content struct {
	CreateTime int           `json:"create_time"`
	UpdateTime int           `json:"update_time"`
	NewsItem   []NewItemList `json:"news_item"`
}

// news item
type NewItemList struct {
	Title        string `json:"title"`
	ThumbMediaId string `json:"thumb_media_id"`
	Digest       string `json:"digest"`
	Url          string `json:"url"`
	ThumbUrl     string `json:"thumb_url"`
	Author       string `json:"author"`
}

/**
total_count	该类型的素材的总数
item_count	本次调用获取的素材的数量
title	图文消息的标题
thumb_media_id	图文消息的封面图片素材id（必须是永久mediaID）
thumb_url  	图片 url
author	作者
digest	图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
// content	图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
url	图文页的URL，或者，当获取的列表是图片素材列表时，该字段是图片的URL
content_source_url	图文消息的原文地址，即点击“阅读原文”后的URL
update_time	这篇图文消息素材的最后更新时间
name	文件名称
*/
