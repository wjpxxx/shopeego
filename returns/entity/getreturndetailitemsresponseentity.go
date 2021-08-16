package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailItemsResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    VariationId	int64	`json:"variation_id"`
    QuantityPurchased	int	`json:"quantity_purchased"`
    OriginalPrice	string	`json:"original_price"`
}
func (g GetReturnDetailItemsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
