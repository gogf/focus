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

// @summary 展示创建内容页面
// @tags    内容
// @produce html
// @router  /content/create [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Create(r *ghttp.Request) {
	var (
		data *model.ContentApiCreateReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	service.View.Render(r, model.View{
		ContentType: data.Type,
	})
}

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
	if err := r.ParseForm(&data); err != nil {
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

// @summary 展示修改内容页面
// @tags    内容
// @produce html
// @param   id query int true "问答ID"
// @router  /content/update [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Update(r *ghttp.Request) {
	var (
		data *model.ContentApiUpdateReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Content.GetDetail(r.Context(), data.Id); err != nil {
		service.View.Render500(r)
	} else {
		service.View.Render(r, model.View{
			ContentType: getDetailRes.Content.Type,
			Data:        getDetailRes,
		})
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
	if err := r.ParseForm(&data); err != nil {
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
	if err := r.ParseForm(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Content.Delete(r.Context(), data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}
