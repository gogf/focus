package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/o1egl/govatar"
)

// 用户管理服务
var User = userService{
	avatarUploadPath:      g.Cfg().GetString(`upload.path`) + `/avatar`,
	avatarUploadUrlPrefix: `/upload/avatar`,
}

type userService struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func init() {
	// 启动时创建头像存储目录
	if !gfile.Exists(User.avatarUploadPath) {
		if err := gfile.Mkdir(User.avatarUploadPath); err != nil {
			g.Log().Fatal(err)
		}
	}
}

// 获得头像上传路径
func (s *userService) GetAvatarUploadPath() string {
	return s.avatarUploadPath
}

// 获得头像上传对应的URL前缀
func (s *userService) GetAvatarUploadUrlPrefix() string {
	return s.avatarUploadUrlPrefix
}

// 执行登录
func (s *userService) Login(ctx context.Context, loginReq *define.UserServiceLoginReq) error {
	userEntity, err := s.GetUserByPassportAndPassword(
		loginReq.Passport,
		s.EncryptPassword(loginReq.Passport, loginReq.Password),
	)
	if err != nil {
		return err
	}
	if userEntity == nil {
		return gerror.New(`账号或密码错误`)
	}
	if err := Session.SetUser(ctx, userEntity); err != nil {
		return err
	}
	// 自动更新上线
	shared.Context.SetUser(ctx, &model.ContextUser{
		Id:       userEntity.Id,
		Passport: userEntity.Passport,
		Nickname: userEntity.Nickname,
		Avatar:   userEntity.Avatar,
	})
	return nil
}

// 注销
func (s *userService) Logout(ctx context.Context) error {
	return Session.RemoveUser(ctx)
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
		dao.User.Columns.Password: password,
	}).One()
}

// 检测给定的账号是否唯一
func (s *userService) CheckPassportUnique(passport string) error {
	n, err := dao.User.Where(dao.User.Columns.Passport, passport).Count()
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
	n, err := dao.User.Where(dao.User.Columns.Nickname, nickname).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, nickname)
	}
	return nil
}

// 用户注册。
func (s *userService) Register(r *define.UserServiceRegisterReq) error {
	var user *model.User
	if err := gconv.Struct(r, &user); err != nil {
		return err
	}
	if err := s.CheckPassportUnique(user.Passport); err != nil {
		return err
	}
	if err := s.CheckNicknameUnique(user.Nickname); err != nil {
		return err
	}
	user.Password = s.EncryptPassword(user.Passport, user.Password)
	// 自动生成头像
	avatarFilePath := fmt.Sprintf(`%s/%s.jpg`, s.avatarUploadPath, user.Passport)
	if err := govatar.GenerateFileForUsername(govatar.MALE, user.Passport, avatarFilePath); err != nil {
		return gerror.Wrapf(err, `自动创建头像失败`)
	}
	user.Avatar = fmt.Sprintf(`%s/%s.jpg`, s.avatarUploadUrlPrefix, user.Passport)
	_, err := dao.User.Data(user).OmitEmpty().Save()
	return err
}

// 修改个人密码
func (s *userService) UpdatePassword(ctx context.Context, r *define.UserApiPasswordReq) error {
	oldPassword := s.EncryptPassword(shared.Context.Get(ctx).User.Passport, r.OldPassword)
	n, err := dao.User.Where(dao.User.Columns.Password, oldPassword).Where(dao.User.Columns.Id, shared.Context.Get(ctx).User.Id).Count()
	if err != nil {
		return err
	}
	if n == 0 {
		return gerror.New(`原始密码错误`)
	}
	newPassword := s.EncryptPassword(shared.Context.Get(ctx).User.Passport, r.NewPassword)
	_, err = dao.User.Data(g.Map{
		dao.User.Columns.Password: newPassword,
	}).Where(dao.User.Columns.Id, shared.Context.Get(ctx).User.Id).Update()
	return err
}

// 获取个人信息
func (s *userService) GetProfileById(ctx context.Context, userId uint) (*define.UserServiceProfileRes, error) {
	getProfile := new(define.UserServiceProfileRes)

	userRecord, err := dao.User.Fields(define.UserServiceProfileRes{}).WherePri(userId).M.One()
	if err != nil {
		return nil, err
	}

	if err := userRecord.Struct(getProfile); err != nil {
		return nil, err
	}

	getProfile.Stats, err = s.GetUserStats(ctx, userId)
	if err != nil {
		return nil, err
	}

	return getProfile, nil
}

// 修改个人资料
func (s *userService) GetProfile(ctx context.Context) (*define.UserServiceProfileRes, error) {
	return s.GetProfileById(ctx, shared.Context.Get(ctx).User.Id)
}

// 修改个人头像
func (s *userService) UpdateAvatar(ctx context.Context, r *define.UserApiUpdateProfileReq) error {
	user := shared.Context.Get(ctx).User
	userId := user.Id
	userServiceUpdateAvatarReq := new(define.UserServiceUpdateAvatarReq)
	err := gconv.Struct(r, &userServiceUpdateAvatarReq)
	if err != nil {
		return err
	}

	_, err = dao.User.Data(userServiceUpdateAvatarReq).Where(dao.User.Columns.Id, userId).Update()
	// 更新登录session Avatar
	if err == nil && user.Avatar != r.Avatar {
		sessionUser := Session.GetUser(ctx)
		sessionUser.Avatar = r.Avatar
		Session.SetUser(ctx, sessionUser)
	}

	return err
}

// 修改个人资料
func (s *userService) UpdateProfile(ctx context.Context, r *define.UserApiUpdateProfileReq) error {
	user := shared.Context.Get(ctx).User
	userId := user.Id
	n, err := dao.User.Where(dao.User.Columns.Nickname, r.Nickname).Where("id <> ?", userId).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return gerror.Newf(`昵称"%s"已被占用`, r.Nickname)
	}

	userServiceUpdateProfileReq := new(define.UserServiceUpdateProfileReq)
	err = gconv.Struct(r, &userServiceUpdateProfileReq)
	if err != nil {
		return err
	}

	_, err = dao.User.Data(userServiceUpdateProfileReq).Where(dao.User.Columns.Id, userId).Update()
	// 更新登录session Nickname
	if err == nil && user.Nickname != r.Nickname {
		sessionUser := Session.GetUser(ctx)
		sessionUser.Nickname = r.Nickname
		Session.SetUser(ctx, sessionUser)
	}

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

// 查询用户内容列表及用户信息
func (s *userService) GetList(ctx context.Context, r *define.UserServiceGetListReq) (*define.UserServiceGetListRes, error) {
	var data *define.ContentServiceGetListReq
	if err := gconv.Struct(r, &data); err != nil {
		return nil, err
	}

	getListRes, err := Content.GetList(ctx, data)
	if err != nil {
		return nil, err
	}

	user, err := User.GetProfileById(ctx, r.UserId)
	if err != nil {
		return nil, err
	}

	res := new(define.UserServiceGetListRes)
	res.Content = getListRes
	res.User = user
	res.Stats, err = s.GetUserStats(ctx, data.UserId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 消息列表
func (s *userService) GetMessageList(ctx context.Context, r *define.UserServiceGetMessageListReq) (*define.UserServiceGetMessageListRes, error) {
	userId := shared.Context.Get(ctx).User.Id
	replyReq := &define.ReplyServiceGetListReq{}
	gconv.Struct(r, replyReq)

	// 管理员看所有的
	if !s.IsAdminShow(ctx, userId) {
		replyReq.UserId = userId
	}

	replyList, err := Reply.GetList(ctx, replyReq)
	if err != nil {
		return nil, err
	}

	getListRes := &define.UserServiceGetMessageListRes{
		Page:  r.Page,
		Size:  r.Size,
		Total: replyList.Total,
	}

	getListRes.List = replyList.List
	getListRes.Stats, err = s.GetUserStats(ctx, userId)
	if err != nil {
		return nil, err
	}

	return getListRes, nil
}

// 获取文章数量
func (s *userService) GetUserStats(ctx context.Context, userId uint) (map[string]int, error) {
	m := dao.Content.Fields(model.ContentListItem{})
	m = m.Fields(dao.Content.Columns.Type, "count(*) total")
	if userId > 0 && !s.IsAdminShow(ctx, userId) {
		m = m.Where(dao.Content.Columns.UserId, userId)
	}
	statsModel := m.Group(dao.Content.Columns.Type)
	statsAll, err := statsModel.M.All()
	if err != nil {
		return nil, err
	}
	statsMap := make(map[string]int)
	for _, v := range statsAll {
		statsMap[v["type"].String()] = v["total"].Int()
	}

	replyModel := dao.Reply.Fields("count(*) total")
	if userId > 0 && !s.IsAdminShow(ctx, userId) {
		replyModel = replyModel.Where(dao.Reply.Columns.UserId, userId)
	}
	record, err := replyModel.M.One()
	if err != nil {
		return nil, err
	}
	statsMap["message"] = record["total"].Int()

	return statsMap, nil
}

// 是否是访问管理员的数据
func (s *userService) IsAdminShow(ctx context.Context, userId uint) bool {
	context := shared.Context.Get(ctx)
	if context == nil {
		return false
	}
	if context.User == nil {
		return false
	}
	if userId != context.User.Id {
		return false
	}
	return context.User.IsAdmin
}
