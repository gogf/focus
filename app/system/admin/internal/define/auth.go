package define

// Service创建权限
type AuthServiceCreateReq struct {
	AuthServiceCreateUpdateBase
}

// Service修改权限
type AuthServiceUpdateReq struct {
	AuthServiceCreateUpdateBase
	Id uint
}

type AuthServiceCreateUpdateBase struct {
	ParentId uint   // 父级菜单
	UserId   uint   // 创建用户ID
	Name     string // 权限名称
	Key      string // 权限键名(用于程序)
	Value    string // 权限键值，部分自定义权限可能有键值存在
	Sort     int    // 排序
	Icon     string // 展示图标
}
