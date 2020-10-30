package boot

import (
	_ "focus/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	// 绑定Swagger Plugin
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
