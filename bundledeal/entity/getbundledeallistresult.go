package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetBundleDealListResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	GetBundleDealListResponseResponseEntity	`json:"response"`
}
func (g GetBundleDealListResult) String() string {
    return lib.ObjectToString(g)
}
