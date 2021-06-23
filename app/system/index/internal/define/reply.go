package define

import (
	"focus/app/model"
)

type ReplyDoCreateReq struct {
	ReplyCreateInput
	ParentId   uint   `v:"required#请输入账号"`    // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string `v:"required#评论内容类型错误"` // 评论类型: topic, ask, article, reply
	TargetId   uint   `v:"required#评论目标ID错误"` // 对应内容ID
	Content    string `v:"required#评论内容不能为空"` // 回复内容
}

// 执行删除内容
type ReplyDoDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}

// 创建内容
type ReplyCreateInput struct {
	Title      string
	ParentId   uint   // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string // 评论类型: topic, ask, article, reply
	TargetId   uint   // 对应内容ID
	Content    string // 回复内容
	UserId     uint
}

// 查询回复列表请求
type ReplyGetListReq struct {
	ReplyGetListInput
}

// 查询回复列表
type ReplyGetListInput struct {
	Page       int    // 分页码
	Size       int    // 分页数量
	TargetType string // 数据类型
	TargetId   int    // 数据ID
	UserId     uint   // 用户ID
}

// 查询列表结果
type ReplyGetListOutput struct {
	List  []ReplyGetListOutputItem `json:"list"`  // 列表
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

// 查询列表结果项
type ReplyGetListOutputItem struct {
	Reply    *model.ReplyListItem           `json:"reply"`
	User     *model.ReplyListUserItem       `json:"user"`
	Content  *model.ContentListItem         `json:"content"`
	Category *model.ContentListCategoryItem `json:"category"`
}
