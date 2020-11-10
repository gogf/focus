package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

var Content = new(contentService)

type contentService struct{}

// 查询内容列表
func (s *contentService) GetList(ctx context.Context, r *model.ContentServiceGetListReq) (*model.ContentServiceGetListRes, error) {
	m := dao.Content.Fields(model.ContentListItem{})
	m = m.Where(dao.Content.Columns.Type, r.Type)
	if r.CategoryId > 0 {
		// 栏目检索
		idArray, err := Category.GetSubIdList(ctx, r.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.Columns.CategoryId, idArray)
	}
	listModel := m.Page(r.Page, r.Size)
	switch r.Sort {
	case model.ContentSortHot:
		listModel = listModel.Order(dao.Content.Columns.ZanCount, "DESC")
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
	getListRes := &model.ContentServiceGetListRes{
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
	return getListRes, nil
}

// 查询详情
func (s *contentService) GetDetail(ctx context.Context, id uint) (*model.ContentServiceGetDetailRes, error) {
	getDetailRes := new(model.ContentServiceGetDetailRes)
	contentEntity, err := dao.Content.Fields(getDetailRes.Content).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	userRecord, err := dao.User.Fields(getDetailRes.User).WherePri(contentEntity.UserId).M.One()
	if err != nil {
		return nil, err
	}
	getDetailRes.Content = contentEntity
	if err := userRecord.Struct(&getDetailRes.User); err != nil {
		return nil, err
	}
	return getDetailRes, nil
}

// 创建
func (s *contentService) Create(ctx context.Context, r *model.ContentServiceCreateReq) error {
	if r.UserId == 0 {
		r.UserId = Context.Get(ctx).User.Id
	}
	_, err := dao.Content.Data(r).Insert()
	return err
}

// 修改
func (s *contentService) Update(ctx context.Context, r *model.ContentServiceUpdateReq) error {
	_, err := dao.Content.Data(r).
		FieldsEx(dao.Content.Columns.Id).
		Where(dao.Content.Columns.UserId, Context.Get(ctx).User.Id).Update()
	return err
}

// 删除
func (s *contentService) Delete(ctx context.Context, id uint) error {
	_, err := dao.Content.Where(g.Map{
		dao.Content.Columns.Id:     id,
		dao.Content.Columns.UserId: Context.Get(ctx).User.Id,
	}).Delete()
	return err
}
