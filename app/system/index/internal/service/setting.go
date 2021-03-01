package service

import (
	"focus/app/dao"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
)

// 持久化Key-Value管理服务
var Setting = settingService{}

type settingService struct{}

// 设置KV。
func (s *settingService) Set(key, value string) error {
	_, err := dao.Setting.Data(model.Setting{
		K: key,
		V: value,
	}).Save()
	return err
}

// 查询KV。
func (s *settingService) Get(key string) (string, error) {
	v, err := s.GetVar(key)
	return v.String(), err
}

// 查询KV，返回泛型，便于转换。
func (s *settingService) GetVar(key string) (*g.Var, error) {
	v, err := dao.Setting.Fields(dao.Setting.Columns.V).Where(dao.Setting.Columns.K, key).Value()
	return v, err
}
