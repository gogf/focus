package define

import (
	"focus/app/model"
)

// ==========================================================================================
// API
// ==========================================================================================
// API查看内容详情
type ContentDetailReq struct {
	Id uint `v:"min:1#请选择查看的内容"`
}

// API展示创建内容页面
type ContentCreateReq struct {
	Type string `v:"required#请选择需要创建的内容类型"`
}

// API展示修改内容页面
type ContentUpdateReq struct {
	Id uint `v:"min:1#请选择需要修改的内容"`
}

// API创建/修改内容基类
type ContentCreateUpdateBase struct {
	ContentServiceCreateUpdateBase
	CategoryId uint   `v:"min:1#请输入栏目ID"`    // 栏目ID
	Title      string `v:"required#请输入内容标题"` // 标题
	Content    string `v:"required#请输入内容内容"` // 内容
}

// API执行创建内容
type ContentDoCreateReq struct {
	ContentCreateInput
}

// API执行修改内容
type ContentDoUpdateReq struct {
	ContentUpdateInput
	Id uint `v:"min:1#请选择需要修改的内容"` // 修改时ID不能为空
}

// API执行删除内容
type ContentDoDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}

// API执行采纳回复
type ContentAdoptReplyReq struct {
	Id      uint `v:"min:1#请选择需要采纳回复的内容"` // 采纳回复时ID不能为空
	ReplyId uint `v:"min:1#请选择需要采纳的回复"`   // 采纳回复时回复ID不能为空
}

type ContentGetListReq struct {
	ContentGetListInput
	CategoryId uint `p:"cate"`                    // 栏目ID
	Page       int  `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size       int  `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
}

type ContentGetListInput struct {
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	UserId     uint   // 要查询的用户ID
}

// Service查询列表结果
type ContentGetListOutput struct {
	List  []*ContentGetListOutputItem `json:"list"`  // 列表
	Page  int                         `json:"page"`  // 分页码
	Size  int                         `json:"size"`  // 分页数量
	Total int                         `json:"total"` // 数据总数
}

// Service搜索列表
type ContentSearchReq struct {
	ContentSearchInput
	CategoryId uint `p:"cate"`                    // 栏目ID
	Page       int  `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size       int  `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
}

// Service搜索列表
type ContentSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// Service搜索列表结果
type ContentSearchOutput struct {
	List  []*ContentSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

type ContentGetListOutputItem struct {
	Content  *model.ContentListItem         `json:"content"`
	Category *model.ContentListCategoryItem `json:"category"`
	User     *model.ContentListUserItem     `json:"user"`
}

type ContentSearchOutputItem struct {
	ContentGetListOutputItem
}

// Service查询详情结果
type ContentGetDetailOutput struct {
	Content *model.Content `json:"content"`
	User    *model.User    `json:"user"`
}

// Service创建/修改内容基类
type ContentServiceCreateUpdateBase struct {
	Type       string   // 内容模型
	CategoryId uint     // 栏目ID
	Title      string   // 标题
	Content    string   // 内容
	Brief      string   // 摘要
	Thumb      string   // 缩略图
	Tags       []string // 标签名称列表，以JSON存储
	Referer    string   // 内容来源，例如github/gitee
}

// Service创建内容
type ContentCreateInput struct {
	ContentServiceCreateUpdateBase
	UserId uint
}

// Service创建内容返回结果
type ContentCreateOutput struct {
	ContentId uint `json:"content_id"`
}

// Service修改内容
type ContentUpdateInput struct {
	ContentServiceCreateUpdateBase
	Id uint
}
