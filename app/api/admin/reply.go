package admin

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 评论管理
var Reply = new(replyApi)

type replyApi struct{}

// @summary 展示评论列表
// @tags    后台-评论
// @produce html
// @router  /admin/reply [GET]
// @success 200 {string} html "页面HTML"
func (a *replyApi) Index(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title: "评论列表",
	})
}

// @summary 创建评论
// @tags    后台-评论
// @produce json
// @router  /admin/reply/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoCreate(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 修改评论
// @tags    后台-评论
// @produce json
// @router  /admin/reply/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoUpdate(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}

// @summary 删除评论
// @tags    后台-评论
// @produce json
// @router  /admin/reply/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoDelete(r *ghttp.Request) {
	response.JsonExit(r, 0, "")
}
