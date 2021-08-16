package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealItemItemListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    Status	int	`json:"status"`
}
func (g GetBundleDealItemItemListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
