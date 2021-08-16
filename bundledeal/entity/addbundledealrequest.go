package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddBundleDealRequest struct{
    RuleType	int	`json:"rule_type"`
    DiscountValue	float32	`json:"discount_value"`
    FixPrice	float32	`json:"fix_price"`
    DiscountPercentage	int	`json:"discount_percentage"`
    MinAmount	int	`json:"min_amount"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    Name	string	`json:"name"`
    PurchaseLimit	int	`json:"purchase_limit"`
}
func (g AddBundleDealRequest) String() string {
    return lib.ObjectToString(g)
}
