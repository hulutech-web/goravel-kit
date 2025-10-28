
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RoleRequest struct {

	Name string `json:"name" form:"name"`

	Label string `json:"label" form:"label"`

	Remark string `json:"remark" form:"remark"`

	IsDisable string `json:"is_disable" form:"is_disable"`

	Sort string `json:"sort" form:"sort"`

	TenantId string `json:"tenant_id" form:"tenant_id"`

	IsAdmin string `json:"is_admin" form:"is_admin"`

}

func (r *RoleRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RoleRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *RoleRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"name": "required",

		"label": "required",

		"remark": "required",

		"is_disable": "required",

		"sort": "required",

		"tenant_id": "required",

		"is_admin": "required",

	}
}

func (r *RoleRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *RoleRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
