package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealItemRequest struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
}
func (g GetBundleDealItemRequest) String() string {
    return lib.ObjectToString(g)
}
