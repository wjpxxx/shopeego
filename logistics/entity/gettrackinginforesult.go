package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetTrackingInfoResult
type GetTrackingInfoResult struct {
	commonentity.Result
	Response GetTrackingInfoResponse `json:"response"`
}

//String
func (g GetTrackingInfoResult) String() string {
	return lib.ObjectToString(g)
}

//GetShippingParameterResponse
type GetTrackingInfoResponse struct {
	OrderSn         string               `json:"order_sn"`
	PackageNumber   string               `json:"package_number"`
	LogisticsStatus string               `json:"logistics_status"`
	TrackingInfo    []TrackingInfoEntity `json:"tracking_info"`
}

//String
func (g GetTrackingInfoResponse) String() string {
	return lib.ObjectToString(g)
}
