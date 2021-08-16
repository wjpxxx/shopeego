package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformancePercentPreOrderListingDataResponseEntity struct{
    Target	string	`json:"target"`
    MyShopPerformance		string	`json:"my_shop_performance	"`
    PenaltyPoints	string	`json:"penalty_points"`
}
func (g ShopPerformancePercentPreOrderListingDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
