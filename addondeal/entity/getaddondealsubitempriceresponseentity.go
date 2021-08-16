package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetAddOnDealSubItemPriceResponseEntity struct{
    PromoInputPrice	float32	`json:"promo_input_price"`
}
func (g GetAddOnDealSubItemPriceResponseEntity) String() string {
    return lib.ObjectToString(g)
}
