package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/define"
	"focus/app/system/admin/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 用户管理
var User = new(userApi)

type userApi struct{}

// @summary 展示用户管理页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}

// @summary 展示用户管理页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user/info [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Info(r *ghttp.Request) {
	if getProfile, err := service.User.GetProfile(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			Data: getProfile,
		})
	}
}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body define.UserApiPasswordReq true "请求参数" required
// @router  /admin/user/update-password [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdatePassword(r *ghttp.Request) {
	var (
		data *define.UserApiPasswordReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.UpdatePassword(r.Context(), data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary AJAX保存个人资料
// @tags    前台-更新头像
// @produce json
// @param   file formData file true "文件域"
// @param   nickname formData string true "请求参数" required
// @router  /admin/user/update-avatar [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateAvatar(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {
		response.JsonExit(r, 1, "请选择需要上传的文件")
	}
	res, err := service.File.Upload(r.Context(), &define.FileServiceUploadReq{
		File:       file,
		RandomName: true,
	})
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var (
		data *define.UserApiUpdateProfileReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	data.Avatar = res.Url

	if err := service.User.UpdateAvatar(r.Context(), data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body define.UserApiUpdateProfileReq true "请求参数" required
// @router  /admin/user/update-profile [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateProfile(r *ghttp.Request) {
	var (
		data *define.UserApiUpdateProfileReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.User.UpdateProfile(r.Context(), data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary 展示用户管理页面
// @tags    后台-用户管理
// @produce html
// @router  /admin/user/logout [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		r.Response.RedirectTo("/admin/login")
	}

}
