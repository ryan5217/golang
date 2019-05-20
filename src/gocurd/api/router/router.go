package router

import (
	"github.com/gin-gonic/gin"
	. "gocurd/api/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.POST("/user", Store)

	router.PUT("/user/:id", Update)

	router.DELETE("/user/:id", Destroy)

	router.Any("/get_user", GetUsers)

	router.GET("get_test", Lists)

	return router
}