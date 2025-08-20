package models

import (
	"goravel/app/models/common"
)

// House 房源模型
type House struct {
	ID         uint   `gorm:"primaryKey;column:id;comment:房源ID" json:"id" form:"id"`
	LandlordID uint   `gorm:"not null;column:landlord_id;comment:房东ID" json:"landlord_id" form:"landlord_id"`
	Type       string `gorm:"type:enum('whole','room');default:'whole';column:type;comment:房源类型" json:"type" form:"type"`

	// 公共信息
	Title       string          `gorm:"size:200;not null;column:title;comment:房源标题" json:"title" form:"title"`
	Address     string          `gorm:"size:255;not null;column:address;comment:详细地址" json:"address" form:"address"`
	Location    common.CoordRes `gorm:"type:geometry;column:location;comment:经纬度" json:"location" form:"location"`
	Traffic     string          `gorm:"type:text;column:traffic;comment:交通信息" json:"traffic" form:"traffic"`
	Shopping    string          `gorm:"type:text;column:shopping;comment:购物信息" json:"shopping" form:"shopping"`
	Video       string          `gorm:"type:text;column:video;comment:视频URL" json:"video" form:"video"`
	Swipers     common.Swipers  `gorm:"type:text;column:swipers;comment:轮播图JSON" json:"swipers" form:"swipers"`
	Features    common.Items    `gorm:"type:text;column:features;comment:特色标签" json:"features" form:"features"`
	Facilities  common.Items    `gorm:"type:text;column:facilities;comment:设施信息" json:"facilities" form:"facilities"`
	PropertyFee float64         `gorm:"column:property_fee;comment:物业费" json:"property_fee" form:"property_fee"`

	// 整租特有字段
	Description string       `gorm:"type:text;column:description;comment:房源描述" json:"description" form:"description"`
	MonthlyRent float64      `gorm:"column:monthly_rent;comment:月租金" json:"monthly_rent" form:"monthly_rent"`
	Deposit     float64      `gorm:"column:deposit;comment:押金" json:"deposit" form:"deposit"`
	CoverImg    string       `gorm:"column:cover_img;comment:封面图" json:"cover_img" form:"cover_img"`
	Poster      string       `gorm:"column:poster;comment:海报图" json:"poster" form:"poster"`
	Albums      common.Items `gorm:"type:text;column:albums;comment:相册" json:"albums" form:"albums"`
	Area        float64      `gorm:"column:area;comment:面积" json:"area" form:"area"`
	HouseType   string       `gorm:"type:enum('一居','二居','三居','四居','五居','六居','其他');column:house_type;comment:户型" json:"house_type" form:"house_type"`
	Orientation string       `gorm:"type:enum('东','南','西','北','南北','东西');column:orientation;comment:朝向" json:"orientation" form:"orientation"`

	// 状态
	Status string `gorm:"type:enum('draft','published','rented','maintenance');default:'draft';column:status;comment:状态" json:"status" form:"status"`

	// 关系
	Landlord User   `gorm:"foreignKey:LandlordID" json:"landlord" form:"-"`
	Rooms    []Room `gorm:"foreignKey:HouseID" json:"rooms" form:"-"`
}
