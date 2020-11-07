package service

import (
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var View = new(viewService)

type viewService struct{}

// 渲染模板页面
func (s *viewService) Render(r *ghttp.Request, data ...model.View) {
	if len(data) > 0 {
		m := gconv.Map(data[0])
		for k, v := range m {
			if g.IsEmpty(v) {
				delete(m, k)
			}
		}
		r.Response.WriteTplDefault(m)
	} else {
		r.Response.WriteTplDefault()
	}
}

// 跳转中间页面
func (s *viewService) Render302(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	view.MainTpl = "web/pages/302.html"
	s.Render(r, view)
}

// 404页面
func (s *viewService) Render404(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	view.MainTpl = "web/pages/404.html"
	s.Render(r, view)
}

// 500页面
func (s *viewService) Render500(r *ghttp.Request, data ...model.View) {
	view := model.View{}
	if len(data) > 0 {
		view = data[0]
	}
	view.MainTpl = "web/pages/500.html"
	s.Render(r, view)
}
