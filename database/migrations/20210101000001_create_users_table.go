package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20210101000001CreateUsersTable struct{}

// Signature The unique signature for the migration.
func (r *M20210101000001CreateUsersTable) Signature() string {
	return "20210101000001_create_users_table"
}

// Up Run the migrations.
func (r *M20210101000001CreateUsersTable) Up() error {
	return facades.Schema().Create("users", func(table schema.Blueprint) {
		table.ID("id")
		table.String("username")
		table.String("sex")
		table.String("password")
		table.String("phone")
		table.String("openid")
		table.String("unionid")
		table.String("realname")
		table.String("id_card_number")
		table.String("avatar")
		table.Text("remark")
		table.DateTime("last_login").UseCurrent().Nullable()
		table.Enum("status", []any{"正常", "暂停", "关闭"}).Default("正常")
		table.TimestampsTz()
	})
}

// Down Reverse the migrations.
func (r *M20210101000001CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
