package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteBundleDealItemRequest struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
    ItemList	[]DeleteBundleDealItemItemListRequestEntity	`json:"item_list"`
}
func (g DeleteBundleDealItemRequest) String() string {
    return lib.ObjectToString(g)
}
