package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type topicApi struct{}

var Topic = new(topicApi)

// @summary 展示话题首页
// @tags    话题
// @produce html
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Index(r *ghttp.Request) {
	var (
		data *model.TopicServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if getListRes, err := service.Topic.GetList(r.Context(), data); err != nil {
		service.View.Render(r)
	} else {
		service.View.Render(r, g.Map{
			"list": getListRes,
		})
	}
}

// @summary 展示话题详情
// @tags    话题
// @produce html
// @param   id path int false "话题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Detail(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary 展示创建话题页面
// @tags    话题
// @produce html
// @router  /topic/create [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Create(r *ghttp.Request) {
	userSession := service.User.GetSessionUser(r)
	if userSession == nil {

	}
	service.View.Render(r)
}

// @summary 创建话题
// @description 客户端AJAX提交，客户端
// @tags    话题
// @produce json
// @param   entity body model.TopicApiCreateReq true "请求参数" required
// @router  /topic/do-create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *topicApi) DoCreate(r *ghttp.Request) {
	var (
		data             *model.TopicApiCreateReq
		serviceCreateReq *model.TopicServiceCreateReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceCreateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Topic.Create(r.Context(), serviceCreateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}

// @summary 展示修改话题页面
// @tags    话题
// @produce html
// @param   id query int true "话题ID"
// @router  /topic/update [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Update(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary 修改话题
// @tags    话题
// @produce json
// @param   entity body model.TopicApiUpdateReq true "请求参数" required
// @router  /topic/do-update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *topicApi) DoUpdate(r *ghttp.Request) {
	var (
		data             *model.TopicApiUpdateReq
		serviceUpdateReq *model.TopicServiceUpdateReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &serviceUpdateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Topic.Update(r.Context(), serviceUpdateReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}

// @summary 删除话题
// @tags    话题
// @produce json
// @param   id formData int true "话题ID"
// @router  /topic/delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *topicApi) Delete(r *ghttp.Request) {
	var (
		data *model.TopicApiDeleteReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Topic.Delete(r.Context(), data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK")
	}
}
