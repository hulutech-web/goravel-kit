package models

import (
	"github.com/goravel/framework/database/orm"
	"goravel/app/models/common"
)

// Room 房间模型
type Room struct {
	orm.Model
	HouseID    uint `gorm:"not null;column:house_id;comment:所属房源ID" json:"house_id" form:"house_id"`
	LandlordID uint `gorm:"not null;column:landlord_id;comment:房东ID" json:"landlord_id" form:"landlord_id"`

	// 房间信息
	Title        string       `gorm:"size:200;not null;column:title;comment:房间标题" json:"title" form:"title"`
	Description  string       `gorm:"type:text;column:description;comment:房间描述" json:"description" form:"description"`
	MonthlyRent  float64      `gorm:"column:monthly_rent;comment:月租金" json:"monthly_rent" form:"monthly_rent"`
	Deposit      float64      `gorm:"column:deposit;comment:押金" json:"deposit" form:"deposit"`
	CoverImg     string       `gorm:"column:cover_img;comment:封面图" json:"cover_img" form:"cover_img"`
	Albums       common.Items `gorm:"type:text;column:albums;comment:相册" json:"albums" form:"albums"`
	Area         float64      `gorm:"column:area;comment:房间面积" json:"area" form:"area"`
	Orientation  string       `gorm:"type:enum('东','南','西','北','南北','东西');column:orientation;comment:朝向" json:"orientation" form:"orientation"`
	RoomFeatures common.Items `gorm:"type:text;column:room_features;comment:房间特色" json:"room_features" form:"room_features"`

	// 状态
	Status string `gorm:"type:enum('draft','published','rented','maintenance');default:'draft';column:status;comment:状态" json:"status" form:"status"`

	// 关系
	House    House `gorm:"foreignKey:HouseID" json:"house" form:"-"`
	Landlord User  `gorm:"foreignKey:LandlordID" json:"landlord" form:"-"`
}
