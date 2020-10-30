package index

import (
	"github.com/gogf/gf/net/ghttp"
)

type C struct{}

// @summary 展示站点首页
// @tags    首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Index(r *ghttp.Request) {

}
