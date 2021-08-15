package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutDetailResponseResponseEntity struct{
    More	bool	`json:"more"`
    PayoutList	[]GetPayoutDetailPayoutListResponseEntity	`json:"payout_list"`
}
func (g GetPayoutDetailResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
