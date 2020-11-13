package router

import (
	"focus/app/api"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Use(
		service.Middleware.CustomCtx,
		service.Middleware.CustomView,
	)
	// 所有路由注册
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Index)                     // 首页
		group.ALL("/login", api.Login)                // 登录
		group.ALL("/category", api.Category)          // 栏目
		group.ALL("/content", api.Content)            // 内容
		group.ALL("/topic", api.Topic)                // 主题
		group.ALL("/topic/:id", api.Topic.Detail)     // 主题 - 详情
		group.ALL("/ask", api.Ask)                    // 问答
		group.ALL("/ask/:id", api.Ask.Detail)         // 问答 - 详情
		group.ALL("/article", api.Article)            // 文章
		group.ALL("/article/:id", api.Article.Detail) // 文章 - 详情
		group.ALL("/reply", api.Reply)                // 回复
		group.ALL("/user", api.User)                  // 用户
		group.ALL("/user/:id", api.User.Index)        // 用户 - 主页
		group.ALL("/file", api.File)                  // 文件
		group.ALL("/captcha", api.Captcha)            // 验证码
	})
	// 权限控制路由
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Auth)
		group.ALL("/topic/{.method}", api.Topic, "Create, Update")
		group.ALL("/ask/{.method}", api.Topic, "Create, Update")
		group.ALL("/article/{.method}", api.Topic, "Create, Update")
	})
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
