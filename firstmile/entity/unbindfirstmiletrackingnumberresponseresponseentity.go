package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnbindFirstMileTrackingNumberResponseResponseEntity struct{
    FirstMileTrackingNumber	string	`json:"first_mile_tracking_number"`
    OrderList	[]UnbindFirstMileTrackingNumberOrderListResponseEntity	`json:"order_list"`
}
func (g UnbindFirstMileTrackingNumberResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
