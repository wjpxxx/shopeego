package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetReturnListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	[]GetReturnListResponseResponseEntity	`json:"response"`
}
func (g GetReturnListResult) String() string {
    return lib.ObjectToString(g)
}
