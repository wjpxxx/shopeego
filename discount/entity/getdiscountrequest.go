package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountRequest struct{
    DiscountId	int64	`json:"discount_id"`
    PageNo	int	`json:"page_no"`
    PageSize	int	`json:"page_size"`
}
func (g GetDiscountRequest) String() string {
    return lib.ObjectToString(g)
}
