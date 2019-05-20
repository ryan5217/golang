package invitaion

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type Uuid struct {
	Code    int64       `json:"code"`
	Message int64       `json:"message"`
	Data    interface{} `json:"data"`
}
type PriInvitationCode struct {
	Id             int64
	StudentId      int64
	InvitationCode string
	InvitedCount   int64
	UsedInvitTimes int64
}
type UserInfoResApi struct {
	Info   string                `json:"info"`
	Status int64                 `json:"status"`
	Data   map[string][]UserData `json:"data"`
}
type UserAccountInfoResApi struct {
	Info   string      `json:"info"`
	Status int64       `json:"status"`
	Data   UserAccount `json:"data"`
}
type UserAccount struct {
	UserId    int64 `json:"UserId"`
	StudentId int64 `json:"StudentId"`
}
type UserRegInfoResApi struct {
	Info   string      `json:"info"`
	Status int64       `json:"status"`
	Data   UserRegData `json:"data"`
}
type UserData struct {
	Uid          int64  `json:"Uid"`
	StudentId    int64  `json:"StudentId"`
	NickName     string `json:"NickName"`
	RealName     string `json:"RealName"`
	Email        string `json:"Email"`
	Phone        string `json:"Phone"`
	PictureUrl   string `json:"PictureUrl"`
	UserStatus   int64  `json:"UserStatus"`
	UserRegTime  int64  `json:"UserRegTime"`
	CreateBy     int64  `json:"CreateBy"`
	UserPartner  int64  `json:"UserPartner"`
	UserWxExtend int64  `json:"UserWxExtend"`
	Status       int64  `json:"Status"`
	Info         string `json:"Info"`
	ErrorCode    int64  `json:"ErrorCode"`
	LoginTime    int64  `json:"LoginTime"`
}
type UserRegData struct {
	UserId     int64  `json:"UserId"`
	StudentId  int64  `json:"StudentId"`
	NickName   string `json:"NickName"`
	PictureUrl string `json:"PictureUrl"`
}

type PhoneCodeResApi struct {
	Info   string `json:"info"`
	Status int64  `json:"status"`
	Data   string `json:"data"`
}
type WeixinVerifyResApi struct {
	Message string     `json:"message"`
	Code    int64      `json:"code"`
	Data    WeixinData `json:"data"`
}
type WeixinData struct {
	HeadUrl      string `json:"head_url"`
	HasPhone     int64  `json:"has_phone"`
	NickName     string `json:"nickname"`
	LoginSource  string `json:"login_source"`
	SalesmanId   int64  `json:"salesman_id"`
	LoginVersion int64  `json:"login_version"`
}
type SendPhoneCodeResApi struct {
	Info   string `json:"info"`
	Status int64  `json:"status"`
	Data   string `json:"data"`
}
type InvitationCodeModel struct {
	*xorm.Engine
	s   *xorm.Session
	pic PriInvitationCode
}

func NewInvitationCodeModel() *InvitationCodeModel {
	return &InvitationCodeModel{Engine: utils.GaodunPrimaryDb}
}
func (b *InvitationCodeModel) WeixinVerify(flag string) (*WeixinVerifyResApi, error) {
	var res WeixinVerifyResApi
	param := req.Param{
		"student_id": flag,
	}
	r, err := req.Post(conf.MTIKU_DOMAIN+"/api/member/getUserInfo", req.Header{}, param)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}

// 获取列表 按条件获取
func (b *InvitationCodeModel) GetCode(studentId int64) (*PriInvitationCode, error) {
	invitation_code := PriInvitationCode{}

	_, err := b.Where("student_id=?", studentId).Get(&invitation_code)
	return &invitation_code, err
}
func (b *InvitationCodeModel) GetInvitCode(invitationCode string) (*PriInvitationCode, error) {
	invitation_code := PriInvitationCode{}

	_, err := b.Where("invitation_code=?", invitationCode).Get(&invitation_code)
	return &invitation_code, err
}
func (b *InvitationCodeModel) IncCount(studentId int64) (int64, error) {
	return b.Where("student_id=?", studentId).Incr("invited_count", 1).Update(&PriInvitationCode{})

}
func (b *InvitationCodeModel) Add(invitC *PriInvitationCode) (int64, error) {
	row, err := b.InsertOne(invitC)
	return row, err
}
func (b *InvitationCodeModel) GetUserInfo(studentId int64) (*UserInfoResApi, error) {
	var res UserInfoResApi
	header := req.Header{"origin": "gaodun.com"}
	r, err := req.Get(conf.SSO_DOMAIN+fmt.Sprintf("/getbaseuserinfo/%d", studentId), header)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}
func (b *InvitationCodeModel) VerifyPhoneCode(phone string, phoneCode string) (*PhoneCodeResApi, error) {
	var res PhoneCodeResApi
	header := req.Header{"origin": "gaodun.com"}
	r, err := req.Get(conf.SSO_DOMAIN+fmt.Sprintf("/api/v1/verifyphonecode?phone=%s&code=%s", phone, phoneCode), header)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}
func (b *InvitationCodeModel) GetUuid() (*Uuid, error) {
	var res Uuid
	r, err := req.Post(conf.MTIKU_DOMAIN + "/api/member/getUuid")
	//fmt.Println(r)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}
func (b *InvitationCodeModel) GetUserInfoByPhone(phone string, flag string, clientIp string) (*UserRegInfoResApi, error) {
	var res UserRegInfoResApi
	header := req.Header{"origin": "gaodun.com"}
	//ip := req.
	param := req.Param{
		"phone":     phone,
		"appid":     conf.APP_ID,
		"sessionid": flag,
		"clientip":  clientIp,
	}
	r, err := req.Post(conf.SSO_DOMAIN+"/registerphone", header, param)
	//fmt.Println(r)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}
func (b *InvitationCodeModel) GetUserInfoByAccount(account string) (*UserAccountInfoResApi, error) {
	var res UserAccountInfoResApi
	header := req.Header{"origin": "gaodun.com"}
	r, err := req.Get(conf.SSO_DOMAIN+"/getdatabyaccount/"+account, header)
	//fmt.Println(r)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}
func (b *InvitationCodeModel) SendPhoneCode(phone string) (*SendPhoneCodeResApi, error) {
	var res SendPhoneCodeResApi
	header := req.Header{"origin": "gaodun.com"}
	r, err := req.Get(conf.SSO_DOMAIN+"/api/v1/sendphonecode?phone="+phone, header)
	//fmt.Println(r)
	if err == nil {
		err := r.ToJSON(&res)
		return &res, err
	} else {
		return &res, err
	}
}
