package api

import "github.com/gogf/gf/net/ghttp"

type loginApi struct{}

var Login = new(loginApi)

// @summary 展示登录页面
// @tags    登录
// @produce html
// @router  /login [GET]
// @success 200 {string} html "页面HTML"
func (api *loginApi) Index(r *ghttp.Request) {

}

// @summary 提交登录
// @description 前面5次不需要验证码，同一个IP登录失败5次之后将会启用验证码校验。
// @tags    登录
// @produce json
// @param   passport    formData string true "账号"
// @param   password    formData string true "密码"
// @param   verify_code formData string false "验证码"
// @router  /login/do [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (api *loginApi) Do(r *ghttp.Request) {

}
