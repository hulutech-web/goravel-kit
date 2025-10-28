package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812181028CreateUserRolesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812181028CreateUserRolesTable) Signature() string {
	return "20250812181028_create_user_roles_table"
}

// Up Run the migrations.
func (r *M20250812181028CreateUserRolesTable) Up() error {
	if !facades.Schema().HasTable("user_roles") {
		return facades.Schema().Create("user_roles", func(table schema.Blueprint) {
			table.UnsignedBigInteger("user_id")
			table.UnsignedBigInteger("role_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812181028CreateUserRolesTable) Down() error {
	return facades.Schema().DropIfExists("user_roles")
}
