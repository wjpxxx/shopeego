package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteBundleDealResponseResponseEntity struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
}
func (g DeleteBundleDealResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
