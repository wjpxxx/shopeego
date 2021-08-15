package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteDiscountItemRequest struct{
    DiscountId	int64	`json:"discount_id"`
    ItemId	int64	`json:"item_id"`
    ModelId	int64	`json:"model_id"`
}
func (g DeleteDiscountItemRequest) String() string {
    return lib.ObjectToString(g)
}
