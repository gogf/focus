package router

import (
	"focus/app/api/admin"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	// 后台系统路由注册
	g.Server().Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.AdminCtx,
			service.Middleware.AdminAuth,
		)
		group.ALL("/", admin.Index) // 首页
	})
}
