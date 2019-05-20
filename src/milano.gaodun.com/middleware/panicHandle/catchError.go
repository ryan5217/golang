package panicHandle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"runtime/debug"
)

type HTTPError interface {
	HTTPStatus() int
}

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				stackInfo := "stack info:" + string(debug.Stack())
				setting.Logger.WithField("stackInfo", stackInfo).Infof("", stackInfo)
				c.JSON(200, error_code.New(error_code.SYSERR, error_code.GetInfo()[error_code.SYSERR]))

				c.Abort()
			}
		}()
		c.Next()
	}
}
