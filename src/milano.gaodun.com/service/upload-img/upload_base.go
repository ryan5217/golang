package upload_img

import (
	"fmt"
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/error-code"
	"net/url"
	"time"
)

type ImgResult struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Result string `json:"result"`
}

// 上传微信图片
func UploadWechatImg(reqUrl string) ImgResult {
	it := ImgResult{Status: error_code.FAIL, Info: "uload_error"}
	ur, _ := url.Parse(reqUrl)
	m, _ := url.ParseQuery(ur.RawQuery)
	var imgName string
	var ok bool
	var im []string
	imgName = ".jpg"
	if im, ok = m["wx_fmt"]; ok && len(im) > 0 {
		imgName = im[0]
	} else {
		fmt.Println(m, "-------------")
	}
	req.SetTimeout(10 * time.Second)
	retry := 3
	i := 0
	var res *req.Resp
	var wxErr error
	wxErr = fmt.Errorf("error")
	// 重试
	for i < retry && wxErr != nil {
		res, wxErr = req.Get(reqUrl)
		i++
		time.Sleep(50 * time.Millisecond)
	}
	if wxErr != nil {
		it.Info = wxErr.Error()
		it.Status = error_code.FAIL
		return it
	}
	header := req.Header{"Accept": "application/json", "Origin": "gaodun.com"}
	param := req.Param{}
	param["file_name"] = "img." + imgName
	param["item_name"] = "fund"
	param["file_type"] = "binary"
	param["body"] = res.String()
	if re, err := req.Post(conf.Upload_Host+"/upload/Home/UploadFileRest", header, param); err == nil {
		re.ToJSON(&it)
	} else {
		it.Info = err.Error()
	}
	return it
}
