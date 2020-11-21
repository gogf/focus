package admin

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 系统管理
var Setting = new(settingApi)

type settingApi struct{}

// @summary 展示系统设置页面
// @tags    后台-系统管理
// @produce html
// @router  /admin/setting/system [GET]
// @success 200 {string} html "页面HTML"
func (a *settingApi) System(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示菜单管理页面
// @tags    后台-系统管理
// @produce html
// @router  /admin/setting/menu [GET]
// @success 200 {string} html "页面HTML"
func (a *settingApi) Menu(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示字典管理页面
// @tags    后台-系统管理
// @produce html
// @router  /admin/setting/setting [GET]
// @success 200 {string} html "页面HTML"
func (a *settingApi) Setting(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示模板管理页面
// @tags    后台-系统管理
// @produce html
// @router  /admin/setting/template [GET]
// @success 200 {string} html "页面HTML"
func (a *settingApi) Template(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
