package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type AddBundleDealItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	AddBundleDealItemResponseResponseEntity	`json:"response"`
}
func (g AddBundleDealItemResult) String() string {
    return lib.ObjectToString(g)
}
