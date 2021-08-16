package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateAddOnDealMainItemMainItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
}
func (g UpdateAddOnDealMainItemMainItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
