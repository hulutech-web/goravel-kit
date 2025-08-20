package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250815095805CreateRoomsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250815095805CreateRoomsTable) Signature() string {
	return "20250815095805_create_rooms_table"
}

// Up Run the migrations.
func (r *M20250815095805CreateRoomsTable) Up() error {
	if !facades.Schema().HasTable("rooms") {
		return facades.Schema().Create("rooms", func(table schema.Blueprint) {
			table.ID()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250815095805CreateRoomsTable) Down() error {
 	return facades.Schema().DropIfExists("rooms")
}
