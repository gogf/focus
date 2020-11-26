package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Login = new(loginApi)

type loginApi struct{}

// @summary 展示登录页面
// @tags    后台-登录
// @produce html
// @router  /admin/login [GET]
// @success 200 {string} html "页面HTML"
func (a *loginApi) Index(r *ghttp.Request) {
	service.View.RenderTpl(r, "admin/login/index.html", model.View{})
}

// @summary 提交登录
// @description 前面5次不需要验证码，同一个IP登录失败5次之后将会启用验证码校验。
// @description 注意提交的密码是明文。
// @description 登录成功后前端引导页面跳转。
// @tags    前台-登录
// @produce json
// @param   passport    formData string true "账号"
// @param   password    formData string true "密码"
// @param   verify_code formData string false "验证码"
// @router  /admin/login/do [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (a *loginApi) Do(r *ghttp.Request) {
	// TODO
	response.JsonExit(r, 0, "")
}
