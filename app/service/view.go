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

func (s *viewService) Render302(r *ghttp.Request) {

}

func (s *viewService) Render404(r *ghttp.Request) {

}

func (s *viewService) Render500(r *ghttp.Request) {

}
