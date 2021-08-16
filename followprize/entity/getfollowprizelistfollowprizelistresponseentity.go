package entity

import (
    "github.com/wjpxxx/letgo/lib"
)
//GetFollowPrizeListFollowPrizeListResponseEntity
type GetFollowPrizeListFollowPrizeListResponseEntity struct{
	CampaignId	int64	`json:"campaign_id"`
    CampaignStatus	string	`json:"campaign_status"`
    FollowPrizeName	string	`json:"follow_prize_name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    UsageQuantity	int	`json:"usage_quantity"`
    Claimed	int	`json:"claimed"`
}

func (g GetFollowPrizeListFollowPrizeListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
