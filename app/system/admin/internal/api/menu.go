package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 菜单管理
var Menu = new(menuApi)

type menuApi struct{}

// @summary 展示菜单管理页面
// @tags    后台-菜单
// @produce html
// @router  /admin/setting/system [GET]
// @success 200 {string} html "页面HTML"
func (a *menuApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
