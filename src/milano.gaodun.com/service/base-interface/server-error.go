package base_interface

import (
	"github.com/apex/log"
	"milano.gaodun.com/pkg/error-code"
)

type ServiceError interface {
	GetErrCode() error_code.CodeTypeInt
	GetErr() error
}

type ServiceErr struct {
	errorCode error_code.CodeTypeInt
	err       error
	L         *log.Entry
}

// 检查 error
func (s *ServiceErr) CheckErr(err error, errCode error_code.CodeTypeInt) bool {
	if err != nil {
		s.L.Info(err.Error())
		s.errorCode = errCode
		s.err = err
		return true
	}

	return false
}

// get error
func (s *ServiceErr) GetErr() error {
	return s.err
}

// get error code
func (s *ServiceErr) GetErrCode() error_code.CodeTypeInt {
	return s.errorCode
}
