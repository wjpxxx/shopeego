package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteDiscountItemResponseResponseEntity struct{
    DiscountId	int64	`json:"discount_id"`
    ErrorList	[]DeleteDiscountItemErrorListResponseEntity	`json:"error_list"`
}
func (g DeleteDiscountItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
