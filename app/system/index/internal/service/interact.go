package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// 交互管理服务
var Interact = interactService{}

type interactService struct{}

const (
	contextMapKeyForMyInteractList = "ContextMapKeyForMyInteractList"
)

// 赞
func (s *interactService) Zan(ctx context.Context, targetType string, targetId uint) error {
	customCtx := shared.Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Data(&model.Interact{
		UserId:     customCtx.User.Id,
		TargetId:   targetId,
		TargetType: targetType,
		Type:       model.InteractTypeZan,
	}).FieldsEx(dao.Interact.Columns.Id).InsertIgnore()
	if err != nil {
		return err
	}

	if n, _ := r.RowsAffected(); n == 0 {
		return gerror.New("您已经赞过啦")
	}
	return s.updateCount(ctx, model.InteractTypeZan, targetType, targetId, 1)
}

// 取消赞
func (s *interactService) CancelZan(ctx context.Context, targetType string, targetId uint) error {
	customCtx := shared.Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, shared.Context.Get(ctx).User.Id,
		dao.Interact.Columns.TargetId, targetId,
		dao.Interact.Columns.TargetType, targetType,
		dao.Interact.Columns.Type, model.InteractTypeZan,
	}).Delete()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return nil
	}
	return s.updateCount(ctx, model.InteractTypeZan, targetType, targetId, -1)
}

// 我是否有对指定内容赞
func (s *interactService) DidIZan(ctx context.Context, targetType string, targetId uint) (bool, error) {
	list, err := s.getMyList(ctx)
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v.TargetId == targetId && v.TargetType == targetType && v.Type == model.InteractTypeZan {
			return true, nil
		}
	}
	return false, nil
}

// 踩
func (s *interactService) Cai(ctx context.Context, targetType string, targetId uint) error {
	customCtx := shared.Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Data(&model.Interact{
		UserId:     customCtx.User.Id,
		TargetId:   targetId,
		TargetType: targetType,
		Type:       model.InteractTypeCai,
	}).FieldsEx(dao.Interact.Columns.Id).InsertIgnore()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return gerror.New("您已经踩过啦")
	}
	return s.updateCount(ctx, model.InteractTypeCai, targetType, targetId, 1)
}

// 取消踩
func (s *interactService) CancelCai(ctx context.Context, targetType string, targetId uint) error {
	customCtx := shared.Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, shared.Context.Get(ctx).User.Id,
		dao.Interact.Columns.TargetId, targetId,
		dao.Interact.Columns.TargetType, targetType,
		dao.Interact.Columns.Type, model.InteractTypeCai,
	}).Delete()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return nil
	}
	return s.updateCount(ctx, model.InteractTypeCai, targetType, targetId, -1)
}

// 我是否有对指定内容踩
func (s *interactService) DidICai(ctx context.Context, targetType string, targetId uint) (bool, error) {
	list, err := s.getMyList(ctx)
	if err != nil {
		return false, err
	}
	for _, v := range list {
		if v.TargetId == targetId && v.TargetType == targetType && v.Type == model.InteractTypeCai {
			return true, nil
		}
	}
	return false, nil
}

// 获得我的互动数据列表，内部带请求上下文缓存
func (s *interactService) getMyList(ctx context.Context) ([]*model.Interact, error) {
	customCtx := shared.Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil, nil
	}
	if v, ok := customCtx.Data[contextMapKeyForMyInteractList]; ok {
		return v.([]*model.Interact), nil
	}
	all, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, customCtx.User.Id,
	}).All()
	if err != nil {
		return nil, err
	}
	customCtx.Data[contextMapKeyForMyInteractList] = all
	return all, err
}

func (s *interactService) updateCount(ctx context.Context, interactType int, targetType string, targetId uint, count int) error {
	defer func() {
		// 清空上下文对应的互动数据缓存
		if customCtx := shared.Context.Get(ctx); customCtx != nil {
			delete(customCtx.Data, contextMapKeyForMyInteractList)
		}
	}()
	var err error
	switch targetType {
	case model.InteractTargetTypeContent:
		switch interactType {
		case model.InteractTypeZan:
			_, err = dao.Content.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.ZanCount, dao.Content.Columns.ZanCount, count)).
				Where(dao.Content.Columns.Id, targetId).Where(dao.Content.Columns.ZanCount + ">=0").Update()
			if err != nil {
				return err
			}

		case model.InteractTypeCai:
			_, err = dao.Content.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.CaiCount, dao.Content.Columns.CaiCount, count)).
				Where(dao.Content.Columns.Id, targetId).Where(dao.Content.Columns.CaiCount + ">=0").Update()
			if err != nil {
				return err
			}
		}

	case model.InteractTargetTypeReply:
		switch interactType {
		case model.InteractTypeZan:
			_, err = dao.Reply.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Reply.Columns.ZanCount, dao.Reply.Columns.ZanCount, count)).
				Where(dao.Reply.Columns.Id, targetId).Where(dao.Reply.Columns.ZanCount + ">=0").Update()
			if err != nil {
				return err
			}

		case model.InteractTypeCai:
			_, err = dao.Reply.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Reply.Columns.CaiCount, dao.Reply.Columns.CaiCount, count)).
				Where(dao.Reply.Columns.Id, targetId).Where(dao.Reply.Columns.CaiCount + ">=0").Update()
			if err != nil {
				return err
			}

		}
	}
	return nil
}
