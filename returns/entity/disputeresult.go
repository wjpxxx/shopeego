package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DisputeResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	DisputeResponseResponseEntity	`json:"response"`
}
func (g DisputeResult) String() string {
    return lib.ObjectToString(g)
}
