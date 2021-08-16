package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetVoucherDetailResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetVoucherDetailResponseResponseEntity	`json:"response"`
}
func (g GetVoucherDetailResult) String() string {
    return lib.ObjectToString(g)
}
