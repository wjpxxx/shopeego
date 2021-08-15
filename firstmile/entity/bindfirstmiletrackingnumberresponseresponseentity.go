package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type BindFirstMileTrackingNumberResponseResponseEntity struct{
    FirstMileTrackingNumber	string	`json:"first_mile_tracking_number"`
    OrderList	[]BindFirstMileTrackingNumberOrderListResponseEntity	`json:"order_list"`
}
func (g BindFirstMileTrackingNumberResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
