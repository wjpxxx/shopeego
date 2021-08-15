package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountRequest struct{
    DiscountName	string	`json:"discount_name"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
}
func (g AddDiscountRequest) String() string {
    return lib.ObjectToString(g)
}
