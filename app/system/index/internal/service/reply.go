package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

// 评论/回复管理服务
var Reply = replyService{}

type replyService struct{}

// 创建回复
func (s *replyService) Create(ctx context.Context, r *define.ReplyServiceCreateReq) error {
	if r.UserId == 0 {
		r.UserId = shared.Context.Get(ctx).User.Id
	}
	_, err := dao.Reply.Data(r).Insert()
	if err == nil {
		_ = Content.AddReplyCount(ctx, r.TargetId, 1)
	}
	return err
}

// 删除
func (s *replyService) Delete(ctx context.Context, id uint) error {
	r, err := dao.Reply.WherePri(id).One()
	if err != nil {
		return err
	}
	_, err = dao.Reply.Where(g.Map{
		dao.Reply.Columns.Id:     id,
		dao.Reply.Columns.UserId: shared.Context.Get(ctx).User.Id,
	}).Delete()
	if err == nil {
		// 回复统计-1
		_ = Content.AddReplyCount(ctx, r.TargetId, -1)
		// 判断回复是否采纳
		c, err := dao.Content.WherePri(r.TargetId).One()
		if err == nil && c != nil && c.AdoptedReplyId == id {
			_ = Content.UnacceptedReply(ctx, r.TargetId)
		}
	}
	return err
}

// 获取回复列表
func (s *replyService) GetList(ctx context.Context, r *define.ReplyServiceGetListReq) (*define.ReplyServiceGetListRes, error) {
	m := dao.Reply.Fields(model.ReplyListItem{})

	if r.TargetType != "" {
		m = m.Where(dao.Reply.Columns.TargetType, r.TargetType)
	}
	if r.TargetId > 0 {
		m = m.Where(dao.Reply.Columns.TargetId, r.TargetId)
	}
	if r.UserId > 0 {
		m = m.Where(dao.Reply.Columns.UserId, r.UserId)
	}

	listModel := m.Page(r.Page, r.Size).Order(dao.Content.Columns.Id, "DESC")
	replyEntities, err := listModel.M.All()
	if err != nil {
		return nil, err
	}
	if replyEntities.IsEmpty() {
		return nil, nil
	}
	getListRes := &define.ReplyServiceGetListRes{
		Page: r.Page,
		Size: r.Size,
	}

	// User
	if err := replyEntities.ScanList(&getListRes.List, "Reply"); err != nil {
		return nil, err
	}
	err = dao.User.
		Fields(model.ReplyListUserItem{}).
		Where(dao.User.Columns.Id, gutil.ListItemValuesUnique(getListRes.List, "Reply", "UserId")).
		ScanList(&getListRes.List, "User", "Reply", "id:UserId")
	if err != nil {
		return nil, err
	}

	// Content
	err = dao.Content.Fields(dao.Content.Columns.Id, dao.Content.Columns.Title, dao.Content.Columns.CategoryId).
		Where(dao.Content.Columns.Id, gutil.ListItemValuesUnique(getListRes.List, "Reply", "TargetId")).
		ScanList(&getListRes.List, "Content", "Reply", "id:TargetId")
	if err != nil {
		return nil, err
	}

	// Category
	err = dao.Category.
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "CategoryId")).
		ScanList(&getListRes.List, "Category", "Content", "id:CategoryId")

	return getListRes, nil
}
