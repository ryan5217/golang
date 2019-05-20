package main

import (
	"fmt"
	_ "milano.gaodun.com/docs"
	"milano.gaodun.com/routers"

	"github.com/fvbock/endless"
	"milano.gaodun.com/conf"
	"syscall"
)

type ResponseObject struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// @BasePath /v1
func main() {

	endless.DefaultReadTimeOut = conf.ReadTimeout
	endless.DefaultWriteTimeOut = conf.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", conf.HTTPPort)

	fmt.Println("start : ", endPoint)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		fmt.Println("Actual pid is %d", syscall.Getpid())
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}

}
