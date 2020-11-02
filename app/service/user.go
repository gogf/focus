package service

import (
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

type userService struct{}

var User = new(userService)

// 登录
func (s *userService) Login(r *ghttp.Request, loginReq *model.UserServiceLoginReq) error {
	userEntity, err := dao.User.GetUserByPassportAndPassword(loginReq.Password, loginReq.Passport)
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New(`账号或密码错误`)
	}
	if err := r.Session.Set(model.UserSessionKey, userEntity); err != nil {
		return err
	}
	if err := Context.SetCtxWithUserEntity(r, userEntity); err != nil {
		return err
	}
	return nil
}

// 注销
func (s *userService) Logout(r *ghttp.Request) error {
	return r.Session.Remove(model.UserSessionKey)
}
