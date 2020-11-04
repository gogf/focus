package router

import (
	"focus/app/api"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Use(service.Middleware.SessionToCtx, service.Middleware.View)
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 首页
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", api.Index)
		})
		// 登录
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/login", api.Login)
		})
		// 分类
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/category", api.Category)
			group.ALL("/category/:id", api.Category.Item)
		})
		// 主题
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/topic", api.Topic)
			group.ALL("/topic/:id", api.Topic.Detail)
		})
		// 问答
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/ask", api.Ask)
			group.ALL("/ask/:id", api.Ask.Detail)
		})
		// 文章
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/article", api.Article)
			group.ALL("/article/:id", api.Article.Detail)
		})
		// 回复
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/reply", api.Reply)
		})
		// 用户
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/user", api.User)
		})
	})
}
