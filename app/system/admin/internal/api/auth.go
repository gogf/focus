package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 权限管理
var Auth = new(authApi)

type authApi struct{}

// @summary 展示权限列表页面
// @tags    后台-用户
// @produce html
// @router  /admin/auth [GET]
// @success 200 {string} html "页面HTML"
func (a *authApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "权限列表",
	})
}
