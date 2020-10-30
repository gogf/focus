package router

import (
	"focus/app/api/index/article"
	"focus/app/api/index/ask"
	"focus/app/api/index/category"
	"focus/app/api/index/index"
	"focus/app/api/index/login"
	"focus/app/api/index/reply"
	"focus/app/api/index/topic"
	"focus/app/api/index/user"
	"focus/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.View)
		// 首页
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", new(index.C))
		})
		// 登录
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/login", new(login.C))
		})
		// 分类
		group.Group("/", func(group *ghttp.RouterGroup) {
			c := new(category.C)
			group.ALL("/category", c)
			group.ALL("/category/:id", c.Item)
		})
		// 主题
		group.Group("/", func(group *ghttp.RouterGroup) {
			c := new(topic.C)
			group.ALL("/topic", c)
			group.ALL("/topic/:id", c.Detail)
		})
		// 问答
		group.Group("/", func(group *ghttp.RouterGroup) {
			c := new(ask.C)
			group.ALL("/ask", c)
			group.ALL("/ask/:id", c.Detail)
		})
		// 文章
		group.Group("/", func(group *ghttp.RouterGroup) {
			c := new(article.C)
			group.ALL("/article", c)
			group.ALL("/article/:id", c.Detail)
		})
		// 回复
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/reply", new(reply.C))
		})
		// 用户
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/user", new(user.C))
		})
	})
}
