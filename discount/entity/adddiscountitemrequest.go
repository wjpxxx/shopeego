package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountItemRequest struct{
    DiscountId	int64	`json:"discount_id"`
    ItemIst	[]AddDiscountItemItemIstRequestEntity	`json:"item_ist"`
}
func (g AddDiscountItemRequest) String() string {
    return lib.ObjectToString(g)
}
