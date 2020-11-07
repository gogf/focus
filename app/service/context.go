package service

import (
	"context"
	"focus/app/model"
	"github.com/gogf/gf/net/ghttp"
)

var Context = new(contextService)

type contextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(r *ghttp.Request) {
	r.SetCtxVar(model.ContextUserKey, new(model.Context))
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextUserKey)
	if value == nil {
		panic("找不到上下文对象，代码流程执行出现问题，检查下中间件注册顺序？")
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextService) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// 设置请求的上下文的提示信息，该提示信息用于模板渲染，随后请求结束后清空。
func (s *contextService) SetMessage(ctx context.Context, ctxMessage *model.ContextMessage) {
	s.Get(ctx).Message = ctxMessage
}
