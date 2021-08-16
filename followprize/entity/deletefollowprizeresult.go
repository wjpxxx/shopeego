package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteFollowPrizeResult struct{
    Response	DeleteFollowPrizeResponseResponseEntity	`json:"response"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g DeleteFollowPrizeResult) String() string {
    return lib.ObjectToString(g)
}

type DeleteFollowPrizeResponseResponseEntity struct{
    CampaignId	int64	`json:"campaign_id"`
}
func (g DeleteFollowPrizeResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
