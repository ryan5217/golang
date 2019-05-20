package user

import (
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
	"strconv"
)

type UserApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type UserResp struct {
	Data   map[string][]UserData `json:"data"`
	Status int64                 `json:"status"`
	Info   string                `json:"info"`
}
type UserData struct {
	Uid         int    `json:"Uid"`
	StudentId   int    `json:"StudentId"`
	NickName    string `json:"NickName"`
	RealName    string `json:"RealName"`
	Phone       string `json:"Phone"`
	PictureUrl  string `json:"PictureUrl"`
	UserRegTime int64  `json:"UserRegTime"`
}

type UserDataOne struct {
	Data struct {
		UserId    int
		StudentId int
	} `json:"data"`
	IsNowReg int    `json:"is_now_reg"`
	Status   int64  `json:"status"`
	Info     string `json:"info"`
}

func NewUserApi(h *utils.HttpClient) *UserApi {
	var p = UserApi{}
	p.HttpClient = h
	return &p
}

func (g *UserApi) GetUserInfo(studentIds string) (map[string]UserData, error) {
	res := UserResp{}
	ulist := map[string]UserData{}
	g.Uri = conf.SSO_DOMAIN + "/getbaseuserinfo"
	header := req.Header{
		"ORIGIN":     "gaodun.com",
		"app_id":     conf.BaseCoinKey["app_id"],
		"app_secret": conf.BaseCoinKey["app_secret"],
	}
	r, err := req.Post(g.Uri, req.Param{"user_id": studentIds}, header)
	if err != nil {
		return ulist, err
	}
	err = r.ToJSON(&res)
	for _, v := range res.Data["UserData"] {
		if v.StudentId != 0 {
			ulist[strconv.Itoa(v.StudentId)] = v
		}

	}
	return ulist, err
}

// 手机号查询并注册
// 如果手机号已存在，则直接返回userid和studentid等基本信息，否则注册新账号（不发送短信）
func (g *UserApi) RegisterAndFindByPhone(phone, nickname, password string) (*UserDataOne, error) {
	res := UserDataOne{}
	g.Uri = conf.SSO_DOMAIN + "/registerphonenew"
	header := req.Header{
		"ORIGIN": "gaodun.com",
	}
	param := req.Param{}
	param["phone"] = phone
	param["nickname"] = nickname
	param["password"] = password
	r, err := req.Post(g.Uri, param, header)
	if err != nil {
		return &res, err
	}
	err = r.ToJSON(&res)
	return &res, err
}
