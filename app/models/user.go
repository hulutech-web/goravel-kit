package models

import (
	"fmt"
	"goravel/database/factories"
	"time"

	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"github.com/goravel/framework/support/json"
)

type User struct {
	orm.Model
	Username string `gorm:"column:username" form:"username" json:"username"`
	Sex      string `gorm:"sex" form:"sex" json:"sex"`

	Phone string `gorm:"phone" form:"phone" json:"phone"`

	Openid       string           `gorm:"column:openid" form:"openid" json:"openid"`
	Unionid      string           `gorm:"column:unionid" form:"unionid" json:"unionid"`
	Password     string           `gorm:"password" form:"password"  json:"password"`
	Avatar       string           `gorm:"avatar" form:"avatar" json:"avatar"`
	Realname     string           `gorm:"realname" form:"realname" json:"realname"`
	IDCardNumber string           `gorm:"id_card_number" form:"id_card_number" json:"id_card_number"`
	Remark       string           `gorm:"remark" form:"remark" json:"remark"`
	LastLogin    *carbon.DateTime `gorm:"column:last_login" form:"last_login" json:"last_login"`
	Status       string           `gorm:"column:status;default:null" form:"status" json:"status"`
	Roles        []Role           `gorm:"many2many:user_roles;" json:"roles" form:"roles"`
}

func (u *User) Factory() factory.Factory {
	return &factories.UserFactory{}
}
func (u *User) GetRoles() string {
	isExist := facades.Cache().Store("redis").Has(fmt.Sprintf("user_%d_roles", u.ID))
	if isExist {
		return facades.Cache().Store("redis").Get(fmt.Sprintf("user_%d_roles", u.ID)).(string)
	} else {
		facades.Cache().Store("redis").Remember(fmt.Sprintf("user_%d_roles", u.ID), 2*time.Hour, func() (any, error) {
			roles := []string{}
			facades.Orm().Query().Model(u).With("Roles").Find(&u)
			for _, role := range u.Roles {
				roles = append(roles, role.Name)
			}
			marshal, _ := json.Marshal(roles)
			return marshal, nil
		})
		return facades.Cache().Store("redis").Get(fmt.Sprintf("user_%d_roles", u.ID)).(string)
	}
}

// 获取权限
func (u *User) GetPermissions() string {
	//如果缓存没有就把数据从数据库查出来，再放到缓存。
	isExist := facades.Cache().Store("redis").Has(fmt.Sprintf("user_%d_permissions", u.ID))
	if isExist {
		return facades.Cache().Store("redis").Get(fmt.Sprintf("user_%d_permissions", u.ID)).(string)
	} else {
		facades.Cache().Store("redis").Remember(fmt.Sprintf("user_%d_permissions", u.ID), 2*time.Hour, func() (any, error) {
			permissions := []string{}
			uniquePermissions := make(map[string]struct{}) // 用于去重的 map
			facades.Orm().Query().Model(&u).With("Roles").Find(&u)
			for _, role := range u.Roles {
				facades.Orm().Query().With("Permissions").Find(&role)
				// 使用 map 检查权限是否存在，避免重复
				for _, permission := range role.Permissions {
					// 使用 map 检查权限是否存在，避免重复
					if _, exists := uniquePermissions[permission.Code]; !exists {
						uniquePermissions[permission.Code] = struct{}{}
						permissions = append(permissions, permission.Code)
					}
				}
			}
			marshal, err := json.Marshal(permissions)
			return marshal, err
		})
		return facades.Cache().Store("redis").Get(fmt.Sprintf("user_%d_permissions", u.ID)).(string)
	}
}

func (u User) GetMenus() string {
	isExist := facades.Cache().Has(fmt.Sprintf("user_%d_menus", u.ID))
	if isExist {
		return facades.Cache().Store("redis").Get(fmt.Sprintf("user_%d_menus", u.ID)).(string)
	} else {
		facades.Cache().Store("redis").Remember(fmt.Sprintf("user_%d_menus", u.ID), 5*time.Hour, func() (any, error) {
			menus := []Menu{}
			menu_ids := []uint{}
			if u.ID == 1 {
				facades.Orm().Query().Model(&Menu{}).Find(&menus)
				res := new(Menu).BuildMenuTree(menus, 0)
				marshal, _ := json.Marshal(res)
				return marshal, nil
			}
			for _, role := range u.Roles {
				facades.Orm().Query().With("Permissions").Find(&role)
				for _, permission := range role.Permissions {
					menu_ids = append(menu_ids, permission.MenuID)
				}
			}
			facades.Orm().Query().Model(&Menu{}).Where("id in ?", menu_ids).Find(&menus)
			marshal, _ := json.Marshal(menus)
			return marshal, nil
		})
		return facades.Cache().Store("redis").Get(fmt.Sprintf("user_%d_menus", u.ID)).(string)
	}
}
