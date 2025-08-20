package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812181002CreateRolePermissionsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812181002CreateRolePermissionsTable) Signature() string {
	return "20250812181002_create_role_permissions_table"
}

// Up Run the migrations.
func (r *M20250812181002CreateRolePermissionsTable) Up() error {
	if !facades.Schema().HasTable("role_permissions") {
		return facades.Schema().Create("role_permissions", func(table schema.Blueprint) {
			table.UnsignedBigInteger("role_id")
			table.UnsignedBigInteger("permission_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812181002CreateRolePermissionsTable) Down() error {
	return facades.Schema().DropIfExists("role_permissions")
}
