package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetPayoutDetailResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetPayoutDetailResponseResponseEntity	`json:"response"`
}
func (g GetPayoutDetailResult) String() string {
    return lib.ObjectToString(g)
}
