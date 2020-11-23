package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 系统设置
var System = new(systemApi)

type systemApi struct{}

// @summary 展示系统设置页面
// @tags    后台-系统
// @produce html
// @router  /admin/setting/setting [GET]
// @success 200 {string} html "页面HTML"
func (a *systemApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
