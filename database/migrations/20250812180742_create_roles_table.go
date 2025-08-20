package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812180742CreateRolesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812180742CreateRolesTable) Signature() string {
	return "20250812180742_create_roles_table"
}

// Up Run the migrations.
func (r *M20250812180742CreateRolesTable) Up() error {
	if !facades.Schema().HasTable("roles") {
		return facades.Schema().Create("roles", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.String("label")
			table.String("remark").Nullable()
			table.Integer("is_disable").Default(0)
			table.Integer("sort").Default(0)
			table.UnsignedBigInteger("tenant_id").Default(0)
			table.Integer("is_admin").Default(0)
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812180742CreateRolesTable) Down() error {
	return facades.Schema().DropIfExists("roles")
}
