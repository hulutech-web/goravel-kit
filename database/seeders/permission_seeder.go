package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type PermissionSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *PermissionSeeder) Signature() string {
	return "PermissionSeeder"
}

// Run executes the seeder logic.
func (s *PermissionSeeder) Run() error {

	// 2. 创建角色
	roles := []models.Role{
		{Name: "admin", Label: "超级管理员", Remark: "超级管理员"},
		{Name: "instructor", Label: "教练", Remark: "教练"},
	}
	if err := facades.Orm().Query().Create(&roles); err != nil {
		return err
	}

	// 3. 为用户分配角色
	if _, err := facades.Orm().Query().Exec("INSERT INTO user_roles (user_id, role_id) VALUES (1, 1), (2, 2)"); err != nil {
		return err
	}

	//4. 初始化权限
	sys_pers := []models.Permission{
		{Name: "首页", Code: "workplace", Type: 1, MenuID: 1},
		{Name: "系统", Code: "system", Type: 1, MenuID: 2},
		{Name: "菜单管理", Code: "system:auth:menu", Type: 1, MenuID: 3},
		{Name: "角色管理", Code: "system:auth:role", Type: 1, MenuID: 4},
		{Name: "权限管理", Code: "system:auth:permission", Type: 1, MenuID: 5},
		{Name: "用户管理", Code: "system:user", Type: 1, MenuID: 6},
		{Name: "附件中心", Code: "system:netdisk", Type: 1, MenuID: 7},
		{Name: "CRUD生成器", Code: "system:crud:index", Type: 1, MenuID: 8},
		{Name: "字段设计", Code: "system:crud:column", Type: 1, MenuID: 9},
	}
	if err := facades.Orm().Query().Create(&sys_pers); err != nil {
		return err
	}

	//5. 为角色分配权限
	if _, err := facades.Orm().Query().Exec("INSERT INTO role_permissions (role_id, permission_id) VALUES (1, 1), (1, 2),(1, 3),(1, 4),(1, 5),(1, 6),(1, 7),(1,8),(1,9)"); err != nil {
		return err
	}
	return nil
}
