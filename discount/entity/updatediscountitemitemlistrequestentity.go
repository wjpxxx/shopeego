package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateDiscountItemItemListRequestEntity struct{
    ItemId	int64	`json:"item_id"`
    ModelList	[]UpdateDiscountItemModelListRequestEntity	`json:"model_list"`
    ItemPromotionPrice	float32	`json:"item_promotion_price"`
    PurchaseLimit	int	`json:"purchase_limit"`
}
func (g UpdateDiscountItemItemListRequestEntity) String() string {
    return lib.ObjectToString(g)
}
