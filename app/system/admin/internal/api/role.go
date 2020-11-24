package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 角色管理
var Role = new(roleApi)

type roleApi struct{}

// @summary 展示角色管理页面
// @tags    后台-角色管理
// @produce html
// @router  /admin/role [GET]
// @success 200 {string} html "页面HTML"
func (a *roleApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "角色管理",
	})
}
