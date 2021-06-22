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
		m = m.Where(dao.Content.C.Type, r.Type)
	} else {
		m = m.Where(dao.Content.C.Type, model.ContentTypeTopic)
	}
	// 栏目检索
	if r.CategoryId > 0 {
		idArray, err := Category.GetSubIdList(ctx, r.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.C.CategoryId, idArray)
	}

	// 管理员查看所有文章
	if r.UserId > 0 && !User.IsAdminShow(ctx, r.UserId) {
		m = m.Where(dao.Content.C.UserId, r.UserId)
	}
	listModel := m.Page(r.Page, r.Size)
	switch r.Sort {
	case model.ContentSortHot:
		listModel = listModel.OrderDesc(dao.Content.C.ViewCount)
	case model.ContentSortActive:
		listModel = listModel.OrderDesc(dao.Content.C.UpdatedAt)
	default:
		listModel = listModel.OrderDesc(dao.Content.C.Id)
	}

	var list []*model.Content
	if err := listModel.Scan(&list); err != nil {
		return nil, err
	}
	// 没有数据
	if len(list) == 0 {
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
	if err := listModel.ScanList(&getListRes.List, "Content"); err != nil {
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
	// User
	err = dao.User.
		Fields(model.ContentListUserItem{}).
		Where(dao.User.C.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "UserId")).
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
	m = m.WhereLike(dao.Content.C.Content, likePattern).WhereOrLike(dao.Content.C.Title, likePattern)
	// 内容类型
	if r.Type != "" {
		m = m.Where(dao.Content.C.Type, r.Type)
	}
	// 栏目检索
	if r.CategoryId > 0 {
		idArray, err := Category.GetSubIdList(ctx, r.CategoryId)
		if err != nil {
			return nil, err
		}
		m = m.Where(dao.Content.C.CategoryId, idArray)
	}
	listModel := m.Page(r.Page, r.Size)
	switch r.Sort {
	case model.ContentSortHot:
		listModel = listModel.Order(dao.Content.C.ViewCount, "DESC")
	case model.ContentSortActive:
		listModel = listModel.Order(dao.Content.C.UpdatedAt, "DESC")
	case model.ContentSortScore:
		listModel = listModel.Order("score", "DESC")
	default:
		listModel = listModel.Order(dao.Content.C.Id, "DESC")
	}
	contentEntities, err := listModel.All()
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
	statsModel := m.Fields(dao.Content.C.Type, "count(*) total").Group(dao.Content.C.Type)
	statsAll, err := statsModel.All()
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
		Where(dao.Category.C.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "CategoryId")).
		ScanList(&getListRes.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return nil, err
	}
	// User
	err = dao.User.
		Fields(model.ContentListUserItem{}).
		Where(dao.User.C.Id, gutil.ListItemValuesUnique(getListRes.List, "Content", "UserId")).
		ScanList(&getListRes.List, "User", "Content", "id:UserId")
	if err != nil {
		return nil, err
	}

	return getListRes, nil
}

// 查询详情
func (s *contentService) GetDetail(ctx context.Context, id uint) (*define.ContentServiceGetDetailRes, error) {
	result := &define.ContentServiceGetDetailRes{}
	if err := dao.Content.Ctx(ctx).WherePri(id).Scan(&result.Content); err != nil {
		return nil, err
	}
	// 没有数据
	if result.Content == nil {
		return nil, nil
	}
	err := dao.User.Ctx(ctx).WherePri(result.Content.UserId).Scan(result.User)
	if err != nil {
		return nil, err
	}
	return result, nil
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
		FieldsEx(dao.Content.C.Id).
		Where(dao.Content.C.Id, r.Id).
		Where(dao.Content.C.UserId, shared.Context.Get(ctx).User.Id).
		Update()
	return err
}

// 删除
func (s *contentService) Delete(ctx context.Context, id uint) error {
	user := shared.Context.Get(ctx).User
	// 管理员直接删除文章和评论
	if user.IsAdmin {
		_, err := dao.Content.Where(g.Map{
			dao.Content.C.Id: id,
		}).Delete()
		if err == nil {
			_, err = dao.Reply.Where(g.Map{
				dao.Reply.C.TargetId: id,
			}).Delete()
		}
		return err
	}

	_, err := dao.Content.Where(g.Map{
		dao.Content.C.Id:     id,
		dao.Content.C.UserId: shared.Context.Get(ctx).User.Id,
	}).Delete()
	// 删除评论
	if err == nil {
		_, err = dao.Reply.Where(g.Map{
			dao.Reply.C.TargetId: id,
			dao.Reply.C.UserId:   shared.Context.Get(ctx).User.Id,
		}).Delete()
	}
	return err
}

// 浏览次数增加
func (s *contentService) AddViewCount(ctx context.Context, id uint, count int) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=%s+%d`, dao.Content.C.ViewCount, dao.Content.C.ViewCount, count)).
		WherePri(id).Update()
	if err != nil {
		return err
	}
	return nil
}

// 回复次数增加
func (s *contentService) AddReplyCount(ctx context.Context, id uint, count int) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=IFNULL(%s,0)+%d`, dao.Content.C.ReplyCount, dao.Content.C.ReplyCount, count)).
		WherePri(id).Update()
	if err != nil {
		return err
	}
	return nil
}

// 采纳回复
func (s *contentService) AdoptReply(ctx context.Context, id uint, replyId uint) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=%d`, dao.Content.C.AdoptedReplyId, replyId)).
		WherePri(id).
		Where(dao.Content.C.UserId, shared.Context.Get(ctx).User.Id).
		Update()
	if err != nil {
		return err
	}
	return nil
}

// 取消采纳回复
func (s *contentService) UnacceptedReply(ctx context.Context, id uint) error {
	_, err := dao.Content.
		Data(fmt.Sprintf(`%s=%d`, dao.Content.C.AdoptedReplyId, 0)).
		WherePri(id).
		Update()
	if err != nil {
		return err
	}
	return nil
}
