package service

import (
	"fmt"
	"focus/app/model"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gmode"
)

// 视图自定义方法管理对象
type ViewBuildIn struct {
	httpRequest *ghttp.Request
}

// 根据性别字段内容返回性别的font。
func (s *ViewBuildIn) GenderFont(gender int) string {
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
func (s *ViewBuildIn) Page(total, size int) string {
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

// 我是否赞了这个内容
func (s *ViewBuildIn) DidIZan(targetType string, targetId uint) bool {
	b, _ := Interact.DidIZan(s.httpRequest.Context(), targetType, targetId)
	return b
}

// 我是否踩了这个内容
func (s *ViewBuildIn) DidICai(targetType string, targetId uint) bool {
	b, _ := Interact.DidICai(s.httpRequest.Context(), targetType, targetId)
	return b
}

// 获取顶部菜单列表
func (s *ViewBuildIn) TopMenus() ([]*model.TopMenuItem, error) {
	topMenus, err := Menu.GetTopMenus()
	if err != nil {
		return nil, err
	}
	if len(topMenus) == 0 {
		return nil, nil
	}
	currentUriWithQueryString := s.httpRequest.URL.String()
	// 处理是否选中, URL，包含QueryString
	for _, v := range topMenus {
		if gstr.Equal(v.Url, currentUriWithQueryString) {
			v.Active = true
			return topMenus, nil
		}
	}
	// 处理是否选中, URI
	for _, v := range topMenus {
		if v.Url == "/" {
			continue
		}
		if gstr.HasPrefix(currentUriWithQueryString, v.Url) {
			v.Active = true
			return topMenus, nil
		}
	}
	// 没有选中的菜单，那么自动识别第一层路由，例如：
	// /topic/1 则选中 /topic 菜单。
	array := gstr.SplitAndTrim(s.httpRequest.URL.Path, "/")
	if len(array) > 1 {
		path := "/" + array[0]
		for _, v := range topMenus {
			if gstr.Equal(v.Url, path) {
				v.Active = true
				return topMenus, nil
			}
		}
	}
	// 最后则自动高亮首页(第一个菜单)
	topMenus[0].Active = true
	return topMenus, nil
}

// 获取当前页面的Url Path.
func (s *ViewBuildIn) UrlPath() string {
	return s.httpRequest.URL.Path
}

// 获得指定的栏目树形对象，当contentType为空时，表示获取所有的栏目树形对象。
func (s *ViewBuildIn) CategoryTree(contentType string) ([]*model.CategoryTree, error) {
	return Category.GetTree(s.httpRequest.Context(), contentType)
}

// 随机数 开发环境时间戳，线上为前端版本号
func (s *ViewBuildIn) Random() string {
	var rand string
	if gmode.IsDevelop() {
		rand = gconv.String(gtime.TimestampMilli())
	} else {
		rand = "1.0.0"
	}
	return rand
}

// FormatTime 格式化时间
func (s *ViewBuildIn) FormatTime(gt *gtime.Time) string {
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
