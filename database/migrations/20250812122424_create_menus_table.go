package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812122424CreateMenusTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812122424CreateMenusTable) Signature() string {
	return "20250812122424_create_menus_table"
}

// Up Run the migrations.
func (r *M20250812122424CreateMenusTable) Up() error {
	if !facades.Schema().HasTable("menus") {
		return facades.Schema().Create("menus", func(table schema.Blueprint) {
			table.ID()
			table.Integer("pid").Comment("父级ID").Default(0)
			table.String("title").Comment("标题")
			table.String("name").Comment("名称")
			table.String("path").Comment("路径").Nullable()
			table.String("component").Comment("组件").Nullable()
			table.String("icon").Default("AlertOutlined").Comment("图标")
			table.String("menu_type").Comment("菜单类型: [page=页面, action=操作]")
			table.Boolean("cacheable").Comment("是否缓存").Default(true)
			table.Boolean("render_menu").Comment("是否渲染菜单").Default(false)
			table.String("permission").Comment("权限标识").Nullable()
			table.Integer("sort").Comment("排序").Nullable()
			table.String("target").Comment("目标").Nullable()
			table.String("badge").Comment("角标").Nullable()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812122424CreateMenusTable) Down() error {
	return facades.Schema().DropIfExists("menus")
}
