package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Article = new(articleApi)

type articleApi struct{}

// @summary 展示文章板块页面
// @tags    文章
// @produce html
// @router  /article [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Index(r *ghttp.Request) {
	// TODO 文章内容查询，展示
	r.Response.WriteTpl("web/layout/layout.html", g.Map{
		"mainTpl":     "web/article/article.html",
		"title":       "gf bbs - 文章",
		"keywords":    "gf bbs - article keywords",
		"description": "gf bbs - article description",
	})
}

// @summary 展示文章内容
// @tags    文章
// @produce html
// @router  /article/publish [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Publish(r *ghttp.Request) {
	// TODO 文章内容查询，展示
	r.Response.WriteTpl("web/layout/layout.html", g.Map{
		"mainTpl":     "web/article/publish.html",
		"title":       "gf bbs - Publish",
		"keywords":    "gf bbs - publish keywords",
		"description": "gf bbs - publish description",
	})
}

// @summary 展示文章内容
// @tags    文章
// @produce html
// @param   id path int false "文章ID"
// @router  /article/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Detail(r *ghttp.Request) {

}

func (a *articleApi) Create(r *ghttp.Request) {

}

func (a *articleApi) Update(r *ghttp.Request) {

}

func (a *articleApi) Delete(r *ghttp.Request) {

}
