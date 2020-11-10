package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var File = new(fileApi)

type fileApi struct{}

// @summary 上传文件
// @tags    文件
// @produce json
// @param   file formData file true "文件域"
// @router  /file/upload [POST]
// @success 200 {object} model.FileApiUploadRes "请求结果"
func (a *fileApi) Upload(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {
		response.JsonExit(r, 1, "请选择需要上传的文件")
	}
	res, err := service.File.Upload(r.Context(), &model.FileServiceUploadReq{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "", &model.FileApiUploadRes{
		Name: res.Name,
		Url:  res.Url,
	})
}
