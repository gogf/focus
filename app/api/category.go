package api

import (
	"focus/app/dao"
	"focus/app/model"
	"focus/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type categoryApi struct{}

var Category = new(categoryApi)

// @summary 获取分类列表，构造成树形结构返回
// @tags    分类
// @produce json
// @param   contentType query string true  "分类类型:topic, ask, article, reply"
// @param   parentId    query int    false "父级分类ID"
// @router  /category/list [GET]
// @success 200 {array} model.CategoryItem "分类列表"
func (a *categoryApi) List(r *ghttp.Request) {
	var (
		data          *model.CategoryApiListReq
		daoGetListReq *model.CategoryDaoGetListReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(data, &daoGetListReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if list, err := dao.Category.GetList(daoGetListReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK", list)
	}
}

// @summary 获取指定分类的详情信息
// @tags    分类
// @produce json
// @param   id query string true "分类ID"
// @router  /category/item [GET]
// @success 200 {object} model.CategoryItem "分类详情"
func (a *categoryApi) Item(r *ghttp.Request) {
	var (
		data *model.CategoryApiItemReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if item, err := dao.Category.GetItem(data.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "OK", item)
	}
}
