package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetWalletTransactionListResult struct{
    Response	GetWalletTransactionListResponseResponseEntity	`json:"response"`
    RequestId	string	`json:"request_id"`
    Message	string	`json:"message"`
    Error	string	`json:"error"`
}
func (g GetWalletTransactionListResult) String() string {
    return lib.ObjectToString(g)
}
