package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPenaltyOngoingPunishmentResponseEntity struct{
    PunishmentTier	int	`json:"punishment_tier"`
    DaysLeft	int	`json:"days_left"`
    PunishmentName	string	`json:"punishment_name"`
}
func (g ShopPenaltyOngoingPunishmentResponseEntity) String() string {
    return lib.ObjectToString(g)
}
