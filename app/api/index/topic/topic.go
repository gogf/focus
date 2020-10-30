package topic

import (
	"github.com/gogf/gf/net/ghttp"
)

type C struct{}

// @summary 展示社区模块页面
// @tags    话题
// @produce html
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Index(r *ghttp.Request) {

}

// @summary 展示社区模块特定话题内容页面
// @tags    话题
// @produce html
// @param   id query int false "话题ID"
// @router  /topic/detail [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Detail(r *ghttp.Request) {

}

func (c *C) Create(r *ghttp.Request) {

}

func (c *C) Update(r *ghttp.Request) {

}

func (c *C) Delete(r *ghttp.Request) {

}
