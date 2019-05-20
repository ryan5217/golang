package routers

import "github.com/gin-gonic/gin"

func Root(c *gin.Context) {
	c.String(200, `{"version":1.0.0}`)
}
