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
// @tags    后台-权限
// @produce html
// @router  /admin/auth [GET]
// @success 200 {string} html "页面HTML"
func (a *authApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "权限列表",
	})
}

// @summary 添加权限
// @tags    后台-权限
// @produce json
// @param   parent_id formData int false "父级ID"
// @param   parent_id formData int false "父级ID"
// @router  /admin/auth/create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *authApi) Create(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "权限列表",
	})
}
