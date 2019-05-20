package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	form "milano.gaodun.com/model/sales"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
	service "milano.gaodun.com/service/sales"
	"time"
)

var SalseApi = NewSalesApi()

func NewSalesApi() *Sales {
	return &Sales{}
}

type Sales struct {
	Base
}

/**
获取列表
*/
func (s Sales) GetList(c *gin.Context) {
	projectId := s.GetInt64(c, "project_id", true)
	if c.GetBool(Verify) {
		return
	}
	var resultList []form.TkYxSalesman
	key := fmt.Sprintf("sales_main_info_project_id_%d", projectId)
	redis := utils.RedisHandle
	salesManList := form.TkYxSalesman{}
	salesManList.ProjectId = projectId
	service := service.NewSalesService(setting.GinLogger(c))
	if list := redis.GetData(key); list != "" {
		json.Unmarshal([]byte(list.(string)), &resultList)
		s.ServerJSONSuccess(c, resultList)
	} else {
		list := service.GetList(&salesManList)
		if len(list) > 0 {
			data, _ := json.Marshal(list)
			redis.SetData(key, data, time.Second*300)
			s.ServerJSONSuccess(c, list)
		} else {
			s.ServerJSONError(c, nil, error_code.NODATA)
		}
	}
	return
}
