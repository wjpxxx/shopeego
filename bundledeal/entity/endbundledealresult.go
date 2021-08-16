package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type EndBundleDealResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	EndBundleDealResponseResponseEntity	`json:"response"`
}
func (g EndBundleDealResult) String() string {
    return lib.ObjectToString(g)
}
