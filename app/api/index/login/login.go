package login

import "github.com/gogf/gf/net/ghttp"

type C struct{}

// @summary 展示登录页面
// @tags    登录
// @produce html
// @router  /login [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Index(r *ghttp.Request) {

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
func (c *C) Do(r *ghttp.Request) {

}
