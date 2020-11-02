package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type topicApi struct{}

var Topic = new(topicApi)

// @summary 展示话题首页
// @tags    话题
// @produce html
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Index(r *ghttp.Request) {

}

// @summary 展示话题详情
// @tags    话题
// @produce html
// @param   id path int false "话题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Detail(r *ghttp.Request) {

}

// @summary 创建话题
// @tags    话题
// @produce json
// @param   entity body model.TopicApiCreateReq true "请求参数" required
// @router  /topic/create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *topicApi) Create(r *ghttp.Request) {

}

// @summary 修改话题
// @tags    话题
// @produce json
// @param   entity body model.TopicApiUpdateReq true "请求参数" required
// @router  /topic/update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *topicApi) Update(r *ghttp.Request) {

}

// @summary 删除话题
// @tags    话题
// @produce json
// @param   id formData int true "话题ID"
// @router  /topic/delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *topicApi) Delete(r *ghttp.Request) {

}
