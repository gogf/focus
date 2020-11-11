package service

import (
	"encoding/json"
	"focus/app/model"
)

// 菜单管理
var Menu = new(menuService)

type menuService struct{}

const (
	settingTopMenusKey = "TopMenus"
)

// 获取顶部菜单
func (s *menuService) SetTopMenus(menus []*model.TopMenuItem) error {
	b, err := json.Marshal(menus)
	if err != nil {
		return err
	}
	return Setting.Set(settingTopMenusKey, string(b))
}

// 获取顶部菜单
func (s *menuService) GetTopMenus() ([]*model.TopMenuItem, error) {
	var topMenus []*model.TopMenuItem
	v, err := Setting.GetVar(settingTopMenusKey)
	if err != nil {
		return nil, err
	}
	err = v.Structs(&topMenus)
	return topMenus, err
}
