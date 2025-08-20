package models

import "github.com/goravel/framework/database/orm"

type FileCate struct {
	orm.Model
	Name     string  `gorm:"column:name;not null;default:'';comment:'类目名称'" form:"name" json:"name"`
	Sort     int     `gorm:"column:sort;not null;default:0;comment:'排序'" form:"sort" json:"sort"`
	Type     string  `gorm:"column:type;not null;default:'';comment:'类目类型: [image=图片, video=视频, audio=音频, file=文件]';index" form:"type" json:"type"`
	PID      uint    `gorm:"column:pid;not null;default:0;comment:'父类目ID'" form:"pid" json:"pid"`
	TenantID uint    `gorm:"column:tenant_id;not null;default:0;comment:'租户ID'" form:"tenant_id" json:"tenant_id"`
	Files    []*File `gorm:"foreignKey:CID;references:ID" json:"files"`
}
