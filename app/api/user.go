package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var User = new(userApi)

type userApi struct{}

// @summary 用户主页
// @tags    用户
// @produce html
// @param   id path int false "用户ID"
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:       "gf bbs - 用户",
		Keywords:    "gf bbs - user keywords",
		Description: "gf bbs - user description",
	})
}

// @summary 展示个人资料页面
// @tags    用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Profile(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary 修改密码页面
// @tags    用户
// @produce html
// @router  /user/passwd [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Passwd(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary 展示注册页面
// @tags    用户
// @produce html
// @router  /user/register [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Register(r *ghttp.Request) {
	//service.View.Render(r)
	// TODO
	r.Response.WriteTpl("web/layout/layout.html", g.Map{
		"mainTpl":     "web/user/register.html",
		"title":       "gf bbs - Register",
		"keywords":    "gf bbs - Register keywords",
		"description": "gf bbs - Register description",
	})
}

// @summary 执行注册提交处理
// @description 注意提交的密码是明文。
// @description 注册成功后自动登录。前端页面引导跳转
// @tags    用户
// @produce json
// @param   entity body model.UserApiRegisterReq true "请求参数" required
// @router  /user/do-register [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) DoRegister(r *ghttp.Request) {
	var (
		data               *model.UserApiRegisterReq
		serviceRegisterReq *model.UserServiceRegisterReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceRegisterReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Register(serviceRegisterReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		// 自动登录
		err := service.User.Login(r.Context(), &model.UserServiceLoginReq{
			Passport: serviceRegisterReq.Passport,
			Password: serviceRegisterReq.Password,
		})
		if err != nil {
			response.JsonExit(r, 1, err.Error())
		}
		response.JsonExit(r, 0, "OK")
	}
}

// @summary AJAX保存个人资料
// @tags    用户
// @produce json
// @param   entity body model.UserApiUpdateProfileReq true "请求参数" required
// @router  /user/update-profile [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateProfile(r *ghttp.Request) {
	var (
		data                    *model.UserApiUpdateProfileReq
		serviceUpdateProfileReq *model.UserServiceUpdateProfileReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceUpdateProfileReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.UpdateProfile(r.Context(), serviceUpdateProfileReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}

// @summary 提交登录
// @description 前面5次不需要验证码，同一个IP登录失败5次之后将会启用验证码校验。
// @description 注意提交的密码是明文。
// @description 登录成功后前端引导页面跳转。
// @tags    用户
// @produce json
// @param   passport    formData string true "账号"
// @param   password    formData string true "密码"
// @param   verify_code formData string false "验证码"
// @router  /user/do-login [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (a *loginApi) DoLogin(r *ghttp.Request) {
	var (
		data            *model.UserApiLoginReq
		serviceLoginReq *model.UserServiceLoginReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceLoginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.Login(r.Context(), serviceLoginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}

// @summary 注销退出
// @description 注销成功后前端引导页面跳转到首页。
// @tags    用户
// @produce json
// @router  /user/logout [GET]
// @success 200 {object} response.JsonRes "执行结果"
func (a *loginApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r.Context()); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}
