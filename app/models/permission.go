package models

import (
	"github.com/goravel/framework/database/orm"
)

type Permission struct {
	orm.Model
	Name        string  `gorm:"column:name;not null;default:'';comment:'权限标识'" form:"name" json:"name"`
	Code        string  `gorm:"column:code;not null;default:'';comment:'权限编码'" form:"code" json:"code"`
	Type        int     `gorm:"column:type;not null;default:1;comment:'权限类型1:菜单2-按钮3-API'" form:"type" json:"type"`
	Description string  `gorm:"column:description;not null;default:'';comment:'权限描述'" form:"description" json:"description"`
	Roles       *[]Role `gorm:"many2many:role_permissions;" form:"roles" json:"roles"`
	MenuID      uint    `gorm:"column:menu_id;index" form:"menu_id" json:"menu_id"`
	Menu        *Menu   `gorm:"foreignKey:MenuID"`
}
