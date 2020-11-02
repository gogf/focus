package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type articleApi struct{}

var Article = new(articleApi)

// @summary 展示文章板块页面
// @tags    文章
// @produce html
// @router  /article [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Index(r *ghttp.Request) {

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
