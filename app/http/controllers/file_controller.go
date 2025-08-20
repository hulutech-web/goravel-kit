package controllers

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/path"
	httpfacade "github.com/hulutech-web/http_result"
	"github.com/spf13/cast"
	"goravel/app/http/requests"
	"goravel/app/models"
	"os"
	"path/filepath"
)

type FileController struct {
	//Dependent services
}

func NewFileController() *FileController {
	return &FileController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/admin/file [get]
func (r *FileController) Index(ctx http.Context) http.Response {
	files := []models.File{}
	queries := ctx.Request().Queries()
	if cast.ToInt(queries["cid"]) == 0 {
		pagination, _ := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&files)
		return pagination

	} else {
		pagination, _ := httpfacade.NewResult(ctx).SearchByParams(queries, map[string]interface{}{
			"cid": queries["cid"],
		}).ResultPagination(&files)
		return pagination
	}
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileList
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/admin/file/list [get]
func (r *FileController) List(ctx http.Context) http.Response {
	files := []models.File{}
	queries := ctx.Request().Queries()
	if queries["cid"] == "" {
		return httpfacade.NewResult(ctx).SearchByParams(queries, nil).Success("", files)

	} else {
		return httpfacade.NewResult(ctx).SearchByParams(queries, map[string]interface{}{
			"cid": queries["cid"],
		}).Success("", files)
	}
}
func (r *FileController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	file := models.File{}
	facades.Orm().Query().Model(&models.File{}).Where("id = ?", id).First(&file)
	return httpfacade.NewResult(ctx).Success("", file)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param fileData body requests.FileRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/file [post]
func (r *FileController) Store(ctx http.Context) http.Response {
	var fileRequest requests.FileRequest
	errors, err := ctx.Request().ValidateRequest(&fileRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	file := models.File{}
	//todo add request values
	facades.Orm().Query().Model(&models.File{}).Create(&file)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param fileData body requests.FileRequest true "用户数据"
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/file/{id} [put]
func (r *FileController) Update(ctx http.Context) http.Response {
	user := models.User{}
	facades.Auth(ctx).User(&user)
	id := ctx.Request().Route("id")
	file := models.File{}
	facades.Orm().Query().Model(&models.File{}).Where("id=?", id).Find(&file)
	var fileRequest requests.FileRequest
	errors, err := ctx.Request().ValidateRequest(&fileRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.File{}).Where("id=?", id).Find(&file)
	file.Name = fileRequest.Name
	file.Engine = fileRequest.Engine
	file.Ext = fileRequest.Ext
	file.Path = fileRequest.Path
	file.Size = fileRequest.Size
	file.Type = fileRequest.Type
	file.TenantID = fileRequest.TenantId
	file.Uri = fileRequest.Uri
	file.UserID = user.ID
	file.CID = fileRequest.Cid
	facades.Orm().Query().Model(&models.File{}).Where("id=?", id).Save(&file)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/file/{id} [delete]
func (r *FileController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.File{}).Delete(&models.File{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileOption
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/file/option [get]
func (r *FileController) Option(ctx http.Context) http.Response {
	files := []models.File{}
	queries := ctx.Request().Queries()
	res, _ := httpfacade.NewResult(ctx).SearchByParams(queries, nil).List(&files)
	return res
}

// DelBatch 批量删除
// @Summary      批量删除
// @Description  批量删除
// @Tags         FileController
// @Accept       json
// @Produce      json
// @Id FileDelBatch
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param fileData body requests.FileDelRequest true "删除数据"
// @Success 200 {string} json {}
// @Router       /api/admin/file_del_batch [post]
func (r *FileController) DelBatch(ctx http.Context) http.Response {
	var fileRequest requests.FileDelRequest
	errors, err := ctx.Request().ValidateRequest(&fileRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	files := []models.File{}
	facades.Orm().Query().Model(&models.File{}).Where("id in ?", fileRequest.DelIDs).Find(&files)
	for _, file := range files {
		del_path := filepath.Join(path.Public("uploads"), file.Path)
		fmt.Println(del_path)
		if file.ID > 0 {
			os.Remove(del_path)
		}
	}
	facades.Orm().Query().Model(&models.File{}).Where("id in ?", fileRequest.DelIDs).Delete()
	return httpfacade.NewResult(ctx).Success("操作成功", nil)
}
