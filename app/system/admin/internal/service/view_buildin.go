package service

import (
	"fmt"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

// 视图自定义方法管理对象
type viewBuildIn struct {
	httpRequest *ghttp.Request
}

// 获取管理后台菜单列表，最多两级菜单
func (s *viewBuildIn) Menus() []*model.MenuItem {
	var (
		adminMenus    = make([]*model.MenuItem, 0)
		menuJsonArray = g.Cfg("admin").GetJsons("menus")
		value         string
		items         []string
		array         []string
	)
	for _, v := range menuJsonArray {
		value = v.GetString("value")
		items = v.GetStrings("items")
		array = gstr.SplitAndTrim(value, ",")
		menuItem := &model.MenuItem{
			Name:   array[0],
			Url:    array[1],
			Icon:   array[2],
			Active: s.isMenuUrlActive(array[1]),
			Items:  make([]*model.MenuItem, 0),
		}
		if len(array) > 3 {
			menuItem.Target = array[3]
		}
		for _, item := range items {
			array = gstr.SplitAndTrim(item, ",")
			item := &model.MenuItem{
				Name:   array[0],
				Url:    array[1],
				Icon:   array[2],
				Active: s.isMenuUrlActive(array[1]),
			}
			if len(array) > 3 {
				item.Target = array[3]
			}
			if item.Active {
				menuItem.Active = true
			}
			menuItem.Items = append(menuItem.Items, item)
		}
		adminMenus = append(adminMenus, menuItem)
	}
	return adminMenus
}

// 判断给定的管理后台URL是否被选中
func (s *viewBuildIn) isMenuUrlActive(url string) bool {
	// 处理是否选中, URL，包含QueryString
	if gstr.Equal(url, s.httpRequest.URL.String()) {
		return true
	}

	// 处理是否选中, URI
	if gstr.Equal(gstr.Split(url, "?")[0], s.httpRequest.URL.Path) {
		return true
	}

	return false
}

// 获得指定的栏目树形对象，当contentType为空时，表示获取所有的栏目树形对象。
func (s *viewBuildIn) CategoryTree(contentType string) ([]*model.CategoryTree, error) {
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
