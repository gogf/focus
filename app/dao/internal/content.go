// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ContentDao is the manager for logic model data accessing and custom defined data operations functions management.
type ContentDao struct {
	gmvc.M                // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      contentColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB         // DB is the raw underlying database management object.
	Table  string         // Table is the underlying table name of the DAO.
}

// ContentColumns defines and stores column names for table gf_content.
type contentColumns struct {
	Id             string // 自增ID
	Key            string // 唯一键名，用于程序硬编码，一般不常用
	Type           string // 内容模型: topic, ask, article等，具体由程序定义
	CategoryId     string // 栏目ID
	UserId         string // 用户ID
	AdoptedReplyId string // 采纳的回复ID，问答模块有效
	Title          string // 标题
	Content        string // 内容
	Sort           string // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief          string // 摘要
	Thumb          string // 缩略图
	Tags           string // 标签名称列表，以JSON存储
	Referer        string // 内容来源，例如github/gitee
	Status         string // 状态 0: 正常, 1: 禁用
	ReplyCount     string // 回复数量
	ViewCount      string // 浏览数量
	ZanCount       string // 赞
	CaiCount       string // 踩
	CreatedAt      string // 创建时间
	UpdatedAt      string // 修改时间
}

// NewContentDao creates and returns a new DAO object for table data access.
func NewContentDao() *ContentDao {
	columns := contentColumns{
		Id:             "id",
		Key:            "key",
		Type:           "type",
		CategoryId:     "category_id",
		UserId:         "user_id",
		AdoptedReplyId: "adopted_reply_id",
		Title:          "title",
		Content:        "content",
		Sort:           "sort",
		Brief:          "brief",
		Thumb:          "thumb",
		Tags:           "tags",
		Referer:        "referer",
		Status:         "status",
		ReplyCount:     "reply_count",
		ViewCount:      "view_count",
		ZanCount:       "zan_count",
		CaiCount:       "cai_count",
		CreatedAt:      "created_at",
		UpdatedAt:      "updated_at",
	}
	return &ContentDao{
		C:     columns,
		M:     g.DB("default").Model("gf_content").Safe(),
		DB:    g.DB("default"),
		Table: "gf_content",
	}
}
