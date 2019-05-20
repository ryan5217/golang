package controller

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/service/invitation"
)

var InvitationApi = NewInvitationApi()

func NewInvitationApi() *Invitation {
	return &Invitation{}
}

type Invitation struct {
	Base
}

func (i Invitation) GetCode(c *gin.Context) {
	studentId := i.GetInt64(c, "student_id")
	bServer := invitation.NewInvitationService(setting.GinLogger(c))
	r := bServer.GetCode(studentId)
	i.ServerJSONSuccess(c, r)
}
func (i Invitation) GetPhoneCode(c *gin.Context) {
	phone := i.GetString(c, "phone")
	bServer := invitation.NewInvitationService(setting.GinLogger(c))
	r := bServer.GetPhoneCode(phone)
	if r.Status == 0 {
		i.ServerJSONSuccess(c, r)
	} else {
		i.ServerJSONError(c, r, int(r.Status))
	}
}
func (i Invitation) GetUuid(c *gin.Context) {
	bServer := invitation.NewInvitationService(setting.GinLogger(c))
	r := bServer.GetUuid()
	if r.Code == 11999999 {
		i.ServerJSONSuccess(c, r)
	} else {
		i.ServerJSONError(c, r, error_code.HTTPERR)
	}

}
func (i Invitation) AddInvitation(c *gin.Context) {
	phone := i.GetString(c, "phone", true)
	phoneCode := i.GetString(c, "phone_code", true)
	flag := i.GetString(c, "flag", true)
	invitationCode := i.GetString(c, "invitation_code", true)
	if c.GetBool(Verify) {
		return
	}
	clientIp := c.ClientIP()
	bServer := invitation.NewInvitationService(setting.GinLogger(c))
	r := bServer.AddInvitation(phone,phoneCode,flag,invitationCode,clientIp)
	if(r.Status == 0){
		i.ServerJSONSuccess(c, "")
	} else {
		i.ServerJSONError(c,r,int(r.Status))
	}
}
func (i Invitation) InvitationList(c *gin.Context) {
	studentId := i.GetInt64(c, "student_id")
	bServer := invitation.NewInvitationService(setting.GinLogger(c))
	r := bServer.InvitationList(studentId)
	i.ServerJSONSuccess(c, r)
}
