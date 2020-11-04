package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var View = new(viewService)

type viewService struct{}

// 渲染模板页面
func (s *viewService) Render(r *ghttp.Request, data ...g.Map) {
	if len(data) > 0 {
		r.Response.WriteTplDefault(data[0])
	} else {
		r.Response.WriteTplDefault()
	}
}

func (s *viewService) Render302(r *ghttp.Request) {

}

func (s *viewService) Render404(r *ghttp.Request) {

}

func (s *viewService) Render500(r *ghttp.Request) {

}
