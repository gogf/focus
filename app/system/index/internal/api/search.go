package api

import (
	"focus/app/model"
	"focus/app/system/index/internal/define"
	"focus/app/system/index/internal/service"
	"github.com/gogf/gf/net/ghttp"
)

var Search = searchApi{}

type searchApi struct{}

// @summary 搜索页面
// @tags    前台-搜索
// @produce html
// @param   key  query string true  "关键字"
// @param   cate query int    false "栏目ID"
// @param   page query int    false "分页号码"
// @param   size query int    false "分页数量"
// @param   sort query string false "排序方式"
// @router  /search [GET]
// @success 200 {string} html "页面HTML"
func (a *searchApi) Index(r *ghttp.Request) {
	var (
		req *define.ContentServiceSearchReq
	)
	if err := r.Parse(&req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	if searchRes, err := service.Content.Search(r.Context(), req); err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	} else {
		service.View.Render(r, model.View{
			Data:  searchRes,
			Title: service.View.GetTitle(r.Context(), &define.ViewServiceGetTitleReq{}),
		})
	}
}
