package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteDiscountItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	DeleteDiscountItemResponseResponseEntity	`json:"response"`
}
func (g DeleteDiscountItemResult) String() string {
    return lib.ObjectToString(g)
}
