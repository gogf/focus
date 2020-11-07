package api

import (
	"focus/app/model"
	"focus/app/service"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Topic = new(topicApi)

type topicApi struct{}

// @summary 展示主题首页
// @tags    主题
// @produce html
// @param   cate query int    false "栏目ID"
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
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getListRes, err := service.Topic.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			Data: getListRes,
		})
	}
}

// @summary 展示主题详情
// @tags    主题
// @produce html
// @param   id path int false "主题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Detail(r *ghttp.Request) {
	var (
		data *model.TopicApiDetailReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Topic.GetDetail(r.Context(), data.Id); err != nil {
		service.View.Render500(r)
	} else {
		service.View.Render(r, model.View{
			Data: getDetailRes,
		})
	}
}

// @summary 展示创建主题页面
// @tags    主题
// @produce html
// @router  /topic/create [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Create(r *ghttp.Request) {
	userSession := service.User.GetSessionUser(r)
	if userSession == nil {

	}
	service.View.Render(r)
}

// @summary 创建主题
// @description 客户端AJAX提交，客户端
// @tags    主题
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

// @summary 展示修改主题页面
// @tags    主题
// @produce html
// @param   id query int true "主题ID"
// @router  /topic/update [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Update(r *ghttp.Request) {
	service.View.Render(r)
}

// @summary 修改主题
// @tags    主题
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

// @summary 删除主题
// @tags    主题
// @produce json
// @param   id formData int true "主题ID"
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
