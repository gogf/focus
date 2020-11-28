package service

import (
	"focus/app/model"
	"focus/app/shared"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gmode"
)

// 中间件管理服务
var Middleware = &middlewareService{
	LoginUrl: "/admin/login",
}

type middlewareService struct {
	LoginUrl string // 登录路由地址
}

// 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	shared.Context.Init(r, customCtx)
	// 开发阶段 - 测试信息
	if gmode.IsDevelop() {
		customCtx.User = &model.ContextUser{
			Id:       1,
			Passport: "root",
			Nickname: "ROOT",
		}
	}
	if userEntity := Session.GetUser(r.Context()); userEntity != nil {
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			Passport: userEntity.Passport,
			Nickname: userEntity.Nickname,
			Avatar:   userEntity.Avatar,
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// 后台系统权限控制，用户必须登录才能访问
func (s *middlewareService) Auth(r *ghttp.Request) {
	r.Middleware.Next()
}
