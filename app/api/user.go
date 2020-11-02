package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type userApi struct{}

var User = new(userApi)

// @summary 访问用户详情首页
// @tags    用户
// @produce html
// @param   id path int false "用户ID"
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (api *userApi) Index(r *ghttp.Request) {

}

// @summary 展示用户自己的信息
// @tags    用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (api *userApi) Profile(r *ghttp.Request) {

}

func (api *userApi) Update(r *ghttp.Request) {

}
