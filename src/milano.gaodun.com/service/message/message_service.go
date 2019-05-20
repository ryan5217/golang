package message

import (
	"github.com/apex/log"
	"milano.gaodun.com/conf"
	mm "milano.gaodun.com/model/message"
	"fmt"
)

type MessageServiceInterface interface {
	ModifySystemMessage(*mm.TkSystemMessage) (int64, error)
	GetSystemMessages(param mm.Param) (*SystemMessageResp, error)
	List(param mm.Param) (*MessageResp, error)
	DeleteSystemMessage(id int64, isdel int) (int64, error)
	DeleteNotifyMessage(id int64, uid int64, msgType int64) (int64, error)
	GetNotifyNum(param mm.Param) (int, error)
	ReadMessage(id int64, uid int64, msgType int64) (int64, error)
}

type MessageParam struct {
	Id        int64
	Url       string
	ProjectId int64
}
type MessageService struct {
	msg    *mm.SystemMessageModel
	msgr   *mm.SystemMessageRelationModel
	fmsg   *mm.FeedbackMessageModel
	logger *log.Entry
}
type MessageResp struct {
	Page    int       `json:"page"`
	Limit   int       `json:"limit"`
	ALlNum  int       `json:"all_num"`
	AllPage int       `json:"all_page"`
	List    []Message `json:"list"`
}
type SystemMessageResp struct {
	Page    int                  `json:"page"`
	Limit   int                  `json:"limit"`
	ALlNum  int                  `json:"all_num"`
	AllPage int                  `json:"all_page"`
	List    []mm.TkSystemMessage `json:"list"`
}
type Message struct {
	Id          int64  `json:"id"`
	Content     string `json:"content"`
	RetContent  string `json:"ret_content"`
	Url         string `json:"url"`
	Pic         string `json:"pic"`
	IsRead      int    `json:"is_read"`
	UpdatedTime string `json:"updated_time"`
	CreatedTime string `json:"created_time"`
}

func NewMessageService(logger *log.Entry) MessageServiceInterface {

	return &MessageService{
		msg:    mm.NewSystemMessageModel(),
		msgr:   mm.NewSystemMessageRelationModel(),
		fmsg:   mm.NewFeedbackMessageModel(),
		logger: logger,
	}
}

func (g *MessageService) ModifySystemMessage(msg *mm.TkSystemMessage) (int64, error) {
	return g.msg.Modify(msg)
}
func (g *MessageService) GetSystemMessages(param mm.Param) (*SystemMessageResp, error) {
	msgResp := SystemMessageResp{List: []mm.TkSystemMessage{}}
	msgs, err := g.msg.GetSystemMessages(param)
	if err != nil {
		return &msgResp, err
	}
	msgResp.Page = param.Page
	msgResp.Limit = param.Limit
	msgResp.ALlNum, err = g.msg.GetSystemMessageCount(param)
	msgResp.AllPage = msgResp.ALlNum / msgResp.Limit
	if msgResp.ALlNum%msgResp.Limit > 0 {
		msgResp.AllPage++
	}
	msgResp.List = msgs
	return &msgResp, err
}
func (g *MessageService) DeleteSystemMessage(id int64, isdel int) (int64, error) {
	return g.msg.DeleteMessage(id, isdel)
}
func (g *MessageService) DeleteNotifyMessage(id int64, uid int64, msgType int64) (int64, error) {
	if msgType == 1 {
		return g.fmsg.DeleteMessage(id, uid)
	} else {
		return g.msgr.DeleteRelation(id, uid)
	}
}
func (g *MessageService) ReadMessage(id int64, uid int64, msgType int64) (int64, error) {
	if msgType == 1 {
		return g.fmsg.ReadMessage(id, uid)
	} else {
		return g.msgr.ReadMessage(id, uid)
	}
}
func (g *MessageService) GetNotifyNum(param mm.Param) (int, error) {
	fmsgNum, err := g.fmsg.GetUnreadFeedbackMessageCount(param)
	if err != nil {
		return 0, err
	}
	param.Page = 1
	param.Limit = 20
	rmsg, err := g.List(param)
	if err != nil {
		return 0, err
	}
	for _, v := range rmsg.List {
		if v.IsRead == 0 {
			fmsgNum++
		}
	}
	return fmsgNum, nil
}
func (g *MessageService) List(param mm.Param) (*MessageResp, error) {
	tmpMsgList := []Message{}
	resp := MessageResp{List: tmpMsgList}
	resp.Page = param.Page
	resp.Limit = param.Limit
	if param.Type == 0 {
		//获取系统消息列表
		paramAll := param
		paramAll.Limit = 0
		msgList, err := g.msg.GetSystemMessages(paramAll)
		if err != nil {
			return &resp, err
		}
		//获取学员关联系统消息列表
		relationList, err := g.msgr.GetSystemMessageRelations(param.Uid)
		if err != nil {
			return &resp, err
		}
		relationMap := map[int64]mm.TkSystemMessageRelation{}
		for _, v := range relationList {
			relationMap[v.MessageId] = v
		}
		//分页获取学员关联系统消息内容，没有关联则创建关联
		for _, v := range msgList {
			r, ok := relationMap[v.Id]
			if ok {
				if r.Isdel == 0 && v.Isdel == 0 {
					msg := Message{
						Id:          v.Id,
						Content:     v.Title,
						Url:         v.Url,
						IsRead:      r.IsRead,
						CreatedTime: v.CreatedTime,
						UpdatedTime: v.UpdatedTime,
					}
					tmpMsgList = append(tmpMsgList, msg)
				}
			} else {
				if v.Isdel == 0 {
					msg := Message{
						Id:          v.Id,
						Content:     v.Title,
						Url:         v.Url,
						IsRead:      r.IsRead,
						CreatedTime: v.CreatedTime,
						UpdatedTime: v.UpdatedTime,
					}
					tmpMsgList = append(tmpMsgList, msg)
					sr := mm.TkSystemMessageRelation{
						Uid:       param.Uid,
						MessageId: v.Id,
						IsRead:    0,
						Isdel:     0,
					}
					g.msgr.Add(&sr)
				}
			}
		}
		start := (param.Page - 1) * param.Limit
		end := start + param.Limit
		for k, v := range tmpMsgList {
			if k >= start && k < end {
				resp.List = append(resp.List, v)
			}
		}
		resp.ALlNum = len(tmpMsgList)
		resp.AllPage = len(tmpMsgList) / resp.Limit
		if len(tmpMsgList)%resp.Limit > 0 {
			resp.AllPage++
		}
	} else {
		fmsgs, err := g.fmsg.GetFeedbackMessages(param)
		if err != nil {
			return &resp, nil
		}
		resp.ALlNum, err = g.fmsg.GetFeedbackMessageCount(param)
		if err != nil {
			return &resp, nil
		}
		for _, v := range fmsgs {
			msg := Message{
				Id:          v.Id,
				Content:     v.Content,
				RetContent:  v.RetContent,
				Url:         conf.MUSES_DOMAIN + fmt.Sprintf("/feedback/list/%d",param.ProjectId),
				Pic:         v.Pic,
				IsRead:      v.IsRead,
				CreatedTime: v.CreatedTime,
				UpdatedTime: v.UpdatedTime,
			}
			resp.List = append(resp.List, msg)
		}
		resp.AllPage = len(fmsgs) / resp.Limit
		if len(fmsgs)%resp.Limit > 0 {
			resp.AllPage++
		}
	}

	return &resp, nil
}
