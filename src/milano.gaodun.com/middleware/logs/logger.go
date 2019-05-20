package logs

import (
	"time"

	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/tool"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"fmt"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		uuid := tool.GetUID()

		logger := setting.GrayLog()
		logger = logger.WithField("X-Response-ID", uuid)
		c.Set("logger", logger)
		c.Header("X-Response-ID", uuid)
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		b, _ := json.Marshal(c.Request.PostForm)
		bind := map[string]interface{}{}
		if err := c.ShouldBindJSON(&bind); err == nil {
			bd, _ := json.Marshal(bind)
			logger = logger.WithField("request_json_body", string(bd))
		}
		logger = logger.WithField("request_url", c.Request.URL.Path)
		logger = logger.WithField("response_time", fmt.Sprintf("%.3f", latency.Seconds()))
		logger.WithField("request_body：", string(b))
		logger.Info(c.GetString("response_body"))
	}
}
