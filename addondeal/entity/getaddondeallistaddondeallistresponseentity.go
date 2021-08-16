package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealListAddOnDealListResponseEntity struct{
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    PromotionType	int	`json:"promotion_type"`
    PurchaseMinSpend	float32	`json:"purchase_min_spend"`
    AddOnDealId	int64	`json:"add_on_deal_id"`
    PerGiftNum	int	`json:"per_gift_num"`
    PromotionPurchaseLimit	int	`json:"promotion_purchase_limit"`
    AddOnDealName	string	`json:"add_on_deal_name"`
    Source	int	`json:"source"`
    SubItemPrioriry	[]int	`json:"sub_item_prioriry"`
}
func (g GetAddOnDealListAddOnDealListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
