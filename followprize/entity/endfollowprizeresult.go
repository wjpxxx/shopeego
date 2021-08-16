package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndFollowPrizeResult struct{
    Response	EndFollowPrizeResponseResponseEntity	`json:"response"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g EndFollowPrizeResult) String() string {
    return lib.ObjectToString(g)
}

type EndFollowPrizeResponseResponseEntity struct{
    CampaignId	int64	`json:"campaign_id"`
}

func (g EndFollowPrizeResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
