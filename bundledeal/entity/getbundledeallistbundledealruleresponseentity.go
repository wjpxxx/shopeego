package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealListBundleDealRuleResponseEntity struct{
    RuleType	int	`json:"rule_type"`
    DiscountValue	float32	`json:"discount_value"`
    FixPrice	float32	`json:"fix_price"`
    DiscountPercentage	int	`json:"discount_percentage"`
    MinAmount	int	`json:"min_amount"`
}
func (g GetBundleDealListBundleDealRuleResponseEntity) String() string {
    return lib.ObjectToString(g)
}
