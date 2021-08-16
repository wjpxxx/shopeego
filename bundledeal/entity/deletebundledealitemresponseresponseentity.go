package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteBundleDealItemResponseResponseEntity struct{
    FailedList	[]DeleteBundleDealItemFailedListResponseEntity	`json:"failed_list"`
    SuccessList	[]int64	`json:"success_list"`
}
func (g DeleteBundleDealItemResponseResponseEntity) String() string {
    return lib.ObjectToString(g)
}
