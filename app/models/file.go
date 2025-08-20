package models

import "github.com/goravel/framework/database/orm"

type File struct {
	orm.Model
	CID      uint   `gorm:"column:cid;not null;default:0;comment:'类目ID';index" form:"cid" json:"cid"`
	UserID   uint   `gorm:"column:user_id;not null;default:0;comment:'用户ID'" form:"user_id" json:"user_id"`
	Type     string `gorm:"not null;default:'image';comment:'文件类型: [image=图片, video=视频, audio=音频, file=文件]';index" form:"type" json:"type"`
	Name     string `gorm:"column:name;not null;default:'';comment:'文件名称''" form:"name" json:"name"`
	Uri      string `gorm:"column:uri;not null;comment:'文件路径'" form:"uri" json:"uri"`
	Ext      string `gorm:"column:ext;not null;default:'';comment:'文件扩展'" form:"ext" json:"ext"`
	Size     int64  `gorm:"column:size;not null;default:0;comment:文件大小" form:"size" json:"size"`
	Engine   string `gorm:"column:engine;not null;default:'';comment:'存储引擎'" form:"engine" json:"engine"`
	Path     string `gorm:"column:path;not null;default:'';comment:'访问路径'" form:"path" json:"path"`
	TenantID uint   `gorm:"column:tenant_id;not null;default:0;comment:'租户ID'" form:"tenant_id" json:"tenant_id"`
}
