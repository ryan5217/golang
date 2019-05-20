package gd_info

import (
	"github.com/apex/log"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
)

type GdInfoInterface interface {
	Get(infoType string, page, limit int64) (string, error)
	GetInfo(fileUrl string) (string, error)
}

type GdInfoService struct {
	BaseUri string
	logger  *log.Entry
}

func NewGdInfoService(l *log.Entry) GdInfoInterface {
	var g = GdInfoService{}
	g.BaseUri = conf.GD_INFO_DOMAIN
	g.logger = l
	return &g
}

func (g *GdInfoService) Get(infoType string, page, limit int64) (string, error) {
	p := req.Param{}
	p["tag"] = "son"
	p["type"] = 2
	p["tid"] = infoType
	p["page"] = page
	p["pagenum"] = limit
	res, err := req.Get(g.BaseUri+"/api/m_article.php", p)
	if err != nil {
		g.logger.Info(err.Error())
	}
	return res.String(), err
}

func (g *GdInfoService) GetInfo(fileUrl string) (string, error) {
	p := req.Param{}
	p["type"] = 1
	p["file"] = fileUrl
	res, err := req.Get(g.BaseUri+"/api/m_article.php", p)
	if err != nil {
		g.logger.Info(err.Error())
	}
	return res.String(), err
}
