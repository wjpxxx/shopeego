package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutDetailPayoutListResponseEntity struct{
    PayoutInfo	GetPayoutDetailPayoutInfoResponseEntity	`json:"payout_info"`
    EscrowList	[]GetPayoutDetailEscrowListResponseEntity	`json:"escrow_list"`
    OfflineAdjustmentList	[]GetPayoutDetailOfflineAdjustmentListResponseEntity	`json:"offline_adjustment_list"`
}
func (g GetPayoutDetailPayoutListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
