// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"focus/app/model/internal"
)

const (
	UserSessionKey     = "UserSessionKey" // 用户信息存放在Session中的Key
	UserDefaultRoleId  = 1                // 默认的用户角色ID
	UserStatusOk       = 0                // 用户状态正常
	UserStatusDisabled = 1                // 用户状态禁用
)

// User is the golang structure for table gf_user.
type User internal.User

// API用户注册
type UserApiRegisterReq struct {
	Passport  string `v:"required#请输入账号"`                           // 账号
	Password  string `v:"required#请输入密码"`                           // 密码(明文)
	Password2 string `v:"required|same:Password#请再次输入密码|两次密码输入不一致"` // 确认密码(明文)
	Nickname  string `v:"required#请输入昵称"`                           // 昵称
}

// API修改个人资料
type UserApiUpdateProfileReq struct {
	Nickname string `v:"required#请输入昵称信息"` // 昵称
	Avatar   string // 头像地址
	Gender   int    // 性别 0: 未设置 1: 男 2: 女
}

// API禁用用户
type UserApiDisableReq struct {
	Id *uint `v:"required#请选择需要禁用的用户"` // 删除时ID不能为空
}

// Api用户登录
type UserApiLoginReq struct {
	Passport string `v:"required#请输入账号"` // 账号
	Password string `v:"required#请输入密码"` // 密码(明文)
}

// Service用户登录
type UserServiceLoginReq struct {
	Passport string `v:"required#请输入账号"` // 账号
	Password string `v:"required#请输入密码"` // 密码(明文)
}

// Service创建用户
type UserServiceRegisterReq struct {
	RoleId   int    // 角色ID，允许负数：< 0 系统使用; > 0 业务使用. 一个用户只有一个角色
	Passport string // 账号
	Password string // 密码(明文)
	Nickname string // 昵称
}

// Service修改用户
type UserServiceUpdateProfileReq struct {
	Id       uint   // 用户ID
	Nickname string // 昵称
	Avatar   string // 头像地址
	Gender   int    // 性别 0: 未设置 1: 男 2: 女
}
