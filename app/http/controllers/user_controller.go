package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/json"
	httpfacade "github.com/hulutech-web/http_result"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

// Index 分页查询，支持搜索，路由参数?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx,等其他任意的查询参数
// @Summary      分页查询
// @Description  分页查询
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserIndex
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param  name  query  string  false  "name"
// @Param  pageSize  query  string  false  "pageSize"
// @Param  currentPage  query  string  false  "currentPage"
// @Param  sort  query  string  false  "sort"
// @Param  order  query  string  false  "order"
// @Success 200 {string} json {}
// @Router       /api/admin/user [get]
func (r *UserController) Index(ctx http.Context) http.Response {
	users := []models.User{}
	queries := ctx.Request().Queries()
	res, err := httpfacade.NewResult(ctx).SearchByParams(queries, nil).ResultPagination(&users)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "", err.Error())
	}
	return res
}

// List 列表查询
// @Summary      列表查询
// @Description  列表查询
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserList
// @Param  username  query  string  false  "username"
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json {}
// @Router       /api/admin/user/list [get]
func (r *UserController) List(ctx http.Context) http.Response {
	users := []models.User{}
	queries := ctx.Request().Queries()
	return httpfacade.NewResult(ctx).SearchByParams(queries, nil).Success("", users)
}
func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")
	user := models.User{}
	facades.Orm().Query().Model(&models.User{}).Where("id = ?", id).First(&user)
	return httpfacade.NewResult(ctx).Success("", user)
}

// Store 新增
// @Summary      新增
// @Description  新增
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserStore
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param userData body requests.UserRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/admin/user [post]
func (r *UserController) Store(ctx http.Context) http.Response {
	var userRequest requests.UserRequest
	errors, err := ctx.Request().ValidateRequest(&userRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	user := models.User{}
	//todo add request values
	facades.Orm().Query().Model(&models.User{}).Create(&user)
	return httpfacade.NewResult(ctx).Success("创建成功", nil)
}

// Update
// @Summary      更新
// @Description  更新
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserUpdate
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param userData body requests.UserRequest true "用户数据"
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/user/{id} [put]
func (r *UserController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	user := models.User{}
	facades.Orm().Query().Model(&models.User{}).Where("id=?", id).Find(&user)
	var userRequest requests.UserRequest
	errors, err := ctx.Request().ValidateRequest(&userRequest)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//todo add request values
	facades.Orm().Query().Model(&models.User{}).Where("id=?", id).Save(&user)
	return httpfacade.NewResult(ctx).Success("修改成功", nil)
}

// Destroy 删除
// @Summary      删除
// @Description  删除
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserDestroy
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Param id path string true "id"  // 关键：指定参数位置为 path
// @Router       /api/admin/user/{id} [delete]
func (r *UserController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facades.Orm().Query().Model(&models.User{}).Delete(&models.User{}, id)
	return httpfacade.NewResult(ctx).Success("删除成功", nil)
}

// Option 选项
// @Summary      选项
// @Description  选项
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserOption
// @Security ApiKeyAuth
// @Param  username  query  string  false  "username"
// @Success 200 {string} json {}
// @Router       /api/admin/user/option [get]
func (r *UserController) Option(ctx http.Context) http.Response {
	type Opt struct {
		Value uint   `json:"value"`
		Label string `json:"label"`
	}
	username := ctx.Request().Query("username")

	var opts []Opt
	if username == "" {
		facades.Orm().Query().Model(&models.User{}).Select("id as value,username as label").Scan(&opts)
	} else {
		facades.Orm().Query().Model(&models.User{}).Where("username like ?", "%"+username+"%").Select("id as value,username as label").Scan(&opts)
	}

	return httpfacade.NewResult(ctx).Success("", opts)
}

// Own 当前登录人信息
// @Summary      选项
// @Description  选项
// @Tags         UserController
// @Accept       json
// @Produce      json
// @Id UserOwn
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/user/own [get]
func (r *UserController) Own(ctx http.Context) http.Response {
	user := models.User{}
	err := facades.Auth(ctx).User(&user)

	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "用户不存在", err.Error())
	}
	facades.Orm().Query().Model(&models.User{}).Find(&user)
	rs := []string{}
	ps := []string{}
	roles := user.GetRoles()
	permissions := ""
	permissions = user.GetPermissions()
	json.Unmarshal([]byte(roles), &rs)
	json.Unmarshal([]byte(permissions), &ps)
	return httpfacade.NewResult(ctx).Success("", http.Json{
		"user":        user,
		"roles":       rs,
		"permissions": ps,
	})
}
