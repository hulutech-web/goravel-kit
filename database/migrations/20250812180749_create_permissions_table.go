package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812180749CreatePermissionsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812180749CreatePermissionsTable) Signature() string {
	return "20250812180749_create_permissions_table"
}

// Up Run the migrations.
func (r *M20250812180749CreatePermissionsTable) Up() error {
	if !facades.Schema().HasTable("permissions") {
		return facades.Schema().Create("permissions", func(table schema.Blueprint) {
			table.ID("id")
			table.String("name").Comment("权限标识")
			table.String("code").Comment("权限标识")
			table.Integer("type").Comment("权限类型: 1-菜单，2-按钮，3-API")
			table.String("description").Comment("描述信息")
			table.Integer("menu_id").Comment("菜单ID")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812180749CreatePermissionsTable) Down() error {
	return facades.Schema().DropIfExists("permissions")
}
