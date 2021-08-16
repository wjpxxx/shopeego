package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealMainItemMainItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
}
func (g AddAddOnDealMainItemMainItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
