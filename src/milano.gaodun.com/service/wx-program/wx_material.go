package wx_program

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/model/wechat-material"
	"milano.gaodun.com/pkg/utils"
	"milano.gaodun.com/service/upload-img"
	"sync"
	"time"
)

// 目前只用到 appid  secret
type WechatApi struct {
	appid     string
	secret    string
	grantType string
	logger    *log.Entry
	wrg       sync.WaitGroup
}

func NewWechatApi(l *log.Entry) *WechatApi {
	w := WechatApi{}
	w.appid = conf.WxAppid
	w.secret = conf.WxSecret
	w.grantType = conf.WxGrantType
	w.logger = l
	return &w
}

// 重置 appid
func (w *WechatApi) InitConfig(appid, secret, grantype string) {
	w.appid = appid
	w.secret = secret
	w.grantType = grantype
}

// 获取图文列表
func (w *WechatApi) BatchGetMaterial(materialType string, offset, count int) {

}

// 刷新 token
func (w *WechatApi) getCacheToken() (string, error) {
	redisClient := utils.RedisHandle.RedisClientHandle
	accessToken := new(AccessToken)

	// if cache exist ，use cache
	if b, err := redisClient.Get(AccessTokenKey).Bytes(); err == nil && b != nil {
		if err := json.Unmarshal(b, accessToken); err != nil {
			w.logger.Info("wx_error_" + err.Error())
			return "", err
		}
		return accessToken.AccessToken, nil
	}

	return "", fmt.Errorf("empty")
}

// 获取 token 并自动刷新 token
func (w *WechatApi) GetAccessToken() (string, error) {
	redisClient := utils.RedisHandle.RedisClientHandle
	accessToken := new(AccessToken)

	// if cache exist ，use cache
	if r, err := w.getCacheToken(); err == nil {
		return r, err
	}

	// 获取
	tokenUrl := WechatHost + "/cgi-bin/token"
	param := req.Param{}
	param["appid"] = w.appid
	param["secret"] = w.secret
	param["grant_type"] = w.grantType
	res, err := req.Get(tokenUrl, param)
	if err != nil {
		w.logger.Info("wx_error_" + err.Error())
		return "", err
	}

	if err := res.ToJSON(accessToken); err != nil {
		w.logger.Info("wx_error_" + err.Error())
		return "", err
	}

	if accessToken.Errcode == 0 {
		redisClient.Set(AccessTokenKey, res.Bytes(), time.Minute*100) // 分钟
		return accessToken.AccessToken, nil
	}

	w.logger.Info("wx_error_" + accessToken.Errmsg)
	return "", fmt.Errorf(accessToken.Errmsg)

}

// 获取图文列表
func (w *WechatApi) GetMaterialList(offset, count int) ([]wechat_material.PriWechatMaterial, error) {

	wechatModel := wechat_material.NewWechatMaterialModel()
	wechatModel.Page(count, offset)
	wms := []wechat_material.PriWechatMaterial{}
	if err := wechatModel.GetSession().OrderBy("update_time desc").Find(&wms); err != nil {
		w.logger.Info("error_wechat_model" + err.Error())
	}
	return wms, nil
}

func (w *WechatApi) RefreshMaterial() {

}

// 初始化 material  把微信素材列表的图片全部下载到数据库中 第一次初始化后，每天拉前50条 。。。 后台有脚本跑
func (w *WechatApi) InitWechatMaterial() error {
	if _, err := w.getWechatMaterialList(0, 1); err == nil {
		tcount := 5
		c := 10
		for i := 0; i <= tcount; i++ {
			materiallist, err := w.getWechatMaterialList(i*c, c)
			if err != nil {
				return err
			}
			for _, item := range materiallist.Item {
				w.wrg.Add(1)
				go w.IntoTable(item)
			}
			w.wrg.Wait()
		}
		return nil
	} else {
		return err
	}

}

// 添加到表中
func (w *WechatApi) IntoTable(item ItemList) {
	wechatModel := wechat_material.NewWechatMaterialModel()
	wm := wechat_material.PriWechatMaterial{}
	wechatModel.WhereAndParam("media_id", item.MediaId)
	if ok, _ := wechatModel.GetSession().Get(&wm); !ok {
		for _, val := range item.Content.NewsItem {
			m := wechat_material.PriWechatMaterial{}
			m.Title = val.Title
			m.MediaId = item.MediaId
			m.Author = val.Author
			m.Digest = val.Digest
			m.ThumbMediaId = val.ThumbMediaId
			m.ThumbUrl = val.ThumbUrl
			m.Url = val.Url
			m.UpdateTime = item.UpdateTime
			imgRes := upload_img.UploadWechatImg(val.ThumbUrl)
			m.ThumbPicUrl = ""
			if imgRes.Status == 0 {
				m.ThumbPicUrl = conf.SImg_Host + "/" + imgRes.Result
			} else {
				w.logger.Info("error_upload" + imgRes.Info)
			}
			wechatModel.GetSession().InsertOne(&m)

		}
	}
	w.wrg.Done()
}

// 获取图文列表  自动把微信文章列表放到数据库中
func (w *WechatApi) getWechatMaterialList(offset, count int) (*MaterialList, error) {
	materialByte := new(MaterialList)
	token, err := w.GetAccessToken() // 先获取 token
	if err != nil {
		return materialByte, err
	}

	materialUrl := WechatHost + "/cgi-bin/material/batchget_material?access_token=" + token
	param := req.Param{}
	param["type"] = "news"
	param["offset"] = offset * count
	param["count"] = count
	res, err := req.Post(materialUrl, req.BodyJSON(param))
	if err != nil {
		w.logger.Info("wx_error_" + err.Error())
		return materialByte, err
	}

	if err := res.ToJSON(materialByte); err != nil {
		w.logger.Info("wx_error_" + err.Error())
		return materialByte, err
	}

	if materialByte.Errcode == 0 {
		return materialByte, nil
	}

	w.logger.Info("wx_error_" + materialByte.Errmsg)
	return materialByte, fmt.Errorf(materialByte.Errmsg)
}


// 刷新 token
func (w *WechatApi) DelMaterial(id int) (int64, error) {
	wechatModel := wechat_material.NewWechatMaterialModel()
	wechatModel.WhereAndParam("id", id)
	wm := wechat_material.PriWechatMaterial{}
	return wechatModel.GetSession().Delete(&wm)
}
