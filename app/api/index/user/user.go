package user

import (
	"github.com/gogf/gf/net/ghttp"
)

type C struct{}

// @summary 访问用户详情首页
// @tags    用户
// @produce html
// @param   id path int false "用户ID"
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Index(r *ghttp.Request) {

}

// @summary 展示用户自己的信息
// @tags    用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Profile(r *ghttp.Request) {

}

func (c *C) Update(r *ghttp.Request) {

}
