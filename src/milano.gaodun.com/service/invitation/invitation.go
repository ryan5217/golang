package invitation

import (
	"github.com/apex/log"
	"milano.gaodun.com/model/invitaion"
	//"encoding/json"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/utils"
)

type InvitationServiceInterface interface {
	GetCode(studentId int64) *invitaion.PriInvitationCode
	AddInvitation(phone string, phoneCode string, flag string, invitationCode string, clientIp string) *invitaion.PriInvitationList
	InvitationList(studentId int64) *[]invitaion.PriInvitationList
	GetUuid() *invitaion.Uuid
	GetPhoneCode(phone string) *invitaion.UserAccountInfoResApi
}

type InvitationService struct {
	InviteCodeM *invitaion.InvitationCodeModel
	InviteListM *invitaion.InvitationlistModel
	logger      *log.Entry
}

func NewInvitationService(logger *log.Entry) InvitationServiceInterface {
	return &InvitationService{InviteCodeM: invitaion.NewInvitationCodeModel(), InviteListM: invitaion.NewInvitationlistModel(), logger: logger}
}
func (i *InvitationService) GetUuid() *invitaion.Uuid {
	row, err := i.InviteCodeM.GetUuid()
	if err != nil {
		i.logger.Error(err.Error())
	}
	return row
}

// api 数据返回
func (i *InvitationService) GetPhoneCode(phone string) *invitaion.UserAccountInfoResApi {
	r, err := i.InviteCodeM.GetUserInfoByAccount(phone)
	if err != nil {
		i.logger.Error(err.Error())
	}
	//手机号已经注册
	if r.Status == 0 {
		r.Status = error_code.PHONEREGERR
		errorInfo := error_code.GetInfo()
		r.Info = errorInfo[error_code.PHONEREGERR]
	}
	//手机号尚未注册，发送验证码
	if r.Status == 1610614548 {
		s, err := i.InviteCodeM.SendPhoneCode(phone)
		if err != nil {
			i.logger.Error(err.Error())
		}
		r.Status = s.Status
		r.Info = s.Info
		if s.Status != 0 {
			r.Status = error_code.SENDPHONECODEERR
		}
	}

	return r
}

// api 数据返回
func (i *InvitationService) GetCode(studentId int64) *invitaion.PriInvitationCode {
	row, err := i.InviteCodeM.GetCode(studentId)
	if err != nil {
		i.logger.Error(err.Error())
	}
	if row.InvitationCode == "" {
		row := invitaion.PriInvitationCode{StudentId: studentId}
		row.InvitationCode = utils.GetRandomString(5)
		i.InviteCodeM.Add(&row)
		newCode, err := i.InviteCodeM.GetCode(studentId)
		if err != nil {
			i.logger.Error(err.Error())
		}
		return newCode
	}
	return row
}
func (i *InvitationService) InvitationList(studentId int64) *[]invitaion.PriInvitationList {
	result, err := i.InviteListM.GetInvitationList(studentId)
	if err != nil {
		i.logger.Error(err.Error())
	}

	return &result
}

func (i *InvitationService) AddInvitation(phone string, phoneCode string, flag string, invitationCode string, clientIp string) *invitaion.PriInvitationList {
	row, err := i.InviteCodeM.GetInvitCode(invitationCode)
	var result invitaion.PriInvitationList
	//查询推荐码报错
	if err != nil {
		i.logger.Error(err.Error())
		result.Status = error_code.DBERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.DBERR]
		return &result
	}
	//邀请码错误
	if invitationCode != row.InvitationCode {
		result.Status = error_code.INVITERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.INVITERR]
		return &result
	}
	//推荐人数达到上限
	if row.InvitedCount >= 20 {
		result.Status = error_code.INVITMAX
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.INVITMAX]
		return &result
	}
	weixinMsg, _ := i.InviteCodeM.WeixinVerify(flag)
	//微信登录校验
	if weixinMsg.Code != 11999999 {
		result.Status = error_code.WEIXINERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.WEIXINERR]
		return &result
	}
	phoneCodeInfo, err := i.InviteCodeM.VerifyPhoneCode(phone, phoneCode)
	//手机号验证码校验失败
	if phoneCodeInfo.Status != 0 {
		result.Status = error_code.PHONECODEERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.PHONECODEERR]
		return &result
	}
	studentRegInfo, err := i.InviteCodeM.GetUserInfoByPhone(phone, flag, clientIp)
	//获取用户信息失败
	if studentRegInfo.Status != 0 {
		result.Status = error_code.USERINFOERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.USERINFOERR]
		return &result
	}

	invitedStudent := studentRegInfo.Data.StudentId
	studentInfo, err := i.InviteCodeM.GetUserInfo(invitedStudent)
	//查询用户信息报错
	if err != nil {
		i.logger.Error(err.Error())
		result.Status = error_code.HTTPERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.HTTPERR]
		return &result
	}
	//查询用户信息报错
	if studentInfo.Data["UserData"][0].Status != 0 {
		result.Status = error_code.USERINFOERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.USERINFOERR]
		return &result
	}
	//无法邀请自己
	if invitedStudent == row.StudentId {
		result.Status = error_code.SELFINVIT
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.SELFINVIT]
		return &result
	}
	result, err = i.InviteListM.GetInvit(invitedStudent)
	//查询推荐人出错
	if err != nil {
		i.logger.Error(err.Error())
		return &result
	}
	//推荐人已添加，无需重复添加
	if result.Id != 0 {
		result.Status = error_code.INVITDUPLICATE
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.INVITDUPLICATE]
		return &result
	}
	result.StudentId = row.StudentId
	result.InvitedStudent = invitedStudent
	s := i.InviteListM.NewSession()
	defer s.Close()
	s.Begin()
	s.InsertOne(&result)
	//查询推荐人报错
	if err != nil {
		i.logger.Error(err.Error())
		result.Status = error_code.DBERR
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.DBERR]
		s.Rollback()
		return &result
	} else {
		_, err = i.InviteCodeM.IncCount(row.StudentId)
		if err != nil {
			i.logger.Error(err.Error())
			result.Status = error_code.DBERR
			errorInfo := error_code.GetInfo()
			result.Message = errorInfo[error_code.DBERR]
			s.Rollback()
			return &result
		}
		s.Commit()
		result, err = i.InviteListM.GetInvit(invitedStudent)
		result.Status = error_code.SUCCESSSTATUS
		errorInfo := error_code.GetInfo()
		result.Message = errorInfo[error_code.SUCCESSSTATUS]
		return &result
	}
}
