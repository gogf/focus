package router

import (
	"focus/app/api/index/article"
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
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/", new(index.C))
			group.ALL("/login", new(login.C))

			group.ALL("/user", new(user.C))
			group.ALL("/reply", new(reply.C))
			group.ALL("/topic", new(topic.C))
			group.ALL("/article", new(article.C))
		})
	})
}
