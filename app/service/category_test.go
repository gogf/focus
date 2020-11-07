package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestCategoryService_GetTree(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tree, err := Category.GetTree(context.TODO(), "")
		t.Assert(err, nil)
		g.Dump(tree)
	})
}

func TestCategoryService_GetSubIdList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		array, err := Category.GetSubIdList(context.TODO(), 5)
		t.Assert(err, nil)
		g.Dump(array)
	})
}
