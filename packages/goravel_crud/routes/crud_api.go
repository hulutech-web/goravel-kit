package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers")

func CrudApi() {
	jwt_middleware_cbk := facades.Config().Get("crud.middleware").(func() http.Middleware)
	prefix := facades.Config().Get("crud.prefix").(string)
	facades.Route().Middleware(jwt_middleware_cbk()).Prefix(prefix).Group(func(router route.Router) {


roleCtrl := controllers.NewRoleController()
router.Resource("role", roleCtrl)
router.Get("role/list", roleCtrl.List)
router.Get("role/option", roleCtrl.Option)


permissionCtrl := controllers.NewPermissionController()
router.Resource("permission", permissionCtrl)
router.Get("permission/list", permissionCtrl.List)
router.Get("permission/option", permissionCtrl.Option)


menuCtrl := controllers.NewMenuController()
router.Resource("menu", menuCtrl)
router.Get("menu/list", menuCtrl.List)
router.Get("menu/option", menuCtrl.Option)


houseCtrl := controllers.NewHouseController()
router.Resource("house", houseCtrl)
router.Get("house/list", houseCtrl.List)
router.Get("house/option", houseCtrl.Option)


fileCtrl := controllers.NewFileController()
router.Resource("file", fileCtrl)
router.Get("file/list", fileCtrl.List)
router.Get("file/option", fileCtrl.Option)


fileCateCtrl := controllers.NewFileCateController()
router.Resource("fileCate", fileCateCtrl)
router.Get("fileCate/list", fileCateCtrl.List)
router.Get("fileCate/option", fileCateCtrl.Option)


contractCtrl := controllers.NewContractController()
router.Resource("contract", contractCtrl)
router.Get("contract/list", contractCtrl.List)
router.Get("contract/option", contractCtrl.Option)

	})
}
