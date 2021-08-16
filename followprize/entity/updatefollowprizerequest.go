package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateFollowPrizeRequest struct{
    FollowPrizeName	string	`json:"follow_prize_name"`
    CampaignId	int64	`json:"campaign_id"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    UsageQuantity	int	`json:"usage_quantity"`
    MinSpend	float32	`json:"min_spend"`
}
func (g UpdateFollowPrizeRequest) String() string {
    return lib.ObjectToString(g)
}
