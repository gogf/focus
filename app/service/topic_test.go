package service

import (
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestTopicService_GetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Topic.GetList(&model.TopicServiceGetListReq{
			Page: 1,
			Size: 4,
			Sort: "",
		})
		t.Assert(err, nil)
		g.Dump(list)
	})
}
