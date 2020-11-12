package api

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/net/ghttp"
)

var Ask = new(askApi)

type askApi struct{}

// @summary 展示问答首页
// @tags    问答
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /ask [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Index(r *ghttp.Request) {
	var (
		data *model.ContentServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	data.Type = model.ContentTypeAsk
	if getListRes, err := service.Content.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeAsk,
			Data:        getListRes,
		})
	}
}

// @summary 展示问答详情
// @tags    问答
// @produce html
// @param   id path int false "问答ID"
// @router  /ask/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Detail(r *ghttp.Request) {
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
			ContentType: model.ContentTypeAsk,
			Data:        getDetailRes,
		})
	}
}

// @summary 展示创建问答页面
// @tags    问答
// @produce html
// @router  /ask/create [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Create(r *ghttp.Request) {
	service.View.Render(r, model.View{
		ContentType: model.ContentTypeAsk,
	})
}

// @summary 展示修改问答页面
// @tags    问答
// @produce html
// @param   id query int true "问答ID"
// @router  /ask/update [GET]
// @success 200 {string} html "页面HTML"
func (a *askApi) Update(r *ghttp.Request) {
	var (
		data *model.ContentApiUpdateReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Content.GetDetail(r.Context(), data.Id); err != nil {
		service.View.Render500(r)
	} else {
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeAsk,
			Data:        getDetailRes,
		})
	}
}
