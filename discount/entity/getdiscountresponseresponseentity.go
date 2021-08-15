package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDiscountResponseResponseEntity struct{
    Status	string	`json:"status"`
    DiscountName	int	`json:"discount_name"`
    ItemList	[]GetDiscountItemListResponseEntity	`json:"item_list"`
    StartTime	int	`json:"start_time"`
    DiscountId	int64	`json:"discount_id"`
    EndTime	int	`json:"end_time"`
    More	bool	`json:"more"`
}
func (g GetDiscountResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
