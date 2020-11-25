package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/admin/internal/define"
	"github.com/gogf/gf/util/gconv"
)

// 权限管理服务
var Auth = new(authService)

type authService struct{}

// 查询权限列表，构造成树形返回
func (s *authService) GetTree(ctx context.Context) ([]*model.AuthTree, error) {
	entities, err := s.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return s.formTree(0, entities)
}

// 构造树形权限列表。
func (s *authService) formTree(parentId uint, entities []*model.Auth) ([]*model.AuthTree, error) {
	tree := make([]*model.AuthTree, 0)
	for _, entity := range entities {
		if entity.ParentId == parentId {
			subTree, err := s.formTree(entity.Id, entities)
			if err != nil {
				return nil, err
			}
			item := &model.AuthTree{
				Items: subTree,
			}
			if err = gconv.Struct(entity, item); err != nil {
				return nil, err
			}
			tree = append(tree, item)
		}
	}
	return tree, nil
}

// 获得所有的权限列表。
func (s *authService) GetList(ctx context.Context) ([]*model.Auth, error) {
	orderBy := fmt.Sprintf(
		`%s ASC, %s ASC`,
		dao.Auth.Columns.Sort,
		dao.Auth.Columns.Id,
	)
	return dao.Auth.Order(orderBy).All()
}

// 查询单个权限信息
func (s *authService) GetItem(ctx context.Context, id uint) (*model.Auth, error) {
	return dao.Auth.FindOne(id)
}

// 创建
func (s *authService) Create(ctx context.Context, r *define.AuthServiceCreateReq) error {
	if r.UserId == 0 {
		r.UserId = shared.Context.Get(ctx).User.Id
	}
	if _, err := dao.Auth.Data(r).Insert(); err != nil {
		return err
	}
	return nil
}

// 修改
func (s *authService) Update(ctx context.Context, r *define.AuthServiceUpdateReq) error {
	_, err := dao.Auth.Data(r).
		FieldsEx(dao.Auth.Columns.Id).
		Where(dao.Auth.Columns.Id, r.Id).
		Update()
	return err
}

// 删除
func (s *authService) Delete(ctx context.Context, id uint) error {
	_, err := dao.Auth.Where(dao.Auth.Columns.Id, id).Delete()
	return err
}
