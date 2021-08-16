package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetItemListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GetItemListResponseResponseEntity	`json:"response"`
}
func (g GetItemListResult) String() string {
    return lib.ObjectToString(g)
}
