package service

import (
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/util/gconv"
)

var Category = new(categoryService)

type categoryService struct{}

// 查询列表
func (s *categoryService) GetList(r *model.CategoryServiceGetListReq) ([]*model.CategoryItem, error) {
	m := dao.Category.Where(dao.Category.Columns.ContentType, r.ContentType)
	if r.ParentId > 0 {
		m = m.Where(dao.Category.Columns.ParentId, r.ParentId)
	}
	m = m.Order(dao.Category.Columns.Sort, "ASC")
	// 查询数据
	all, err := m.All()
	if err != nil {
		return nil, err
	}
	// 结构体转换
	list := make([]*model.CategoryItem, len(all))
	for i, v := range all {
		list[i], err = s.entityToItem(v)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

// 查询详情
func (s *categoryService) GetItem(id uint) (*model.CategoryItem, error) {
	entity, err := dao.Category.FindOne(id)
	if err != nil {
		return nil, err
	}
	return s.entityToItem(entity)
}

// 将ORM Entity装换为Item对象返回
func (s *categoryService) entityToItem(entity *model.Category) (*model.CategoryItem, error) {
	if entity == nil {
		return nil, nil
	}
	item := &model.CategoryItem{}
	if err := gconv.Struct(entity, item); err != nil {
		return nil, err
	}
	return item, nil
}
