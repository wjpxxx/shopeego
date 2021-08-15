package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountItemListResponseEntity struct{
    ItemPromotionPrice	float32	`json:"item_promotion_price"`
    ItemName	string	`json:"item_name"`
    ModelList	[]GetDiscountModelListResponseEntity	`json:"model_list"`
    ItemId	int64	`json:"item_id"`
    PurchaseLimit	int	`json:"purchase_limit"`
    ItemOriginalPrice	float32	`json:"item_original_price"`
    NormalStock	int	`json:"normal_stock"`
    ItemInflatedPriceOfOriginalPrice	float32	`json:"item_inflated_price_of_original_price"`
    ItemInflatedPriceOfPromotionPrice	float32	`json:"item_inflated_price_of_promotion_price"`
    ItemPromotionStock	int	`json:"item_promotion_stock"`
}
func (g GetDiscountItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
