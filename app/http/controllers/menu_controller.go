package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httpfacade "github.com/hulutech-web/http_result"
)

type MenuController struct {
	//Dependent services
}

func NewMenuController() *MenuController {
	return &MenuController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         MenuController
// @Accept       json
// @Produce      json
// @Id MenuIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/admin/menu [get]
func (r *MenuController) Index(ctx http.Context) http.Response {
	menus := []models.Menu{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&menus)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         MenuController
// @Accept       json
// @Produce      json
// @Id MenuList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/admin/menu/list [get]
func (r *MenuController) List(ctx http.Context) http.Response {
	menus := []models.Menu{}
	queries := ctx.Request().Queries()
	return httpfacade.NewResult(ctx).SearchByParams(queries, nil).Success("", menus)
}
func (r *MenuController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	menu := models.Menu{}
	facades.Orm().Query().Model(&models.Menu{}).Where("id = ?", id).First(&menu)
	return httpfacade.NewResult(ctx).Success("", menu)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         MenuController
// @Accept       json
// @Produce      json
// @Id MenuStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param menuData body requests.MenuRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/menu [post]
func (r *MenuController) Store(ctx http.Context) http.Response {
	var menuRequest requests.MenuRequest
	errors, err := ctx.Request().ValidateRequest(&menuRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	menu := models.Menu{}
	//todo add request values
	facades.Orm().Query().Model(&models.Menu{}).Create(&menu)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         MenuController
// @Accept       json
// @Produce      json
// @Id MenuUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param menuData body requests.MenuRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/menu/{id} [put]
func (r *MenuController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	menu := models.Menu{}
	facades.Orm().Query().Model(&models.Menu{}).Where("id=?", id).Find(&menu)
	var menuRequest requests.MenuRequest
	errors, err := ctx.Request().ValidateRequest(&menuRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.Menu{}).Where("id=?", id).Save(&menu)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         MenuController
// @Accept       json
// @Produce      json
// @Id MenuDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/menu/{id} [delete]
func (r *MenuController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.Menu{}).Delete(&models.Menu{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         MenuController
// @Accept       json
// @Produce      json
// @Id MenuOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/menu/option [get]
func (r *MenuController) Option(ctx http.Context) http.Response {
	type Opt struct {
		Value uint   `json:"value"`
		Label string `json:"label"`
	}
	//todo change fields name follow model struct
	//eg name value
	var opts []Opt
	facades.Orm().Query().Model(&models.Menu{}).Select("id as value,title as label").Scan(&opts)
	return httpfacade.NewResult(ctx).Success("", opts)
}
