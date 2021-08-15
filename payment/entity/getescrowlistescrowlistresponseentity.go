package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowListEscrowListResponseEntity struct{
    OrderSn	string	`json:"order_sn"`
    PayoutAmount	float32	`json:"payout_amount"`
    EscrowReleaseTime	int	`json:"escrow_release_time"`
}
func (g GetEscrowListEscrowListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
