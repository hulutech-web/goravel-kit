package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httpfacade "github.com/hulutech-web/http_result"
	"github.com/spf13/cast"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type FileCateController struct {
	//Dependent services
}

func NewFileCateController() *FileCateController {
	return &FileCateController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         FileCateController
// @Accept       json
// @Produce      json
// @Id FileCateIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/admin/file_cate [get]
func (r *FileCateController) Index(ctx http.Context) http.Response {
	file_cates := []models.FileCate{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&file_cates)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         FileCateController
// @Accept       json
// @Produce      json
// @Id FileCateList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/admin/file_cate/list [get]
func (r *FileCateController) List(ctx http.Context) http.Response {
	file_cates := []models.FileCate{}
	queries := ctx.Request().Queries()
	queries["pageSize"] = "20"
	httpfacade.NewResult(ctx).SearchByParams(queries, nil).List(&file_cates)
	return httpfacade.NewResult(ctx).Success("", file_cates)
}
func (r *FileCateController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	file_cate := models.FileCate{}
	facades.Orm().Query().Model(&models.FileCate{}).Where("id = ?", id).First(&file_cate)
	return httpfacade.NewResult(ctx).Success("", file_cate)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         FileCateController
// @Accept       json
// @Produce      json
// @Id FileCateStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param file_cateData body requests.FileCateRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/file_cate [post]
func (r *FileCateController) Store(ctx http.Context) http.Response {
	var file_cateRequest requests.FileCateRequest
	errors, err := ctx.Request().ValidateRequest(&file_cateRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	file_cate := models.FileCate{
		Name:     file_cateRequest.Name,
		PID:      cast.ToUint(file_cateRequest.Pid),
		Type:     file_cateRequest.Type,
		TenantID: 1,
	}
	//todo add request values
	facades.Orm().Query().Model(&models.FileCate{}).Create(&file_cate)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         FileCateController
// @Accept       json
// @Produce      json
// @Id FileCateUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param file_cateData body requests.FileCateRequest true "用户数据"
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/file_cate/{id} [put]
func (r *FileCateController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	file_cate := models.FileCate{}
	facades.Orm().Query().Model(&models.FileCate{}).Where("id=?", id).Find(&file_cate)
	var file_cateRequest requests.FileCateRequest
	errors, err := ctx.Request().ValidateRequest(&file_cateRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.FileCate{}).Where("id=?", id).Save(&file_cate)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         FileCateController
// @Accept       json
// @Produce      json
// @Id FileCateDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/file_cate/{id} [delete]
func (r *FileCateController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.FileCate{}).Delete(&models.FileCate{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         FileCateController
// @Accept       json
// @Produce      json
// @Id FileCateOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/file_cate/option [get]
func (r *FileCateController) Option(ctx http.Context) http.Response {
	file_cates := []models.FileCate{}
	queries := ctx.Request().Queries()
	res, _ := httpfacade.NewResult(ctx).SearchByParams(queries, nil).List(&file_cates)
	return res
}

func (r *FileCateController) Files(ctx http.Context) http.Response {
	files := []models.File{}
	queries := ctx.Request().Queries()
	res, _ := httpfacade.NewResult(ctx).SearchByParams(queries, nil).List(&files)
	return res
}
