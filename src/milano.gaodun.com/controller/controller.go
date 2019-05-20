package controller

import (
	"github.com/gin-gonic/gin"
	"milano.gaodun.com/pkg/setting"
)

type TestApi struct {
	Base
}

func (myc *TestApi) TestGraylog(c *gin.Context) {
	// println(reflect.TypeOf(new(project_service_impl.ProjectServiceStruct)))
	setting.GinLogger(c).Info("abc")

}

type TestPro struct {
	Base
}

func (pc *TestPro) Get(c *gin.Context) {

}
