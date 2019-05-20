package tiku

import (
	"github.com/apex/log"
	"milano.gaodun.com/model/api/item"
	"milano.gaodun.com/pkg/utils"
	"strings"
)

// 查询所有
var All = []string{
	"title",
	"option",
	"select",
	"partnum",
	"rank",
	"icid",
	"type",
	"answer",
	"analysis",
	"pid",
	"isvideo",
	"videoa",
	"app_video",
	"flag",
	"item_id",
	"rightnum",
	"wrongnum",
	"finishnum",
	"favoritenum",
	"fromsource",
	"translate",
}

type ItemService struct {
	*item.ItemApi
	logger *log.Entry
}

func NewItemService(h *utils.HttpClient, logger *log.Entry) *ItemService {
	var i ItemService
	i.logger = logger
	i.ItemApi = item.NewItemApi(h)
	return &i
}

// 根据 id 获取题目
func (i *ItemService) GetOne(pkId int64) (*item.ItemInfo, error) {
	return i.ItemApi.SetBaseApi(i.SetParam("field", strings.Join(All, ","))).GetOne(pkId)
}
