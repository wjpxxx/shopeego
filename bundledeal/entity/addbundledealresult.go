package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddBundleDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	AddBundleDealResponseResponseEntity	`json:"response"`
}
func (g AddBundleDealResult) String() string {
    return lib.ObjectToString(g)
}
