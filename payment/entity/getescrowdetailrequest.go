package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowDetailRequest struct{
    OrderSn	string	`json:"order_sn"`
}
func (g GetEscrowDetailRequest) String() string {
    return lib.ObjectToString(g)
}
