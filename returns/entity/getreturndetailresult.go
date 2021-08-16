package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnDetailResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GetReturnDetailResponseResponseEntity	`json:"response"`
}
func (g GetReturnDetailResult) String() string {
    return lib.ObjectToString(g)
}
