package boot

import (
	_ "focus/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func init() {
	g.SetDebug(true)

	// 绑定Swagger Plugin
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	// 静态目录设置
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		g.Log().Fatal("文件上传配置路径不能为空")
	}
	s.AddStaticPath("/upload", uploadPath)
}
