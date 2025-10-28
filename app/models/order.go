package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
)

type Order struct {
	orm.Model
	TargetType string          `gorm:"type:enum('house','room');column:target_type;comment:目标类型" json:"target_type" form:"target_type"`
	HouseID    *uint           `gorm:"column:house_id;comment:房源ID" json:"house_id" form:"house_id"`
	RoomID     *uint           `gorm:"column:room_id;comment:房间ID" json:"room_id" form:"room_id"`
	TenantID   uint            `gorm:"not null;column:tenant_id;comment:租客ID" json:"tenant_id" form:"tenant_id"`
	StartDate  carbon.DateTime `gorm:"type:date;column:start_date;comment:租赁开始日期" json:"start_date" form:"start_date"`
	EndDate    carbon.DateTime `gorm:"type:date;column:end_date;comment:租赁结束日期" json:"end_date" form:"end_date"`
	Amount     float64         `gorm:"column:amount;comment:订单金额" json:"amount" form:"amount"`
	Status     string          `gorm:"default:'待支付';column:status;comment:状态(1待支付 2已签约 3执行中 4已结束)" json:"status" form:"status"`

	// 关系
	House  House `gorm:"foreignKey:HouseID" json:"house" form:"-"`
	Room   Room  `gorm:"foreignKey:RoomID" json:"room" form:"-"`
	Tenant User  `gorm:"foreignKey:TenantID" json:"tenant" form:"-"`
}
