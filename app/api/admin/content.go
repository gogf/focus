package admin

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 内容管理
var Content = new(contentApi)

type contentApi struct{}

// @summary 展示内容列表页面
// @tags    后台-内容管理
// @produce html
// @router  /admin/content/list [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) List(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示内容栏目页面
// @tags    后台-内容管理
// @produce html
// @router  /admin/content/category [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Category(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示评论管理页面
// @tags    后台-内容管理
// @produce html
// @router  /admin/content/reply [GET]
// @success 200 {string} html "页面HTML"
func (a *contentApi) Reply(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
