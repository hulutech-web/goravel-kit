
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type FileRequest struct {

	Cid uint `json:"cid" form:"cid"`

	Engine string `json:"engine" form:"engine"`

	Ext string `json:"ext" form:"ext"`

	Name string `json:"name" form:"name"`

	Path string `json:"path" form:"path"`

	Size int64 `json:"size" form:"size"`

	TenantId uint `json:"tenant_id" form:"tenant_id"`

	Type string `json:"type" form:"type"`

	Uri string `json:"uri" form:"uri"`

	UserId uint `json:"user_id" form:"user_id"`

}

func (r *FileRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *FileRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *FileRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"cid": "required",

		"engine": "required",

		"ext": "required",

		"name": "required",

		"path": "required",

		"size": "required",

		"tenant_id": "required",

		"type": "string",

		"uri": "string",

		"user_id": "required",

	}
}

func (r *FileRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
