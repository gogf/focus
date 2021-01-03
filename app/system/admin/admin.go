package admin

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 后台系统初始化
func Init() {
	// 后台系统路由注册
	g.Server().Group("/admin", func(group *ghttp.RouterGroup) {
		// 暂未开放。
	})
}
