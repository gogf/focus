package main

import (
	"fmt"
	"focus/app/model"
	"focus/app/service"
	"github.com/gogf/gf/frame/g"
)

func main() {
	defaultTopMenus := []*model.TopMenuItem{
		{
			Name: "首页",
			Url:  "/",
		},
		{
			Name: "主题",
			Url:  "/topic",
		},
		{
			Name: "问答",
			Url:  "/ask",
		},
		{
			Name: "文章",
			Url:  "/article",
		},
		{
			Name:   "教程",
			Url:    "#",
			Target: "",
			Items: []*model.TopMenuItem{
				{
					Name:   "GoFrame源码",
					Url:    "https://github.com/gogf/gf",
					Target: "_blank",
				},
				{
					Name:   "GoFrame官网",
					Url:    "https://goframe.org",
					Target: "_blank",
				},
				{
					Name:   "GoFrame教程",
					Url:    "https://github.com/gogf/gf",
					Target: "_blank",
				},
			},
		},
	}
	err := service.Menu.SetTopMenus(defaultTopMenus)
	fmt.Println(err)
	g.Dump(service.Menu.GetTopMenus())
}
