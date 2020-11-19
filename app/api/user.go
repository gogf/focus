package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var User = new(userApi)

type userApi struct{}

// @summary 用户主页
// @tags    用户
// @produce html
// @param   entity body model.UserServiceGetListReq true "请求参数" required
// @router  /user/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {
	var (
		data *model.UserServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}

	if getListRes, err := service.User.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := "gf bbs - 用户 " + getListRes.User.Nickname + " 主页"
		service.View.Render(r, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getListRes,
		})
	}

}

// @summary 展示个人资料页面
// @tags    用户
// @produce html
// @router  /user/profile [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Profile(r *ghttp.Request) {
	if getProfile, err := service.User.GetProfile(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := "gf bbs - 用户 " + getProfile.Nickname + " 资料"
		service.View.Render(r, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
	}
}

// @summary 修改头像页面
// @tags    用户
// @produce html
// @router  /user/avatar [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Avatar(r *ghttp.Request) {
	if getProfile, err := service.User.GetProfile(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		title := "gf bbs - 用户 " + getProfile.Nickname + " 头像"
		service.View.Render(r, model.View{
			Title:       title,
			Keywords:    title,
			Description: title,
			Data:        getProfile,
		})
	}
}

// @summary 修改密码页面
// @tags    用户
// @produce html
// @router  /user/password [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Password(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary 我的文章页面
// @tags    用户
// @produce html
// @router  /user/article [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Article(r *ghttp.Request) {
	a.getContentList(r, model.ContentTypeArticle)
}

// @summary 我的主题页面
// @tags    用户
// @produce html
// @router  /user/topic [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Topic(r *ghttp.Request) {
	a.getContentList(r, model.ContentTypeTopic)
}

// @summary 我的问答页面
// @tags    用户
// @produce html
// @router  /user/ask [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Ask(r *ghttp.Request) {
	a.getContentList(r, model.ContentTypeAsk)
}

// 获取内容列表 参数contentType
func (a *userApi) getContentList(r *ghttp.Request, contentType string) {
	var (
		data *model.ContentServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	data.Type = contentType
	// 设置UserID
	data.UserId = service.Context.Get(r.Context()).User.Id

	if getListRes, err := service.Content.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: data.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(r.Context(), &model.ViewServiceGetTitleReq{
				ContentType: data.Type,
				CategoryId:  data.CategoryId,
			}),
		})
	}
}

// @summary 我的消息页面
// @tags    用户
// @produce html
// @router  /user/message [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Message(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary AJAX保存个人资料
// @tags    用户
// @produce json
// @param   entity body model.UserApiPasswordReq true "请求参数" required
// @router  /user/update-password [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdatePassword(r *ghttp.Request) {
	var (
		data *model.UserApiPasswordReq
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
// @tags    用户
// @produce json
// @param   entity body model.UserApiUpdateProfileReq true "请求参数" required
// @router  /user/update-profile [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *userApi) UpdateProfile(r *ghttp.Request) {
	var (
		data *model.UserApiUpdateProfileReq
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
// @tags    用户
// @produce json
// @router  /user/logout [GET]
// @success 200 {object} response.JsonRes "执行结果"
func (a *userApi) Logout(r *ghttp.Request) {
	if err := service.User.Logout(r.Context()); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		r.Response.RedirectTo(service.Middleware.LoginUrl)
	}
}
