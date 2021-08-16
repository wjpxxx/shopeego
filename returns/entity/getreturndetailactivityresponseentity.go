package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailActivityResponseEntity struct{
    ActivityId	int64	`json:"activity_id"`
    ActivityType	string	`json:"activity_type"`
    OriginalPrice	string	`json:"original_price"`
    DiscountedPrice	string	`json:"discounted_price"`
    Items	[]GetReturnDetailItemsResponseEntity	`json:"items"`
}
func (g GetReturnDetailActivityResponseEntity) String() string {
    return lib.ObjectToString(g)
}
