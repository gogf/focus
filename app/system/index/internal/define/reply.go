package define

import (
	"focus/app/model"
)

// API用户注册
type ReplyApiCreateUpdateBase struct {
	Title      string
	ParentId   uint   `v:"required#请输入账号"`    // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string `v:"required#评论内容类型错误"` // 评论类型: topic, ask, article, reply
	TargetId   uint   `v:"required#评论目标ID错误"` // 对应内容ID
	Content    string `v:"required#评论内容不能为空"` // 回复内容
}

// API执行删除内容
type ReplyApiDoDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}

// Service创建内容
type ReplyServiceCreateReq struct {
	ReplyApiCreateUpdateBase
	UserId uint
}

// Service查询列表结果
type ReplyServiceGetListReq struct {
	Page       int    `json:"page"`        // 分页码
	Size       int    `json:"size"`        // 分页数量
	TargetType string `json:"target_type"` // 数据类型
	TargetId   int    `json:"target_id"`   // 数据ID
	UserId     uint   `json:"user_id"`     // 用户ID
}

// Service查询列表结果
type ReplyServiceGetListRes struct {
	List  []*ReplyServiceGetListResItem `json:"list"`  // 列表
	Page  int                           `json:"page"`  // 分页码
	Size  int                           `json:"size"`  // 分页数量
	Total int                           `json:"total"` // 数据总数
}

// Service查询列表结果项
type ReplyServiceGetListResItem struct {
	Reply    *model.ReplyListItem           `json:"reply"`
	User     *model.ReplyListUserItem       `json:"user"`
	Content  *model.ContentListItem         `json:"content"`
	Category *model.ContentListCategoryItem `json:"category"`
}
