package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddAddOnDealMainItemMainItemListRequestEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
}
func (g AddAddOnDealMainItemMainItemListRequestEntity) String() string {
    return lib.ObjectToString(g)
}
