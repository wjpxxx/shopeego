package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteDiscountItemErrorListResponseEntity struct{
    ItemId	int64	`json:"item_id"`
    ModleId	int64	`json:"modle_id"`
    FailMessage	string	`json:"fail_message"`
    FailError	string	`json:"fail_error"`
}
func (g DeleteDiscountItemErrorListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
