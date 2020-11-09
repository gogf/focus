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

// 自定义上下文对象
func (s *middlewareService) CustomCtx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	Context.Init(r)
	customCtx := Context.Get(r.Context())
	customCtx.Session = r.Session
	// 开发环境使用，设置测试用户信息
	if gmode.IsDevelop() {
		Context.SetUser(r.Context(), &model.ContextUser{
			Id:       1,
			Passport: "root",
			Nickname: "ROOT",
		})
	}
	if userEntity := User.GetSessionUser(r); userEntity != nil {
		Context.SetUser(r.Context(), &model.ContextUser{
			Id:       userEntity.Id,
			Passport: userEntity.Passport,
			Nickname: userEntity.Nickname,
		})
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context":   r.Context(),
		"CustomCtx": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
	// 清理请求自定义消息
	if customCtx.Message != nil {
		customCtx.Message = nil
	}
}

// 该中间件用于根据URL.Path自动设置mainTpl模板变量
func (s *middlewareService) CustomView(r *ghttp.Request) {
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
		// 内置变量
		r.Assigns(g.Map{
			"BuildIn": &ViewBuildIn{httpRequest: r},
		})
	}
	r.Middleware.Next()
}
