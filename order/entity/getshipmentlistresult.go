package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/shopeego/commonentity"
)

//GetShipmentListResult
type GetShipmentListResult struct {
	commonentity.Result
	Response GetShipmentListResponse `json:"response"`
}

//String
func (g GetShipmentListResult) String() string {
	return lib.ObjectToString(g)
}

//GetShipmentListResponse
type GetShipmentListResponse struct {
	More       bool             `json:"more"`
	NextCursor string           `json:"next_cursor"`
	OrderList  []ShipmentOrders `json:"order_list"`
}

//String
func (g GetShipmentListResponse) String() string {
	return lib.ObjectToString(g)
}

//ShipmentOrders
type ShipmentOrders struct {
	OrderSN       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
}

//String
func (g ShipmentOrders) String() string {
	return lib.ObjectToString(g)
}
