package category

type ListRequest struct {
	ContentType string `v:"required#请输入分类类型"` // 分类类型
	ParentId    int    // 父级ID
}

// 查询详情请求
type ItemRequest struct {
	Id uint `json:"id" v:"min:1#请输入分类ID"` // 分类ID
}
