package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Content = new(contentApi)

type contentApi struct{}

// @summary 创建内容
// @description 客户端AJAX提交，客户端
// @tags    内容
// @produce json
// @param   entity body model.ContentApiDoCreateReq true "请求参数" required
// @router  /content/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoCreate(r *ghttp.Request) {
	var (
		data             *model.ContentApiDoCreateReq
		serviceCreateReq *model.ContentServiceCreateReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Content.Create(r.Context(), serviceCreateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary 修改内容
// @tags    内容
// @produce json
// @param   entity body model.ContentApiDoUpdateReq true "请求参数" required
// @router  /content/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoUpdate(r *ghttp.Request) {
	var (
		data             *model.ContentApiDoUpdateReq
		serviceUpdateReq *model.ContentServiceUpdateReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Content.Update(r.Context(), serviceUpdateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary 删除内容
// @tags    内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoDelete(r *ghttp.Request) {
	var (
		data *model.ContentApiDoDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Content.Delete(r.Context(), data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary 赞-内容
// @tags    内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) Zan(r *ghttp.Request) {
	var (
		data *model.ContentApiZanReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.Zan(r.Context(), model.ZanCaiContentTypeContent, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消赞-内容
// @tags    内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/cancel-zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) CancelZan(r *ghttp.Request) {
	var (
		data *model.ContentApiCancelZanReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.CancelZan(r.Context(), model.ZanCaiContentTypeContent, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 踩-内容
// @tags    内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) Cai(r *ghttp.Request) {
	var (
		data *model.ContentApiCaiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.Cai(r.Context(), model.ZanCaiContentTypeContent, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消踩-内容
// @tags    内容
// @produce json
// @param   id formData int true "内容ID"
// @router  /content/cancel-cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) CancelCai(r *ghttp.Request) {
	var (
		data *model.ContentApiCancelCaiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.CancelCai(r.Context(), model.ZanCaiContentTypeContent, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}
