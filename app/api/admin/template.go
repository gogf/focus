package admin

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

// 模板管理
var Template = new(templateApi)

type templateApi struct{}

// @summary 展示模板管理页面
// @tags    后台-模板
// @produce html
// @router  /admin/setting/setting [GET]
// @success 200 {string} html "页面HTML"
func (a *templateApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
