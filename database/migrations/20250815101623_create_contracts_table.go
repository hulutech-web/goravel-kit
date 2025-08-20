package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250815101623CreateContractsTable struct{}

// Signature The unique signature for the migration.
func (r *M20250815101623CreateContractsTable) Signature() string {
	return "20250815101623_create_contracts_table"
}

// Up Run the migrations.
func (r *M20250815101623CreateContractsTable) Up() error {
	if !facades.Schema().HasTable("contracts") {
		return facades.Schema().Create("contracts", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("order_id").Default(0)
			table.UnsignedBigInteger("landlord_id").Default(0)
			table.UnsignedBigInteger("tenant_id").Default(0)
			table.Enum("type", []any{"代理合同", "租赁合同"}).Default("代理合同")
			table.Text("content")
			table.String("tenant_sign")
			table.String("landlord_sign")
			table.DateTime("signed_time").UseCurrent()
			table.String("signed_location").Comment("签约地点")
			table.String("paper_contract").Nullable().Comment("纸质合同")

			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250815101623CreateContractsTable) Down() error {
	return facades.Schema().DropIfExists("contracts")
}
