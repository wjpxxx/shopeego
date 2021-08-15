package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type BindFirstMileTrackingNumberWarningResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
}
func (g BindFirstMileTrackingNumberWarningResponseEntity) String() string {
    return lib.ObjectToString(g)
}
