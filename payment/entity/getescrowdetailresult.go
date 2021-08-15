package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetEscrowDetailResult struct{
    Response	GetEscrowDetailResponseResponseEntity	`json:"response"`
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
}
func (g GetEscrowDetailResult) String() string {
    return lib.ObjectToString(g)
}
