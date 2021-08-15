package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	AddDiscountItemResponseResponseEntity	`json:"response"`
}
func (g AddDiscountItemResult) String() string {
    return lib.ObjectToString(g)
}
