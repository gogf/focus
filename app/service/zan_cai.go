package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

// 管理赞、踩、浏览操作
var ZanCai = &zanCaiService{}

type zanCaiService struct{}

// 赞
func (s *zanCaiService) Zan(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.ZanCai.Data(&model.ZanCai{
		UserId:      customCtx.User.Id,
		ContentId:   contentId,
		ContentType: contentType,
		Type:        model.ZanCaiTypeZan,
	}).FieldsEx(dao.ZanCai.Columns.Id).InsertIgnore()
	if err != nil {
		return err
	}

	if n, _ := r.RowsAffected(); n == 0 {
		return gerror.New("您已经赞过啦")
	}
	return s.updateContentCount(model.ZanCaiTypeZan, contentType, contentId, 1)
}

// 取消赞
func (s *zanCaiService) CancelZan(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.ZanCai.Where(g.Slice{
		dao.ZanCai.Columns.UserId, Context.Get(ctx).User.Id,
		dao.ZanCai.Columns.ContentId, contentId,
		dao.ZanCai.Columns.ContentType, contentType,
		dao.ZanCai.Columns.Type, model.ZanCaiTypeZan,
	}).Delete()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return nil
	}
	return s.updateContentCount(model.ZanCaiTypeZan, contentType, contentId, -1)
}

// 我是否有对指定内容赞
func (s *zanCaiService) DidIZan(ctx context.Context, contentType string, contentId uint) (bool, error) {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return false, nil
	}
	n, err := dao.ZanCai.Where(g.Slice{
		dao.ZanCai.Columns.UserId, Context.Get(ctx).User.Id,
		dao.ZanCai.Columns.ContentId, contentId,
		dao.ZanCai.Columns.ContentType, contentType,
		dao.ZanCai.Columns.Type, model.ZanCaiTypeZan,
	}).Count()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// 踩
func (s *zanCaiService) Cai(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.ZanCai.Data(&model.ZanCai{
		UserId:      customCtx.User.Id,
		ContentId:   contentId,
		ContentType: contentType,
		Type:        model.ZanCaiTypeCai,
	}).FieldsEx(dao.ZanCai.Columns.Id).InsertIgnore()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return gerror.New("您已经踩过啦")
	}
	return s.updateContentCount(model.ZanCaiTypeCai, contentType, contentId, 1)
}

// 取消踩
func (s *zanCaiService) CancelCai(ctx context.Context, contentType string, contentId uint) error {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return nil
	}
	r, err := dao.ZanCai.Where(g.Slice{
		dao.ZanCai.Columns.UserId, Context.Get(ctx).User.Id,
		dao.ZanCai.Columns.ContentId, contentId,
		dao.ZanCai.Columns.ContentType, contentType,
		dao.ZanCai.Columns.Type, model.ZanCaiTypeCai,
	}).Delete()
	if err != nil {
		return err
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return nil
	}
	return s.updateContentCount(model.ZanCaiTypeCai, contentType, contentId, -1)
}

// 我是否有对指定内容赞
func (s *zanCaiService) DidICai(ctx context.Context, contentType string, contentId uint) (bool, error) {
	customCtx := Context.Get(ctx)
	if customCtx == nil || customCtx.User == nil {
		return false, nil
	}
	n, err := dao.ZanCai.Where(g.Slice{
		dao.ZanCai.Columns.UserId, Context.Get(ctx).User.Id,
		dao.ZanCai.Columns.ContentId, contentId,
		dao.ZanCai.Columns.ContentType, contentType,
		dao.ZanCai.Columns.Type, model.ZanCaiTypeCai,
	}).Count()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func (s *zanCaiService) updateContentCount(zanCaiType int, contentType string, contentId uint, count int) error {
	var err error
	switch contentType {
	case model.ZanCaiContentTypeContent:
		switch zanCaiType {
		case model.ZanCaiTypeZan:
			_, err = dao.Content.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.ZanCount, dao.Content.Columns.ZanCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}

		case model.ZanCaiTypeCai:
			_, err = dao.Content.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.CaiCount, dao.Content.Columns.CaiCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}
		}

	case model.ZanCaiContentTypeReply:
		switch zanCaiType {
		case model.ZanCaiTypeZan:
			_, err = dao.Reply.
				Data(fmt.Sprintf(`%s=%s+%d`, dao.Reply.Columns.ZanCount, dao.Reply.Columns.ZanCount, count)).
				WherePri(contentId).Update()
			if err != nil {
				return err
			}

		case model.ZanCaiTypeCai:
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
