package index

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Index = new(indexApi)

type indexApi struct{}

// @summary 展示站点首页
// @tags    首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (a *indexApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "gf bbs - 首页",
	})
}

// @summary 搜索
// @tags    首页
// @produce html
// @router  /search [GET]
// @success 200 {string} html "页面HTML"
func (a *indexApi) Search(r *ghttp.Request) {
	q := r.GetString("q")

	service.View.Render(r, model.View{
		Title: "gf bbs - 搜索",
		Data: g.Map{
			"q": q,
		},
	})
}
