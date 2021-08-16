package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceNumberOfDaysPreOrderListingExceedsTargetDaResponseEntity struct{
    Target	string	`json:"target"`
    MyShopPerformance	string	`json:"my_shop_performance"`
    PenaltyPoints	string	`json:"penalty_points"`
}
func (g ShopPerformanceNumberOfDaysPreOrderListingExceedsTargetDaResponseEntity) String() string {
    return lib.ObjectToString(g)
}
