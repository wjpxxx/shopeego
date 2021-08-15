package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnbindFirstMileTrackingNumberRequest struct{
    FirstMileTrackingNumber	string	`json:"first_mile_tracking_number"`
    OrderList	[]UnbindFirstMileTrackingNumberOrderListRequestEntity	`json:"order_list"`
}
func (g UnbindFirstMileTrackingNumberRequest) String() string {
    return lib.ObjectToString(g)
}
