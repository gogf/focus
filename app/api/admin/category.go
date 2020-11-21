package admin

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 栏目管理
var Category = new(categoryApi)

type categoryApi struct{}

// @summary 展示栏目里列表页面
// @tags    后台-栏目
// @produce html
// @router  /admin/category [GET]
// @success 200 {string} html "页面HTML"
func (a *categoryApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 创建栏目
// @tags    后台-栏目
// @produce json
// @router  /admin/category/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *categoryApi) DoCreate(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 修改栏目
// @tags    后台-栏目
// @produce json
// @router  /admin/category/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *categoryApi) DoUpdate(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 删除栏目
// @tags    后台-栏目
// @produce json
// @router  /admin/category/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *categoryApi) DoDelete(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}
