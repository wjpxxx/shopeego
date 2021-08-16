package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealRequest struct{
    AddOnDealName	string	`json:"add_on_deal_name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    PromotionType	int	`json:"promotion_type"`
    PurchaseMinSpend	float32	`json:"purchase_min_spend"`
    PerGiftNum	int	`json:"per_gift_num"`
    PromotionPurchaseLimit	int	`json:"promotion_purchase_limit"`
}
func (g AddAddOnDealRequest) String() string {
    return lib.ObjectToString(g)
}
