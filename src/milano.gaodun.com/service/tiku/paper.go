package tiku

import (
	"fmt"
	"github.com/apex/log"
	"milano.gaodun.com/model/api/tiku"
	"milano.gaodun.com/pkg/utils"
)

type PaperEType int

const (
	PaperEType_Smart    PaperEType = 1 + iota // 智能
	PaperEType_Point                          // 知识点
	PaperEType_BeReal                         // 真题
	PaperEType_Multiple                       // 组卷
	PaperEType_All                            // 全部
)

// 试卷
type PaperService struct {
	*tiku.PaperApi
	logger *log.Entry
	pm     *tiku.Paper
}

func NewPaperService(h *utils.HttpClient, logger *log.Entry) *PaperService {
	var p PaperService
	p.logger = logger
	p.PaperApi = tiku.NewPaperApi(h)
	return &p
}

func (p *PaperService) GetPaper(pkId int64) {

	pm, err := p.SetBaseApi(p.SetParam("is_need_item", "y")).GetPaperOne(pkId)
	fmt.Println(pm, err)
	p.pm = pm
}

// todo
//func (p PaperService) GetPaperItem(e PaperEType) {
//	if p.pm != nil {
//		for i := range p.pm.ItemData {
//
//		}
//	}
//}
