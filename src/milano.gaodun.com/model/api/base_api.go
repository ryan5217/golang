package api

import (
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
)

type BaseApi struct {
	WhereParam map[string]interface{}
	ApiHeader  map[string]string
	HttpClient *utils.HttpClient
	Uri        string
	Key        string
}

func NewBaseApi(h *utils.HttpClient) *BaseApi {
	var b = BaseApi{}
	b.HttpClient = h
	return &b
}

func (b *BaseApi) SetUri(path string) *BaseApi {
	b.Uri = conf.BASE_DOMAIN + path
	return b
}

func (b *BaseApi) SetKey(key string) *BaseApi {
	b.Key = key
	return b
}

func (b *BaseApi) SetParam(key string, val interface{}) *BaseApi {
	if b.WhereParam == nil {
		b.WhereParam = make(map[string]interface{})
	}

	b.WhereParam[key] = val
	return b
}

func (b *BaseApi) SetApiHeader(key string, val string) *BaseApi {
	if b.ApiHeader == nil {
		b.ApiHeader = make(map[string]string)
	}

	b.WhereParam[key] = val
	return b
}

func (b *BaseApi) ResetParam() {
	b.WhereParam = nil
}
