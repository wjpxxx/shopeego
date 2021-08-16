package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateBundleDealItemRequest struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
    ItemList	[]UpdateBundleDealItemItemListRequestEntity	`json:"item_list"`
}
func (g UpdateBundleDealItemRequest) String() string {
    return lib.ObjectToString(g)
}
