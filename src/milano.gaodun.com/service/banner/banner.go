package banner

import (
	"encoding/json"
	"github.com/apex/log"
	"milano.gaodun.com/model/banner"
)

type BannerServiceInterface interface {
	List() []banner.PriAppBanner
	Edit(gab *banner.PriAppBanner) (int64, error)
	Add(gab *banner.PriAppBanner) (int64, error)
	Delete() int64
}

type BannerService struct {
	BannerM *banner.BannerModel
	logger  *log.Entry
}

func NewBannerService(model *banner.BannerParam, logger *log.Entry) BannerServiceInterface {
	return &BannerService{BannerM: banner.NewBannerModel(model), logger: logger}
}

// api 数据返回
func (b *BannerService) List() []banner.PriAppBanner {
	rows, err := b.BannerM.List()
	if err != nil {
		b.logger.Error(err.Error())
	}
	b.string2jason(rows)
	return rows
}

func (b *BannerService) string2jason(g []banner.PriAppBanner) {
	for k, v := range g {
		if len(v.ShowTableDict) > 0 {
			if err := json.Unmarshal([]byte(v.ShowTableDict), &g[k].ShowTableDictJASON); err != nil {
				b.logger.Info(err.Error())
			}
		}

		if len(v.Remark) > 0 {
			if err := json.Unmarshal([]byte(v.Remark), &g[k].RemarkJSON); err != nil {
				b.logger.Info(err.Error())
			}
		}
	}

}

//保存更新内容
func (b *BannerService) Edit(gab *banner.PriAppBanner) (int64, error) {
	row, err := b.BannerM.Edit(gab)
	if err != nil {
		b.logger.Error(err.Error())
	}
	b.BannerM.Id(gab.Id).Get(gab)
	return row, err
}

//保存更新内容
func (b *BannerService) Add(gab *banner.PriAppBanner) (int64, error) {
	row, err := b.BannerM.Add(gab)
	if err != nil {
		b.logger.Error(err.Error())
		return row, err
	}
	b.BannerM.Id(gab.Id).Get(gab)
	return row, err
}
func (b *BannerService) Delete() int64 {
	r, err := b.BannerM.WhereParam().Update(banner.PriAppBanner{Status: 2})
	if err != nil {
		b.logger.Error(err.Error())
	}
	return r
}
