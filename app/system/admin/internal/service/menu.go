package service

import (
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// 菜单管理服务
var Menu = new(menuService)

type menuService struct{}

// 获取侧边菜单
func (s *menuService) GetMenus(r *ghttp.Request) []*model.MenuItem {
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
			Active: s.isMenuUrlActive(r, array[1]),
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
				Active: s.isMenuUrlActive(r, array[1]),
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
func (s *menuService) isMenuUrlActive(r *ghttp.Request, url string) bool {
	// 处理是否选中, URL，包含QueryString
	if gstr.Equal(url, r.URL.String()) {
		return true
	}

	// 处理是否选中, URI
	if gstr.Equal(gstr.Split(url, "?")[0], r.URL.Path) {
		return true
	}

	return false
}

// 根据URL自动识别当前选中的侧边菜单，并返回级联标题，用于html title展示
func (s *menuService) GetCurrentTitle(r *ghttp.Request) string {
	return s.getTitleFromMenus(s.GetMenus(r))
}

// 递归根据选中的菜单获取级联title
func (s *menuService) getTitleFromMenus(menus []*model.MenuItem) string {
	title := ""
	for _, menu := range menus {
		if menu.Active {
			if len(menu.Items) > 0 {
				if str := s.getTitleFromMenus(menu.Items); str != "" {
					if title == "" {
						title = str
					} else {
						title = title + " - " + str
					}
				}
			}
			if title == "" {
				title = menu.Name
			} else {
				title = title + " - " + menu.Name
			}
		}
	}
	return title
}
