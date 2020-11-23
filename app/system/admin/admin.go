package admin

import (
	"focus/app/system/admin/internal/api"
	"focus/app/system/admin/internal/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 后台系统初始化
func Init() {
	// 后台系统路由注册
	g.Server().Group("/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.Auth,
		)
		group.ALL("/", api.Index)             // 后台首页
		group.ALL("/system", api.System)      // 系统设置
		group.ALL("/setting", api.Setting)    // 字段管理
		group.ALL("/template", api.Template)  // 模板管理
		group.ALL("/menu", api.Menu)          // 菜单管理
		group.ALL("/content", api.Content)    // 内容管理
		group.ALL("/category", api.Category)  // 栏目管理
		group.ALL("/reply", api.Reply)        // 评论管理
		group.ALL("/user", api.User)          // 用户管理
		group.ALL("/user-role", api.UserRole) // 用户角色
	})
}
