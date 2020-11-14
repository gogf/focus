package service

import (
	"context"
	"focus/app/model"
)

// Session管理
var Session = new(sessionService)

type sessionService struct{}

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "SessionKeyUser"

	// Referer存储，当已存在该session时不会更新。
	// 用于用户未登录时引导用户登录，并在登录后跳转到登录前页面。
	sessionKeyLoginReferer = "SessionKeyReferer"

	// 存放在Session中的提示信息，往往使用后则删除
	sessionKeyNotice = "SessionKeyNotice"
)

// 设置用户Session.
func (s *sessionService) SetUser(ctx context.Context, user *model.User) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sessionService) GetUser(ctx context.Context) *model.User {
	value := Context.Get(ctx).Session.Get(sessionKeyUser)
	if value != nil {
		if userEntity, ok := value.(*model.User); ok {
			return userEntity
		}
	}
	return nil
}

// 删除用户Session。
func (s *sessionService) RemoveUser(ctx context.Context) error {
	return Context.Get(ctx).Session.Remove(sessionKeyUser)
}

// 设置LoginReferer.
func (s *sessionService) SetLoginReferer(ctx context.Context, referer string) error {
	if s.GetLoginReferer(ctx) == "" {
		return Context.Get(ctx).Session.Set(sessionKeyLoginReferer, referer)
	}
	return nil
}

// 获取LoginReferer.
func (s *sessionService) GetLoginReferer(ctx context.Context) string {
	return Context.Get(ctx).Session.GetString(sessionKeyLoginReferer)
}

// 删除LoginReferer.
func (s *sessionService) RemoveLoginReferer(ctx context.Context) error {
	return Context.Get(ctx).Session.Remove(sessionKeyLoginReferer)
}

// 设置Notice
func (s *sessionService) SetNotice(ctx context.Context, message *model.SessionNotice) error {
	return Context.Get(ctx).Session.Set(sessionKeyNotice, message)
}

// 获取Notice
func (s *sessionService) GetNotice(ctx context.Context) (*model.SessionNotice, error) {
	v := Context.Get(ctx).Session.GetVar(sessionKeyNotice)
	if v != nil {
		var message *model.SessionNotice
		err := v.Struct(&message)
		return message, err
	}
	return nil, nil
}

// 删除Notice
func (s *sessionService) RemoveNotice(ctx context.Context) error {
	return Context.Get(ctx).Session.Remove(sessionKeyNotice)
}
