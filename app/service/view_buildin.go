package service

import (
	"focus/app/model"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// 视图自定义方法管理对象
type ViewBuildIn struct {
	httpRequest *ghttp.Request
}

// 创建分页HTML内容
func (s *ViewBuildIn) Page(total, size int) string {
	page := s.httpRequest.GetPage(total, size)
	page.LinkStyle = "page-link"
	page.SpanStyle = "page-link active"
	content := page.GetContent(4)
	content = gstr.ReplaceByMap(content, map[string]string{
		"<span":     "<li class=\"page-item\"><span",
		"/span>":    "/span></li>",
		"<a":        "<li class=\"page-item\"><a",
		"/a>":       "/a></li>",
		"GPageSpan": "GPageSpan page-link",
	})
	return content
}

// 获得指定的栏目树形对象，当contentType为空时，表示获取所有的栏目树形对象。
func (s *ViewBuildIn) CategoryTree(contentType string) ([]*model.CategoryTree, error) {
	return Category.GetTree(s.httpRequest.Context(), contentType)
}
