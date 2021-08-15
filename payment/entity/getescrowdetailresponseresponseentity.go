package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowDetailResponseResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    BuyerUserName	string	`json:"buyer_user_name"`
    ReturnOrderSnList	[]string	`json:"return_order_sn_list"`
    OrderIncome	GetEscrowDetailOrderIncomeResponseEntity	`json:"order_income"`
}
func (g GetEscrowDetailResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
