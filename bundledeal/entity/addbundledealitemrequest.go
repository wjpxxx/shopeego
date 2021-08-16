package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddBundleDealItemRequest struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
    ItemList	[]AddBundleDealItemItemListRequestEntity	`json:"item_list"`
}
func (g AddBundleDealItemRequest) String() string {
    return lib.ObjectToString(g)
}
