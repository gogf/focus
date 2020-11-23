package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 内容管理
var Content = new(contentApi)

type contentApi struct{}

// @summary 展示内容列表页面
// @tags    后台-内容
// @produce html
// @router  /admin/content [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "内容管理",
	})
}

// @summary 显示创建内容
// @tags    后台-内容
// @produce html
// @router  /admin/content/create [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Create(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 创建内容
// @tags    后台-内容
// @produce json
// @router  /admin/content/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoCreate(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 显示修改内容
// @tags    后台-内容
// @produce html
// @router  /admin/content/update [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Update(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 修改内容
// @tags    后台-内容
// @produce json
// @router  /admin/content/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoUpdate(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 删除内容
// @tags    后台-内容
// @produce json
// @router  /admin/content/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *contentApi) DoDelete(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}
