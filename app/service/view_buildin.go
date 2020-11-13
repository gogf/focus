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

// 获取顶部菜单列表
func (s *ViewBuildIn) TopMenus() ([]*model.TopMenuItem, error) {
	return Menu.GetTopMenus()
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
