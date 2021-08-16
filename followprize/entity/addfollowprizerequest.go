package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddFollowPrizeRequest struct{
    FollowPrizeName	string	`json:"follow_prize_name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    UsageQuantity	int	`json:"usage_quantity"`
    MinSpend	float32	`json:"min_spend"`
    RewardType	int	`json:"reward_type"`
    DiscountAmount	float32	`json:"discount_amount"`
    Percentage	int	`json:"percentage"`
    MaxPrice	float32	`json:"max_price"`
}
func (g AddFollowPrizeRequest) String() string {
    return lib.ObjectToString(g)
}
