package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteDiscountRequest struct{
    DiscountId	int64	`json:"discount_id"`
}
func (g DeleteDiscountRequest) String() string {
    return lib.ObjectToString(g)
}
