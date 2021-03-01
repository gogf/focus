package service

import (
	"context"
	"fmt"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"
	"github.com/gogf/gf/encoding/ghtml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

// 内容管理服务
var Content = contentService{}

type contentService struct{}

// 查询内容列表
func (s *contentService) GetList(ctx context.Context, r *define.ContentServiceGetListReq) (*define.ContentServiceGetListRes, error) {
	m := dao.Content.Fields(model.ContentListItem{})

	// 默认查询topic
	if r.Type != "" {
		m = m.Where(dao.Content.Columns.Type, r.Type)
	} else {
		m = m.Where(dao.Content.Columns.Type, model.ContentTypeTopic)
	}

	if r.CategoryId > 0 {
		// 栏目检索
		idArray, err := Category.GetSubIdList(ctx, r.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.Columns.CategoryId, idArray)
	}

	// 管理员查看所有文章
	if r.UserId > 0 && !User.IsAdminShow(ctx, r.UserId) {
		m = m.Where(dao.Content.Columns.UserId, r.UserId)
	}
	listModel := m.Page(r.Page, r.Size)
	switch r.Sort {
	case model.ContentSortHot:
		listModel = listModel.Order(dao.Content.Columns.ViewCount, "DESC")
	case model.ContentSortActive:
		listModel = listModel.Order(dao.Content.Columns.UpdatedAt, "DESC")
	default:
		listModel = listModel.Order(dao.Content.Columns.Id, "DESC")
	}
	contentEntities, err := listModel.M.All()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if contentEntities.IsEmpty() {
		return nil, nil
	}
	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &define.ContentServiceGetListRes{
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

// 搜索内容列表
func (s *contentService) Search(ctx context.Context, r *define.ContentServiceSearchReq) (*define.ContentServiceSearchRes, error) {
	likePattern := `%` + r.Key + `%`
	m := dao.Content.Fields(model.ContentListItem{})
	m = m.Where(dao.Content.Columns.Content+" LIKE ?", likePattern).Or(dao.Content.Columns.Title+" LIKE ?", likePattern)
	// 内容类型
	if r.Type != "" {
		m = m.Where(dao.Content.Columns.Type, r.Type)
	}
	// 栏目检索
	if r.CategoryId > 0 {
		idArray, err := Category.GetSubIdList(ctx, r.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.Columns.CategoryId, idArray)
	}
	listModel := m.Page(r.Page, r.Size)
	switch r.Sort {
	case model.ContentSortHot:
		listModel = listModel.Order(dao.Content.Columns.ViewCount, "DESC")
	case model.ContentSortActive:
		listModel = listModel.Order(dao.Content.Columns.UpdatedAt, "DESC")
	case model.ContentSortScore:
		listModel = listModel.Order("score", "DESC")
	default:
		listModel = listModel.Order(dao.Content.Columns.Id, "DESC")
	}
	contentEntities, err := listModel.M.All()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if contentEntities.IsEmpty() {
		return &define.ContentServiceSearchRes{}, nil
	}
	countModel := m.Fields("*")
	total, err := countModel.Count()
	if err != nil {
		return nil, err
	}
	// 搜索统计
	statsModel := m.Fields(dao.Content.Columns.Type, "count(*) total").
		Group(dao.Content.Columns.Type)
	statsAll, err := statsModel.M.All()
	if err != nil {
		return nil, err
	}
	statsMap := make(map[string]int)
	for _, v := range statsAll {
		statsMap[v["type"].String()] = v["total"].Int()
	}
	getListRes := &define.ContentServiceSearchRes{
		Page:  r.Page,
		Size:  r.Size,
		Total: total,
		Stats: statsMap,
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
func (s *contentService) GetDetail(ctx context.Context, id uint) (*define.ContentServiceGetDetailRes, error) {
	getDetailRes := new(define.ContentServiceGetDetailRes)
	contentEntity, err := dao.Content.Fields(getDetailRes.Content).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	// 没有数据
	if contentEntity == nil {
		return nil, nil
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
func (s *contentService) Create(ctx context.Context, r *define.ContentServiceCreateReq) (*define.ContentServiceCreateRes, error) {
	if r.UserId == 0 {
		r.UserId = shared.Context.Get(ctx).User.Id
	}
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(r); err != nil {
		return nil, err
	}
	result, err := dao.Content.Data(r).Insert()
	if err != nil {
		return nil, err
	}
	n, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &define.ContentServiceCreateRes{ContentId: uint(n)}, err
}

// 修改
func (s *contentService) Update(ctx context.Context, r *define.ContentServiceUpdateReq) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(r); err != nil {
		return err
	}
	_, err := dao.Content.Data(r).
		FieldsEx(dao.Content.Columns.Id).
		Where(dao.Content.Columns.Id, r.Id).
		Where(dao.Content.Columns.UserId, shared.Context.Get(ctx).User.Id).
		Update()
	return err
}

// 删除
func (s *contentService) Delete(ctx context.Context, id uint) error {
	user := shared.Context.Get(ctx).User
	// 管理员直接删除文章和评论
	if user.IsAdmin {
		_, err := dao.Content.Where(g.Map{
			dao.Content.Columns.Id: id,
		}).Delete()
		if err == nil {
			_, err = dao.Reply.Where(g.Map{
				dao.Reply.Columns.TargetId: id,
			}).Delete()
		}
		return err
	}

	_, err := dao.Content.Where(g.Map{
		dao.Content.Columns.Id:     id,
		dao.Content.Columns.UserId: shared.Context.Get(ctx).User.Id,
	}).Delete()
	// 删除评论
	if err == nil {
		_, err = dao.Reply.Where(g.Map{
			dao.Reply.Columns.TargetId: id,
			dao.Reply.Columns.UserId:   shared.Context.Get(ctx).User.Id,
		}).Delete()
	}
	return err
}

// 浏览次数增加
func (s *contentService) AddViewCount(ctx context.Context, id uint, count int) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.Columns.ViewCount, dao.Content.Columns.ViewCount, count)).
		WherePri(id).Update()
	if err != nil {
		return err
	}
	return nil
}

// 回复次数增加
func (s *contentService) AddReplyCount(ctx context.Context, id uint, count int) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=IFNULL(%s,0)+%d`, dao.Content.Columns.ReplyCount, dao.Content.Columns.ReplyCount, count)).
		WherePri(id).Update()
	if err != nil {
		return err
	}
	return nil
}

// 采纳回复
func (s *contentService) AdoptReply(ctx context.Context, id uint, replyId uint) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=%d`, dao.Content.Columns.AdoptedReplyId, replyId)).
		WherePri(id).
		Where(dao.Content.Columns.UserId, shared.Context.Get(ctx).User.Id).
		Update()
	if err != nil {
		return err
	}
	return nil
}

// 取消采纳回复
func (s *contentService) UnacceptedReply(ctx context.Context, id uint) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=%d`, dao.Content.Columns.AdoptedReplyId, 0)).
		WherePri(id).
		Update()
	if err != nil {
		return err
	}
	return nil
}
