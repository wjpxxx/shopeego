package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealListResponseResponseEntity struct{
    BundleDealList	[]GetBundleDealListBundleDealListResponseEntity	`json:"bundle_deal_list"`
    More	bool	`json:"more"`
}
func (g GetBundleDealListResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
