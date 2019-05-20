package student_flag

import (
	"milano.gaodun.com/model/invitaion"
	"milano.gaodun.com/service/base-interface"
	"milano.gaodun.com/pkg/error-code"
	"errors"
	"github.com/apex/log"
)

type StudentFlagInterface interface {
	base_interface.ServiceError
	VerifyFlag() *invitaion.WeixinData
}

type StudentFlag struct {
	flag string
	InviteCodeM *invitaion.InvitationCodeModel
	base_interface.ServiceErr
}

func NewStudentFlag(f string, logger *log.Entry) *StudentFlag {
	s := StudentFlag{}
	s.flag = f
	s.L = logger
	s.InviteCodeM = invitaion.NewInvitationCodeModel()

	return &s
}

// 验证 flag 是否合法
func (s *StudentFlag) VerifyFlag() *invitaion.WeixinData {
	res, err := s.InviteCodeM.WeixinVerify(s.flag)
	s.CheckErr(err, error_code.FAIL)

	if res.Code != 11999999 {
		s.CheckErr(errors.New(res.Message), error_code.CodeTypeInt(res.Code))
	}

	return &res.Data
}



