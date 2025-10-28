package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"
	"github.com/goravel/framework/support/json"
	httpfacade "github.com/hulutech-web/http_result"
	"goravel/app/http/controllers/common"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type AuthController struct {
	*common.WechatService
}

func NewAuthController() *AuthController {
	return &AuthController{
		//Inject services
	}
}

// Login
// @Summary      后台登录
// @Description  后台登录
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthLogin
// @Security ApiKeyAuth
// @Param userData body requests.LoginRequest true "登录数据"
// @Success 200 {string} json {}
// @Router       /api/admin/auth/login [post]
func (r *AuthController) Login(ctx http.Context) http.Response {
	var user models.User
	var login_request requests.LoginRequest
	errors, err := ctx.Request().ValidateRequest(&login_request)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	//手机号密码验证
	facades.Orm().Query().Model(&user).Where("phone", login_request.Account).
		OrWhere("username", login_request.Account).First(&user)

	if user.ID == 0 {
		ctx.Request().AbortWithStatusJson(500, http.Json{
			"message": "error",
			"fail":    "用户不存在,请点击注册",
		})
		return nil
	}
	var user_exist models.User
	facades.Orm().Query().Model(&user).Where("phone", login_request.Account).
		OrWhere("username", login_request.Account).First(&user_exist)
	//解密
	if user_exist.ID == 0 {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"message": "无权登录",
		})
	}

	if !facades.Hash().Check(login_request.Password, user_exist.Password) {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
			"message": "密码错误",
		})
	} else {
		//	生成token
		token, err1 := facades.Auth(ctx).Login(&user_exist)
		if err1 != nil {
			return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
				"message": "token生成失败",
			})
		}
		facades.Orm().Query().Model(&user_exist).Update("last_login", carbon.NewDateTime(carbon.Now()))
		return ctx.Response().Status(http.StatusOK).Json(http.Json{
			"message": "登录成功",
			"data": struct {
				Token string      `json:"token"`
				User  models.User `json:"user"`
			}{
				Token: token,
				User:  user_exist,
			},
		})
	}
}

// Menu 当前登录人信息
// @Summary      选项
// @Description  选项
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthMenu
// @Security ApiKeyAuth
// @Success 200 {string} json {}
// @Router       /api/admin/menu/route [get]
func (r *AuthController) Menu(ctx http.Context) http.Response {
	user := models.User{}
	facades.Auth(ctx).User(&user)
	facades.Orm().Query().With("Roles").Find(&user)
	menus := user.GetMenus()
	ms := []models.Menu{}
	//menus转换为[]byte,并json解析到ms中
	json.Unmarshal([]byte(menus), &ms)
	return httpfacade.NewResult(ctx).Success("", map[string]any{
		"menus": ms,
	})
}

// Logout 退出登录
func (r *AuthController) Logout(ctx http.Context) http.Response {
	return nil
}

// MiniLogin 小程序登录
// @Summary      登录
// @Description  登录
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthMiniLogin
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param eventData body interface{} true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/mini/login [post]
func (r *AuthController) MiniLogin(ctx http.Context) http.Response {
	var user models.User
	openid := ctx.Request().Input("openid", "")
	unionid := ctx.Request().Input("unionid", "")
	facades.Orm().Query().Model(&models.User{}).Where("openid=?", openid).With("Coach").First(&user)
	if user.ID == 0 {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "用户不存在,请点击注册", "")

	} else {
		//更新user
		if _, err := facades.Orm().Query().Model(models.User{}).Where("id=?", user.ID).Update(models.User{Openid: openid, Unionid: unionid, LastLogin: carbon.NewDateTime(carbon.Now())}); err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, http.Json{
				"error": "更新用户信息失败",
			})
		}
		if token, err2 := facades.Auth(ctx).Login(&user); err2 != nil {
			return httpfacade.NewResult(ctx).Error(http.StatusUnprocessableEntity, "用户授权失败", "")

		} else {
			return httpfacade.NewResult(ctx).Success("登录成功", http.Json{
				"token": token,
				"user":  user,
			})
		}
	}
}

// MiniLogout 小程序退出
// @Summary      登录
// @Description  登录
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthMiniLogout
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param eventData body interface{} true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/mini/logout [get]
func (r *AuthController) MiniLogout(ctx http.Context) http.Response {
	facades.Auth(ctx).Logout()
	return httpfacade.NewResult(ctx).Success("退出成功", nil)
}

// MiniOpenid 小程序退出
// @Summary      登录
// @Description  登录
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthMiniOpenid
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param code body string true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/mini/openid [post]
func (r *AuthController) MiniOpenid(ctx http.Context) http.Response {

	code := ctx.Request().Input("code")
	openid, unionid, err := r.GetOpenidByCode(code)
	if err != nil {
		ctx.Request().AbortWithStatusJson(500, http.Json{
			"message": "获取openid失败" + err.Error(),
		})
		return nil
	}
	return httpfacade.NewResult(ctx).Success("", http.Json{
		"openid":  openid,
		"unionid": unionid,
	})
}

// MiniRegist 小程序注册
// @Summary      登录
// @Description  登录
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthMiniRegist
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param code body models.User true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/mini/regist [post]
func (r *AuthController) MiniRegist(ctx http.Context) http.Response {
	var user models.User
	ctx.Request().Bind(&user)
	//if user.Phone == "" {
	//	return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{
	//		"error": "手机号不能为空",
	//	})
	//}
	facades.Orm().Query().Model(&user).Where("openid", user.Openid).First(&user)
	if user.ID > 0 {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "用户已存在", "")
	}
	facades.Orm().Query().Create(&user)
	//直接登录
	token, _ := facades.Auth(ctx).Login(&user)

	return httpfacade.NewResult(ctx).Success("注册成功", http.Json{
		"token": token,
		"user":  user,
	})
}

// MiniPhone 获取手机号
// @Summary      获取手机号
// @Description  获取手机号
// @Tags         AuthController
// @Accept       json
// @Produce      json
// @Id AuthMiniPhone
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param courseData body requests.AuthCodeRequest true "用户数据"
// @Success 200 {string} json {}
// @Router       /api/mini/phone [post]
func (r *AuthController) MiniPhone(ctx http.Context) http.Response {
	var user models.User
	facades.Auth(ctx).User(&user)
	var request requests.AuthCodeRequest
	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "数据错误", err.Error())
	}
	if errors != nil {
		return httpfacade.NewResult(ctx).ValidError("", errors.All())
	}
	phone, err := r.GetPhoneNumberByCode(request.Code)
	if err != nil {
		return httpfacade.NewResult(ctx).Error(http.StatusInternalServerError, "获取手机号失败", "")
	}
	user.Phone = phone
	facades.Orm().Query().Save(&user)
	return httpfacade.NewResult(ctx).Success("获取手机号成功", user)
}
