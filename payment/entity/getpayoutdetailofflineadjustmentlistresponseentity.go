package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutDetailOfflineAdjustmentListResponseEntity struct{
    AdjustmentAmount	float32	`json:"adjustment_amount"`
    Module	string	`json:"module"`
    Remark	string	`json:"remark"`
    Scenario	string	`json:"scenario"`
    AdjustmentLevel	string	`json:"adjustment_level"`
    OrderSn	string	`json:"order_sn"`
}
func (g GetPayoutDetailOfflineAdjustmentListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
