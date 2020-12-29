package service

import (
	"fmt"
	"focus/app/model"
	"focus/app/system/admin/internal/define"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

// 视图自定义方法管理对象
type viewBuildIn struct {
	httpRequest *ghttp.Request // 务必使用小写，否则dump方法会出问题
}

// 获取管理后台菜单列表，最多两级菜单
func (s *viewBuildIn) SideMenus() []*define.SideMenuItem {
	return Menu.GetMenus(s.httpRequest)
}

// 获得指定的栏目树形对象，当contentType为空时，表示获取所有的栏目树形对象。
func (s *viewBuildIn) CategoryTree(contentType string) ([]*model.CategoryTreeItem, error) {
	return Category.GetTree(s.httpRequest.Context(), contentType)
}

// 根据性别字段内容返回性别的font。
func (s *viewBuildIn) GenderFont(gender int) string {
	switch gender {
	case model.UserGenderMale:
		return "&#xe651;"
	case model.UserGenderFemale:
		return "&#xe636;"
	default:
		return "&#xead2;"
	}
}

// 根据性别字段内容返回性别。
func (s *viewBuildIn) Gender(gender int) string {
	switch gender {
	case model.UserGenderMale:
		return "男"
	case model.UserGenderFemale:
		return "女"
	default:
		return "未知"
	}
}

// 创建分页HTML内容
func (s *viewBuildIn) Page(total, size int) string {
	page := s.httpRequest.GetPage(total, size)
	page.LinkStyle = "page-link"
	page.SpanStyle = "page-link"
	page.PrevPageTag = "«"
	page.NextPageTag = "»"
	content := page.PrevPage() + page.PageBar() + page.NextPage()
	content = gstr.ReplaceByMap(content, map[string]string{
		"<span":  "<li class=\"page-item disabled\"><span",
		"/span>": "/span></li>",
		"<a":     "<li class=\"page-item\"><a",
		"/a>":    "/a></li>",
	})
	return content
}

// 获取当前页面的Url Path.
func (s *viewBuildIn) UrlPath() string {
	return s.httpRequest.URL.Path
}

// FormatTime 格式化时间
func (s *viewBuildIn) FormatTime(gt *gtime.Time) string {
	if gt == nil {
		return ""
	}
	n := gtime.Now().Timestamp()
	t := gt.Timestamp()

	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1

	var rs string

	d := n - t
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d秒前", int(d/ss))
	default:
		rs = "刚刚"
	}

	return rs
}
