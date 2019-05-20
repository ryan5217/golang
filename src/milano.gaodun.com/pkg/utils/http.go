package utils

import (
	"time"

	"strconv"

	"github.com/imroc/req"
	"milano.gaodun.com/conf"
)

type HttpClient struct {
	// HttpClientHandle *req.Req
	Debug bool
}

var HttpHandle = initHttpClient()

func initHttpClient() *HttpClient {
	var httpClient = new(HttpClient)
	// httpClient.HttpClientHandle = req.New()
	httpClient.Debug = false
	req.SetTimeout(5 * time.Second)
	return httpClient
}

func (hc *HttpClient) handle(paramData map[string]interface{}, headerParam map[string]string) (req.Param, req.Header) {
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	if headerParam != nil {
		for k, v := range headerParam {
			header[k] = v
		}
	}
	if hc.Debug {
		req.Debug = true
	}
	param := req.Param(paramData)
	return param, headerParam
}

func (hc *HttpClient) Post(url string, paramData map[string]interface{}, headerParam map[string]string, appKey ...string) (*req.Resp, error) {
	if len(appKey) > 0 {
		header, err := hc.prepareDefaultOptions(appKey[0])
		if err != nil {
			return nil, err
		}
		headerParam = ContactStrMap(headerParam, header)
	}

	param, header := hc.handle(paramData, headerParam)
	r, err := req.Post(url, param, header)
	return r, err
}

func (hc *HttpClient) Put(url string, paramData map[string]interface{}, headerParam map[string]string, appKey ...string) (*req.Resp, error) {
	if len(appKey) > 0 {
		header, err := hc.prepareDefaultOptions(appKey[0])
		if err != nil {
			return nil, err
		}
		headerParam = ContactStrMap(headerParam, header)
	}

	param, header := hc.handle(paramData, headerParam)
	r, err := req.Put(url, param, header)
	return r, err
}

func (hc *HttpClient) Delete(url string, paramData map[string]interface{}, headerParam map[string]string, appKey ...string) (*req.Resp, error) {
	if len(appKey) > 0 {
		header, err := hc.prepareDefaultOptions(appKey[0])
		if err != nil {
			return nil, err
		}
		headerParam = ContactStrMap(headerParam, header)
	}

	param, header := hc.handle(paramData, headerParam)
	r, err := req.Delete(url, param, header)
	return r, err
}

func (hc *HttpClient) Get(url string, paramData map[string]interface{}, headerParam map[string]string, appKey ...string) (*req.Resp, error) {
	if len(appKey) > 0 {
		header, err := hc.prepareDefaultOptions(appKey[0])
		if err != nil {
			return nil, err
		}
		headerParam = ContactStrMap(headerParam, header)
	}
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Get(url, param, header)
	if err != nil {
		return nil, err
	}
	return r, err
}

func (hc *HttpClient) PostBodyJson(url string, body interface{}) (*req.Resp, error) {
	r, err := req.Post(url, req.BodyJSON(&body))
	return r, err
}

func (hc *HttpClient) PostBodyXml(url string, body interface{}) (*req.Resp, error) {
	r, err := req.Post(url, req.BodyXML(&body))

	return r, err
}

func (hc *HttpClient) prepareDefaultOptions(appKey string) (map[string]string, error) {
	headerParam := make(map[string]string)
	now := time.Now().Unix()

	secret := conf.GetAppSecret(appKey)

	randStr := GenRandStr(10)
	nowStr := strconv.FormatInt(now, 10)
	signature := Signature(nowStr, randStr, secret)

	headerParam["App-Id-Key"] = appKey
	headerParam["Accept"] = "application/json"
	headerParam["App-Signature"] = signature
	headerParam["App-Timestamp"] = nowStr
	headerParam["App-nonce"] = randStr
	return headerParam, nil
}
