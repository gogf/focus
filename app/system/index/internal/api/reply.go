package api

import (
	"focus/app/model"
	"focus/app/system/index/internal/define"
	"focus/app/system/index/internal/service"
	"focus/library/response"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 回复控制器
var Reply = replyApi{}

type replyApi struct{}

// @summary 创建回复
// @description 客户端AJAX提交，客户端
// @tags    前台-回复
// @produce json
// @param   entity body define.ContentApiDoCreateReq true "请求参数" required
// @router  /reply/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoCreate(r *ghttp.Request) {
	var (
		data             *define.ReplyApiCreateUpdateBase
		serviceCreateReq *define.ReplyServiceCreateReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Reply.Create(r.Context(), serviceCreateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

func (a *replyApi) Index(r *ghttp.Request) {
	var (
		data *define.ReplyServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if getListRes, err := service.Reply.GetList(r.Context(), data); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		service.View.RenderTpl(r, "index/reply.html", model.View{Data: getListRes})
	}
}

// @summary 删除回复
// @tags    前台-回复
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/do-delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoDelete(r *ghttp.Request) {
	var (
		data *define.ReplyApiDoDeleteReq
	)
	if err := r.ParseForm(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Reply.Delete(r.Context(), data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}
