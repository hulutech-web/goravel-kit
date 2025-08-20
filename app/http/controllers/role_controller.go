package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httpfacade "github.com/hulutech-web/http_result"
)

type RoleController struct {
	//Dependent services
}

func NewRoleController() *RoleController {
	return &RoleController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/admin/role [get]
func (r *RoleController) Index(ctx http.Context) http.Response {
	roles := []models.Role{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&roles)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/admin/role/list [get]
func (r *RoleController) List(ctx http.Context) http.Response {
	roles := []models.Role{}
	queries := ctx.Request().Queries()
	return httpfacade.NewResult(ctx).SearchByParams(queries, nil).Success("", roles)
}
func (r *RoleController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	role := models.Role{}
	facades.Orm().Query().Model(&models.Role{}).Where("id = ?", id).First(&role)
	return httpfacade.NewResult(ctx).Success("", role)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param roleData body requests.RoleRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/role [post]
func (r *RoleController) Store(ctx http.Context) http.Response {
	var roleRequest requests.RoleRequest
	errors, err := ctx.Request().ValidateRequest(&roleRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	role := models.Role{}
	//todo add request values
	facades.Orm().Query().Model(&models.Role{}).Create(&role)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param roleData body requests.RoleRequest true "用户数据"
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/role/{id} [put]
func (r *RoleController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	role := models.Role{}
	facades.Orm().Query().Model(&models.Role{}).Where("id=?", id).Find(&role)
	var roleRequest requests.RoleRequest
	errors, err := ctx.Request().ValidateRequest(&roleRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.Role{}).Where("id=?", id).Save(&role)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/role/{id} [delete]
func (r *RoleController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.Role{}).Delete(&models.Role{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/role/option [get]
func (r *RoleController) Option(ctx http.Context) http.Response {
	roles := []models.Role{}
	queries := ctx.Request().Queries()
	res, _ := httpfacade.NewResult(ctx).SearchByParams(queries, nil).List(&roles)
	return res
}

// Permissions 获取角色的权限
// @Summary      选项
// @Description  选项
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RolePermissions
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/role/{id}/permissions [get]
func (r *RoleController) Permissions(ctx http.Context) http.Response {
	pers := []models.Permission{}
	id := ctx.Request().Route("id")
	role := models.Role{}
	facades.Orm().Query().Model(&models.Role{}).Where("id=?", id).First(&role)
	facades.Orm().Query().Model(&role).
		Association("Permissions").Find(&pers)
	return httpfacade.NewResult(ctx).Success("", pers)
}

// SyncPermissions
// @Summary      同步权限
// @Description  同步权限
// @Tags         RoleController
// @Accept       json
// @Produce      json
// @Id RoleSyncPermissions
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param roleData body requests.SyncRolePerRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/role/{id}/sync_permissions [post]
func (r *RoleController) SyncPermissions(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var syncRolePerRequest requests.SyncRolePerRequest
	query := facades.Orm().Query()
	var role models.Role

	// 查询角色
	if err := query.Model(&models.Role{}).With("Permissions").Where("id = ?", id).First(&role); err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusNotFound, "角色未找到", "")
	}
	existPers := role.Permissions

	// 验证请求参数
	errors, err := ctx.Request().ValidateRequest(&syncRolePerRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	permissions := []models.Permission{}
	query.Model(&models.Permission{}).Where("id in ?", syncRolePerRequest.FormIDs).Find(&permissions)

	err = query.Model(&role).Association("Permissions").Delete(&existPers)

	err = query.Model(&role).Association("Permissions").Append(&permissions)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "权限同步错误", err.Error())
	}
	return httpfacade.NewResult(ctx).Success("权限同步成功", nil)
}
