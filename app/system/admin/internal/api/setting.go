package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 系统管理
var Setting = new(settingApi)

type settingApi struct{}

// @summary 展示字典管理页面
// @tags    后台-字典
// @produce html
// @router  /admin/setting/setting [GET]
// @success 200 {string} html "页面HTML"
func (a *settingApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
