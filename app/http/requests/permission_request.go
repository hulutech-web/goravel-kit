
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type PermissionRequest struct {

	Name string `json:"name" form:"name"`

	Code string `json:"code" form:"code"`

	Type string `json:"type" form:"type"`

	Description string `json:"description" form:"description"`

	MenuId string `json:"menu_id" form:"menu_id"`

}

func (r *PermissionRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *PermissionRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *PermissionRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"name": "required",

		"code": "required",

		"type": "required",

		"description": "",

		"menu_id": "",

	}
}

func (r *PermissionRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PermissionRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
