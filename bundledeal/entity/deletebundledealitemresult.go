package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteBundleDealItemResult struct{
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Response	DeleteBundleDealItemResponseResponseEntity	`json:"response"`
}
func (g DeleteBundleDealItemResult) String() string {
    return lib.ObjectToString(g)
}
