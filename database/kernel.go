package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"goravel/database/migrations"
	"goravel/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateUsersTable{},
		&migrations.M20210101000002CreateJobsTable{},
		&migrations.M20250812104711CreateHousesTable{},
		&migrations.M20250812122145CreateFileCatesTable{},
		&migrations.M20250812122233CreateFilesTable{},
		&migrations.M20250812122424CreateMenusTable{},
		&migrations.M20250812180742CreateRolesTable{},
		&migrations.M20250812180749CreatePermissionsTable{},
		&migrations.M20250812181002CreateRolePermissionsTable{},
		&migrations.M20250812181028CreateUserRolesTable{},
		&migrations.M20250815095805CreateRoomsTable{},
		&migrations.M20250815100241CreateOrdersTable{},
		&migrations.M20250815101623CreateContractsTable{},
		&migrations.M20250815101745CreateContractTemplatesTable{},
		&migrations.M20250816105218CreatePdfGensTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
		&seeders.UserSeeder{},
		&seeders.TplSeeder{},
		&seeders.MenuSeeder{},
		&seeders.PermissionSeeder{},
	}
}
