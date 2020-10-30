package category

import (
	"focus/app/model/category"
	"github.com/gogf/gf/util/gconv"
)

// 查询列表
func GetList(param *GetListParam) ([]*Item, error) {
	model := category.Model
	model = model.Where(category.Columns.ContentType, param.ContentType)
	if param.ParentId > 0 {
		model = model.Where(category.Columns.ParentId, param.ParentId)
	}
	model = model.Order(category.Columns.Sort, "ASC")
	// 查询数据
	all, err := model.All()
	if err != nil {
		return nil, err
	}
	// 结构体转换
	list := make([]*Item, len(all))
	for i, v := range all {
		list[i], err = entityToItem(v)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

// 查询详情
func GetItem(id uint) (*Item, error) {
	entity, err := category.Model.FindOne(id)
	if err != nil {
		return nil, err
	}
	return entityToItem(entity)
}

// 将ORM Entity装换为Item对象返回
func entityToItem(entity *category.Entity) (*Item, error) {
	if entity == nil {
		return nil, nil
	}
	item := &Item{}
	if err := gconv.Struct(entity, item); err != nil {
		return nil, err
	}
	return item, nil
}
