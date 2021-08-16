package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateFollowPrizeResult struct{
    Response	UpdateFollowPrizeResponseResponseEntity	`json:"response"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g UpdateFollowPrizeResult) String() string {
    return lib.ObjectToString(g)
}

type UpdateFollowPrizeResponseResponseEntity struct{
    CampaignId	int	`json:"campaign_id"`
}
func (g UpdateFollowPrizeResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}