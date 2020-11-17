package service

import (
	"context"
	"focus/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestContentService_GetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Content.GetList(context.TODO(), &model.ContentServiceGetListReq{
			Page: 1,
			Size: 4,
		})
		t.Assert(err, nil)
		g.Dump(list)
	})
}

func TestContentService_Search(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Content.Search(context.TODO(), &model.ContentServiceSearchReq{
			Key:  "goframe",
			Page: 1,
			Size: 4,
		})
		t.Assert(err, nil)
		g.Dump(list)
	})
}

func TestTopicService_GetDetail(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Content.GetDetail(context.TODO(), 1)
		t.Assert(err, nil)
		g.Dump(list)
	})
}
