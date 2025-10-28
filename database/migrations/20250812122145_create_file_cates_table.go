package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812122145CreateFileCatesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812122145CreateFileCatesTable) Signature() string {
	return "20250812122145_create_file_cates_table"
}

// Up Run the migrations.
func (r *M20250812122145CreateFileCatesTable) Up() error {
	if !facades.Schema().HasTable("file_cates") {
		return facades.Schema().Create("file_cates", func(table schema.Blueprint) {
			table.ID()
			table.String("name").Comment("类目名称")
			table.Integer("sort").Comment("排序")
			table.String("type").Comment("类目类型: [image=图片, video=视频, audio=音频, file=文件]")
			table.BigInteger("pid").Comment("父类目ID")
			table.BigInteger("tenant_id").Comment("租户ID")
			table.TimestampsTz()
		})
	}
	return nil
}

// Down Reverse the migrations.
func (r *M20250812122145CreateFileCatesTable) Down() error {
	return facades.Schema().DropIfExists("file_cates")
}
