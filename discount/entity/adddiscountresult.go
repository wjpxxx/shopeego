package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddDiscountResult struct{
    Message	string	`json:"message"`
    Error	string	`json:"error"`
    RequestId	string	`json:"request_id"`
    Response	AddDiscountResponseResponseEntity	`json:"response"`
}
func (g AddDiscountResult) String() string {
    return lib.ObjectToString(g)
}
