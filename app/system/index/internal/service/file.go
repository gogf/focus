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

// 同一上传文件
func (s *fileService) Upload(ctx context.Context, r *define.FileServiceUploadReq) (*define.FileServiceUploadRes, error) {
	uploadPath := g.Cfg().GetString("upload.path")
	if uploadPath == "" {
		return nil, gerror.New("上传文件路径配置不存在")
	}
	if r.Name != "" {
		r.File.Filename = r.Name
	}
	// 同一用户1分钟之内只能上传10张图片
	count, err := dao.File.
		Where(dao.File.Columns.UserId, shared.Context.Get(ctx).User.Id).
		Where(dao.File.Columns.CreatedAt+">=?", gtime.Now().Add(time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= model.FileMaxUploadCountMinute {
		return nil, gerror.New("您上传得太频繁，请稍后再操作")
	}
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := r.File.Save(gfile.Join(uploadPath, dateDirName), r.RandomName)
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
	result, err := dao.File.Data(data).OmitEmpty().Insert()
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &define.FileServiceUploadRes{
		Id:   uint(id),
		Name: data.Name,
		Path: data.Src,
		Url:  data.Url,
	}, nil
}
