package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndVoucherResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	EndVoucherResponseResponseEntity	`json:"response"`
}
func (g EndVoucherResult) String() string {
    return lib.ObjectToString(g)
}
