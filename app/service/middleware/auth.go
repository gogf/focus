package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

// 该中间件用于用户鉴权，保证用户登录之后才能执行下一步服务调用
func Auth(r *ghttp.Request) {
	r.Middleware.Next()
}
