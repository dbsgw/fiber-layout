package models

import (
	"fiber-layout/initalize"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Uid       string `gorm:"not null;comment:用户uid，唯一id" json:"uid"`
	Username  string `gorm:"comment:用户名" json:"username"`
	Nickname  string `gorm:"comment:昵称" json:"nickname"`
	Password  string `gorm:"comment:密码" json:"password"`
	Telephone string `gorm:"comment:手机号" json:"telephone"`
	Email     string `gorm:"comment:邮箱" json:"email"`
	HeadImg   string `gorm:"comment:头像" json:"headImg"`
	Sex       string `gorm:"comment:性别，0:男 1:女" json:"sex"`
	Age       string `gorm:"comment:年龄" json:"age"`
	Birthday  string `gorm:"comment:生日" json:"birthday"`
	Status    int    `gorm:"default:1;comment:0:禁用 1：正常" json:"status"`
}

func NewSysUser() *SysUser {
	return &SysUser{}
}

func (t *SysUser) GetList() ([]SysUser, error) {
	var sys []SysUser
	if err := initalize.DB.Raw("select * from sys_user").Find(&sys).Error; err != nil {
		return nil, err
	}
	return sys, nil
}
