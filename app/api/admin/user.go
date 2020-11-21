package admin

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 用户管理
var User = new(userApi)

type userApi struct{}

// @summary 展示用户列表页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user/list [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) List(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示用户角色页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user/role [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Role(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
