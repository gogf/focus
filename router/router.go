package router

import (
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmode"
)

func init() {
	s := g.Server()
	// 中间件
	s.Use(
		service.Middleware.CustomCtx,
	)
	// HOOK, 开发阶段禁止浏览器缓存,方便调试
	if gmode.IsDevelop() {
		s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
			r.Response.Header().Set("Cache-Control", "no-store")
		})
	}
	// 错误页面
	s.BindStatusHandler(401, func(r *ghttp.Request) {
		service.View.Render401(r)
	})
	s.BindStatusHandler(403, func(r *ghttp.Request) {
		service.View.Render403(r)
	})
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		service.View.Render404(r)
	})
	s.BindStatusHandler(500, func(r *ghttp.Request) {
		service.View.Render404(r)
	})

}
