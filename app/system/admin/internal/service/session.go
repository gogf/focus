package service

import (
	"context"
	"focus/app/model"
	"focus/app/shared"
)

// Session管理服务
var Session = new(sessionService)

type sessionService struct{}

const (
	sessionKeyUser   = "SessionKeyUserOfAdmin"   // 用户信息存放在Session中的Key，注意和前台系统的不一样哦
	sessionKeyNotice = "SessionKeyNoticeOfAdmin" // 存放在Session中的提示信息，往往使用后则删除
)

// 设置用户Session.
func (s *sessionService) SetUser(ctx context.Context, user *model.User) error {
	return shared.Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sessionService) GetUser(ctx context.Context) *model.User {
	v := shared.Context.Get(ctx).Session.GetVar(sessionKeyUser)
	if !v.IsNil() {
		var user *model.User
		_ = v.Struct(&user)
		return user
	}
	return nil
}

// 删除用户Session。
func (s *sessionService) RemoveUser(ctx context.Context) error {
	return shared.Context.Get(ctx).Session.Remove(sessionKeyUser)
}

// 设置Notice
func (s *sessionService) SetNotice(ctx context.Context, message *model.SessionNotice) error {
	return shared.Context.Get(ctx).Session.Set(sessionKeyNotice, message)
}

// 获取Notice
func (s *sessionService) GetNotice(ctx context.Context) (*model.SessionNotice, error) {
	v := shared.Context.Get(ctx).Session.GetVar(sessionKeyNotice)
	if v != nil {
		var message *model.SessionNotice
		err := v.Struct(&message)
		return message, err
	}
	return nil, nil
}

// 删除Notice
func (s *sessionService) RemoveNotice(ctx context.Context) error {
	return shared.Context.Get(ctx).Session.Remove(sessionKeyNotice)
}
