package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetDetailResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GetDetailResponseResponseEntity	`json:"response"`
}
func (g GetDetailResult) String() string {
    return lib.ObjectToString(g)
}
