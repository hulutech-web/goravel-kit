package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type FileRequest struct {
	Cid uint `json:"cid" form:"cid"`

	UserId string `json:"user_id" form:"user_id"`

	Type string `json:"type" form:"type"`

	Name string `json:"name" form:"name"`

	Uri string `json:"uri" form:"uri"`

	Ext string `json:"ext" form:"ext"`

	Size int64 `json:"size" form:"size"`

	Engine string `json:"engine" form:"engine"`

	Path string `json:"path" form:"path"`

	TenantId uint `json:"tenant_id" form:"tenant_id"`
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

		"user_id": "required",

		"type": "required",

		"name": "required",

		"uri": "required",

		"ext": "required",

		"size": "required",

		"engine": "required",

		"path": "required",

		"tenant_id": "",
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
