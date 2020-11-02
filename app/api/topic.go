package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type topicApi struct{}

var Topic = new(topicApi)

// @summary 展示社区模块页面
// @tags    话题
// @produce html
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (api *topicApi) Index(r *ghttp.Request) {

}

// @summary 展示社区模块特定话题内容页面
// @tags    话题
// @produce html
// @param   id path int false "话题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (api *topicApi) Detail(r *ghttp.Request) {

}

func (api *topicApi) Create(r *ghttp.Request) {

}

func (api *topicApi) Update(r *ghttp.Request) {

}

func (api *topicApi) Delete(r *ghttp.Request) {

}
