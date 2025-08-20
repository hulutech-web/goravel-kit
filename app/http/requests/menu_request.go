
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type MenuRequest struct {

	Pid string `json:"pid" form:"pid"`

	Title string `json:"title" form:"title"`

	Name string `json:"name" form:"name"`

	Path string `json:"path" form:"path"`

	Component string `json:"component" form:"component"`

	Icon string `json:"icon" form:"icon"`

	MenuType string `json:"menu_type" form:"menu_type"`

	Cacheable string `json:"cacheable" form:"cacheable"`

	RenderMenu string `json:"render_menu" form:"render_menu"`

	Permission string `json:"permission" form:"permission"`

	Sort string `json:"sort" form:"sort"`

	Target string `json:"target" form:"target"`

	Badge string `json:"badge" form:"badge"`

}

func (r *MenuRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *MenuRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *MenuRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"pid": "required",

		"title": "required",

		"name": "required",

		"path": "required",

		"component": "required",

		"icon": "",

		"menu_type": "required",

		"cacheable": "",

		"render_menu": "",

		"permission": "",

		"sort": "",

		"target": "",

		"badge": "",

	}
}

func (r *MenuRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *MenuRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *MenuRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
