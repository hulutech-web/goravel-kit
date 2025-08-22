
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type FileCateRequest struct {

	Name string `json:"name" form:"name"`

	Pid uint `json:"pid" form:"pid"`

	Sort int `json:"sort" form:"sort"`

	TenantId uint `json:"tenant_id" form:"tenant_id"`

	Type string `json:"type" form:"type"`

}

func (r *FileCateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *FileCateRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *FileCateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"name": "string",

		"pid": "required",

		"sort": "required",

		"tenant_id": "required",

		"type": "string",

	}
}

func (r *FileCateRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileCateRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileCateRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
