package models

import "github.com/goravel/framework/database/orm"

type Role struct {
	orm.Model
	Name        string       `gorm:"column:name;not null;default:'';comment:'角色名称''" form:"name" json:"name"`
	Label       string       `gorm:"column:label;type:varchar(255);" json:"label" form:"label"`
	Remark      string       `gorm:"column:remark;default:'';comment:'备注信息'" form:"remark" json:"remark"`
	IsDisable   int          `gorm:"column:is_disable;default:0;comment:'是否禁用: 0=否, 1=是'" form:"is_disable" json:"is_disable"`
	Sort        int          `gorm:"column:sort;default:0;comment:'角色排序'" form:"sort" json:"sort"`
	TenantID    uint         `gorm:"column:tenant_id;default:0;comment:'租户ID'" form:"tenant_id" json:"tenant_id"`
	IsAdmin     bool         `gorm:"column:is_admin;default:0;comment:'是否管理员'" form:"is_admin" json:"is_admin"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions" form:"permissions"`
}
