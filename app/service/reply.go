package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/util/gutil"
)

// 评论/回复管理服务
var Reply = new(replyService)

type replyService struct{}

// 创建回复
func (s *replyService) Create(ctx context.Context, r *model.ReplyServiceCreateReq) error {
	if r.UserId == 0 {
		r.UserId = Context.Get(ctx).User.Id
	}
	_, err := dao.Reply.Data(r).Insert()
	return err
}

// 获取回复列表
func (s *replyService) GetList(ctx context.Context, r *model.ReplyServiceGetListReq) (*model.ReplyServiceGetListRes, error) {
	m := dao.Reply.Fields(model.ReplyListItem{})
	if r.TargetType != "" {
		m = m.Where(dao.Reply.Columns.TargetType, r.TargetType)
	}
	if r.TargetId > 0 {
		m = m.Where(dao.Reply.Columns.TargetId, r.TargetId)
	}
	listModel := m.Page(r.Page, r.Size).Order(dao.Content.Columns.Id, "DESC")
	replyEntities, err := listModel.M.All()
	if err != nil {
		return nil, err
	}
	if replyEntities.IsEmpty() {
		return nil, nil
	}
	getListRes := &model.ReplyServiceGetListRes{
		Page: r.Page,
		Size: r.Size,
	}
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
	return getListRes, nil
}
