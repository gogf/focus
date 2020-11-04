package service

import (
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gmode"
)

type middlewareService struct{}

var Middleware = new(middlewareService)

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
	if r.Method == "GET" {
		// 内容模板变量自动设置仅对GET请求有效
		array := gstr.SplitAndTrim(r.URL.Path, "/")
		if len(array) >= 3 {
			r.Assigns(g.Map{
				"mainTpl": gfile.Join(array[1], array[2]) + ".html",
			})
		} else if len(array) >= 2 {
			r.Assigns(g.Map{
				"mainTpl": gfile.Join(array[0], array[1]) + ".html",
			})
		} else if len(array) == 1 {
			r.Assigns(g.Map{
				"mainTpl": array[0] + "/index.html",
			})
		}
	}
	r.Middleware.Next()
}
