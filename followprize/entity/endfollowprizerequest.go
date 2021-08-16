package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndFollowPrizeRequest struct{
    CampaignId	int64	`json:"campaign_id"`
}
func (g EndFollowPrizeRequest) String() string {
    return lib.ObjectToString(g)
}
