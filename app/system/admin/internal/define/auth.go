package define

// API创建/修改权限基类
type AuthApiCreateUpdateBase struct {
	AuthServiceCreateUpdateBase
	Name  string `v:"required#请输入权限名称"`
	Key   string `v:"required-without:Value#权限标识和权限路由不能同时为空"` // 权限标识
	Value string `v:"required-without:Key#权限标识和权限路由不能同时为空"`   // 权限路由
}

// API执行创建权限
type AuthApiCreateReq struct {
	AuthApiCreateUpdateBase
}

// API执行修改权限
type AuthApiUpdateReq struct {
	AuthApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的权限"` // 修改时ID不能为空
}

// API执行删除权限
type AuthApiDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的权限"` // 删除时ID不能为空
}

// Service创建权限
type AuthServiceCreateReq struct {
	AuthServiceCreateUpdateBase
	UserId uint
}

// Service修改权限
type AuthServiceUpdateReq struct {
	AuthServiceCreateUpdateBase
	Id uint
}

type AuthServiceCreateUpdateBase struct {
	ParentId uint   // 父级菜单
	Name     string // 权限名称
	Key      string // 权限键名(用于程序)
	Value    string // 权限键值，部分自定义权限可能有键值存在
	Sort     int    // 排序
	Icon     string // 展示图标
}
