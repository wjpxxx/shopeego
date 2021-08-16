package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealResponseResponseEntity struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
    Name	string	`json:"name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    BundleDealRule	GetBundleDealBundleDealRuleResponseEntity	`json:"bundle_deal_rule"`
    PurchaseLimit	int	`json:"purchase_limit"`
}
func (g GetBundleDealResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
