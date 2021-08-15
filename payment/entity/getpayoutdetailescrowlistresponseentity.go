package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutDetailEscrowListResponseEntity struct{
    EscrowAmount	float32	`json:"escrow_amount"`
    Currency	string	`json:"currency"`
    OrderSn	string	`json:"order_sn"`
}
func (g GetPayoutDetailEscrowListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
