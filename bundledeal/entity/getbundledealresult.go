package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetBundleDealResponseResponseEntity	`json:"response"`
}
func (g GetBundleDealResult) String() string {
    return lib.ObjectToString(g)
}
