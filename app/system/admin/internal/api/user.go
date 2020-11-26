package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 用户管理
var User = new(userApi)

type userApi struct{}

// @summary 展示用户管理页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示用户管理页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user/info [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Info(r *ghttp.Request) {
	// TODO 个人信息
	service.View.Render(r, model.View{
		Title: "个人中心",
	})
}

// @summary 展示用户管理页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user/logout [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Logout(r *ghttp.Request) {
	// TODO 登出逻辑
	r.Response.RedirectTo("/admin/login")
}
