package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndDiscountRequest struct{
    DiscountId	int64	`json:"discount_id"`
}
func (g EndDiscountRequest) String() string {
    return lib.ObjectToString(g)
}
