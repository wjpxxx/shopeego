package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateTopPicksItemListResponseEntity struct{
    ItemName	string	`json:"item_name"`
    ItemId	int64	`json:"item_id"`
    CurrentPrice	float32	`json:"current_price"`
    InflatedPriceOfCurrentPrice	float32	`json:"inflated_price_of_current_price"`
    Sales	int	`json:"sales"`
}
func (g UpdateTopPicksItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
