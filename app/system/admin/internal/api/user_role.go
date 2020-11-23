package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 用户角色管理
var UserRole = new(userRoleApi)

type userRoleApi struct{}

// @summary 展示用户角色列表页面
// @tags    后台-用户橘色
// @produce html
// @router  /admin/user-role [GET]
// @success 200 {string} html "页面HTML"
func (a *userRoleApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
