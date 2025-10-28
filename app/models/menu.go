package models

import (
	"github.com/goravel/framework/database/orm"
	"sort"
)

type Menu struct {
	orm.Model
	PID        uint   `gorm:"column:pid" form:"pid" json:"pid"`
	Title      string `gorm:"column:title" form:"title" json:"title"`
	Name       string `gorm:"column:name" form:"name" json:"name"`
	Path       string `gorm:"column:path" form:"path" json:"path"`
	Component  string `gorm:"column:component" form:"component" json:"component"`
	Icon       string `gorm:"column:icon;" form:"icon" json:"icon"`
	MenuType   string `json:"menu_type" gorm:"column:menu_type;index" form:"menu_type"` // page, action
	Cacheable  bool   `json:"cacheable" gorm:"column:cacheable" form:"cacheable"`
	RenderMenu bool   `json:"renderMenu" gorm:"column:render_menu" form:"renderMenu"` //遵循step-template的规范
	Permission string `gorm:"column:permission;index" form:"permission" json:"permission"`
	Sort       int    `json:"sort" gorm:"column:sort" form:"sort"`
	Target     string `json:"target" gorm:"column:target" form:"target"`
	Badge      string `json:"badge" gorm:"column:badge" form:"badge"`
	//设置外键 子级菜单
	Children []Menu `gorm:"foreignKey:PID;references:ID" json:"children"`
}

// BuildMenuTree 构建菜单树形结构
// menus 是所有的菜单列表
// parentID 是当前层级的父ID，顶层菜单的parentID通常为0
func (m *Menu) BuildMenuTree(menus []Menu, parentID uint) []Menu {
	var tree []Menu

	for _, menu := range menus {
		if menu.PID == parentID {
			// 递归查找子菜单
			children := m.BuildMenuTree(menus, menu.ID)
			if len(children) > 0 {
				menu.Children = children
			}
			tree = append(tree, menu)
		}
	}

	// 按Sort字段排序
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].Sort < tree[j].Sort
	})

	return tree
}
