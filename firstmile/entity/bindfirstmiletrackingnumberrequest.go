package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type BindFirstMileTrackingNumberRequest struct{
    FirstMileTrackingNumber	string	`json:"first_mile_tracking_number"`
    ShipmentMethod	string	`json:"shipment_method"`
    Region	string	`json:"region"`
    LogisticsChannelId	int64	`json:"logistics_channel_id"`
    Volume	float32	`json:"volume"`
    Weight	float32	`json:"weight"`
    Width	float32	`json:"width"`
    Length	float32	`json:"length"`
    Height	float32	`json:"height"`
    OrderList	[]BindFirstMileTrackingNumberOrderListRequestEntity	`json:"order_list"`
}
func (g BindFirstMileTrackingNumberRequest) String() string {
    return lib.ObjectToString(g)
}
