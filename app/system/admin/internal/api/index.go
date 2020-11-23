package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 后台首页
var Index = new(indexApi)

type indexApi struct{}

// @summary 展示后台首页
// @tags    后台-首页
// @produce html
// @router  /admin [GET]
// @success 200 {string} html "页面HTML"
func (a *indexApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
