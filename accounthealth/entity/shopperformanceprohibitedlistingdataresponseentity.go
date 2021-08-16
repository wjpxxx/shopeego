package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceProhibitedListingDataResponseEntity struct{
    Target	string	`json:"target"`
    MyShopPerformance		string	`json:"my_shop_performance	"`
    PenaltyPoints	string	`json:"penalty_points"`
}
func (g ShopPerformanceProhibitedListingDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
