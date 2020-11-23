package define

// Service获取面包屑请求
type ViewServiceGetBreadCrumbReq struct {
	ContentId   uint   // (可选)内容ID
	ContentType string // (可选)内容类型
	CategoryId  uint   // (可选)栏目ID
}

// Service获取title请求
type ViewServiceGetTitleReq struct {
	ContentType string // (可选)内容类型
	CategoryId  uint   // (可选)栏目ID
	CurrentName string // (可选)当前名称
}
