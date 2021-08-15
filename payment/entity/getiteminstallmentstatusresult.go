package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetItemInstallmentStatusResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetItemInstallmentStatusResponseResponseEntity	`json:"response"`
}
func (g GetItemInstallmentStatusResult) String() string {
    return lib.ObjectToString(g)
}
