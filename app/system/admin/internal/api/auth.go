package api

import (
	"focus/app/model"
	"focus/app/system/admin/internal/define"
	"focus/app/system/admin/internal/service"
	"focus/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 权限管理
var Auth = new(authApi)

type authApi struct{}

// @summary 展示权限列表页面
// @tags    后台-权限
// @produce html
// @router  /admin/auth [GET]
// @success 200 {string} html "页面HTML"
func (a *authApi) Index(r *ghttp.Request) {
	authTree, err := service.Auth.GetTree(r.Context())
	if err != nil {
		service.View.Render500(r, model.View{
			Error: err.Error(),
		})
	}
	service.View.Render(r, model.View{
		Data: g.Map{
			"authTree": authTree,
		},
	})
}

// @summary 创建权限
// @tags    后台-权限
// @produce json
// @router  /admin/auth/create [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *authApi) Create(r *ghttp.Request) {
	var (
		apiReq     *define.AuthApiCreateReq
		serviceReq *define.AuthServiceCreateReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Auth.Create(r.Context(), serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary 修改权限
// @tags    后台-权限
// @produce json
// @router  /admin/auth/update [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *authApi) Update(r *ghttp.Request) {
	var (
		apiReq     *define.AuthApiUpdateReq
		serviceReq *define.AuthServiceUpdateReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Auth.Update(r.Context(), serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}

// @summary 删除权限
// @tags    后台-权限
// @produce json
// @router  /admin/auth/delete [POST]
// @success 200 {object} response.JsonRes "请求结果"
func (a *authApi) Delete(r *ghttp.Request) {
	var (
		apiReq *define.AuthApiDeleteReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Auth.Delete(r.Context(), apiReq.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "")
	}
}
