package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPenaltyResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    PenaltyPoints	ShopPenaltyPenaltyPointsResponseEntity	`json:"penalty_points"`
    OngoingPunishment	[]ShopPenaltyOngoingPunishmentResponseEntity	`json:"ongoing_punishment"`
}
func (g ShopPenaltyResult) String() string {
    return lib.ObjectToString(g)
}
