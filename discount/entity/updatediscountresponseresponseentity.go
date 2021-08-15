package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateDiscountResponseResponseEntity struct{
    DiscountId	int64	`json:"discount_id"`
    ModifyTime	int	`json:"modify_time"`
}
func (g UpdateDiscountResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
