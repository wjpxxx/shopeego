package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFollowPrizeDetailResult struct{
    Response	GetFollowPrizeDetailResponseResponseEntity	`json:"response"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetFollowPrizeDetailResult) String() string {
    return lib.ObjectToString(g)
}

type GetFollowPrizeDetailResponseResponseEntity struct{
    CampaignStatus	string	`json:"campaign_status"`
    CampaginId	int64	`json:"campagin_id"`
    UsageQuantity	int	`json:"usage_quantity"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    MinSpend	float32	`json:"min_spend"`
    RewardType	int	`json:"reward_type"`
    FollowPrizeName	string	`json:"follow_prize_name"`
    DiscountAmount	float32	`json:"discount_amount"`
    Percentage	int	`json:"percentage"`
    MaxPrice	float32	`json:"max_price"`
}

func (g GetFollowPrizeDetailResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}