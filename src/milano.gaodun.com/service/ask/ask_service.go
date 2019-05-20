package ask

import (
	"fmt"
	"github.com/apex/log"
	simpleJson "github.com/bitly/go-simplejson"
	serialize "github.com/techoner/gophp/serialize"
	"milano.gaodun.com/model/api/ask"
	upApi "milano.gaodun.com/model/api/upload"
	userApi "milano.gaodun.com/model/api/user"
	askM "milano.gaodun.com/model/ask"
	"milano.gaodun.com/pkg/utils"
	"strings"
	"time"
)

type AskService struct {
	aa     *ask.AskApi
	up     *upApi.UploadApi
	user   *userApi.UserApi
	am     *askM.AskQuestionsModel
	askm    *askM.AskModel
	logger *log.Entry
}
type AskView struct {
	AskList []ask.Ask `json:"questions"`
	Page    string    `json:"page"`
	Limit   string    `json:"ps"`
	Total   string    `json:"total"`
}
type SyllabusItem struct {
	List map[string]Item
}
type Item struct {
	Id       string
	Name     string
	ParentId string
}
type AskDetailResp struct {
	Id            string      `json:"id"`
	CourseId      string      `json:"course_id"`
	StudentId     string      `json:"student_id"`
	CourseName    string      `json:"course_name"`
	SourceName    string      `json:"source_name"`
	Avatar        string      `json:"avatar"`
	ResourcesId   string      `json:"resources_id"`
	Content       string      `json:"content"`
	Img           interface{} `json:"img"`
	NickName      string      `json:"nick_name"`
	VideoTime     string      `json:"video_time"`
	OfficialReply Reply       `json:"officialreply"`
	AskRegdate    string      `json:"ask_regdate"`
}
type Reply struct {
	Id        string        `json:"id"`
	Content   string        `json:"content"`
	Regdate   string        `json:"regdate"`
	StudentId string        `json:"student_id"`
	Status    string        `json:"status"`
	Zhuiwen   []AnswerReAsk `json:"zhuiwen"`
	Iszhui    string        `json:"iszhui"`
	Nickname  string        `json:"nickname"`
	Avatar    string        `json:"avatar"`
	AskId     string        `json:"ask_id"`
}
type AnswerReAsk struct {
	Id        string      `json:"id"`
	Content   string      `json:"content"`
	Regdate   string      `json:"regdate"`
	StudentId string      `json:"student_id"`
	Status    string      `json:"status"`
	Isteach   string      `json:"isteach"`
	Img       interface{} `json:"img"`
	Nickname  string      `json:"nickname"`
	Avatar    string      `json:"avatar"`
}

func NewAskService(logger *log.Entry) *AskService {
	h := utils.HttpHandle
	var g AskService
	g.logger = logger
	g.aa = ask.NewAskApi(h)
	g.up = upApi.NewUploadApi(h)
	g.am = askM.NewAskQuestionsModel()
	g.askm =askM.NewAskModel()
	g.user = userApi.NewUserApi(h)
	return &g
}
func (g *AskService) GetAskDetail(askId int64) (AskDetailResp, error) {
	resp := AskDetailResp{}
	as, err := g.aa.GetAskDetail(askId)
	if err != nil {
		return resp, err
	}
	studentIds := ""
	asJson, err := simpleJson.NewJson([]byte(as))
	questions := asJson.Get("result")
	resp.Id, _ = questions.Get("id").String()
	resp.CourseId, _ = questions.Get("course_id").String()
	resp.StudentId, _ = questions.Get("student_id").String()
	if resp.StudentId != "" {
		studentIds = resp.StudentId + ","
	}
	resp.CourseName, _ = questions.Get("course_name").String()
	resp.SourceName, _ = questions.Get("coursewarepart_name").String()
	resp.ResourcesId, _ = questions.Get("resources_id").String()
	resp.Content, _ = questions.Get("content").String()
	resp.VideoTime, _ = questions.Get("video_time").String()
	resp.AskRegdate, _ = questions.Get("regdate").String()
	fileUrl, _ := questions.Get("file_url").String()
	resp.Img, _ = serialize.UnMarshal([]byte(fileUrl))
	resp.OfficialReply.Id, _ = questions.Get("answer").Get("id").String()
	resp.OfficialReply.Content, _ = questions.Get("answer").Get("content").String()
	resp.OfficialReply.Status, _ = questions.Get("answer").Get("status").String()
	resp.OfficialReply.Regdate, _ = questions.Get("answer").Get("regdate").String()
	resp.OfficialReply.StudentId, _ = questions.Get("answer").Get("student_id").String()
	if resp.OfficialReply.StudentId != "" {
		studentIds = studentIds + resp.OfficialReply.StudentId + ","
	}
	resp.OfficialReply.AskId, _ = questions.Get("answer").Get("ask_id").String()
	traceList := questions.Get("answer").Get("trace_list")
	if len(traceList.MustArray()) > 0 {
		resp.OfficialReply.Iszhui = "1"
	} else {
		resp.OfficialReply.Iszhui = "0"
	}
	for i := 0; i < len(traceList.MustArray()); i++ {
		uid, _ := traceList.GetIndex(i).Get("student_id").String()
		if uid != "" {
			studentIds = studentIds + uid + ","
		}
	}
	if studentIds != "" {
		studentIds = studentIds[0 : len(studentIds)-1]

	}
	userList, _ := g.user.GetUserInfo(studentIds)
	resp.NickName = userList[resp.StudentId].NickName
	resp.Avatar = utils.GetAvatarUrl(userList[resp.StudentId].PictureUrl)
	resp.OfficialReply.Nickname = userList[resp.OfficialReply.StudentId].NickName
	resp.OfficialReply.Avatar = utils.GetAvatarUrl(userList[resp.OfficialReply.StudentId].PictureUrl)
	for i := 0; i < len(traceList.MustArray()); i++ {
		one := AnswerReAsk{}
		one.Id, _ = traceList.GetIndex(i).Get("id").String()
		one.Content, _ = traceList.GetIndex(i).Get("content").String()
		one.Regdate, _ = traceList.GetIndex(i).Get("regdate").String()
		one.StudentId, _ = traceList.GetIndex(i).Get("student_id").String()
		one.Nickname = userList[one.StudentId].NickName
		one.Avatar = utils.GetAvatarUrl(userList[one.StudentId].PictureUrl)
		one.Status, _ = traceList.GetIndex(i).Get("status").String()
		one.Isteach, _ = traceList.GetIndex(i).Get("isteach").String()
		zhuiwenUrl, _ := traceList.GetIndex(i).Get("file_url").String()
		one.Img, _ = serialize.UnMarshal([]byte(zhuiwenUrl))
		resp.OfficialReply.Zhuiwen = append(resp.OfficialReply.Zhuiwen, one)
	}
	return resp, err
}
func (g *AskService) AnswerAsk(askQuestion *askM.GdAskQuestions) (int64, error) {
	askQuestion.Regdate = time.Now().Unix()
	askQuestion.Modifydate = time.Now().Unix()
	if askQuestion.FileUrl != "" {
		var imgUrlList []map[string]string
		fileList := strings.Split(askQuestion.FileUrl, ",")
		for _, v := range fileList {
			oneImg := map[string]string{}
			oneImg["o"] = v
			oneImg["t"] = v
			imgUrlList = append(imgUrlList, oneImg)
		}
		imgByte, _ := serialize.Marshal(imgUrlList)
		if len(imgUrlList) > 0 {
			askQuestion.FileUrl = string(imgByte)
		}
	}
	num, err := g.am.Add(askQuestion)
	ask := askM.GdAsk{Id:askQuestion.AskId,Status:3}
	g.askm.Edit(&ask)
	return num, err
}
func (g *AskService) GetAskList(params *ask.AskParam) (*AskView, error) {
	askView := AskView{AskList: []ask.Ask{}}
	//获取答疑列表
	as, err := g.aa.GetAskList(params)
	asJson, err := simpleJson.NewJson([]byte(as))
	if err != nil {
		g.logger.Error(err.Error())
		return &askView, err
	}
	userInfo, _ := g.user.GetUserInfo(fmt.Sprintf("%d", params.Uid))
	questions := asJson.Get("result").Get("questions")
	askView.Page, _ = asJson.Get("result").Get("page").String()
	askView.Limit, _ = asJson.Get("result").Get("ps").String()
	askView.Total, _ = asJson.Get("result").Get("total").String()
	for i := 0; i < len(questions.MustArray()); i++ {
		one := ask.Ask{}
		one.Id, _ = questions.GetIndex(i).Get("id").String()
		one.CourseId, _ = questions.GetIndex(i).Get("course_id").String()
		one.CourseName, _ = questions.GetIndex(i).Get("course_name").String()
		one.ResourceId, _ = questions.GetIndex(i).Get("resources_id").String()
		one.ResourceName, _ = questions.GetIndex(i).Get("coursewarepart_name").String()
		one.StudentId, _ = questions.GetIndex(i).Get("student_id").String()
		one.Picture = utils.GetAvatarUrl(userInfo[one.StudentId].PictureUrl)
		one.NickName = userInfo[one.StudentId].NickName
		one.Content, _ = questions.GetIndex(i).Get("content").String()
		one.SourceType, _ = questions.GetIndex(i).Get("source_type").String()
		one.VideoTime, _ = questions.GetIndex(i).Get("video_time").String()
		one.CourseSyllabusItemId, _ = questions.GetIndex(i).Get("entry_id").String()
		one.CreatedAt,_ = questions.GetIndex(i).Get("regdate").String()
		one.Reply.Id, _ = questions.GetIndex(i).Get("reply").Get("id").String()
		one.Reply.Content, _ = questions.GetIndex(i).Get("reply").Get("content").String()
		one.Reply.StudentId, _ = questions.GetIndex(i).Get("reply").Get("student_id").String()
		one.Reply.CreatedAt, _ = questions.GetIndex(i).Get("reply").Get("regdate").String()
		one.Reply.UpdatedAt, _ = questions.GetIndex(i).Get("reply").Get("modifydate").String()
		urlStr, _ := questions.GetIndex(i).Get("file_url").String()
		one.Images = g.SerializeToStringSlice(urlStr)
		replyurlStr, _ := questions.GetIndex(i).Get("reply").Get("file_url").String()
		one.Reply.Images = g.SerializeToStringSlice(replyurlStr)
		askView.AskList = append(askView.AskList, one)
	}
	return &askView, err
}
func (g *AskService) SerializeToStringSlice(serializeStr string) []string {
	result := []string{}
	if serializeStr == "" {
		return result
	}
	fileList, _ := serialize.UnMarshal([]byte(serializeStr))
	fl := fileList.([]interface{})
	for _, v := range fl {
		mp := v.(map[string]interface{})
		result = append(result, mp["o"].(string))
	}
	return result
}
