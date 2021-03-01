package api

import (
	"focus/app/model"
	"focus/app/system/index/internal/define"
	"focus/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Topic = topicApi{}

type topicApi struct{}

// @summary 展示主题首页
// @tags    前台-主题
// @produce html
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /topic [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Index(r *ghttp.Request) {
	var (
		data *define.ContentServiceGetListReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	data.Type = model.ContentTypeTopic
	if getListRes, err := service.Content.GetList(r.Context(), data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			ContentType: data.Type,
			Data:        getListRes,
			Title: service.View.GetTitle(r.Context(), &define.ViewServiceGetTitleReq{
				ContentType: data.Type,
				CategoryId:  data.CategoryId,
			}),
		})
	}
}

// @summary 展示主题详情
// @tags    前台-主题
// @produce html
// @param   id path int false "主题ID"
// @router  /topic/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *topicApi) Detail(r *ghttp.Request) {
	var (
		data *define.ContentApiDetailReq
	)
	if err := r.Parse(&data); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if getDetailRes, err := service.Content.GetDetail(r.Context(), data.Id); err != nil {
		service.View.Render500(r)
	} else {
		if getDetailRes == nil {
			service.View.Render404(r)
		}
		service.Content.AddViewCount(r.Context(), data.Id, 1)
		service.View.Render(r, model.View{
			ContentType: model.ContentTypeTopic,
			Data:        getDetailRes,
			Title: service.View.GetTitle(r.Context(), &define.ViewServiceGetTitleReq{
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
				CurrentName: getDetailRes.Content.Title,
			}),
			BreadCrumb: service.View.GetBreadCrumb(r.Context(), &define.ViewServiceGetBreadCrumbReq{
				ContentId:   getDetailRes.Content.Id,
				ContentType: getDetailRes.Content.Type,
				CategoryId:  getDetailRes.Content.CategoryId,
			}),
		})
	}
}
