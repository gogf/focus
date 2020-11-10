package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Reply = new(replyApi)

type replyApi struct{}

// @summary 赞-回复
// @tags    内容
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) Zan(r *ghttp.Request) {
	var (
		data *model.ReplyApiZanReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.Zan(r.Context(), model.ZanCaiContentTypeReply, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消赞-回复
// @tags    内容
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/cancel-zan [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) CancelZan(r *ghttp.Request) {
	var (
		data *model.ReplyApiCancelZanReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.CancelZan(r.Context(), model.ZanCaiContentTypeReply, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 踩-回复
// @tags    内容
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) Cai(r *ghttp.Request) {
	var (
		data *model.ReplyApiCaiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.Cai(r.Context(), model.ZanCaiContentTypeReply, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}

// @summary 取消踩-回复
// @tags    内容
// @produce json
// @param   id formData int true "回复ID"
// @router  /reply/cancel-cai [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *replyApi) CancelCai(r *ghttp.Request) {
	var (
		data *model.ReplyApiCancelCaiReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.ZanCai.CancelCai(r.Context(), model.ZanCaiContentTypeReply, data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "")
}
