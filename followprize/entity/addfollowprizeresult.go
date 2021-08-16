package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddFollowPrizeResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	AddFollowPrizeResponseResponseEntity	`json:"response"`
}
func (g AddFollowPrizeResult) String() string {
    return lib.ObjectToString(g)
}

type AddFollowPrizeResponseResponseEntity struct{
    CampaignId	int64	`json:"campaign_id"`
}

func (g AddFollowPrizeResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}