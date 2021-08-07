package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetShippingParameterResult
type GetShippingParameterResult struct {
	commonentity.Result
	Response GetShippingParameterResponse `json:"response"`
}

//String
func (g GetShippingParameterResult) String() string {
	return lib.ObjectToString(g)
}

//GetShippingParameterResponse
type GetShippingParameterResponse struct {
	InfoNeeded InfoNeededEntity `json:"info_needed"`
	Dropoff    DropoffEntity    `json:"dropoff"`
	Pickup     PickupEntity     `json:"pickup"`
}

//String
func (g GetShippingParameterResponse) String() string {
	return lib.ObjectToString(g)
}
