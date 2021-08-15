package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutDetailRequest struct{
    PageSize	int	`json:"page_size"`
    PageNo	int	`json:"page_no"`
    PayoutTimeFrom	int	`json:"payout_time_from"`
    PayoutTimeTo	int	`json:"payout_time_to"`
}
func (g GetPayoutDetailRequest) String() string {
    return lib.ObjectToString(g)
}
