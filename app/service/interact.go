package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// 交互管理器
var Interact = &interactService{}

type interactService struct{}

// 赞
func (s *interactService) Zan(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Data(&model.Interact{
		UserId:      customCtx.User.Id,
		ContentId:   contentId,
		ContentType: contentType,
		Type:        model.InteractTypeZan,
	}).FieldsEx(dao.Interact.Columns.Id).InsertIgnore()
	if err != nil {
		return err
	}

	if n, _ := r.RowsAffected(); n == 0 {
		return gerror.New("您已经赞过啦")
	}
	return s.updateContentCount(model.InteractTypeZan, contentType, contentId, 1)
}

// 取消赞
func (s *interactService) CancelZan(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, Context.Get(ctx).User.Id,
		dao.Interact.Columns.ContentId, contentId,
		dao.Interact.Columns.ContentType, contentType,
		dao.Interact.Columns.Type, model.InteractTypeZan,
	}).Delete()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return nil
	}
	return s.updateContentCount(model.InteractTypeZan, contentType, contentId, -1)
}

// 我是否有对指定内容赞
func (s *interactService) DidIZan(ctx context.Context, contentType string, contentId uint) (bool, error) {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return false, nil
	}
	n, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, Context.Get(ctx).User.Id,
		dao.Interact.Columns.ContentId, contentId,
		dao.Interact.Columns.ContentType, contentType,
		dao.Interact.Columns.Type, model.InteractTypeZan,
	}).Count()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// 踩
func (s *interactService) Cai(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Data(&model.Interact{
		UserId:      customCtx.User.Id,
		ContentId:   contentId,
		ContentType: contentType,
		Type:        model.InteractTypeCai,
	}).FieldsEx(dao.Interact.Columns.Id).InsertIgnore()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return gerror.New("您已经踩过啦")
	}
	return s.updateContentCount(model.InteractTypeCai, contentType, contentId, 1)
}

// 取消踩
func (s *interactService) CancelCai(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, Context.Get(ctx).User.Id,
		dao.Interact.Columns.ContentId, contentId,
		dao.Interact.Columns.ContentType, contentType,
		dao.Interact.Columns.Type, model.InteractTypeCai,
	}).Delete()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return nil
	}
	return s.updateContentCount(model.InteractTypeCai, contentType, contentId, -1)
}

// 我是否有对指定内容赞
func (s *interactService) DidICai(ctx context.Context, contentType string, contentId uint) (bool, error) {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return false, nil
	}
	n, err := dao.Interact.Where(g.Slice{
		dao.Interact.Columns.UserId, Context.Get(ctx).User.Id,
		dao.Interact.Columns.ContentId, contentId,
		dao.Interact.Columns.ContentType, contentType,
		dao.Interact.Columns.Type, model.InteractTypeCai,
	}).Count()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func (s *interactService) updateContentCount(interactType int, contentType string, contentId uint, count int) error {
	var err error
	switch contentType {
	case model.InteractContentTypeContent:
		switch interactType {
		case model.InteractTypeZan:
			_, err = dao.Content.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.ZanCount, dao.Content.Columns.ZanCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}

		case model.InteractTypeCai:
			_, err = dao.Content.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.CaiCount, dao.Content.Columns.CaiCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}
		}

	case model.InteractContentTypeReply:
		switch interactType {
		case model.InteractTypeZan:
			_, err = dao.Reply.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Reply.Columns.ZanCount, dao.Reply.Columns.ZanCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}

		case model.InteractTypeCai:
			_, err = dao.Reply.
				Data(fmt.Sprintf(`%s=%s%d`, dao.Reply.Columns.CaiCount, dao.Reply.Columns.CaiCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}

		}
	}
	return nil
}
