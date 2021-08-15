package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountListDiscountListResponseEntity struct{
    Status	string	`json:"status"`
    DiscountName	int	`json:"discount_name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    DiscountId	int64	`json:"discount_id"`
}
func (g GetDiscountListDiscountListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
