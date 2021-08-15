package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnbindFirstMileTrackingNumberWarningResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    PackageNumber	string	`json:"package_number"`
}
func (g UnbindFirstMileTrackingNumberWarningResponseEntity) String() string {
    return lib.ObjectToString(g)
}
