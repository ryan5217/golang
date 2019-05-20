package resources

import (
	"github.com/imroc/req"
	"milano.gaodun.com/pkg/error-code"
	"milano.gaodun.com/pkg/setting"
	"milano.gaodun.com/pkg/utils"
)

type Resource struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type SendResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (s Resource) Get(url string) (string, int) {
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
		"ORIGIN":       "gaodun.com",
	}
	r, err := req.Get(url, header)
	if err != nil {
		setting.Logger.Infof("SendSMS%s", err.Error())
		return "", error_code.TREELISTAPI
	}
	return r.String(), 0
}
