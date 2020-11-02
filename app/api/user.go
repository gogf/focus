package api

import (
	"focus/app/dao"
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type userApi struct{}

var User = new(userApi)

// @summary 访问用户详情首页
// @tags    用户
// @produce html
// @param   id path int false "用户ID"
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {

}

// @summary 展示用户自己的信息
// @tags    用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Profile(r *ghttp.Request) {

}

// @summary AJAX执行注册提交
// @description 注意提交的密码是明文。
// @description 注册成功后自动登录。前端页面引导跳转
// @tags    用户
// @produce json
// @param   entity body model.UserApiRegisterReq true "请求参数" required
// @router  /user/register [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) Register(r *ghttp.Request) {
	var (
		data           *model.UserApiRegisterReq
		daoRegisterReq *model.UserDaoRegisterReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &daoRegisterReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := dao.User.Register(daoRegisterReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		// 自动登录
		err := service.User.Login(r, &model.UserServiceLoginReq{
			Passport: daoRegisterReq.Passport,
			Password: daoRegisterReq.Password,
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
		data                *model.UserApiUpdateProfileReq
		daoUpdateProfileReq *model.UserDaoUpdateProfileReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &daoUpdateProfileReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := dao.User.UpdateProfile(daoUpdateProfileReq); err != nil {
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
// @router  /user/login [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (a *loginApi) Login(r *ghttp.Request) {
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
	if err := service.User.Login(r, serviceLoginReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}

// @summary 提交登录
// @description 前面5次不需要验证码，同一个IP登录失败5次之后将会启用验证码校验。
// @description 注意提交的密码是明文。
// @description 登录成功后前端引导页面跳转到首页。
// @tags    用户
// @produce json
// @param   passport    formData string true "账号"
// @param   password    formData string true "密码"
// @param   verify_code formData string false "验证码"
// @router  /user/logout [POST]
// @success 200 {object} response.JsonRes "执行结果"
func (a *loginApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}
