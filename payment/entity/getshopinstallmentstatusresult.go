package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetShopInstallmentStatusResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetShopInstallmentStatusResponseResponseEntity	`json:"response"`
}
func (g GetShopInstallmentStatusResult) String() string {
    return lib.ObjectToString(g)
}
