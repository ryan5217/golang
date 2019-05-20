package controller

import (
	"milano.gaodun.com/pkg/error-code"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

const Verify = "verify"

type ResponseObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Base 基类
type Base struct {
}

func (bs *Base) ServerJSON(c *gin.Context, v interface{}, errorCode int) {
	if c.GetBool(Verify) {
		return
	}
	var rep ResponseObject
	rep.Code = errorCode
	rep.Message = error_code.INFO[rep.Code]
	rep.Result = v
	c.Set("response_body", fmt.Sprintf("%+v", rep))
	c.JSON(http.StatusOK, rep)
}

//ServerJSONSuccess 服务器返回
func (bs *Base) ServerJSONOther(c *gin.Context, v interface{}, errorCode error_code.CodeTypeInt, str string) {
	if c.GetBool(Verify) {
		return
	}
	var rep ResponseObject
	rep.Code = int(errorCode)
	rep.Message = str
	rep.Result = v
	c.Set("response_body", fmt.Sprintf("%+v", rep))
	c.JSON(http.StatusOK, rep)
}

//ServerJSONSuccess 服务器返回
func (bs *Base) ServerJSONSuccess(c *gin.Context, v interface{}) {
	bs.ServerJSON(c, v, error_code.SUCCESSSTATUS)
}

//ServerJSONError 服务器返回
func (bs *Base) ServerJSONError(c *gin.Context, v interface{}, errorCode int) {
	bs.ServerJSON(c, v, errorCode)
}

// int 64
func (bs Base) GetInt64(c *gin.Context, key string, b ...bool) (i int64) {
	k, get := c.GetQuery(key)
	l := c.PostForm(key)
	if get {
		h, _ := strconv.ParseInt(k, 10, 64)
		i = h
	} else {
		h, _ := strconv.ParseInt(l, 10, 64)
		i = h
	}
	if len(b) > 0 && b[0] {
		bs.verifyField(c, i)
	}

	return
}

// int 32
func (bs Base) GetInt32(c *gin.Context, key string, b ...bool) (i int32) {
	k, get := c.GetQuery(key)
	l := c.PostForm(key)
	if get {
		h, _ := strconv.ParseInt(k, 10, 32)
		i = int32(h)
	} else {
		h, _ := strconv.ParseInt(l, 10, 32)
		i = int32(h)
	}
	if len(b) > 0 && b[0] {
		bs.verifyField(c, i)
	}

	return
}

// string
func (bs Base) GetString(c *gin.Context, key string, b ...bool) (v string) {
	s, get := c.GetQuery(key)
	t := c.PostForm(key)
	if get {
		v = s
	} else {
		v = t
	}
	if len(b) > 0 && b[0] {
		bs.verifyField(c, v)
	}

	return
}

// int
func (bs Base) GetInt(c *gin.Context, key string, b ...bool) (i int) {
	k, get := c.GetQuery(key)
	l := c.PostForm(key)
	if get {
		h, _ := strconv.ParseInt(k, 10, 32)
		i = int(h)
	} else {
		m, _ := strconv.ParseInt(l, 10, 32)
		i = int(m)
	}
	if len(b) > 0 && b[0] {
		bs.verifyField(c, i)
	}

	return
}

func (bs Base) GetMustCols(c *gin.Context, b ...bool) []string {
	a := c.Request.URL.Query()
	var s []string
	for i := range a {
		if len(a[i]) > 0 {
			s = append(s, i)
		}
	}
	if len(b) > 0 && b[0] {
		bs.verifyField(c, s)
	}

	return s
}

// 必须要修改的字段
func (bs Base) PostMustCols(c *gin.Context, b ...bool) []string {
	a := c.Request.PostForm
	var s []string
	for i := range a {
		if len(a[i]) > 0 {
			s = append(s, i)
		}
	}
	if len(b) > 0 && b[0] {
		bs.verifyField(c, s)
	}

	return s
}

// 参数必填
func (bs Base) verifyField(c *gin.Context, v interface{}) {
	if c.GetBool(Verify) {
		return
	}

	f := false
	switch vv := v.(type) {
	case float64:
		if vv == 0 {
			f = true
		}
	case int:
		if vv == 0 {
			f = true
		}
	case int32:
		if vv == 0 {
			f = true
		}
	case int64:
		if vv == 0 {
			f = true
		}
	case string:
		if len(vv) == 0 {
			f = true
		}
	case []string:
		if len(vv) == 0 {
			f = true
		}
	}

	if f {
		bs.ServerJSONError(c, nil, error_code.MUSTFIELD)
		c.Set(Verify, true)
		c.Abort()
	}

	return
}
