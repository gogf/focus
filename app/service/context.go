package service

import (
	"context"
	"focus/app/model"
	"github.com/gogf/gf/net/ghttp"
)

type contextService struct{}

var Context = new(contextService)

// 将上下文信息设置到上下文请求中
func (s *contextService) SetCtx(r *ghttp.Request, localCtx *model.Context) error {
	r.SetCtxVar(model.ContextUserKey, localCtx)
	return nil
}

// 设置用户信息到Context中
func (s *contextService) SetCtxWithUserEntity(r *ghttp.Request, userEntity *model.User) error {
	localCtx := s.GetCtx(r.Context())
	if localCtx == nil {
		localCtx = &model.Context{
			UserId:       userEntity.Id,
			UserPassport: userEntity.Passport,
			UserNickname: userEntity.Nickname,
		}
	} else {
		localCtx.UserId = userEntity.Id
		localCtx.UserPassport = userEntity.Passport
		localCtx.UserNickname = userEntity.Nickname
	}
	return s.SetCtx(r, localCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) GetCtx(ctx context.Context) *model.Context {
	if ctx == nil {
		return nil
	}
	value := ctx.Value(model.ContextUserKey)
	if value != nil {
		if localCtx, ok := value.(*model.Context); ok {
			return localCtx
		}
	}
	return nil
}
