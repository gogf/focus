package router

import (
	"focus/app/api/index/article"
	"focus/app/api/index/question"
	"focus/app/api/index/reply"
	"focus/app/api/index/user"
	"focus/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.View)
		group.ALL("/user", new(user.C))
		group.ALL("/reply", new(reply.C))
		group.ALL("/article", new(article.C))
		group.ALL("/question", new(question.C))
	})
}
