package api

import (
	"github.com/gogf/gf/net/ghttp"
)

type askApi struct{}

var Ask = new(askApi)

// @summary 展示问答模块页面
// @tags    问答
// @produce html
// @router  /ask [GET]
// @success 200 {string} html "页面HTML"
func (api *askApi) Index(r *ghttp.Request) {

}

// @summary 展示问答模块特定问题内容页面
// @tags    问答
// @produce html
// @param   id path int false "话题ID"
// @router  /ask/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (api *askApi) Detail(r *ghttp.Request) {

}

func (api *askApi) Create(r *ghttp.Request) {

}

func (api *askApi) Update(r *ghttp.Request) {

}

func (api *askApi) Delete(r *ghttp.Request) {

}
