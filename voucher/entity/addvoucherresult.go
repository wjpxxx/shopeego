package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddVoucherResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	AddVoucherResponseResponseEntity	`json:"response"`
}
func (g AddVoucherResult) String() string {
    return lib.ObjectToString(g)
}
