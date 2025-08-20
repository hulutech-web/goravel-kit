package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type FileDelRequest struct {
	DelIDs []int64 `form:"del_ids" json:"del_ids"`
}

func (r *FileDelRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *FileDelRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileDelRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"del_ids": "required",
	}
}

func (r *FileDelRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileDelRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *FileDelRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
