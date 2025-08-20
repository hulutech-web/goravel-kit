package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250815101745CreateContractTemplatesTable struct{}

// Signature The unique signature for the migration.
func (r *M20250815101745CreateContractTemplatesTable) Signature() string {
	return "20250815101745_create_contract_templates_table"
}

// Up Run the migrations.
func (r *M20250815101745CreateContractTemplatesTable) Up() error {
	if !facades.Schema().HasTable("contract_templates") {
		return facades.Schema().Create("contract_templates", func(table schema.Blueprint) {
			table.ID()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250815101745CreateContractTemplatesTable) Down() error {
 	return facades.Schema().DropIfExists("contract_templates")
}
