package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetUnbindOrderListOrderListResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    PackageNumber	string	`json:"package_number"`
    LogisticsStatus	string	`json:"logistics_status"`
}
func (g GetUnbindOrderListOrderListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
