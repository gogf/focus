package api

import (
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"
	"focus/app/system/index/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct{}

// @summary 用户主页
// @tags    前台-用户
// @produce html
// @param   entity body define.UserServiceGetListReq true "请求参数" required
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {
	var (
		data *define.UserServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}

	a.getContentList(r, data.Type, data.Id)

}

// @summary 展示个人资料页面
// @tags    前台-用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Profile(r *ghttp.Request) {
	if getProfile, err := service.User.GetProfile(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := "用户 " + getProfile.Nickname + " 资料"
		service.View.Render(r, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
	}
}

// @summary 修改头像页面
// @tags    前台-用户
// @produce html
// @router  /user/avatar [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Avatar(r *ghttp.Request) {
	if getProfile, err := service.User.GetProfile(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := "用户 " + getProfile.Nickname + " 头像"
		service.View.Render(r, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
	}
}

// @summary 修改密码页面
// @tags    前台-用户
// @produce html
// @router  /user/password [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Password(r *ghttp.Request) {
	if getProfile, err := service.User.GetProfile(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := "用户 " + getProfile.Nickname + " 修改密码"
		service.View.Render(r, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
	}
}

// @summary 我的文章页面
// @tags    前台-用户
// @produce html
// @router  /user/article [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Article(r *ghttp.Request) {
	a.getContentList(r, model.ContentTypeArticle, shared.Context.Get(r.Context()).User.Id)
}

// @summary 我的主题页面
// @tags    前台-用户
// @produce html
// @router  /user/topic [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Topic(r *ghttp.Request) {
	a.getContentList(r, model.ContentTypeTopic, shared.Context.Get(r.Context()).User.Id)
}

// @summary 我的问答页面
// @tags    前台-用户
// @produce html
// @router  /user/ask [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Ask(r *ghttp.Request) {
	a.getContentList(r, model.ContentTypeAsk, shared.Context.Get(r.Context()).User.Id)
}

// 获取内容列表 参数contentType,用户信息
func (a *userApi) getContentList(r *ghttp.Request, contentType string, userId uint) {
	var (
		data *define.UserServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	data.Type = contentType
	// 设置UserID
	data.UserId = userId

	if getListRes, err := service.User.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: data.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(r.Context(), &define.ViewServiceGetTitleReq{
				ContentType: data.Type,
				CategoryId:  data.CategoryId,
			}),
		})
	}
}

// @summary 我的消息页面
// @tags    前台-用户
// @produce html
// @router  /user/message [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Message(r *ghttp.Request) {
	var (
		data *define.UserServiceGetMessageListReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if getListRes, err := service.User.GetMessageList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: data.TargetType,
			Data:        getListRes,
		})
	}

}

// @summary AJAX保存个人资料
// @tags    前台-用户
// @produce json
// @param   entity body define.UserApiPasswordReq true "请求参数" required
// @router  /user/update-password [POST]
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
// @router  /user/update-avatar [POST]
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
// @router  /user/update-profile [POST]
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

// @summary 注销退出
// @description 注销成功后前端引导页面跳转到首页。
// @tags    前台-用户
// @produce json
// @router  /user/logout [GET]
// @success 200 {object} response.JsonRes "执行结果"
func (a *userApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		r.Response.RedirectTo(service.Middleware.GetLoginUrl())
	}
}
