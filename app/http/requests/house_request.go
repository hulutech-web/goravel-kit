
package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type HouseRequest struct {

	LandlordId string `json:"landlord_id" form:"landlord_id"`

	Title string `json:"title" form:"title"`

	Description string `json:"description" form:"description"`

	Addres string `json:"address" form:"address"`

	MonthlyRent string `json:"monthly_rent" form:"monthly_rent"`

	Deposit string `json:"deposit" form:"deposit"`

	HeaderImg string `json:"header_img" form:"header_img"`

	Poster string `json:"poster" form:"poster"`

	Album string `json:"albums" form:"albums"`

	Location string `json:"location" form:"location"`

	Area string `json:"area" form:"area"`

	Facilitie string `json:"facilities" form:"facilities"`

	PropertyFee string `json:"property_fee" form:"property_fee"`

	Traffic string `json:"traffic" form:"traffic"`

	Shopping string `json:"shopping" form:"shopping"`

	Video string `json:"video" form:"video"`

	Statu string `json:"status" form:"status"`

	Swiper string `json:"swipers" form:"swipers"`

}

func (r *HouseRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *HouseRequest) Filters(ctx http.Context) error {
	return nil
}


func (r *HouseRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{

		"landlord_id": "required",

		"title": "required",

		"description": "required",

		"address": "required",

		"monthly_rent": "required",

		"deposit": "required",

		"header_img": "required",

		"poster": "required",

		"albums": "required",

		"location": "required",

		"area": "required",

		"facilities": "required",

		"property_fee": "required",

		"traffic": "required",

		"shopping": "required",

		"video": "required",

		"status": "required",

		"swipers": "required",

	}
}

func (r *HouseRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *HouseRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *HouseRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
