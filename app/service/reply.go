package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
)

var Reply = new(replyService)

type replyService struct{}


// 创建
func (s *replyService) Create(ctx context.Context, r *model.ReplyServiceCreateReq) error {
	if r.UserId == 0 {
		r.UserId = Context.Get(ctx).User.Id
	}
	_, err := dao.Reply.Data(r).Insert()
	return err
}