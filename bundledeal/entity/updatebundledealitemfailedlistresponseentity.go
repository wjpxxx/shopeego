package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateBundleDealItemFailedListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    FailError	string	`json:"fail_error"`
    FailMessage	string	`json:"fail_message"`
}
func (g UpdateBundleDealItemFailedListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
