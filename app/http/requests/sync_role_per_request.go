package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type SyncRolePerRequest struct {
	FormIDs []uint `json:"formIDs" form:"formIDs"`
}

func (r *SyncRolePerRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *SyncRolePerRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SyncRolePerRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"formIDs": "slice",
	}
}

func (r *SyncRolePerRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SyncRolePerRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SyncRolePerRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
