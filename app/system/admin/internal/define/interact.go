package define

// API赞
type InteractApiZanReq struct {
	Id   uint   `v:"min:1#请选择需要赞的内容"`
	Type string `v:"required#请提交需要赞的内容类型"` // content, reply
}

// API取消赞
type InteractApiCancelZanReq struct {
	Id   uint   `v:"min:1#请选择需要取消赞的内容"`
	Type string `v:"required#请提交需要取消赞的内容类型"` // content, reply
}

// API踩
type InteractApiCaiReq struct {
	Id   uint   `v:"min:1#请选择需要踩的内容"`
	Type string `v:"required#请提交需要踩的内容类型"` // content, reply
}

// API取消踩
type InteractApiCancelCaiReq struct {
	Id   uint   `v:"min:1#请选择需要取消踩的内容"`
	Type string `v:"required#请提交需要取消踩的内容类型"` // content, reply
}
