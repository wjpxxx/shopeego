package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UnbindFirstMileTrackingNumberOrderListResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    PackageNumber	string	`json:"package_number"`
    FailError	string	`json:"fail_error"`
    FailMessage	string	`json:"fail_message"`
}
func (g UnbindFirstMileTrackingNumberOrderListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
