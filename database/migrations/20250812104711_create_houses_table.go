package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812104711CreateHousesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812104711CreateHousesTable) Signature() string {
	return "20250812104711_create_houses_table"
}

// Up Run the migrations.
func (r *M20250812104711CreateHousesTable) Up() error {
	if !facades.Schema().HasTable("houses") {
		return facades.Schema().Create("houses", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("landlord_id")
			table.String("title")
			table.String("description")
			table.String("address")
			table.Float("monthly_rent")
			table.Float("deposit")
			table.String("header_img")
			table.String("poster")
			table.Text("albums")
			table.Column("location", "geometry").Nullable().Comment("经纬度")
			table.Float("area")
			table.Text("facilities")
			table.Float("property_fee")
			table.Text("traffic")
			table.Text("shopping")
			table.Text("video")
			table.String("status")
			table.Text("swipers")
			table.Foreign("landlord_id").References("id").On("users")
			table.Index("landlord_id")

			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812104711CreateHousesTable) Down() error {
	return facades.Schema().DropIfExists("houses")
}
