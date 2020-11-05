package service

import (
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmode"
)

var Middleware = new(middlewareService)

type middlewareService struct{}

// 获取session中的相关信息，写入到上下文变量中。
func (s *middlewareService) SessionToCtx(r *ghttp.Request) {
	if gmode.IsDevelop() {
		Context.SetCtx(r, &model.Context{
			UserId:       1,
			UserPassport: "root",
			UserNickname: "ROOT",
		})
	}
	if userEntity := User.GetSessionUser(r); userEntity != nil {
		Context.SetCtxWithUserEntity(r, userEntity)
	}
	r.Middleware.Next()
}

// 该中间件用于根据URL.Path自动设置mainTpl模板变量
func (s *middlewareService) View(r *ghttp.Request) {
	// 内容模板变量自动设置仅对GET请求有效
	if r.Method == "GET" {
		var (
			prefix = "web/"
			array  = gstr.SplitAndTrim(r.URL.Path, "/")
		)
		switch {
		case len(array) == 2:
			// 如果2级路由为数字，那么为模块的详情页面，那么路由固定为/xxx/detail。
			// 如果需要定制化内容模板，请在具体路由方法中设置MainTpl。
			if gstr.IsNumeric(array[1]) {
				array[1] = "detail"
			}
			r.Assigns(g.Map{
				"MainTpl": prefix + gfile.Join(array[0], array[1]) + ".html",
			})
		case len(array) == 1:
			r.Assigns(g.Map{
				"MainTpl": prefix + array[0] + "/index.html",
			})
		default:
			// 默认首页内容
			r.Assigns(g.Map{
				"MainTpl": prefix + "index/index.html",
			})
		}
	}
	r.Middleware.Next()
}
