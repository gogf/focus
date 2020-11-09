package api

import (
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Article = new(articleApi)

type articleApi struct{}

// @summary 展示文章板块页面
// @tags    文章
// @produce html
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
			Data: getListRes,
		})
	}
}

// @summary 展示文章内容
// @tags    文章
// @produce html
// @param   id path int false "文章ID"
// @router  /article/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Detail(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:       "gf bbs - 文章详情",
		Keywords:    "gf bbs - article keywords",
		Description: "gf bbs - article description",
		Data: g.Map{"test":"test"},
	})
}

// @summary 展示发布文章页面
// @tags    文章
// @produce html
// @router  /article/publish [GET]
// @success 200 {string} html "页面HTML"
func (a *articleApi) Publish(r *ghttp.Request) {
	service.View.Render(r, model.View{
		Title:       "gf bbs - Publish",
		Keywords:    "gf bbs - Publish keywords",
		Description: "gf bbs - Publish description",
	})
}

func (a *articleApi) DoPublish(r *ghttp.Request) {

}

func (a *articleApi) Update(r *ghttp.Request) {

}

func (a *articleApi) Delete(r *ghttp.Request) {

}
