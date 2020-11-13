package api

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

var Article = new(articleApi)

type articleApi struct{}

// @summary 展示文章首页
// @tags    文章
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /article [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Index(r *ghttp.Request) {
	var (
		data *model.ContentServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	data.Type = model.ContentTypeArticle
	if getListRes, err := service.Content.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeArticle,
			Data:        getListRes,
		})
	}
}

// @summary 展示文章详情
// @tags    文章
// @produce html
// @param   id path int false "文章ID"
// @router  /article/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Detail(r *ghttp.Request) {
	var (
		data *model.ContentApiDetailReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	// 浏览次数增加
	service.Content.AddViewCount(r.Context(), data.Id, 1)
	if getDetailRes, err := service.Content.GetDetail(r.Context(), data.Id); err != nil {
		service.View.Render500(r)
	} else {
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeArticle,
			Data:        getDetailRes,
		})
	}
}
