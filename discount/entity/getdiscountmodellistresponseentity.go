package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountModelListResponseEntity struct{
    ModelName	string	`json:"model_name"`
    ModelNormalStock	int	`json:"model_normal_stock"`
    ModelOriginalPrice	float32	`json:"model_original_price"`
    ModelPromotionPrice	float32	`json:"model_promotion_price"`
    ModelId	int64	`json:"model_id"`
    ModelInflatedPriceOfOriginalPrice	float32	`json:"model_inflated_price_of_original_price"`
    ModelInflatedPriceOfPromotionPrice	float32	`json:"model_inflated_price_of_promotion_price"`
    ModelPromotionStock	int	`json:"model_promotion_stock"`
}
func (g GetDiscountModelListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
