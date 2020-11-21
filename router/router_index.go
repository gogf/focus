package router

import (
	"focus/app/api/index"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	// 前台系统路由注册
	g.Server().Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", index.Index)                     // 首页
		group.ALL("/login", index.Login)                // 登录
		group.ALL("/register", index.Register)          // 注册
		group.ALL("/category", index.Category)          // 栏目
		group.ALL("/topic", index.Topic)                // 主题
		group.ALL("/topic/:id", index.Topic.Detail)     // 主题 - 详情
		group.ALL("/ask", index.Ask)                    // 问答
		group.ALL("/ask/:id", index.Ask.Detail)         // 问答 - 详情
		group.ALL("/article", index.Article)            // 文章
		group.ALL("/article/:id", index.Article.Detail) // 文章 - 详情
		group.ALL("/reply", index.Reply)                // 回复
		group.ALL("/captcha", index.Captcha)            // 验证码
		group.ALL("/user/:id", index.User.Index)        // 用户 - 主页
		group.ALL("/search", index.Search)              // 搜索
		// 权限控制路由
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.IndexAuth)
			group.ALL("/user", index.User)         // 用户
			group.ALL("/content", index.Content)   // 内容
			group.ALL("/interact", index.Interact) // 交互
			group.ALL("/file", index.File)         // 文件
		})
	})
}
