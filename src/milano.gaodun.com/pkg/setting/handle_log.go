package setting

import (
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"gitlab.gaodun.com/golib/graylog"
)

func GinLogger(c *gin.Context) *log.Entry {
	if l, ok := c.Get("logger"); ok {
		return l.(*log.Entry)
	} else {
		return GrayLog()
	}
}

var alg alilog = "aliyun log handel"

type alilog string

func (*alilog) HandleLog(e *log.Entry) error {
	graylog.GdLog(e.Message, e.Fields)
	return nil
}

func init()  {
	graylog.GRAY_ITEM = "milano.gaodun.com"
	graylog.ProjectName = "force"
	graylog.LogStoreName = "goservice_micro"
}
