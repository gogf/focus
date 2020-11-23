package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

// 用户管理
var User = new(userApi)

type userApi struct{}

// @summary 展示用户列表页面
// @tags    后台-用户
// @produce html
// @router  /admin/user [GET]
// @success 200 {string} html "页面HTML"
func (a *userApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "后台首页",
	})
}
