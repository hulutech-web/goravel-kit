package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UserRequest struct {
	Unionid string `json:"unionid" form:"unionid"`

	Openid string `json:"openid" form:"openid"`

	Is_member string `json:"is_member" form:"is_member"`

	Avatar string `json:"avatar" form:"avatar"`

	Address string `json:"address" form:"address"`

	Remark string `json:"remark" form:"remark"`
	Sex    string `json:"sex" form:"sex"`

	Username string `json:"username" form:"username"`

	Phone string `json:"phone" form:"phone"`

	Password string `json:"password" form:"password"`

	Salt string `json:"salt" form:"salt"`

	Email string `json:"email" form:"email"`

	Role_id string `json:"role_id" form:"role_id"`

	Is_multipoint string `json:"is_multipoint" form:"is_multipoint"`

	Is_disable string `json:"is_disable" form:"is_disable"`

	Tenant_id string `json:"tenant_id" form:"tenant_id"`
}

func (r *UserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"avatar":   "required",
		"username": "required",

		"phone": "required",
		"sex":   "required",
	}
}

func (r *UserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
