package service

import (
	"context"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestTopicService_GetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Topic.GetList(context.TODO(), &model.TopicServiceGetListReq{
			Page: 1,
			Size: 4,
			Sort: "",
		})
		t.Assert(err, nil)
		g.Dump(list)
	})
}

func TestTopicService_GetDetail(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Topic.GetDetail(context.TODO(), 1)
		t.Assert(err, nil)
		g.Dump(list)
	})
}
