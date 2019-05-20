package apis

import (
	"github.com/gin-gonic/gin"
	"gocurd/api/models"
	"net/http"
)

func Lists(c *gin.Context) {
	var list models.Test

	result, err := list.List()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"message": "没有找到",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "success",
		"data": result,
	})
}