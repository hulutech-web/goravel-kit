package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250812122233CreateFilesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250812122233CreateFilesTable) Signature() string {
	return "20250812122233_create_files_table"
}

// Up Run the migrations.
func (r *M20250812122233CreateFilesTable) Up() error {
	if !facades.Schema().HasTable("files") {
		return facades.Schema().Create("files", func(table schema.Blueprint) {
			table.ID()
			table.Integer("cid").Comment("类目ID")
			table.BigInteger("user_id").Comment("用户ID")
			table.String("type").Comment("文件类型: [image=图片, video=视频, audio=音频, file=文件]")
			table.String("name").Comment("文件名称")
			table.String("uri").Comment("文件路径")
			table.String("ext").Comment("文件扩展")
			table.Integer("size").Comment("文件大小")
			table.String("engine").Comment("存储引擎")
			table.String("path").Comment("访问路径")
			table.Integer("tenant_id").Comment("租户ID")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250812122233CreateFilesTable) Down() error {
	return facades.Schema().DropIfExists("files")
}
