package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddBundleDealResponseResponseEntity struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
}
func (g AddBundleDealResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
