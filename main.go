package main

import (
	_ "focus/boot"
	_ "focus/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
