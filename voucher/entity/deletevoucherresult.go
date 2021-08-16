package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteVoucherResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	DeleteVoucherResponseResponseEntity	`json:"response"`
}
func (g DeleteVoucherResult) String() string {
    return lib.ObjectToString(g)
}
