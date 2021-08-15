package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountListResponseResponseEntity struct{
    DiscountList	[]GetDiscountListDiscountListResponseEntity	`json:"discount_list"`
}
func (g GetDiscountListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
