package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountItemModelListRequestEntity struct{
    ModelId	int64	`json:"model_id"`
    ModelPromotionPrice	float32	`json:"model_promotion_price"`
    ModelPromotionStock	int	`json:"model_promotion_stock"`
}
func (g AddDiscountItemModelListRequestEntity) String() string {
    return lib.ObjectToString(g)
}
