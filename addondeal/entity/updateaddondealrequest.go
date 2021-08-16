package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealRequest struct{
    AddOnDealId	int64	`json:"add_on_deal_id"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    PurchaseMinSpend	float32	`json:"purchase_min_spend"`
    PerGiftNum	int	`json:"per_gift_num"`
    PromotionPurchaseLimit	int	`json:"promotion_purchase_limit"`
    SubItemPriority	[]int	`json:"sub_item_priority"`
    AddOnDealName	string	`json:"add_on_deal_name"`
}
func (g UpdateAddOnDealRequest) String() string {
    return lib.ObjectToString(g)
}
