package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/service/student_flag"
	"milano.gaodun.com/pkg/setting"
)

type VerifyFlag struct {
	Base
}

var VerifyF = NewVerifyFlag()

func NewVerifyFlag() *VerifyFlag {
	return &VerifyFlag{}
}

func (v *VerifyFlag) VerifyFlag( cxt *gin.Context) {
	flag := v.GetString(cxt, "flag", true)
	isHandle := v.GetString(cxt, "is_handle")
	if isHandle == "no" {
		return
	}

	if cxt.GetBool(Verify) {
		return
	}

	s := student_flag.NewStudentFlag(flag, setting.GinLogger(cxt))
	d := s.VerifyFlag()
	if s.GetErr() != nil {
		v.ServerJSONOther(cxt, d, s.GetErrCode(), s.GetErr().Error())
		cxt.Abort()
	}

}
