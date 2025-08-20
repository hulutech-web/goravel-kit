package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type MenuSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *MenuSeeder) Signature() string {
	return "MenuSeeder"
}

// Run executes the seeder logic.
func (s *MenuSeeder) Run() error {
	menus := []models.Menu{
		{
			PID:        0,
			Title:      "首页",
			Name:       "workplace",
			Path:       "/workplace",
			Component:  "@/pages/workplace/index.vue",
			Icon:       "HomeOutlined",
			MenuType:   "",
			Cacheable:  false,
			RenderMenu: true,
			Permission: "workplace",
			Sort:       0,
			Target:     "",
			Badge:      "",
		},
		{
			PID:        0,
			Title:      "系统",
			Name:       "system",
			Path:       "/system",
			Component:  "@/components/layout/BlankView.vue",
			Icon:       "ControlOutlined",
			MenuType:   "",
			Cacheable:  false,
			RenderMenu: true,
			Permission: "system",
			Sort:       100,
			Target:     "",
			Badge:      "",
		},
		{
			PID:        2,
			Title:      "菜单管理",
			Name:       "system.auth.menu",
			Path:       "/system/auth/menu",
			Component:  "@/pages/system/auth/menu/index.vue",
			Icon:       "SettingOutlined",
			MenuType:   "",
			Cacheable:  false,
			RenderMenu: true,
			Permission: "system:auth:menu",
			Sort:       100,
			Target:     "",
			Badge:      "",
		},
		{
			PID:        2,
			Title:      "角色管理",
			Name:       "system.auth.role",
			Path:       "/system/auth/role",
			Component:  "@/pages/system/auth/role/index.vue",
			Icon:       "UserSwitchOutlined",
			MenuType:   "",
			Cacheable:  false,
			RenderMenu: true,
			Permission: "system:auth:role",
			Sort:       100,
			Target:     "",
			Badge:      "",
		},
		{
			PID:        2,
			Title:      "权限管理",
			Name:       "permission",
			Path:       "/system/auth/permission",
			Component:  "@/pages/system/auth/permission/index.vue",
			Icon:       "VerifiedOutlined",
			MenuType:   "",
			Cacheable:  false,
			RenderMenu: true,
			Permission: "system:auth:permission",
			Sort:       100,
			Target:     "",
			Badge:      "",
		},
		{
			PID:        2,
			Title:      "用户管理",
			Name:       "user",
			Path:       "/system/user",
			Component:  "@/pages/system/user/index.vue",
			Icon:       "UserOutlined",
			MenuType:   "",
			Cacheable:  false,
			RenderMenu: true,
			Permission: "system:user",
			Sort:       1,
			Target:     "",
			Badge:      "",
		},
		{
			PID:        2,
			Title:      "附件中心",
			Name:       "netdisk",
			Path:       "/system/netdisk",
			Component:  "@/pages/system/netdisk/index.vue",
			Icon:       "FolderOutlined",
			MenuType:   "page",
			Cacheable:  true,
			RenderMenu: true,
			Permission: "system:netdisk",
			Sort:       0,
			Target:     "_self",
			Badge:      "",
		},
	}
	return facades.Orm().Query().Create(&menus)
}
