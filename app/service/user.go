package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type userService struct{}

var User = new(userService)

// 获取当前登录的用户ID，如果用户未登录返回nil。
func (s *userService) GetSessionUser(r *ghttp.Request) *model.User {
	value := r.Session.Get(model.UserSessionKey)
	if value != nil {
		if userEntity, ok := value.(*model.User); ok {
			return userEntity
		}
	}
	return nil
}

// 登录
func (s *userService) Login(r *ghttp.Request, loginReq *model.UserServiceLoginReq) error {
	userEntity, err := s.GetUserByPassportAndPassword(loginReq.Password, loginReq.Passport)
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

// 将密码按照内部算法进行加密
func (s *userService) EncryptPassword(passport, password string) string {
	return gmd5.MustEncrypt(passport + password)
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *userService) GetUserByPassportAndPassword(passport, password string) (*model.User, error) {
	return dao.User.Where(g.Map{
		dao.User.Columns.Passport: passport,
		dao.User.Columns.Password: s.EncryptPassword(passport, password),
	}).One()
}

// 检测给定的账号是否唯一
func (s *userService) CheckPassportUnique(passport string) error {
	n, err := dao.User.Data(dao.User.Columns.Password, passport).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`账号"%s"已被占用`, passport)
	}
	return nil
}

// 检测给定的昵称是否唯一
func (s *userService) CheckNicknameUnique(nickname string) error {
	n, err := dao.User.Data(dao.User.Columns.Nickname, nickname).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, nickname)
	}
	return nil
}

// 用户注册
func (s *userService) Register(r *model.UserServiceRegisterReq) error {
	if r.RoleId == 0 {
		r.RoleId = model.UserDefaultRoleId
	}
	if err := s.CheckPassportUnique(r.Passport); err != nil {
		return err
	}
	if err := s.CheckNicknameUnique(r.Nickname); err != nil {
		return err
	}
	r.Password = s.EncryptPassword(r.Passport, r.Password)
	_, err := dao.User.Data(r).Save()
	return err
}

// 修改个人资料
func (s *userService) UpdateProfile(ctx context.Context, r *model.UserServiceUpdateProfileReq) error {
	if r.Id == 0 {
		return gerror.New("用户ID不能为空")
	}
	if err := s.CheckNicknameUnique(r.Nickname); err != nil {
		return err
	}
	_, err := dao.User.Data(r).Where(dao.User.Columns.Id, Context.GetCtx(ctx).UserId).Save()
	return err
}

// 禁用指定用户
func (s *userService) Disable(id uint) error {
	_, err := dao.User.
		Data(dao.User.Columns.Status, model.UserStatusDisabled).
		Where(dao.User.Columns.Id, id).
		Update()
	return err
}
