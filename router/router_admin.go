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
		group.ALL("/", admin.Index)             // 后台首页
		group.ALL("/system", admin.System)      // 系统设置
		group.ALL("/setting", admin.Setting)    // 字段管理
		group.ALL("/template", admin.Template)  // 模板管理
		group.ALL("/menu", admin.Menu)          // 菜单管理
		group.ALL("/content", admin.Content)    // 内容管理
		group.ALL("/category", admin.Category)  // 栏目管理
		group.ALL("/reply", admin.Reply)        // 评论管理
		group.ALL("/user", admin.User)          // 用户管理
		group.ALL("/user-role", admin.UserRole) // 用户角色
	})
}
