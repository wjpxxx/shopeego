package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountItemItemIstRequestEntity struct{
    ItemId	int64	`json:"item_id"`
    ModelList	[]AddDiscountItemModelListRequestEntity	`json:"model_list"`
    ItemPromotionPrice	float32	`json:"item_promotion_price"`
    PurchaseLimit	int	`json:"purchase_limit"`
    ItemPromotionStock	int	`json:"item_promotion_stock"`
}
func (g AddDiscountItemItemIstRequestEntity) String() string {
    return lib.ObjectToString(g)
}
