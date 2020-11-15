package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

var User = &userService{}

type userService struct{}

// 执行登录
func (s *userService) Login(ctx context.Context, loginReq *model.UserServiceLoginReq) error {
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
	Context.SetUser(ctx, &model.ContextUser{
		Id:       userEntity.Id,
		Passport: userEntity.Passport,
		Nickname: userEntity.Nickname,
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

// 用户注册，注意这里是值传参，因为内部会修改参数的属性，防止对输入参数造成影响。
func (s *userService) Register(r model.UserServiceRegisterReq) error {
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
	_, err := dao.User.Data(r).Where(dao.User.Columns.Id, Context.Get(ctx).User.Id).Save()
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

// 查询用户内容列表
func (s *userService) GetList(ctx context.Context, r *model.UserServiceGetListReq) (*model.UserServiceGetListRes, error) {
	m := dao.Content.Fields(model.ContentListItem{})
	if r.Type != "" {
		m = m.Where(dao.Content.Columns.Type, r.Type)
	}
	if r.CategoryId > 0 {
		// 栏目检索
		idArray, err := Category.GetSubIdList(ctx, r.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.Columns.CategoryId, idArray)
	}
	m = m.Where(dao.Content.Columns.UserId, r.Id)

	listModel := m.Page(r.Page, r.Size)
	switch r.Sort {
	case model.ContentSortHot:
		listModel = listModel.Order(dao.Content.Columns.ViewCount, "DESC")
	case model.ContentSortActive:
		listModel = listModel.Order(dao.Content.Columns.UpdatedAt, "DESC")
	default:
		listModel = listModel.Order(dao.Content.Columns.Id, "DESC")
	}
	contentEntities, err := listModel.M.All()
	if err != nil {
		return nil, err
	}
	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &model.UserServiceGetListRes{
		Page:  r.Page,
		Size:  r.Size,
		Total: total,
	}
	// Content
	if err := contentEntities.ScanList(&getListRes.List, "Content"); err != nil {
		return nil, err
	}
	// Category
	err = dao.Category.
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "CategoryId")).
		ScanList(&getListRes.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}
	// User
	err = dao.User.
		Fields(model.ContentListUserItem{}).
		Where(dao.User.Columns.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "UserId")).
		ScanList(&getListRes.List, "User", "Content", "id:UserId")
	if err != nil {
		return nil, err
	}

	userRecord, err := dao.User.Fields(getListRes.User).WherePri(r.Id).M.One()
	if err != nil {
		return nil, err
	}
	if err := userRecord.Struct(&getListRes.User); err != nil {
		return nil, err
	}

	return getListRes, nil
}
