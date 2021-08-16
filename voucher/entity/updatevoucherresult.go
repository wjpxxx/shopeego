package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type UpdateVoucherResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	UpdateVoucherResponseResponseEntity	`json:"response"`
}
func (g UpdateVoucherResult) String() string {
    return lib.ObjectToString(g)
}
