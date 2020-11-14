package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 回复控制器
var Reply = new(replyApi)

type replyApi struct{}

// @summary 创建回复
// @description 客户端AJAX提交，客户端
// @tags    内容
// @produce json
// @param   entity body model.ContentApiDoCreateReq true "请求参数" required
// @router  /content/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) DoCreate(r *ghttp.Request) {
	var (
		data             *model.ReplyApiCreateUpdateBase
		serviceCreateReq *model.ReplyServiceCreateReq
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