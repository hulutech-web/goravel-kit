
package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httpfacade "github.com/hulutech-web/http_result"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type HouseController struct {
	//Dependent services
}

func NewHouseController() *HouseController {
	return &HouseController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         HouseController
// @Accept       json
// @Produce      json
// @Id HouseIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/house [get]
func (r *HouseController) Index(ctx http.Context) http.Response {
	houses := []models.House{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&houses)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         HouseController
// @Accept       json
// @Produce      json
// @Id HouseList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/house/list [get]
func (r *HouseController) List(ctx http.Context) http.Response {
	houses := []models.House{}
	queries := ctx.Request().Queries()
	return httpfacade.NewResult(ctx).SearchByParams(queries, nil).Success("", houses)
}
func (r *HouseController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	house := models.House{}
	facades.Orm().Query().Model(&models.House{}).Where("id = ?", id).First(&house)
	return httpfacade.NewResult(ctx).Success("", house)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         HouseController
// @Accept       json
// @Produce      json
// @Id HouseStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param houseData body requests.HouseRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/house [post]
func (r *HouseController) Store(ctx http.Context) http.Response {
	var houseRequest requests.HouseRequest
	errors, err := ctx.Request().ValidateRequest(&houseRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	house := models.House{}
	//todo add request values
	facades.Orm().Query().Model(&models.House{}).Create(&house)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         HouseController
// @Accept       json
// @Produce      json
// @Id HouseUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param houseData body requests.HouseRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/house/{id} [put]
func (r *HouseController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	house := models.House{}
	facades.Orm().Query().Model(&models.House{}).Where("id=?", id).Find(&house)
	var houseRequest requests.HouseRequest
	errors, err := ctx.Request().ValidateRequest(&houseRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.House{}).Where("id=?", id).Save(&house)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         HouseController
// @Accept       json
// @Produce      json
// @Id HouseDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/house/{id} [delete]
func (r *HouseController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.House{}).Delete(&models.House{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         HouseController
// @Accept       json
// @Produce      json
// @Id HouseOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/house/option [get]
func (r *HouseController) Option(ctx http.Context) http.Response {
	type Opt struct {
		Value uint   `json:"value"`
		Label string `json:"label"`
	}
	//todo change fields name follow model struct 
	//eg name value
	var opts []Opt
	facades.Orm().Query().Model(&models.House{}).Select("id as value,title as label").Scan(&opts)
	return httpfacade.NewResult(ctx).Success("", opts)
}
