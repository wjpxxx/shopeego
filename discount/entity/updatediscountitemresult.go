package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateDiscountItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Warning	string	`json:"warning"`
    RequestId	string	`json:"request_id"`
    Response	UpdateDiscountItemResponseResponseEntity	`json:"response"`
}
func (g UpdateDiscountItemResult) String() string {
    return lib.ObjectToString(g)
}
