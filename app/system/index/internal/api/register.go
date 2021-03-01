package api

import (
	"focus/app/model"
	"focus/app/system/index/internal/define"
	"focus/app/system/index/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 注册控制器
var Register = registerApi{}

type registerApi struct{}

// @summary 展示注册页面
// @tags    前台-注册
// @produce html
// @router  /register [GET]
// @success 200 {string} html "页面HTML"
func (a *registerApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{})
}

// @summary 执行注册提交处理
// @description 注意提交的密码是明文。
// @description 注册成功后自动登录。前端页面引导跳转
// @tags    前台-注册
// @produce json
// @param   entity body define.UserApiRegisterReq true "请求参数" required
// @router  /register/do [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *registerApi) Do(r *ghttp.Request) {
	var (
		data               *define.UserApiRegisterReq
		serviceRegisterReq *define.UserServiceRegisterReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if !service.Captcha.VerifyAndClear(r, model.CaptchaDefaultName, data.Captcha) {
		response.JsonExit(r, 1, "请输入正确的验证码")
	}
	if err := gconv.Struct(data, &serviceRegisterReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Register(serviceRegisterReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		// 自动登录
		err := service.User.Login(r.Context(), &define.UserServiceLoginReq{
			Passport: serviceRegisterReq.Passport,
			Password: serviceRegisterReq.Password,
		})
		if err != nil {
			response.JsonExit(r, 1, err.Error())
		}
		response.JsonExit(r, 0, "")
	}
}
