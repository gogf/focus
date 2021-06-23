package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"
	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

// 评论/回复管理服务
var Reply = replyService{}

type replyService struct{}

// 创建回复
func (s *replyService) Create(ctx context.Context, r *define.ReplyServiceCreateReq) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if r.UserId == 0 {
			r.UserId = shared.Context.Get(ctx).User.Id
		}
		_, err := dao.Reply.Ctx(ctx).Data(r).Insert()
		if err == nil {
			_ = Content.AddReplyCount(ctx, r.TargetId, 1)
		}
		return err
	})
}

// 删除
func (s *replyService) Delete(ctx context.Context, id uint) error {
	return dao.Reply.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var reply *model.Reply
		err := dao.Reply.Ctx(ctx).WherePri(id).Scan(&reply)
		if err != nil {
			return err
		}
		_, err = dao.Reply.Ctx(ctx).Where(g.Map{
			dao.Reply.C.Id:     id,
			dao.Reply.C.UserId: shared.Context.Get(ctx).User.Id,
		}).Delete()
		if err == nil {
			// 回复统计-1
			_ = Content.AddReplyCount(ctx, reply.TargetId, -1)
			// 判断回复是否采纳
			var content *model.Content
			err := dao.Content.WherePri(reply.TargetId).Scan(&content)
			if err == nil && content != nil && content.AdoptedReplyId == id {
				_ = Content.UnacceptedReply(ctx, reply.TargetId)
			}
		}
		return err
	})
}

// 获取回复列表
func (s *replyService) GetList(ctx context.Context, r *define.ReplyServiceGetListReq) (*define.ReplyServiceGetListRes, error) {
	var result = &define.ReplyServiceGetListRes{}
	m := dao.Reply.Ctx(ctx).Fields(model.ReplyListItem{})

	if r.TargetType != "" {
		m = m.Where(dao.Reply.C.TargetType, r.TargetType)
	}
	if r.TargetId > 0 {
		m = m.Where(dao.Reply.C.TargetId, r.TargetId)
	}
	if r.UserId > 0 {
		m = m.Where(dao.Reply.C.UserId, r.UserId)
	}

	err := m.Page(r.Page, r.Size).OrderDesc(dao.Content.C.Id).ScanList(&result.List, "Reply")
	if err != nil {
		return nil, err
	}
	if len(result.List) == 0 {
		return nil, nil
	}
	getListRes := &define.ReplyServiceGetListRes{
		Page: r.Page,
		Size: r.Size,
	}

	// User
	if err := m.ScanList(&getListRes.List, "Reply"); err != nil {
		return nil, err
	}
	err = dao.User.
		Fields(model.ReplyListUserItem{}).
		Where(dao.User.C.Id, gutil.ListItemValuesUnique(getListRes.List, "Reply", "UserId")).
		ScanList(&getListRes.List, "User", "Reply", "id:UserId")
	if err != nil {
		return nil, err
	}

	// Content
	err = dao.Content.Fields(dao.Content.C.Id, dao.Content.C.Title, dao.Content.C.CategoryId).
		Where(dao.Content.C.Id, gutil.ListItemValuesUnique(getListRes.List, "Reply", "TargetId")).
		ScanList(&getListRes.List, "Content", "Reply", "id:TargetId")
	if err != nil {
		return nil, err
	}

	// Category
	err = dao.Category.
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.C.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "CategoryId")).
		ScanList(&getListRes.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}

	return getListRes, nil
}
