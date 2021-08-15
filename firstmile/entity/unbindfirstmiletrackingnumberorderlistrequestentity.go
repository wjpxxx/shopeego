package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnbindFirstMileTrackingNumberOrderListRequestEntity struct{
    OrderSn	string	`json:"order_sn"`
    PackageNumber	string	`json:"package_number"`
}
func (g UnbindFirstMileTrackingNumberOrderListRequestEntity) String() string {
    return lib.ObjectToString(g)
}
