package service

import (
	"context"
	"focus/app/dao"
	"focus/app/model"
	"focus/app/shared"
	"focus/app/system/index/internal/define"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"time"
)

// 文件管理服务
var File = fileService{}

type fileService struct{}

// 统一上传文件
func (s *fileService) Upload(ctx context.Context, input define.FileUploadInput) (*define.FileUploadOutput, error) {
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		return nil, gerror.New("上传文件路径配置不存在")
	}
	if input.Name != "" {
		input.File.Filename = input.Name
	}
	// 同一用户1分钟之内只能上传10张图片
	count, err := dao.File.Ctx(ctx).
		Where(dao.File.C.UserId, shared.Context.Get(ctx).User.Id).
		WhereGTE(dao.File.C.CreatedAt, gtime.Now().Add(time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= model.FileMaxUploadCountMinute {
		return nil, gerror.New("您上传得太频繁，请稍后再操作")
	}
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := input.File.Save(gfile.Join(uploadPath, dateDirName), input.RandomName)
	if err != nil {
		return nil, err
	}
	// 记录到数据表
	data := model.File{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dateDirName, fileName),
		Url:    "/upload/" + dateDirName + "/" + fileName,
		UserId: shared.Context.Get(ctx).User.Id,
	}
	result, err := dao.File.Ctx(ctx).Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &define.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Path: data.Src,
		Url:  data.Url,
	}, nil
}
