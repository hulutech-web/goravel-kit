package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type AuthCodeRequest struct {
	Code string `form:"code" json:"code"`
}

func (r *AuthCodeRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *AuthCodeRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthCodeRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"code": "required",
	}
}

func (r *AuthCodeRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthCodeRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *AuthCodeRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
