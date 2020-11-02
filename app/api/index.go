package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type indexApi struct{}

var Index = new(indexApi)

// @summary 展示站点首页
// @tags    首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (api *indexApi) Index(r *ghttp.Request) {

}
