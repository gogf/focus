package service

import (
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmode"
)

var View = new(viewService)

type viewService struct{}

// 渲染模板页面
func (s *viewService) Render(r *ghttp.Request, data ...model.View) {
	var viewData g.Map
	if len(data) > 0 {
		viewData = gconv.Map(data[0])
		for k, v := range viewData {
			if g.IsEmpty(v) {
				delete(viewData, k)
			}
		}
		r.Response.WriteTplDefault(viewData)
	} else {
		r.Response.WriteTplDefault()
	}
	// 开发模式下，在页面最下面打印所有的模板变量
	if r.Method == "GET" && gmode.IsDevelop() {
		r.Response.WriteTplContent(`${dump .}`, viewData)
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
