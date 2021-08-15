package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWalletTransactionListPayOrderListResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    ShopName	string	`json:"shop_name"`
}
func (g GetWalletTransactionListPayOrderListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
