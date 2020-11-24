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
		group.ALLMap(g.Map{
			"/":         api.Index,    // 后台首页
			"/system":   api.System,   // 系统设置
			"/setting":  api.Setting,  // 字段管理
			"/template": api.Template, // 模板管理
			"/menu":     api.Menu,     // 菜单管理
			"/content":  api.Content,  // 内容管理
			"/category": api.Category, // 栏目管理
			"/reply":    api.Reply,    // 评论管理
			"/auth":     api.Auth,     // 权限管理
			"/role":     api.Role,     // 角色管理
			"/user":     api.User,     // 用户管理
		})
	})
}
