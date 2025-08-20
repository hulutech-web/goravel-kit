package routes

import (
	"goravel/app/http/controllers"
	"goravel/app/http/middleware"

	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Api() {

	//登录接口
	authCtrl := controllers.NewAuthController()
	facades.Route().Post("api/admin/auth/login", authCtrl.Login)

	//后台管理
	facades.Route().Middleware(middleware.Jwt()).Prefix("/api/admin").Group(func(router route.Router) {
		//系统菜单
		router.Get("menu/route", authCtrl.Menu)

		userCtrl := controllers.NewUserController()
		router.Get("user/own", userCtrl.Own)
		router.Get("user/option", userCtrl.Option)
		router.Get("user", userCtrl.Index)

		roleCtrl := controllers.NewRoleController()
		router.Resource("role", roleCtrl)
		router.Get("role/list", roleCtrl.List)
		router.Get("role/option", roleCtrl.Option)
		router.Get("role/{id}/permissions", roleCtrl.Permissions)
		router.Post("role/{id}/sync_permissions", roleCtrl.SyncPermissions)

		permissionCtrl := controllers.NewPermissionController()
		router.Resource("permission", permissionCtrl)
		router.Get("permission/list", permissionCtrl.List)
		router.Get("permission/option", permissionCtrl.Option)

		menuCtrl := controllers.NewMenuController()
		router.Resource("menu", menuCtrl)
		router.Get("menu/list", menuCtrl.List)
		router.Get("menu/option", menuCtrl.Option)

		fileCtrl := controllers.NewFileController()
		router.Resource("file", fileCtrl)
		router.Get("file/list", fileCtrl.List)
		router.Get("file/option", fileCtrl.Option)
		router.Post("file_del_batch", fileCtrl.DelBatch)

		file_cateCtrl := controllers.NewFileCateController()
		router.Resource("file_cate", file_cateCtrl)
		router.Get("file_cate/list", file_cateCtrl.List)
		router.Get("file_cate/option", file_cateCtrl.Option)
		router.Post("file_cate/{id}/files", file_cateCtrl.Files)

	})

}
