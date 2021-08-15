package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetShopInstallmentStatusResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	SetShopInstallmentStatusResponseResponseEntity	`json:"response"`
}
func (g SetShopInstallmentStatusResult) String() string {
    return lib.ObjectToString(g)
}
