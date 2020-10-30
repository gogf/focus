package ask

import (
	"github.com/gogf/gf/net/ghttp"
)

type C struct{}

// @summary 展示问答模块页面
// @tags    问答
// @produce html
// @router  /ask [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Index(r *ghttp.Request) {

}

// @summary 展示问答模块特定问题内容页面
// @tags    问答
// @produce html
// @param   id path int false "话题ID"
// @router  /ask/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Detail(r *ghttp.Request) {

}

func (c *C) Create(r *ghttp.Request) {

}

func (c *C) Update(r *ghttp.Request) {

}

func (c *C) Delete(r *ghttp.Request) {

}
