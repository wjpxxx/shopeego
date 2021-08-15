package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowListResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetEscrowListResponseResponseEntity	`json:"response"`
}
func (g GetEscrowListResult) String() string {
    return lib.ObjectToString(g)
}
