package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateBundleDealItemResponseResponseEntity struct{
    FailedList	[]UpdateBundleDealItemFailedListResponseEntity	`json:"failed_list"`
    SuccessList	[]int64	`json:"success_list"`
}
func (g UpdateBundleDealItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
