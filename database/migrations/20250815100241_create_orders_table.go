package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250815100241CreateOrdersTable struct{}

// Signature The unique signature for the migration.
func (r *M20250815100241CreateOrdersTable) Signature() string {
	return "20250815100241_create_orders_table"
}

// Up Run the migrations.
func (r *M20250815100241CreateOrdersTable) Up() error {
	if !facades.Schema().HasTable("orders") {
		return facades.Schema().Create("orders", func(table schema.Blueprint) {
			table.ID()
			table.Enum("target_type", []any{"house", "room"}).Comment("目标类型")
			table.UnsignedBigInteger("house_id").Default(0)
			table.UnsignedBigInteger("room_id").Default(0)
			table.UnsignedBigInteger("tenant_id")
			table.Date("start_date")
			table.Date("end_date")
			table.Float("monthly_rent").Default(0)
			table.Enum("status", []any{"待支付", "已支付", "执行中", "已结束"}).Default("待支付")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250815100241CreateOrdersTable) Down() error {
	return facades.Schema().DropIfExists("orders")
}
