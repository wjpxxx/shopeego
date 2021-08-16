package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type DeleteItemListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	DeleteItemListResponseResponseEntity	`json:"response"`
}
func (g DeleteItemListResult) String() string {
    return lib.ObjectToString(g)
}
