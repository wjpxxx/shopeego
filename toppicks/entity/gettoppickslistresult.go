package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetTopPicksListResult struct{
    RequestId	string	`json:"request_id"`
    Error	string	`json:"error"`
    Message	string	`json:"message"`
    Response	GetTopPicksListResponseResponseEntity	`json:"response"`
}
func (g GetTopPicksListResult) String() string {
    return lib.ObjectToString(g)
}
