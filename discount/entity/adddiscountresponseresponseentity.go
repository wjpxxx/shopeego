package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountResponseResponseEntity struct{
    DiscountId	int64	`json:"discount_id"`
}
func (g AddDiscountResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
