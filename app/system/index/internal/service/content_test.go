package service

import (
	"context"
	"focus/app/system/index/internal/define"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestContentService_GetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Content.GetList(context.TODO(), &define.ContentServiceGetListReq{
			Page: 1,
			Size: 4,
		})
		t.Assert(err, nil)
		g.Dump(list)
	})
}

func TestContentService_Search(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		list, err := Content.Search(context.TODO(), &define.ContentServiceSearchReq{
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
