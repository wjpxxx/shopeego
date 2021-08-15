package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateDiscountRequest struct{
    DiscountId	int64	`json:"discount_id"`
    DiscountName	string	`json:"discount_name"`
    EndTime	int	`json:"end_time"`
    StartTime	int	`json:"start_time"`
}
func (g UpdateDiscountRequest) String() string {
    return lib.ObjectToString(g)
}
