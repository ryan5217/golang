package upload

import (
	"github.com/imroc/req"
	"milano.gaodun.com/conf"
	"milano.gaodun.com/pkg/utils"
	"mime/multipart"
)

type UploadApi struct {
	HttpClient *utils.HttpClient
	Uri        string
}

type Upload struct {
	Status                   int64
	Info             string
	Result           string
}
//例子：FileName:test_upload.png; FilePath:ep/upload;
// FileType:img/file/binary; IsCut:裁剪，仅限图片，1裁剪，0，不裁剪
// Thumb:30_30 (30像素*30像素)
type Param struct {
	FileName string
	FilePath string
	FileType string
	IsCut  int64
	Thumb string
	FileHeader *multipart.FileHeader
}
func NewUploadApi(h *utils.HttpClient) *UploadApi {
	var p = UploadApi{}
	p.HttpClient = h
	return &p
}

func (g *UploadApi) UploadFile(upParam *Param) (*Upload, error) {
	res := Upload{}
	file, err := upParam.FileHeader.Open()
	if(err != nil){
		return &res, err
	}
	fu := req.FileUpload{
		File:      file,
		FieldName: "file", // FieldName is form field name
		FileName:  upParam.FileHeader.Filename,
	}
	param := req.Param{
		"file_name": upParam.FileName,
		"item_name": upParam.FilePath,
		"file_type": upParam.FileType,
		"encrypt":   "1",
		"is_cut":    upParam.IsCut,
		"thumb":     upParam.Thumb,
	}
	g.Uri = conf.UPLOAD_DOMAIN + "/upload/Home/UploadFileRest"

	r, err := req.Post(g.Uri, param, fu)
	if(err != nil){
		return &res, err
	}
	err = r.ToJSON(&res)
	return &res, err
}
