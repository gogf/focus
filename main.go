package main

import (
	_ "focus/boot"
	_ "focus/router"

	"github.com/gogf/gf/frame/g"
)

// @title       GoFrame社区API
// @version     1.0
// @description GoFrame社区API
// @schemes     http https
func main() {
	g.Server().Run()
}
