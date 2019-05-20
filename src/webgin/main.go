package main

import (
	"github.com/gin-gonic/gin"
	_ "webgin/modules/log" // 日志
	// _ "webgin/modules/schedule" // 定时任务
	"runtime"
	"webgin/config"
	"webgin/modules/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := initRouter()

	server.Run(router)
}
