package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type loginApi struct{}

var Login = new(loginApi)

// @summary 展示登录页面
// @tags    登录
// @produce html
// @router  /login [GET]
// @success 200 {string} html "页面HTML"
func (a *loginApi) Index(r *ghttp.Request) {

}
