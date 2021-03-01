package api

import (
	"focus/app/system/index/internal/define"
	"focus/app/system/index/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 赞踩控制器
var Interact = interactApi{}

type interactApi struct{}

// @summary 赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) Zan(r *ghttp.Request) {
	var (
		data *define.InteractApiZanReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.Zan(r.Context(), data.Type, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消赞
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) CancelZan(r *ghttp.Request) {
	var (
		data *define.InteractApiCancelZanReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.CancelZan(r.Context(), data.Type, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) Cai(r *ghttp.Request) {
	var (
		data *define.InteractApiCaiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.Cai(r.Context(), data.Type, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消踩
// @tags    前台-交互
// @produce json
// @param   id   formData int    true "内容ID"
// @param   type formData string true "内容类型:content,reply"
// @router  /interact/cancel-cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *interactApi) CancelCai(r *ghttp.Request) {
	var (
		data *define.InteractApiCancelCaiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Interact.CancelCai(r.Context(), data.Type, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}
