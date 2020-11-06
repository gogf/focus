package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gutil"
)

var Topic = new(topicService)

type topicService struct{}

// 查询列表
func (s *topicService) GetList(ctx context.Context, r *model.TopicServiceGetListReq) (*model.TopicServiceGetListRes, error) {
	if r.Size > model.TopicListMaxSize {
		r.Size = model.TopicListMaxSize
	}
	m := dao.Topic.Fields(model.TopicListItem{})
	m = m.Page(r.Page, r.Size)
	switch gstr.ToLower(r.Sort) {
	case "zan":
		m = m.Order(dao.Topic.Columns.ZanCount, "DESC")
	case "active":
		m = m.Order(dao.Topic.Columns.UpdatedAt, "DESC")
	default:
		m = m.Order(dao.Topic.Columns.Id, "DESC")
	}
	topicEntities, err := m.M.All()
	if err != nil {
		return nil, err
	}
	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &model.TopicServiceGetListRes{
		Page:  r.Page,
		Size:  r.Size,
		Total: total,
	}
	if err := topicEntities.ScanList(&getListRes.List, "Topic"); err != nil {
		return nil, err
	}
	err = dao.User.
		Fields(model.TopicUserItem{}).
		Where(dao.User.Columns.Id, gutil.ListItemValuesUnique(getListRes.List, "Topic", "UserId")).
		ScanList(&getListRes.List, "User", "Topic", "id:UserId")
	if err != nil {
		return nil, err
	}
	return getListRes, nil
}

// 查询详情
func (s *topicService) GetDetail(ctx context.Context, id uint) (*model.TopicServiceGetDetailRes, error) {
	getDetailRes := new(model.TopicServiceGetDetailRes)
	topicEntity, err := dao.Topic.Fields(getDetailRes.Topic).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	userRecord, err := dao.User.Fields(getDetailRes.User).WherePri(topicEntity.UserId).M.One()
	if err != nil {
		return nil, err
	}
	getDetailRes.Topic = topicEntity
	if err := userRecord.Struct(&getDetailRes.User); err != nil {
		return nil, err
	}
	return getDetailRes, nil
}

// 创建
func (s *topicService) Create(ctx context.Context, r *model.TopicServiceCreateReq) error {
	if r.UserId == 0 {
		r.UserId = Context.GetCtx(ctx).UserId
	}
	_, err := dao.Topic.Data(r).Save()
	return err
}

// 修改
func (s *topicService) Update(ctx context.Context, r *model.TopicServiceUpdateReq) error {
	_, err := dao.Topic.Data(r).Where(
		dao.Topic.Columns.UserId, Context.GetCtx(ctx).UserId,
	).Save()
	return err
}

// 删除
func (s *topicService) Delete(ctx context.Context, id uint) error {
	_, err := dao.Topic.Where(g.Map{
		dao.Topic.Columns.Id:     id,
		dao.Topic.Columns.UserId: Context.GetCtx(ctx).UserId,
	}).Delete()
	return err
}
