package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndBundleDealRequest struct{
    BundleDealId	int64	`json:"bundle_deal_id"`
}
func (g EndBundleDealRequest) String() string {
    return lib.ObjectToString(g)
}
