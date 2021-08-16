package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ShopPerformanceTotalDataResponseEntity struct{
    Target	string	`json:"target"`
    MyShopPerformance	string	`json:"my_shop_performance"`
    PenaltyPoints	string	`json:"penalty_points"`
}
func (g ShopPerformanceTotalDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
