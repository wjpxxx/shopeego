package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealItemResponseResponseEntity struct{
    ItemList	[]GetBundleDealItemItemListResponseEntity	`json:"item_list"`
    TotalCount	int	`json:"total_count"`
}
func (g GetBundleDealItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
