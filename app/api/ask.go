package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Ask = new(askApi)

type askApi struct{}

// @summary 展示问答模块页面
// @tags    问答
// @produce html
// @router  /ask [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Index(r *ghttp.Request) {

}

// @summary 展示问答模块特定问题内容页面
// @tags    问答
// @produce html
// @param   id path int false "主题ID"
// @router  /ask/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Detail(r *ghttp.Request) {

}

func (a *askApi) Create(r *ghttp.Request) {

}

func (a *askApi) Update(r *ghttp.Request) {

}

func (a *askApi) Delete(r *ghttp.Request) {

}
