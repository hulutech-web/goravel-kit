
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type FileCateRequest struct {

	Name string `json:"name" form:"name"`

	Sort string `json:"sort" form:"sort"`

	Type string `json:"type" form:"type"`

	Pid string `json:"pid" form:"pid"`

	TenantId string `json:"tenant_id" form:"tenant_id"`

}

func (r *FileCateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *FileCateRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *FileCateRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"name": "required",

		"sort": "required",

		"type": "required",

		"pid": "required",

		"tenant_id": "required",

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
