package category

type Item struct {
	Id       uint   `json:"id"`        // 分类ID，自增主键
	ParentId uint   `json:"parent_id"` // 父级分类ID，用于层级管理
	Name     string `json:"name"`      // 分类名称
	Thumb    string `json:"thumb"`     // 封面图
	Brief    string `json:"brief"`     // 简述
}

type GetListParam struct {
	ParentId    uint   // 父级分类ID，用于层级管理
	ContentType string // 内容类型：topic, question, article
}
