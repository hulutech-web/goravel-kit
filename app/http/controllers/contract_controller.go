
package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httpfacade "github.com/hulutech-web/http_result"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type ContractController struct {
	//Dependent services
}

func NewContractController() *ContractController {
	return &ContractController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         ContractController
// @Accept       json
// @Produce      json
// @Id ContractIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/contract [get]
func (r *ContractController) Index(ctx http.Context) http.Response {
	contracts := []models.Contract{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&contracts)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         ContractController
// @Accept       json
// @Produce      json
// @Id ContractList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/contract/list [get]
func (r *ContractController) List(ctx http.Context) http.Response {
	contracts := []models.Contract{}
	queries := ctx.Request().Queries()
	return httpfacade.NewResult(ctx).SearchByParams(queries, nil).Success("", contracts)
}
func (r *ContractController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	contract := models.Contract{}
	facades.Orm().Query().Model(&models.Contract{}).Where("id = ?", id).First(&contract)
	return httpfacade.NewResult(ctx).Success("", contract)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         ContractController
// @Accept       json
// @Produce      json
// @Id ContractStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param contractData body requests.ContractRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/contract [post]
func (r *ContractController) Store(ctx http.Context) http.Response {
	var contractRequest requests.ContractRequest
	errors, err := ctx.Request().ValidateRequest(&contractRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	contract := models.Contract{}
	//todo add request values
	facades.Orm().Query().Model(&models.Contract{}).Create(&contract)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         ContractController
// @Accept       json
// @Produce      json
// @Id ContractUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param contractData body requests.ContractRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/contract/{id} [put]
func (r *ContractController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	contract := models.Contract{}
	facades.Orm().Query().Model(&models.Contract{}).Where("id=?", id).Find(&contract)
	var contractRequest requests.ContractRequest
	errors, err := ctx.Request().ValidateRequest(&contractRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.Contract{}).Where("id=?", id).Save(&contract)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         ContractController
// @Accept       json
// @Produce      json
// @Id ContractDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/contract/{id} [delete]
func (r *ContractController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.Contract{}).Delete(&models.Contract{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         ContractController
// @Accept       json
// @Produce      json
// @Id ContractOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/contract/option [get]
func (r *ContractController) Option(ctx http.Context) http.Response {
	type Opt struct {
		Value uint   `json:"value"`
		Label string `json:"label"`
	}
	//todo change fields name follow model struct 
	//eg name value
	var opts []Opt
	facades.Orm().Query().Model(&models.Contract{}).Select("id as value,title as label").Scan(&opts)
	return httpfacade.NewResult(ctx).Success("", opts)
}
