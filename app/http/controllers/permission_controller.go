package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httpfacade "github.com/hulutech-web/http_result"
)

type PermissionController struct {
	//Dependent services
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         PermissionController
// @Accept       json
// @Produce      json
// @Id PermissionIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/permission [get]
func (r *PermissionController) Index(ctx http.Context) http.Response {
	permissions := []models.Permission{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&permissions)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         PermissionController
// @Accept       json
// @Produce      json
// @Id PermissionList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/admin/permission/list [get]
func (r *PermissionController) List(ctx http.Context) http.Response {
	permissions := []models.Permission{}
	queries := ctx.Request().Queries()
	_, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).List(&permissions)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return httpfacade.NewResult(ctx).Success("", permissions)
}
func (r *PermissionController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	permission := models.Permission{}
	facades.Orm().Query().Model(&models.Permission{}).Where("id = ?", id).First(&permission)
	return httpfacade.NewResult(ctx).Success("", permission)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         PermissionController
// @Accept       json
// @Produce      json
// @Id PermissionStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param permissionData body requests.PermissionRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/permission [post]
func (r *PermissionController) Store(ctx http.Context) http.Response {
	var permissionRequest requests.PermissionRequest
	errors, err := ctx.Request().ValidateRequest(&permissionRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	permission := models.Permission{
		Code:        permissionRequest.Code,
		Name:        permissionRequest.Name,
		Description: permissionRequest.Description,
		Type:        permissionRequest.Type,
		MenuID:      permissionRequest.MenuID,
	}
	//todo add request values
	facades.Orm().Query().Model(&models.Permission{}).Create(&permission)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         PermissionController
// @Accept       json
// @Produce      json
// @Id PermissionUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param permissionData body requests.PermissionRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/permission/{id} [put]
func (r *PermissionController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	permission := models.Permission{}
	facades.Orm().Query().Model(&models.Permission{}).Where("id=?", id).Find(&permission)
	var permissionRequest requests.PermissionRequest
	errors, err := ctx.Request().ValidateRequest(&permissionRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	permission = models.Permission{
		Code:        permissionRequest.Code,
		Name:        permissionRequest.Name,
		Description: permissionRequest.Description,
		Type:        permissionRequest.Type,
		MenuID:      permissionRequest.MenuID,
	}
	//todo add request values
	facades.Orm().Query().Model(&models.Permission{}).Where("id=?", id).Update(&permission)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         PermissionController
// @Accept       json
// @Produce      json
// @Id PermissionDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/permission/{id} [delete]
func (r *PermissionController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.Permission{}).Delete(&models.Permission{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         PermissionController
// @Accept       json
// @Produce      json
// @Id PermissionOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/permission/option [get]
func (r *PermissionController) Option(ctx http.Context) http.Response {
	type Opt struct {
		Value uint   `json:"value"`
		Label string `json:"label"`
		Code  string `json:"code"`
	}
	//todo change fields name follow model struct
	//eg name value
	var opts []Opt
	facades.Orm().Query().Model(&models.Permission{}).Select("id as value,name as label,code as code").Scan(&opts)
	return httpfacade.NewResult(ctx).Success("", opts)
}
