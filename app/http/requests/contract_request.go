package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/support/carbon"
)

type ContractRequest struct {
	Content string `json:"content" form:"content"`

	LandlordId uint `json:"landlord_id" form:"landlord_id"`

	LandlordSign string `json:"landlord_sign" form:"landlord_sign"`

	OrderId uint `json:"order_id" form:"order_id"`

	PaperContract string `json:"paper_contract" form:"paper_contract"`

	SignedLocation string `json:"signed_location" form:"signed_location"`

	SignedTime carbon.DateTime `json:"signed_time" form:"signed_time"`

	TenantId uint `json:"tenant_id" form:"tenant_id"`

	TenantSign string `json:"tenant_sign" form:"tenant_sign"`

	Type string `json:"type" form:"type"`
}

func (r *ContractRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *ContractRequest) Filters(ctx http.Context) error {
	return nil
}

func (r *ContractRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"content": "required",

		"landlord_id": "required",

		"landlord_sign": "required",

		"order_id": "required",

		"paper_contract": "",

		"signed_location": "",

		"signed_time": "",

		"tenant_id": "",

		"tenant_sign": "",

		"type": "",
	}
}

func (r *ContractRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ContractRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ContractRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
