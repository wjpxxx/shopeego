package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceCancellationRateDataResponseEntity struct{
    Target	string	`json:"target"`
    MyShopPerformance	string	`json:"my_shop_performance"`
    PenaltyPoints	string	`json:"penalty_points"`
}
func (g ShopPerformanceCancellationRateDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
