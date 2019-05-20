package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()

	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

var secrets = gin.H{
	"foo": gin.H{"email":"foo@bookedu.com", "phone":"1245"},
	"austin": gin.H{"email":"foo@bookedu.com", "phone":"1245"},
	"lena": gin.H{"email":"foo@bookedu.com", "phone":"1245"},
}

func main() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	route.GET("/hello", func(h *gin.Context) {
		h.JSON(401, gin.H{
			"message" : "hello world",
		})
	})

	route.POST("/post", func(h *gin.Context) {
		message := h.PostForm("message")
		nick := h.DefaultPostForm("nick","我是默认值")

		h.JSON(http.StatusOK, gin.H{
			"status": "post",
			"message" : message,
			"nick": nick,
		})
	})

	route.LoadHTMLGlob("templates/*")
	route.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title" : "Main website 666",
		})
	})

	route.LoadHTMLGlob("templates/**/*")
	route.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title" : "post",
		})
	})

	route.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title" : "users",
		})
	})



	//e := Employee{
	//	Name: "ryan",
	//	Jobs: {
	//		Employer: "GetYourGuide",
	//		Position: "Software Engineer",
	//	},
	//}

	//fmt.Println(e)

	go heartbeat()

	//中间件

	route.Run(":8081") // listen and serve on 0.0.0.0:8080
}

type Employee struct {
	Name string
	Jobs Job
}

type Job struct {
	Employer string
	Position string
}

func heartbeat() {
	for {
		time.Sleep(time.Second)
		fmt.Println("I'm still running....")
	}
}

